package models

import (
	"github.com/jinzhu/gorm"
)

// Series stores a set of events that are all interrelated.
type Series struct {
	gorm.Model

	Name        string
	Description string
	Archived    bool

	Events []Event
}

// Event stores a single event that may be part of a series.
type Event struct {
	gorm.Model

	Title       string
	Description string

	Questions []Question

	SeriesID uint
}

// A Question is part of an event and is asked in response to it.  It
// can have comments, but it also have votes.
type Question struct {
	gorm.Model

	Votes    []Vote
	Comments []Comment

	EventID uint
}

// A Vote is a positive or negative stance on a question, and is used
// to make an ordered set of votes.
type Vote struct {
	gorm.Model

	// Value can be set to -1 or +1 to express a down or up
	// sentiment.  This requires slightly more work to detect, but
	// it does ultimately work.
	Value int

	QuestionID uint
}

// A Comment is part of a discussion on a Question.  Comments belong
// to a user and are ordered by date.
type Comment struct {
	gorm.Model

	Text string

	QuestionID uint
}
