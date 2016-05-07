package accessors

type Transaction struct {
	ID        int
	Timestamp string
	User      string
	Amount    int64
	Recipient string
	Note      string
}
