package publisher

const (
	Direct Relationship = iota
	Reseller
	Other
)

type Relationship int

type Row struct {
	Domain       string       `json:"domain"`
	AccountId    string       `json:"account_id"`
	Relationship Relationship `json:"relationship"`
	Authority    string       `json:"authority"`
}
