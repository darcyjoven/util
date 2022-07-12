package util

import "time"

func GetShortDate() string {
	return time.Now().Format("20060102")
}
func GetLongDate() string {
	return time.Now().Format("20060102150405000")
}
