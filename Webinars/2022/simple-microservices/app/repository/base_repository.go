package repository

import (
	"gorm.io/gorm"
)

type baseRepository struct {
	db *gorm.DB
}

type BaseRepository interface {
	GetDB() *gorm.DB
	BeginTx()
	CommitTx()
	RollbackTx()
	FindByid(id int, model interface{}) (interface{}, error)
	Create(input interface{}) (interface{}, error)
	UpdateByid(id int, input map[string]interface{}, entity interface{}) error
	DeleteByid(id int, entity interface{}) error
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &baseRepository{db}
}

func (br *baseRepository) GetDB() *gorm.DB {
	return br.db
}

func (br *baseRepository) BeginTx() {
	br.db = br.GetDB().Begin()
}

func (br *baseRepository) CommitTx() {
	br.GetDB().Commit()
}

func (br *baseRepository) RollbackTx() {
	br.GetDB().Rollback()
}

func (br *baseRepository) FindByid(id int, entity interface{}) (interface{}, error) {
	err := br.GetDB().
		Where("id=?", id).
		First(entity).Error
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (br *baseRepository) Create(input interface{}) (interface{}, error) {
	err := br.GetDB().Create(input).Error
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (br *baseRepository) UpdateByid(id int, input map[string]interface{}, entity interface{}) error {
	err := br.GetDB().Model(entity).
		Where("id=?", id).
		Updates(input).Error
	if err != nil {
		return err
	}

	return nil
}

func (br *baseRepository) DeleteByid(id int, entity interface{}) error {
	err := br.GetDB().
		Where("id = ?", id).
		Delete(entity).
		Error
	if err != nil {
		return err
	}

	return nil
}
