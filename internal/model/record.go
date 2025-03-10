package model

import "time"

type Record struct {
	ID             int       `gorm:"primaryKey;uniqueIndex;not null" json:"id"`
	ExcavatorName  string    `gorm:"not null" json:"excavator_name"`
	Date           time.Time `gorm:"not null" json:"date"`
	Shift          string    `gorm:"not null" json:"shift"`
	ShiftTime      int       `gorm:"not null" json:"shift_time"`
	LoadTime       float64   `gorm:"not null" json:"load_time"`
	CycleTime      int       `gorm:"not null" json:"cycle_time"`
	ApproachTime   int       `gorm:"not null" json:"approach_time"`
	ActualTrucks   float64   `gorm:"not null" json:"actual_trucks"`
	Productivity   int       `gorm:"not null" json:"productivity"`
	RequiredTrucks float64   `gorm:"not null" json:"required_trucks"`
	PlanVolume     float64   `gorm:"not null" json:"plan_volume"`
	ForecastVolume float64   `gorm:"not null" json:"forecast_volume"`
	Downtime       float64   `gorm:"not null" json:"downtime"`
	UserID         int       `gorm:"not null" json:"user_id"`
}

type RecordForAnalise struct {
	ID             int       `json:"id"`
	ExcavatorName  string    `json:"excavator_name"`
	Date           time.Time `json:"date"`
	Shift          string    `json:"shift"`
	ShiftTime      int       `json:"shift_time"`
	LoadTime       float64   `json:"load_time"`
	CycleTime      int       `json:"cycle_time"`
	ApproachTime   int       `json:"approach_time"`
	ActualTrucks   float64   `json:"actual_trucks"`
	Productivity   int       `json:"productivity"`
	RequiredTrucks float64   `json:"required_trucks"`
	PlanVolume     float64   `json:"plan_volume"`
	ForecastVolume float64   `json:"forecast_volume"`
	Downtime       float64   `json:"downtime"`
	UserName       string    `json:"user_name"`
	UserSurname    string    `json:"user_surname"`
	Company        string    `json:"company"`
	Section        string    `json:"section"`
	JobTitle       string    `json:"job_title"`
}
