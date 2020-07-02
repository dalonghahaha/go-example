package services

import (
	"gin-demo/models/comjia"
)

type BlackService struct {
}

func NewBlackService() *BlackService {
	return &BlackService{}
}

func (s *BlackService) GetAll() []*comjia.CjBlackWord {
	list, err := new(comjia.CjBlackWord).GetAll()
	if err != nil {
		return nil
	}
	return list
}

func (s *BlackService) GetBlackWord(id int64) *comjia.CjBlackWord {
	blackWord := new(comjia.CjBlackWord)
	err := blackWord.GetByID(id)
	if err != nil {
		return nil
	}
	return blackWord
}

func (s *BlackService) CreateBlackWord(blackWord *comjia.CjBlackWord) bool {
	return blackWord.Create() == nil
}

func (s *BlackService) UpdateBlackWord(blackWord *comjia.CjBlackWord) bool {
	return blackWord.Update() == nil
}

func (s *BlackService) DeleteBlackWord(id int64) bool {
	blackWord := new(comjia.CjBlackWord)
	blackWord.ID = id
	return blackWord.Delete() == nil
}
