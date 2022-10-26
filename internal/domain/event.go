package domain

import "time"

type EventsList []Event

type Event struct {
	ID        int64      `db:"id" json:"id"  description:""`
	CreatedAt time.Time  `db:"created_at" json:"created_at"  description:""`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"  description:""`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"  description:""`
	Name      string     `db:"name" json:"name"  description:""`
	Date      string     `db:"date" json:"date"  description:""`
	Park      string     `db:"park" json:"park"  description:""`
	Weather   Weather    `db:"weather" json:"weather" description:""`
}

type Weather struct {
	Kind          string        `db:"kind" json:"kind" description:""`
	Temperature   float32       `db:"temperature" json:"temperature" description:""`
	Precipitation Precipitation `db:"recipitation" json:"recipitation" description:""`
	Wind          int64         `db:"wind" json:"wind" description:""`
}

type Precipitation struct {
	Type      int64   `db:"type" json:"type" description:""`
	Amount    float32 `db:"amount" json:"amount" description:""`
	Intensity int64   `db:"intensity" json:"intensity" description:""`
}
