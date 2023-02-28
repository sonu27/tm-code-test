package service

import (
	"errors"
	"fmt"
	"testing"
	"tm/internal/auction"

	. "github.com/stretchr/testify/assert"
)

func TestAuctionSvc_MoveCompleted(t *testing.T) {
	svc := NewAuctionSvc()
	Equal(t, 0, len(svc.auctions))
	Equal(t, 0, len(svc.completed))

	svc.AddAuction(auction.NewAuction("item1", 1, 5, 1, 5))
	Equal(t, 1, len(svc.auctions))
	Equal(t, 0, len(svc.completed))

	svc.MoveCompleted(4)
	Equal(t, 1, len(svc.auctions))
	Equal(t, 0, len(svc.completed))

	svc.MoveCompleted(5)
	Equal(t, 0, len(svc.auctions))
	Equal(t, 1, len(svc.completed))
}

func TestAuctionSvc_Bid_failsWhenAuctionDoNotExist(t *testing.T) {
	svc := NewAuctionSvc()
	err := svc.Bid(auction.NewBid("auctionThatDoesNotExist", 5, 1, 10))

	if NotNil(t, err, fmt.Sprintf("%s error should have fired", ErrAuctionDoesNotExist)) {
		if !errors.Is(err, ErrAuctionDoesNotExist) {
			t.Fail()
		}
	}
}

func TestAuctionSvc_ErrAuctionAlreadyExists(t *testing.T) {
	svc := NewAuctionSvc()
	a1 := auction.NewAuction("item1", 1, 5, 1, 5)
	a2 := auction.NewAuction("item1", 1, 5, 1, 5)

	svc.AddAuction(a1)

	err := svc.AddAuction(a2)
	if NotNil(t, err, fmt.Sprintf("%s error should have fired", ErrAuctionAlreadyExists)) {
		if !errors.Is(err, ErrAuctionAlreadyExists) {
			t.Fail()
		}
	}
}

func TestSort(t *testing.T) {
	a := NewAuctionSvc()
	a.AddAuction(auction.NewAuction("item3", 1, 4, 1, 1))
	a.AddAuction(auction.NewAuction("item1", 1, 2, 1, 1))
	a.AddAuction(auction.NewAuction("item4", 1, 5, 1, 1))
	a.AddAuction(auction.NewAuction("item5", 1, 6, 1, 1))
	a.AddAuction(auction.NewAuction("item2", 1, 3, 1, 1))

	front := a.ongoing.Front()
	Equal(t, "item1", front.Value.(*auction.Auction).Item())
	front = front.Next()
	Equal(t, "item2", front.Value.(*auction.Auction).Item())
	front = front.Next()
	Equal(t, "item3", front.Value.(*auction.Auction).Item())
	front = front.Next()
	Equal(t, "item4", front.Value.(*auction.Auction).Item())
	front = front.Next()
	Equal(t, "item5", front.Value.(*auction.Auction).Item())
}
