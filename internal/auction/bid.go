package auction

type Bid struct {
	item   string
	time   int
	userID int
	amount float64
}

func (b *Bid) Item() string {
	return b.item
}

func NewBid(item string, time, userId int, amount float64) *Bid {
	return &Bid{
		item:   item,
		time:   time,
		userID: userId,
		amount: amount,
	}
}
