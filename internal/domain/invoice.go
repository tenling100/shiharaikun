package domain

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	Status    string    `gorm:"type:enum('unpaid', 'paid', 'overdue')"; not null`
	Amount    uint64    `gorm:"type:decimal(10,2)"; not null`
	RepayDate time.Time `type:datetime`
	CreateAt  time.Time `gorm:"autoCreateTime"`
	UpdateAt  time.Time `gorm:"autoUpdateTime"`
}
