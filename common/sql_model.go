package common

import "time"

type SQLModel struct {
	Id int `json:"id" gorm:"column:id"`
	//FakedId   int `json:"faked_id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}
