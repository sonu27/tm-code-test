package auction

type Result struct {
	EndTime       int
	Item          string
	UserID        int
	Status        string
	PricePaid     float64
	ValidBidCount int
	HighestBid    float64
	LowestBid     float64
}
