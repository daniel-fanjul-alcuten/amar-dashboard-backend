package amar

import "time"

type MyStuff struct {
	Name      string
	Total     int
	Inventory int
	House     int
	Shared    int
	Guild     int
	Link      string
}

type MyStuffPage struct {
	Time  time.Time
	Stuff map[string]MyStuff
}
