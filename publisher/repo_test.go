package publisher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadByWebsite(t *testing.T) {
	repo := &Repo{}
	db := repo.DB()
	website := "wordpress"
	model := Model{
		Website:      website,
		Domain:       "google.com",
		AccountId:    "123",
		Relationship: Direct,
		Authority:    "abc",
	}

	db.Delete(&model)

	assert.True(t, db.NewRecord(model))
	db.Create(&model)
	assert.False(t, db.NewRecord(model))

	models, err := repo.ReadByWebsite(website)
	assert.Nil(t, err)
	assert.Equal(t, model.ID, (*models)[0].ID)
}

func TestWriteWithDomain(t *testing.T) {
	repo := &Repo{}
	db := repo.DB()
	website := "wordpress"
	row := Row{
		Domain:       "google.com",
		AccountId:    "123",
		Relationship: Reseller,
		Authority:    "abc",
	}

	err := repo.WriteWithWebsite(website, row)

	assert.Nil(t, err)

	models, err := repo.ReadByWebsite(website)

	assert.Nil(t, err)

	model := (*models)[0]

	assert.Equal(t, model.Website, website)

	db.Delete(&model)
}
