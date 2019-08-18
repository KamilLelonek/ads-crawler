package publisher

import (
	"reflect"
	"strings"
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
		{
			row: " example.com, 1  ,Reseller,a # a, b, c",
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

func TestNewList(t *testing.T) {
	cases := []struct {
		list     string
		expected []Row
	}{
		{
			list: " # CNN.com/ads.txt\n # \n # DOMESTIC\ngoogle.com, pub-7439281311086140, DIRECT, f08c47fec0942fa0 # banner, video, native",
			expected: []Row{
				{
					Domain:       "google.com",
					AccountId:    "pub-7439281311086140",
					Relationship: Direct,
					Authority:    "f08c47fec0942fa0",
				},
			},
		},
	}

	for i, c := range cases {
		list, err := NewList(strings.NewReader(c.list))

		if err != nil {
			t.Errorf("(#%d) parse ads.txt failed: %s", i, err)
		}

		if !reflect.DeepEqual(c.expected, list) {
			t.Errorf("want %v, got %v", c.expected, list)
		}
	}
}
