package auction

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestAuction_Item(t *testing.T) {
	a := NewAuction("item1", 1, 10, 1, 5)
	Equal(t, "item1", a.Item())
}

func TestAuction_EndTime(t *testing.T) {
	a := NewAuction("item1", 1, 10, 1, 5)
	Equal(t, 10, a.EndTime())
}

func TestBid_happyPath(t *testing.T) {
	a := NewAuction("item1", 1, 10, 1, 5)
	b := NewBid("item1", 2, 2, 5)
	if err := a.Bid(b); err != nil {
		t.Error(err)
	}
	Equal(t, 1, a.validBidCount)
	Equal(t, b, a.highestBid)
}

func TestValidBidCount_amountLessThanPriceIsValid(t *testing.T) {
	a := NewAuction("item1", 1, 10, 1, 5)
	b1 := NewBid("item1", 2, 2, 4)
	_ = a.Bid(b1)
	Equal(t, 1, a.validBidCount)
}

func TestValidBidCount_amountMustBeMoreThanUsersHighestBid(t *testing.T) {
	a := NewAuction("item1", 1, 10, 1, 5)
	b1 := NewBid("item1", 2, 2, 5)
	_ = a.Bid(b1)
	b2 := NewBid("item1", 2, 2, 5)
	err := a.Bid(b2)
	if NotNil(t, err, fmt.Sprintf("%s error should have fired", ErrBidMustBeHigherThanUsersPrevHighest)) {
		if !errors.Is(err, ErrBidMustBeHigherThanUsersPrevHighest) {
			t.Fail()
		}
	}
	Equal(t, 1, a.validBidCount)
}

func TestValidBidCount_bidAtOrBeforeStartIsNotValid(t *testing.T) {
	a := NewAuction("item1", 2, 10, 1, 5)
	b1 := NewBid("item1", 0, 2, 5)
	_ = a.Bid(b1)
	b2 := NewBid("item1", 1, 2, 6)
	_ = a.Bid(b2)
	Equal(t, 0, a.validBidCount)
}

func TestValidBidCount_bidAfterAuctionIsNotValid(t *testing.T) {
	a := NewAuction("item1", 2, 10, 1, 5)
	b1 := NewBid("item1", 10, 2, 5)
	_ = a.Bid(b1)
	b2 := NewBid("item1", 11, 3, 6)
	_ = a.Bid(b2)
	Equal(t, 1, a.validBidCount)
}

func TestResult_pricePaidIsReservePriceIfOnlyOneBid(t *testing.T) {
	a := NewAuction("item1", 1, 10, 1, 50)
	b1 := NewBid("item1", 2, 2, 60)
	_ = a.Bid(b1)
	Equal(t, float64(50), a.Result().PricePaid)
}

func TestResult_pricePaidIsZeroIfReservePriceIsNotReached(t *testing.T) {
	a := NewAuction("item1", 1, 10, 1, 50)
	b1 := NewBid("item1", 2, 2, 15)
	_ = a.Bid(b1)
	Equal(t, float64(0), a.Result().PricePaid)
}

func TestResult_pricePaidIsOfSecondHighestBid(t *testing.T) {
	a := NewAuction("item1", 1, 10, 1, 5)
	b1 := NewBid("item1", 2, 2, 15)
	_ = a.Bid(b1)
	b2 := NewBid("item1", 4, 3, 20)
	_ = a.Bid(b2)
	Equal(t, float64(15), a.Result().PricePaid)
}

func TestSecondHighestBid(t *testing.T) {
	a := NewAuction("item1", 1, 10, 1, 5)
	b1 := NewBid("item1", 2, 2, 15)
	_ = a.Bid(b1)
	b2 := NewBid("item1", 4, 3, 20)
	_ = a.Bid(b2)
	Equal(t, b1, a.secondHighestBid)
	b3 := NewBid("item1", 5, 2, 18)
	_ = a.Bid(b3)
	Equal(t, b3, a.secondHighestBid)
}
