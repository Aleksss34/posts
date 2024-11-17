package domain

import "time"

type Post struct {
	Id     int
	Author string
	Post   string
	Time   time.Time
}
