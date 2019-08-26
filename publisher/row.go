package publisher

type Relationship string

const (
	Direct   Relationship = "DIRECT"
	Reseller Relationship = "RESELLER"
	Other    Relationship = "OTHER"
)

type Row struct {
	Domain       string       `json:"domain"`
	AccountId    string       `json:"account_id"`
	Relationship Relationship `json:"relationship"`
	Authority    string       `json:"authority"`
}
