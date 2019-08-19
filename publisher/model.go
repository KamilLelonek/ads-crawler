package publisher

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Model struct {
	ID           string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Website      string    `gorm:"not null" json:"website"`
	Domain       string    `gorm:"not null" json:"domain"`
	AccountId    string    `gorm:"not null" json:"account_id"`
	Relationship string    `gorm:"not null;type:relationship" json:"relationship"`
	Authority    string    `json:"authority"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Model{})
}

func (Model) TableName() string {
	return "publishers"
}
