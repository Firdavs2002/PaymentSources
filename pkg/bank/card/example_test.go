package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 10_000_00,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 0000",
			Balance: -10_000_00,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 1111",
			Balance: 15_000_00,
			Active:  false,
		},
	}
	paymentSources := PaymentSources(cards)

	for _, v := range paymentSources {
		fmt.Println(v.Number)
	}

	//Output:  5058 xxxx xxxx 8888

}
