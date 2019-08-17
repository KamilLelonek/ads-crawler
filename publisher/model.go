package publisher

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Model struct {
	ID           string `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Website      string `gorm:"not null"`
	Domain       string `gorm:"not null"`
	AccountId    string `gorm:"not null"`
	Relationship string `gorm:"not null;type:relationship"`
	Authority    string
	CreatedAt    time.Time `gorm:"not null"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Model{})
}

func (Model) TableName() string {
	return "publishers"
}
