package test

import (
	"testing"
	"tm/internal/auction"

	. "github.com/stretchr/testify/assert"
)

func TestAuctionBid_happyPath(t *testing.T) {
	a := auction.NewAuction("item1", 1, 10, 1, 5)
	b := auction.NewBid("item1", 2, 2, 6)
	if err := a.Bid(b); err != nil {
		t.Error(err)
	}

	want := auction.Result{
		EndTime:       10,
		Item:          "item1",
		UserID:        2,
		Status:        "SOLD",
		PricePaid:     5,
		ValidBidCount: 1,
		HighestBid:    6,
		LowestBid:     6,
	}
	got := a.Result()
	Equal(t, want, got)
}

func TestAuctionBid_happyPathComplex(t *testing.T) {
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

	want := auction.Result{
		EndTime:       8,
		Item:          "item1",
		UserID:        4,
		Status:        "SOLD",
		PricePaid:     6,
		ValidBidCount: 4,
		HighestBid:    6,
		LowestBid:     4,
	}
	got := a.Result()
	Equal(t, want, got)
}
