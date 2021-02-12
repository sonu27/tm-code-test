package test

import (
	"testing"
	"tm/internal/auction"

	. "github.com/stretchr/testify/assert"
)

func TestBid_happyPath(t *testing.T) {
	a := auction.NewAuction("item1", 1, 10, 1, 5)
	b := auction.NewBid("item1", 2, 2, 5)
	if err := a.Bid(b); err != nil {
		t.Error(err)
	}

	r := a.Result()
	Equal(t, 10, r.EndTime)
	Equal(t, "item1", r.Item)
	Equal(t, 2, r.UserID)
	Equal(t, "SOLD", r.Status)
	Equal(t, 1, r.ValidBidCount)
	Equal(t, float64(5), r.HighestBid)
	Equal(t, float64(5), r.LowestBid)
}

func TestBid_happyPathComplex(t *testing.T) {
	a := auction.NewAuction("item1", 1, 8, 1, 5)
	b1 := auction.NewBid("item1", 2, 2, 4)
	_ = a.Bid(b1)

	b2 := auction.NewBid("item1", 3, 3, 5)
	_ = a.Bid(b2)

	b3 := auction.NewBid("item1", 4, 4, 6)
	_ = a.Bid(b3)

	b4 := auction.NewBid("item1", 5, 2, 6)
	_ = a.Bid(b4)

	b5 := auction.NewBid("item1", 6, 2, 6)
	_ = a.Bid(b5)

	b6 := auction.NewBid("item1", 7, 2, 5)
	_ = a.Bid(b6)

	b7 := auction.NewBid("item1", 9, 3, 7)
	_ = a.Bid(b7)

	r := a.Result()
	Equal(t, 8, r.EndTime)
	Equal(t, "item1", r.Item)
	Equal(t, 4, r.UserID)
	Equal(t, "SOLD", r.Status)
	Equal(t, 4, r.ValidBidCount)
	Equal(t, float64(6), r.HighestBid)
	Equal(t, float64(4), r.LowestBid)
}
