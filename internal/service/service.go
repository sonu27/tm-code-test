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
	if _, ok := a.auctions[au.Item()]; ok {
		return ErrAuctionAlreadyExists
	}
	a.auctions[au.Item()] = au

	for e := a.ongoing.Back(); e != nil; e = e.Prev() {
		if e.Value.(*auction.Auction).EndTime() <= au.EndTime() {
			a.ongoing.InsertAfter(au, e)
			return nil
		}
	}

	a.ongoing.PushFront(au)
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
	for front := a.ongoing.Front(); front != nil; front = front.Next() {
		if front.Value.(*auction.Auction).EndTime() > time {
			return
		}

		auc := front.Value.(*auction.Auction)
		a.completed = append(a.completed, *auc)
		delete(a.auctions, auc.Item())
		a.ongoing.Remove(a.ongoing.Front())
	}
}

func (a *AuctionSvc) GetResults() []auction.Result {
	results := make([]auction.Result, len(a.completed))
	for i, au := range a.completed {
		results[i] = au.Result()
	}
	return results
}
