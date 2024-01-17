package model

import (
	"regexp"
	"strings"
	"time"

	"github.com/leighmacdonald/steamid/v3/steamid"
)

type FilterAction int

const (
	Kick FilterAction = iota
	Mute
	Ban
)

type Filter struct {
	FilterID     int64          `json:"filter_id"`
	AuthorID     steamid.SID64  `json:"author_id"`
	Pattern      string         `json:"pattern"`
	IsRegex      bool           `json:"is_regex"`
	IsEnabled    bool           `json:"is_enabled"`
	Action       FilterAction   `json:"action"`
	Duration     string         `json:"duration"`
	Regex        *regexp.Regexp `json:"-"`
	TriggerCount int64          `json:"trigger_count"`
	Weight       int            `json:"weight"`
	CreatedOn    time.Time      `json:"created_on"`
	UpdatedOn    time.Time      `json:"updated_on"`
}

func (f *Filter) Init() {
	if f.IsRegex {
		f.Regex = regexp.MustCompile(f.Pattern)
	}
}

func (f *Filter) Match(value string) bool {
	if f.IsRegex {
		return f.Regex.MatchString(strings.ToLower(value))
	}

	return f.Pattern == strings.ToLower(value)
}
