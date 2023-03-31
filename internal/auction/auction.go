package auction

import (
	"errors"
	"fmt"
)

var (
	ErrBidTimeInvalid                      = errors.New("bid time invalid for auction")
	ErrBidMustBeHigherThanUsersPrevHighest = errors.New("bid must be higher than users previous highest bid")
)

func NewAuction(item string, startTime, endTime, userId int, price float64) *Auction {
	return &Auction{
		item:      item,
		startTime: startTime,
		endTime:   endTime,
		userID:    userId,
		price:     price,
		users:     map[int]Bid{},
	}
}

type Auction struct {
	item             string
	startTime        int
	endTime          int
	userID           int
	price            float64
	validBidCount    int
	highestBid       *Bid
	secondHighestBid *Bid
	lowestBid        *Bid
	users            map[int]Bid
}

func (a *Auction) EndTime() int {
	return a.endTime
}

func (a *Auction) Item() string {
	return a.item
}

func (a *Auction) Bid(bid *Bid) error {
	if a.startTime >= bid.time || bid.time > a.endTime {
		return fmt.Errorf("%w, auction id: %s, bid time: %d", ErrBidTimeInvalid, bid.item, bid.time)
	}

	if v, ok := a.users[bid.userID]; ok && bid.amount <= v.amount {
		return fmt.Errorf("%w, prev amount: %f, bid amount: %f", ErrBidMustBeHigherThanUsersPrevHighest, v.amount, bid.amount)
	}

	// set the highest bid and move the current highest to second highest if applicable
	if a.highestBid == nil {
		a.highestBid = bid
	} else if bid.amount > a.highestBid.amount {
		a.secondHighestBid = a.highestBid
		a.highestBid = bid
	} else if a.secondHighestBid == nil || bid.amount > a.secondHighestBid.amount {
		a.secondHighestBid = bid
	}

	if a.lowestBid == nil || bid.amount < a.lowestBid.amount {
		a.lowestBid = bid
	}

	a.validBidCount++
	a.users[bid.userID] = *bid

	return nil
}

func (a *Auction) Result() Result {
	result := Result{
		EndTime:       a.endTime,
		Item:          a.item,
		ValidBidCount: a.validBidCount,
		Status:        "UNSOLD",
	}

	if a.highestBid != nil {
		if a.highestBid.amount >= a.price {
			result.UserID = a.highestBid.userID
			result.Status = "SOLD"
			result.PricePaid = a.price

			if a.secondHighestBid != nil {
				result.PricePaid = a.secondHighestBid.amount
			}
		}
		result.HighestBid = a.highestBid.amount
	}

	if a.lowestBid != nil {
		result.LowestBid = a.lowestBid.amount
	}

	return result
}
