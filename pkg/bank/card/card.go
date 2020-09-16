package card

import (
	"bank/pkg/bank/types"
)

//PaymentSources func
func PaymentSources(cards []types.Card) []types.PaymentSource {

	var payment []types.PaymentSource
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance < 0 {
			continue
		}

		payment = append(payment, types.PaymentSource{
			Type:    "card",
			Number:  string(card.PAN),
			Balance: card.Balance,
		})
	}
	return payment
}
