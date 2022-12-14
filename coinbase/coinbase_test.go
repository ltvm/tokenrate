package coinbase

import (
	"testing"
	"time"
)

const cbName = "coinbase"

func TestCoinBase(t *testing.T) {
	cb := New()
	rate, err := cb.Rate("BTC", "USD", time.Now())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("current BTC/USD price: %f", rate)

	rate, err = cb.USDRate(time.Now())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("current ETH/USD rate: %f", rate)

	rate, err = cb.USDRate(time.Now().AddDate(0, 0, -7))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("last week ETH/USD rate: %f", rate)

	if name := cb.Name(); name != cbName {
		t.Fatal(err)
	}
}
