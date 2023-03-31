package app

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"tm/internal/auction"
	"tm/internal/service"
)

func Start(r io.Reader) ([]string, error) {
	svc := service.NewAuctionSvc()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		action, err := ParseLine(scanner.Text())
		if err != nil {
			return nil, err
		}

		switch action.Type {
		case ActionSell:
			if au, err := newAuction(action); err == nil {
				_ = svc.AddAuction(au)
			}
		case ActionBid:
			if bid, err := newBid(action); err == nil {
				_ = svc.Bid(bid)
			}
		}

		svc.MoveCompleted(action.Timestamp)
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
