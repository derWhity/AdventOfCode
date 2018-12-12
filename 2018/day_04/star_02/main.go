package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const timeFormat = "2006-01-02 15:04"

var (
	entryRegEx   = regexp.MustCompile(`^\[([^\]]+)\]\s(.*)$`)
	guardIDRegEx = regexp.MustCompile(`#([0-9]+)`)
)

type guardRecord struct {
	AsleepHeatMap []uint
}

type logEntry struct {
	Timestamp time.Time
	Awake     bool
	GuardID   int
}

type log []*logEntry

// Implementing sort.Interface

func (l log) Len() int {
	return len(l)
}

func (l log) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l log) Less(i, j int) bool {
	return l[i].Timestamp.Before(l[j].Timestamp)
}

func newGuardRecord() *guardRecord {
	return &guardRecord{
		AsleepHeatMap: make([]uint, 60), // Initialize empty heatmap
	}
}

func unmarshalEntry(str string) *logEntry {
	match := entryRegEx.FindStringSubmatch(str)
	t, err := time.Parse(timeFormat, match[1])
	if err != nil {
		panic("Failed to parse time: " + err.Error())
	}
	entry := logEntry{
		Timestamp: t,
	}
	line := match[2]
	switch line {
	case "falls asleep":
		// Nothing to do
	case "wakes up":
		entry.Awake = true
	default:
		// This is the start of the shift - get the ID
		match = guardIDRegEx.FindStringSubmatch(line)
		id, _ := strconv.ParseInt(match[1], 10, strconv.IntSize)
		entry.GuardID = int(id)
	}
	return &entry
}

func exec(input string) (int, int) {
	rows := strings.Split(input, "\n")
	logEntries := log{}
	for _, row := range rows {
		entry := unmarshalEntry(row)
		logEntries = append(logEntries, entry)
	}
	sort.Sort(logEntries)
	guards := make(map[int]*guardRecord)
	var currentRecord *guardRecord
	var currentMinute int
	var wasAsleep bool
	var displayStr string
	fmt.Println("                000000000011111111112222222222333333333344444444445555555555")
	fmt.Println("                012345678901234567890123456789012345678901234567890123456789")
	for _, row := range logEntries {
		if row.GuardID > 0 {
			// Starting a shift
			// Finish the last shift first
			fmt.Println(fillRecord(currentRecord, currentMinute, 59, wasAsleep, displayStr))
			// Prepare the new guard
			record, ok := guards[row.GuardID]
			if !ok {
				record = newGuardRecord()
				guards[row.GuardID] = record
			}
			currentRecord = record
			if row.Timestamp.Hour() != 0 {
				// Before midnight
				currentMinute = 0
				displayStr = fmt.Sprintf("%2d-%2d [%6d]: ", row.Timestamp.Day()+1, row.Timestamp.Month(), row.GuardID)
			} else {
				// After midnight
				currentMinute = row.Timestamp.Minute()
				// Add the missed minutes to the display string
				displayStr = fmt.Sprintf("%2d-%2d [%6d]: ", row.Timestamp.Day(), row.Timestamp.Month(), row.GuardID)
				displayStr += strings.Repeat("-", row.Timestamp.Minute())
			}
		} else {
			// Mid-shift event
			// fill up until now
			displayStr = fillRecord(currentRecord, currentMinute, row.Timestamp.Minute()-1, wasAsleep, displayStr)
			// prepare continuation
			currentMinute = row.Timestamp.Minute()
			wasAsleep = !row.Awake
		}
	}
	// Finalize the last shift and output it
	fmt.Println(fillRecord(currentRecord, currentMinute, 59, wasAsleep, displayStr))
	// Find the guard that has the highest sleep time on a specific minute
	candidateID := 0
	bestMinute := 0
	bestCount := 0
	for id, row := range guards {
		for minute, count := range row.AsleepHeatMap {
			if bestCount < int(count) {
				fmt.Printf("Candidate #%4d @ %d (%d)\n", id, minute, count)
				candidateID = id
				bestMinute = minute
				bestCount = int(count)
			}
		}
	}

	return candidateID, bestMinute
}

func fillRecord(record *guardRecord, from, to int, asleep bool, displayStr string) string {
	//fmt.Printf("fill: from(%d), to(%d), asleep(%t)\n", from, to, asleep)
	if record != nil && from < 60 && to >= from {
		if asleep {
			for i := from; i <= to; i++ {
				record.AsleepHeatMap[i]++
			}
			displayStr += strings.Repeat("X", to-from+1)
		} else {
			displayStr += strings.Repeat(".", to-from+1)
		}
	}
	return displayStr
}

func main() {
	input, err := ioutil.ReadFile(filepath.Join("..", "input.txt"))
	if err != nil {
		panic(err)
	}
	guardID, minute := exec(string(input))
	fmt.Printf("Guard #%d best sleeping on minute %d - Result: %d\n", guardID, minute, guardID*minute)
}
