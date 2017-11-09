package core

import "time"

// -- Time & date helpers

func GetJodaTime() string {
	return time.Now().Format("2006-01-02T15:04:05.999Z")
}
