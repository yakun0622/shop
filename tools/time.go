package tools

import (
	"time"
)

func GetTime() uint {
	return uint(time.Now().UnixNano() / 1e9)
}

func GetTimeString(types uint8) string {
	switch types {
	case 1:
		return time.Now().Format("2006.01.02 15:04:05")
	case 2:
		return time.Now().Format("20060102150405")
	case 3:
		return time.Now().Format("2006.01.02")
	case 4:
		return time.Now().Format("2006-01-02")
	default:
		return time.Now().Format("2006-01-02 15:04:05")
	}
}
