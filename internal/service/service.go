package service

import (
	"errors"
	"fmt"
	"tm/internal/auction"
)

var (
	ErrAuctionDoesNotExist  = errors.New("auction does not exist")
	ErrAuctionAlreadyExists = errors.New("auction already exists")
)

func NewAuctionSvc() *AuctionSvc {
	return &AuctionSvc{
		auctions:  map[string]*auction.Auction{},
		completed: []*auction.Auction{},
	}
}

type AuctionSvc struct {
	auctions  map[string]*auction.Auction
	completed []*auction.Auction
}

func (a *AuctionSvc) AddAuction(auction *auction.Auction) error {
	if _, exists := a.auctions[auction.Item()]; exists {
		return ErrAuctionAlreadyExists
	}
	a.auctions[auction.Item()] = auction
	return nil
}

func (a *AuctionSvc) Bid(bid *auction.Bid) error {
	au, ok := a.auctions[bid.Item()]
	if !ok {
		return fmt.Errorf("%w, auction id: %s", ErrAuctionDoesNotExist, bid.Item())
	}
	return au.Bid(bid)
}

func (a *AuctionSvc) MoveCompleted(time int) {
	for i, au := range a.auctions {
		if au.EndTime() <= time {
			a.completed = append(a.completed, au)
			delete(a.auctions, i)
		}
	}
}

func (a *AuctionSvc) GetResults() []auction.Result {
	var results []auction.Result
	for _, au := range a.completed {
		results = append(results, au.Result())
	}
	return results
}
