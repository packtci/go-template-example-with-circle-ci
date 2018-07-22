package main

import (
	"testing"
)

func TestParseTemplate(t *testing.T) {
	newSoldier := Soldier{
		Name:          "Luke Cage",
		Rank:          "SGT",
		TimeInService: 4,
	}
	txt := parseTemplate(newSoldier, templateText)
	expectedTxt := `
Name is Luke Cage
Rank is SGT
Time in service is 4
`
	if txt.String() != expectedTxt {
		t.Error("The text returned should match")
	}
}
