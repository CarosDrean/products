package currencylayer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_adapterConverter(t *testing.T) {
	tests := []struct {
		name       string
		currencies map[string]float64
		from       string
		to         string
		amount     float64
		want       float64
	}{
		{
			name: "PEN_COP",
			currencies: map[string]float64{
				"USDPEN": 3.99,
				"USDCOP": 4062.67,
			},
			from:   "PEN",
			to:     "COP",
			amount: 10,
			want:   10182.13,
		},
		{
			name: "PEN_CLP",
			currencies: map[string]float64{
				"USDPEN": 3.99,
				"USDCLP": 852,
			},
			from:   "PEN",
			to:     "CLP",
			amount: 10,
			want:   2135.34,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			price := adapterConverter(tt.currencies, tt.from, tt.to, tt.amount)
			assert.Equal(t, tt.want, price)
		})
	}
}
