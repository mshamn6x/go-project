package db

import "new/test/project/api/model"

type RAM interface {
	Insert(*model.RAM) (*model.RAM, error)
	GetAll() ([]*model.RAM, error)
}

type RamManger struct {
}

func NewRamManger() RAM {
	return &RamManger{}
}

func (cm *RamManger) Insert(data *model.RAM) (*model.RAM, error) {
	result := db.Create(data)
	return data, result.Error
}

func (cm *RamManger) GetAll() ([]*model.RAM, error) {
	usage := []*model.RAM{}
	result := db.Find(&usage)
	return usage, result.Error
}
