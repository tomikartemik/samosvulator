package service

import (
	"samosvulator/internal/model"
	"samosvulator/internal/repository"
	"strconv"
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

func (s *RecordService) GetAllRecords() ([]model.Record, error) {
	return s.repo.GetAllRecords()
}

func (s *RecordService) GetRecordsByUserID(idStr string) ([]model.Record, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	records, err := s.repo.GetRecordsByUserID(id)
	if err != nil {
		return nil, err
	}
	return records, nil
}
