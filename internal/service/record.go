package service

import (
	"samosvulator/internal/model"
	"samosvulator/internal/repository"
)

type RecordService struct {
	repo repository.Record
}

func NewRecordService(repo repository.Record) *RecordService {
	return &RecordService{repo: repo}
}

func (s *RecordService) CreateRecord(record model.Record) error {
	return s.repo.CreateRecord(record)
}
