package pt

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type DigitalFormat uint8
type DigitalResolution uint8

const (
	Blueray DigitalFormat = iota
	HDTV
	WebDL
	UHDTV
	UnknownDigitalFormat
)

const (
	FHD DigitalResolution = iota
	HD
	UHD4K
	UnknownResolution
)

var DigitalFormatMap = map[DigitalFormat][]string{
	Blueray: []string{"bluray", "blu-ray", "blueray", "bd"},
	HDTV:    []string{"hdtv"},
	WebDL:   []string{"webdl", "web-dl"},
	UHDTV:   []string{"uhdtv"},
}

var DigitalResolutionMap = map[DigitalResolution][]string{
	FHD:   []string{"1080", "1080p", "1080i"},
	HD:    []string{"720", "720p"},
	UHD4K: []string{"4k"},
}

type MovieInfo struct {
	Title      string
	Year       int
	Group      string
	Source     DigitalFormat
	Resolution DigitalResolution
}

func ParseHDCTitle(title string) MovieInfo {
	m := MovieInfo{}
	if title == "" {
		return m
	}

	fields := split(title)

	year, yearIndex := findYear(fields)
	source, sourceIndex := findSource(fields)
	resolution, resIndex := findResolution(fields)
	group := findGroup(fields)

	minIndex := minPositive(yearIndex, sourceIndex, resIndex)

	movieTitle := strings.Join(fields[:minIndex], " ")

	return MovieInfo{
		movieTitle, year, group, source, resolution,
	}
}

func split(title string) []string {
	return strings.FieldsFunc(title, func(r rune) bool {
		return (r == '.' || r == ' ')
	})
}

func findYear(fields []string) (int, int) {
	for i, f := range fields {
		if year, err := tryParseYear(f); err == nil {
			return year, i
		}
	}
	return -1, -1
}

func tryParseYear(yyyy string) (int, error) {
	if len(yyyy) != 4 {
		return -1, fmt.Errorf("%s is not 4 digit year", yyyy)
	}

	if yyyy[0] != '1' && yyyy[0] != '2' {
		return -1, fmt.Errorf("%s is not 1xxx or 2xxx, not supported year range", yyyy)
	}

	return strconv.Atoi(yyyy)
}

func findSource(fields []string) (DigitalFormat, int) {
	for i, field := range fields {
		for format, names := range DigitalFormatMap {
			if contains(names, strings.ToLower(field)) {
				return format, i
			}
		}
	}
	return UnknownDigitalFormat, -1
}

func findResolution(fields []string) (DigitalResolution, int) {
	for i, field := range fields {
		for format, names := range DigitalResolutionMap {
			if contains(names, strings.ToLower(field)) {
				return format, i
			}
		}
	}
	return UnknownResolution, -1
}

func findGroup(fields []string) string {
	if len(fields) <= 0 {
		return ""
	}

	last := fields[len(fields)-1]

	if i := strings.LastIndex(last, "-"); i >= 0 {
		group := last[i+1:]
		if ii := strings.LastIndex(group, "@"); ii >= 0 {
			group = group[ii+1:]
		}
		return group
	}

	return ""
}

func contains(array []string, s string) bool {
	for _, e := range array {
		if e == s {
			return true
		}
	}
	return false
}

func minPositive(ints ...int) int {
	m := math.MaxInt32
	for _, i := range ints {
		if i >= 0 && i < m {
			m = i
		}
	}
	return m
}
