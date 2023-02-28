package app_test

import (
	. "github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"tm/internal/app"
)

func TestApp(t *testing.T) {
	str := `
10|1|SELL|toaster_1|10.00|20
12|8|BID|toaster_1|7.50
13|5|BID|toaster_1|12.50
15|8|SELL|tv_1|250.00|20
16
17|8|BID|toaster_1|20.00
18|1|BID|tv_1|150.00
19|3|BID|tv_1|200.00
20
21|3|BID|tv_1|300.00
`
	r := strings.NewReader(strings.TrimSpace(str))
	output, err := app.Start(r)
	Nil(t, err)

	s1 := "20|toaster_1|8|SOLD|12.50|3|20.00|7.50"
	s2 := "20|tv_1||UNSOLD|0.00|2|200.00|150.00"

	Equal(t, 2, len(output))
	Equal(t, s1, output[0])
	Equal(t, s2, output[1])
}
