package publisher

import (
	"reflect"
	"testing"
)

func TestNewRow(t *testing.T) {
	cases := []struct {
		row      string
		expected Row
	}{
		{
			row: "example.com,1,DIRECT",
			expected: Row{
				Domain:       "example.com",
				AccountId:    "1",
				Relationship: Direct,
			},
		},
		{
			row: "example.com,1,Reseller,a",
			expected: Row{
				Domain:       "example.com",
				AccountId:    "1",
				Relationship: Reseller,
				Authority:    "a",
			},
		},
		{
			row: " example.com,  1,Reseller, a",
			expected: Row{
				Domain:       "example.com",
				AccountId:    "1",
				Relationship: Reseller,
				Authority:    "a",
			},
		},
		{
			row: "example.com,1,RESELLER # video",
			expected: Row{
				Domain:       "example.com",
				AccountId:    "1",
				Relationship: Reseller,
			},
		},
		{
			row: "example.com,1,Reseller,a # audio",
			expected: Row{
				Domain:       "example.com",
				AccountId:    "1",
				Relationship: Reseller,
				Authority:    "a",
			},
		},
	}

	for i, c := range cases {
		row, err := NewRow(c.row)

		if err != nil {
			t.Errorf("(#%d) parse ads.txt failed: %s", i, err)
		}

		if !reflect.DeepEqual(c.expected, row) {
			t.Errorf("want %v, got %v", c.expected, row)
		}
	}
}
