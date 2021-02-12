package auction

import (
	"reflect"
	"testing"
)

func TestBid_Item(t *testing.T) {
	bid := Bid{
		item:   "test",
		time:   100,
		userID: 99,
		amount: 77,
	}

	want := "test"

	if got := bid.Item(); got != want {
		t.Errorf("Item() = %v, want %v", got, want)
	}
}

func TestNewBid(t *testing.T) {
	type args struct {
		item   string
		time   int
		userId int
		amount float64
	}

	a := args{
		item:   "test",
		time:   100,
		userId: 99,
		amount: 77,
	}

	actual := Bid{
		item:   "test",
		time:   100,
		userID: 99,
		amount: 77,
	}

	if got := NewBid(a.item, a.time, a.userId, a.amount); !reflect.DeepEqual(*got, actual) {
		t.Errorf("NewBid() = %v, want %v", got, actual)
	}
}
