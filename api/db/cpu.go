package db

import "new/test/project/api/model"

type CPU interface {
	Insert(*model.CPU) (*model.CPU, error)
	GetAll() ([]*model.CPU, error)
}

type CPUManger struct {
}

func NewCPUManger() CPU {
	return &CPUManger{}
}

func (cm *CPUManger) Insert(data *model.CPU) (*model.CPU, error) {
	result := db.Create(data)
	return data, result.Error
}

func (cm *CPUManger) GetAll() ([]*model.CPU, error) {
	usage := []*model.CPU{}
	result := db.Find(&usage)
	return usage, result.Error
}
