package database

// Transaction structure of the transaction table model.
type Transaction struct {
	ID          int
	Amount      float32
	ReferenceID string
	ExternalID  string
	Status      string
	Created     string
}
