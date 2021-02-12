package test

import (
	"log"
	"os"
	"testing"
	"tm/internal"

	. "github.com/stretchr/testify/assert"
)

func TestBootstrap(t *testing.T) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output, err := internal.Bootstrap(file)
	Nil(t, err)

	s1 := "20|toaster_1|8|SOLD|12.50|3|20.00|7.50"
	s2 := "20|tv_1||UNSOLD|0.00|2|200.00|150.00"

	if len(output) != 2 {
		t.Fail()
	}

	if output[0] != s1 && output[0] != s2 {
		t.Fail()
	}

	if output[1] != s1 && output[1] != s2 {
		t.Fail()
	}
}
