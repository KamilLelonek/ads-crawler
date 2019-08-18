package publisher

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestFetchList(t *testing.T) {
	defer gock.Off()

	gock.New("https://random.com").
		Get("/ads.txt").
		Reply(200).
		BodyString("example.com,1,DIRECT")

	result, err := FetchList("random")
	row := []Row{Row{
		Domain:       "example.com",
		AccountId:    "1",
		Relationship: Direct,
	}}

	assert.Nil(t, err)
	assert.Equal(t, row, result)
}
