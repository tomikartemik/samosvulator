package model

import "time"

type Record struct {
	ID            int       `gorm:"primaryKey;uniqueIndex;not null" json:"id"`
	ExcavatorName string    `gorm:"not null" json:"excavator_name"`
	Date          time.Time `gorm:"not null" json:"date"`
	Shift         string    `gorm:"not null" json:"shift"`
	ShiftTime     int       `gorm:"not null" json:"shift_time"`
	LoadTime      float64   `gorm:"not null" json:"load_time"`
	CycleTime     int       `gorm:"not null" json:"cycle_time"`
	ApproachTime  int       `gorm:"not null" json:"approach_time"`
	ActualTrucks  float64   `gorm:"not null" json:"actual_trucks"`
	Productivity  int       `gorm:"not null" json:"8"`
	UserID        int       `gorm:"not null" json:"user_id"`
}
