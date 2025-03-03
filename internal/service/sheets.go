package service

import (
	"samosvulator/internal/model"
	"samosvulator/internal/repository"
	"samosvulator/internal/utils"
)

type SheetsService struct {
	repoR repository.Record
	repoU repository.User
}

func NewSheetsService(repoR repository.Repository, repoU repository.User) *SheetsService {
	return &SheetsService{repoR: repoR, repoU: repoU}
}

func (s *SheetsService) GetRecordsForAnalise() ([]model.RecordForAnalise, error) {
	recordForAnalise := []model.RecordForAnalise{}
	records, err := s.repoR.GetAllRecords()
	if err != nil {
		return nil, err
	}
	for _, record := range records {
		user, err := s.repoU.GetUserByID(record.UserID)
		if err != nil {
			return nil, err
		}
		recordForAnalise = append(recordForAnalise, utils.ConvertToRecordForAnalise(record, user))
	}
	return recordForAnalise, nil
}
