package internal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"tm/internal/auction"
	"tm/internal/service"
)

var (
	ErrNotEnoughFields = errors.New("not enough fields")
)

func Bootstrap(r io.Reader) ([]string, error) {
	svc := service.NewAuctionSvc()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "|")

		if len(fields) < 5 {
			if i, err := strconv.Atoi(fields[0]); err == nil {
				svc.MoveCompleted(i)
			}
			continue
		}

		switch fields[2] {
		case "SELL":
			if au, err := newAuction(fields); err == nil {
				svc.AddAuction(au)
			}

		case "BID":
			if bid, err := newBid(fields); err == nil {
				_ = svc.Bid(bid)
				//todo log err
			}
		}
	}

	return resultsToStrings(svc.GetResults()), nil
}

func resultsToStrings(results []auction.Result) []string {
	var output []string
	for _, v := range results {
		userID := ""
		if v.UserID > 0 {
			userID = strconv.Itoa(v.UserID)
		}
		s := fmt.Sprintf("%d|%s|%s|%s|%0.2f|%d|%0.2f|%0.2f", v.EndTime, v.Item, userID, v.Status, v.PricePaid, v.ValidBidCount, v.HighestBid, v.LowestBid)
		output = append(output, s)
	}
	return output
}

func newAuction(fields []string) (*auction.Auction, error) {
	if len(fields) != 6 {
		return nil, fmt.Errorf("%w", ErrNotEnoughFields)
	}

	item := fields[3]

	i, err := strconv.ParseInt(fields[0], 10, 64)
	if err != nil {
		return nil, err
	}
	startTime := int(i)

	i, err = strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		return nil, err
	}
	userID := int(i)

	i, err = strconv.ParseInt(fields[5], 10, 64)
	if err != nil {
		return nil, err
	}
	endTime := int(i)

	f, err := strconv.ParseFloat(fields[4], 64)
	if err != nil {
		return nil, err
	}
	price := f

	au := auction.NewAuction(item, startTime, endTime, userID, price)

	return au, nil
}

func newBid(fields []string) (*auction.Bid, error) {
	if len(fields) != 5 {
		return nil, fmt.Errorf("%w", ErrNotEnoughFields)
	}

	item := fields[3]

	i, err := strconv.ParseInt(fields[0], 10, 64)
	if err != nil {
		return nil, err
	}
	time := int(i)

	i, err = strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		return nil, err
	}
	userID := int(i)

	f, err := strconv.ParseFloat(fields[4], 64)
	if err != nil {
		return nil, err
	}
	amount := f

	bid := auction.NewBid(item, time, userID, amount)

	return bid, nil
}
