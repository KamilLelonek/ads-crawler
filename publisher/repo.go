package publisher

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "crawler_development"
)

var psqlInfo = fmt.Sprintf(
	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname,
)

type Repo struct {
	db   *gorm.DB
	once struct {
		db sync.Once
	}
}

func (repo *Repo) DB() *gorm.DB {
	repo.once.db.Do(func() {
		db, err := gorm.Open("postgres", psqlInfo)
		db.LogMode(true)

		if err != nil {
			log.Fatalf("Failed to open database connection: %s", err)
		}

		repo.db = db
	})

	return repo.db
}
func (repo *Repo) CloseDB() {
	repo.DB().Close()
}

func (repo *Repo) Migrate() {
	repo.DB().AutoMigrate(&Model{})
}

func (repo *Repo) ReadByWebsite(website string) (*[]Model, error) {
	models := new([]Model)
	err := repo.DB().Find(&models, Model{Website: website}).Error

	return models, err
}

func (repo *Repo) WriteWithWebsite(website string, row Row) error {
	model := &Model{
		Website:      website,
		Domain:       row.Domain,
		AccountId:    row.AccountId,
		Relationship: row.Relationship,
		Authority:    row.Authority,
	}

	return repo.
		DB().
		Create(&model).
		Error
}
