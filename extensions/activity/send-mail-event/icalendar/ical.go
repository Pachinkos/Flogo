package icalendar

import (
	"github.com/google/uuid"
	"time"
)

const DATE_FORMAT string = "20060102T150405Z"

type IcsCreator struct {
	header []string
	events []Event
	footer []string
}

func (ic *IcsCreator) Init() *IcsCreator {
	ic.header = []string{
		"BEGIN:VCALENDAR",
		"VERSION:2.0",
		"PRODID:cloud.tibco.com",
	}

	ic.footer = []string{
		"END:VCALENDAR",
	}

	return ic
}

func (ic *IcsCreator) WithEvent(event Event) *IcsCreator {
	ic.events = append(ic.events, event)
	return ic
}

func (ic *IcsCreator) GenerateIcs() []string {
	lines := []string{}
	lines = append(lines, ic.header...)

	for _, e := range ic.events {
		lines = append(lines, ic.generateEvent(e)...)
	}

	lines = append(lines, ic.footer...)

	return lines
}

func (ic *IcsCreator) generateEvent(event Event) []string {
	lines := []string{}
	lines = append(lines, "BEGIN:VEVENT")
	lines = append(lines, "SUMMARY:"+event.Summary)
	lines = append(lines, "UID:"+event.Uid)
	lines = append(lines, "DESCRIPTION:"+event.Description)
	lines = append(lines, "LOCATION:"+event.Location)
	startTime := event.StartTime.UTC().Format(DATE_FORMAT)
	lines = append(lines, "DTSTAMP:"+startTime)
	lines = append(lines, "DTSTART:"+startTime)
	lines = append(lines, "DTEND:"+event.EndTime.UTC().Format(DATE_FORMAT))
	lines = append(lines, "END:VEVENT")

	return lines
}

type Event struct {
	Uid         string
	Summary     string
	Description string
	Location    string
	StartTime   time.Time
	EndTime     time.Time
}

func (e *Event) Init() *Event {
	id, err := uuid.NewRandom()
	if err != nil {
		e.Uid = time.Now().UTC().String()
	} else {
		e.Uid = id.String()
	}

	return e
}
