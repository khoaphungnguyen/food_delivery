package restaurantstorage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

func NewSQlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
