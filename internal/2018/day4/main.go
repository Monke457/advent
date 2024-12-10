package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"strconv"
	"strings"
)

type action int

const(
	START action = iota
	SLEEP	
	WAKE
)

type timestamp struct {
	year int
	month int 
	day int
	hour int
	minute int
}

func (t timestamp) after(t2 timestamp) bool {
	if t.year == t2.year {
		if t.month == t2.month {
			if t.day == t2.day{
				if t.hour == t2.hour {
					return t.minute > t2.minute
				}
				return t.hour > t2.hour
			}
			return t.day > t2.day
		}
		return t.month > t2.month
	}
	return t.year > t2.year
}

func (t timestamp) daysOfMonth() int {
	switch t.month {
	case 1: return 31
	case 2:	return 28
	case 3:	return 31
	case 4:	return 30
	case 5:	return 31
	case 6:	return 30
	case 7:	return 31
	case 8:	return 31
	case 9:	return 30
	case 10: return 31
	case 11: return 30
	case 12: return 31
	}
	return 0
}

func (t timestamp) difference(t2 timestamp) int {
	if t.after(t2) {
		return t2.difference(t)
	}
	ydiff := (t2.year - t.year) * 365 * 24 * 60
	mdiff := (t2.month - t.month) * t.daysOfMonth() * 24 * 60
	ddiff := (t2.day - t.day) * 24 * 60
	hdiff := (t2.hour - t.hour) * 60
	mindiff := (t2.minute - t.minute)

	return ydiff + mdiff + ddiff + hdiff + mindiff
}

func (t timestamp) add(mins int) timestamp {
	newT := timestamp{t.year, t.month, t.day, t.hour, t.minute}
	newT.minute += mins
	if newT.minute > 59 {
		newT.hour++
		newT.minute = 0
	}
	if newT.hour > 23 {
		newT.day++
		newT.hour = 0
	}
	if newT.day > newT.daysOfMonth() {
		newT.month++
		newT.day = 1
	}
	if newT.month > 12 {
		newT.year++
		newT.month = 1
	}
	return newT 
}

func main() {
	data := reader.FileToArray("data/2018/day4.txt")

	times, actions := parseList(data)
	order := getSortOrder(times)
	minutesSleeping := mapMinutesSleeping(times, actions, order)

	sleepiestGuard := 0
	sleepytime := 0
	for id, st := range minutesSleeping {
		total := 0
		for _, m := range st {
			total += m
		}
		if total > sleepytime {
			sleepiestGuard = id
			sleepytime = total
		}
	}

	sleepiestMinute := 0
	sleepytime = 0
	for minute, count := range minutesSleeping[sleepiestGuard] {
		if count > sleepytime {
			sleepiestMinute = minute
			sleepytime = count
		}
	}

	fmt.Println("sleepiest guard:", sleepiestGuard)
	fmt.Println("sleepiest minute:", sleepiestMinute)
	fmt.Println("First:", sleepiestGuard * sleepiestMinute)

	oldReliable := 0
	frequency := 0
	sleepiestMinute = 0
	for id, minuteMap := range minutesSleeping {
		for minute, count := range minuteMap {
			if count > frequency {
				oldReliable = id
				frequency = count
				sleepiestMinute = minute
			}
		}
	}

	fmt.Println("sleepiest guard:", oldReliable)
	fmt.Println("sleepiest minute:", sleepiestMinute)
	fmt.Println("Second:", oldReliable * sleepiestMinute)

}

func mapMinutesSleeping(times map[int]timestamp, actions map[int]string, order []int) map[int]map[int]int {
	minutesSleeping := map[int]map[int]int{}
	guard := 0
	var sleepbegin timestamp

	for _, idx := range order {
		act, next := parseAction(actions[idx], guard)
		guard = next

		if act == SLEEP {
			sleepbegin = times[idx]
		}

		if act == WAKE {
			diff := sleepbegin.difference(times[idx]) 
			if _, ok := minutesSleeping[guard]; !ok {
				minutesSleeping[guard] = map[int]int{}
			}
			for i := range diff {
				m := sleepbegin.add(i).minute
				minutesSleeping[guard][m]++
			}
		}
	}
	return minutesSleeping
}

func getSortOrder(ts map[int]timestamp) []int {
	order := make([]int, len(ts))
	for i := range len(order) {
		order[i] = i
	}

	for {
		ordered := true
		for i := 0; i < len(order)-1; i++ {
			if ts[order[i]].after(ts[order[i+1]]) {
				order[i] = order[i] ^ order[i+1]
				order[i+1] = order[i] ^ order[i+1]
				order[i] = order[i] ^ order[i+1]
				ordered = false
			}
		}

		if ordered {
			break
		}
	}
	return order
}

func parseList(list []string) (map[int]timestamp, map[int]string) {
	t := map[int]timestamp{}
	a := map[int]string{}

	for i, line := range list {
		timestamp, action, _ := strings.Cut(line[1:], "] ")

		t[i] = parseTime(timestamp)
		a[i] = action
	}

	return t, a
}

func parseAction(action string, guard int) (action, int) {
	if idx := strings.Index(action, "#"); idx > -1 {
		id, _, _ := strings.Cut(action[idx+1:], " ")
		guard, _ = strconv.Atoi(id)
		return START, guard
	}

	if strings.Contains(action, "asleep") {
		return SLEEP, guard
	}
	if strings.Contains(action, "wakes") {
		return WAKE, guard
	}
	panic("big skill issues")
}

func parseTime(ts string) timestamp {
	date, time, _ := strings.Cut(ts, " ")

	dateParts := strings.Split(date, "-")
	year, _ := strconv.Atoi(dateParts[0])
	month, _ := strconv.Atoi(dateParts[1])
	day, _ := strconv.Atoi(dateParts[2])

	hourstr, minutestr, _ := strings.Cut(time, ":")
	hour, _ := strconv.Atoi(hourstr)
	minute, _ := strconv.Atoi(minutestr)

	return timestamp{year, month, day, hour, minute}
}
