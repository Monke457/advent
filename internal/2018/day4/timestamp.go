package main

type Timestamp struct {
	year int
	month int 
	day int
	hour int
	minute int
}

//-1 earlier, 0 same, 1 later
func (t Timestamp) Compare(t2 Timestamp) int {
	if t.year > t2.year {
		return 1 
	}
	if t.year < t2.year {
		return -1
	}
	if t.month > t2.month {
		return 1
	}
	if t.month < t2.month {
		return -1
	}
	if t.day > t2.day{
		return 1
	}
	if t.day < t2.day{
		return -1
	}
	if t.hour > t2.hour {
		return 1
	}
	if t.hour < t2.hour {
		return -1 
	}
	if t.minute > t2.minute {
		return 1
	}
	if t.minute < t2.minute {
		return -1 
	}
	return 0
}
