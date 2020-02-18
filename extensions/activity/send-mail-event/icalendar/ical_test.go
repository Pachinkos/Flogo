package icalendar

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var ical = &IcsCreator{}

func TestIcsCreator(t *testing.T) {
	lines := ical.Init().GenerateIcs()
	assert.NotNil(t, lines)
	assert.True(t, len(lines)>3)
}