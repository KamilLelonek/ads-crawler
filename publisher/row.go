package publisher

const (
	Direct Relationship = iota
	Reseller
	Other
)

type Relationship int

type Row struct {
	Domain       string
	AccountId    string
	Relationship Relationship
	Authority    string
}
