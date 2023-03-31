package service

import (
	"container/list"
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
		ongoing:   list.New(),
		completed: []auction.Auction{},
	}
}

type AuctionSvc struct {
	auctions  map[string]*auction.Auction
	ongoing   *list.List
	completed []auction.Auction
}

func (a *AuctionSvc) AddAuction(au *auction.Auction) error {
	if _, exists := a.auctions[au.Item()]; exists {
		return ErrAuctionAlreadyExists
	}
	a.auctions[au.Item()] = au
	z := a.ongoing.Front()
	for z != nil {
		if z.Value.(*auction.Auction).EndTime() > au.EndTime() {
			a.ongoing.InsertBefore(au, z)
			return nil
		}
		z = z.Next()
	}
	a.ongoing.PushBack(au)
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
	front := a.ongoing.Front()
	for front != nil && front.Value.(*auction.Auction).EndTime() <= time {
		auc := front.Value.(*auction.Auction)
		a.completed = append(a.completed, *auc)
		delete(a.auctions, auc.Item())
		a.ongoing.Remove(a.ongoing.Front())

		front = a.ongoing.Front()
	}
}

func (a *AuctionSvc) GetResults() []auction.Result {
	results := make([]auction.Result, len(a.completed))
	for i, au := range a.completed {
		results[i] = au.Result()
	}
	return results
}
