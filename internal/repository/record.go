package repository

import (
	"gorm.io/gorm"
	"samosvulator/internal/model"
)

type RecordRepository struct {
	db *gorm.DB
}

func NewRecordRepository(db *gorm.DB) *RecordRepository {
	return &RecordRepository{db: db}
}

func (r *RecordRepository) CreateRecord(record model.Record) error {
	return r.db.Create(&record).Error
}

func (r *RecordRepository) GetAllRecords() ([]model.Record, error) {
	var records []model.Record
	result := r.db.Find(&records)
	return records, result.Error
}
