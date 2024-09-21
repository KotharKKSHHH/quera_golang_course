package main

import "strconv"

func ConvertToDigitalFormat(hour, minute, second int) string {
	H := strconv.Itoa(hour)
	M := strconv.Itoa(minute)
	S := strconv.Itoa(second)
	if len(H) < 2 {
		H = "0" + H
	}
	if len(M) < 2 {
		M = "0" + M
	}
	if len(S) < 2 {
		S = "0" + S
	}
	re := H + ":" + M + ":" + S
	return re
}

func ExtractTimeUnits(seconds int) (int, int, int) {
	hour := seconds / 3600
	seconds = seconds % 3600
	minute := seconds / 60
	seconds = seconds % 60
	return hour, minute, seconds
}
