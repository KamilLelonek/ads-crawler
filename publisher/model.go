package publisher

import (
	"database/sql/driver"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

func (r *Relationship) Scan(value interface{}) error {
	*r = Relationship(value.(int))
	return nil
}

func (r Relationship) Value() (driver.Value, error) {
	return string(r), nil
}

type Model struct {
	ID           string       `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Website      string       `gorm:"not null" json:"website"`
	Domain       string       `gorm:"not null" json:"domain"`
	AccountId    string       `gorm:"not null" json:"account_id"`
	Relationship Relationship `gorm:"not null;type:relationship" json:"relationship"`
	Authority    string       `json:"authority"`
	CreatedAt    time.Time    `gorm:"not null" json:"created_at"`
}

func (Model) TableName() string {
	return "publishers"
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()

	if err != nil {
		return err
	}

	return scope.SetColumn("ID", uuid.String())
}
