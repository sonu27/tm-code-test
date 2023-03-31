package app

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"tm/internal/auction"
)

var (
	ErrNotEnoughFields = errors.New("not enough fields")
)

type ActionType int

const (
	ActionUnknown ActionType = iota
	ActionHeartbeat
	ActionSell
	ActionBid
)

type Action struct {
	Timestamp int
	Type      ActionType
	Data      []string
}

func ParseLine(line string) (Action, error) {
	fields := strings.Split(line, "|")

	timestamp, err := strconv.Atoi(fields[0])
	if err != nil {
		return Action{}, err
	}

	if len(fields) == 1 {
		return Action{
			Timestamp: timestamp,
			Type:      ActionHeartbeat,
		}, nil
	}

	if len(fields) < 5 {
		return Action{}, ErrNotEnoughFields
	}

	switch fields[2] {
	case "SELL":
		return Action{
			Timestamp: timestamp,
			Type:      ActionSell,
			Data:      fields,
		}, nil
	case "BID":
		return Action{
			Timestamp: timestamp,
			Type:      ActionBid,
			Data:      fields,
		}, nil
	}

	return Action{}, nil
}

func newAuction(action Action) (*auction.Auction, error) {
	if len(action.Data) != 6 {
		return nil, fmt.Errorf("%w", ErrNotEnoughFields)
	}

	item := action.Data[3]

	startTime := action.Timestamp

	i, err := strconv.ParseInt(action.Data[1], 10, 64)
	if err != nil {
		return nil, err
	}
	userID := int(i)

	i, err = strconv.ParseInt(action.Data[5], 10, 64)
	if err != nil {
		return nil, err
	}
	endTime := int(i)

	f, err := strconv.ParseFloat(action.Data[4], 64)
	if err != nil {
		return nil, err
	}
	price := f

	au := auction.NewAuction(item, startTime, endTime, userID, price)

	return au, nil
}

func newBid(action Action) (*auction.Bid, error) {
	if len(action.Data) != 5 {
		return nil, fmt.Errorf("%w", ErrNotEnoughFields)
	}

	item := action.Data[3]

	time := action.Timestamp

	i, err := strconv.ParseInt(action.Data[1], 10, 64)
	if err != nil {
		return nil, err
	}
	userID := int(i)

	f, err := strconv.ParseFloat(action.Data[4], 64)
	if err != nil {
		return nil, err
	}
	amount := f

	bid := auction.NewBid(item, time, userID, amount)

	return bid, nil
}
