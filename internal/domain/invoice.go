package domain

import (
	"time"
)

type Invoice struct {
	ID        uint       `gorm:"primaryKey"`
	Status    string     `gorm:"type:enum('unpaid', 'paid', 'overdue')"; not null`
	Amount    uint64     `gorm:"type:decimal(10,2)"; not null`
	RepayDate time.Time  `type:datetime`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm: null`
}
