package music

import (
	"fmt"
	"math"
)

type TemperedInterval float64

func FromCents(cents float64) TemperedInterval {
	return TemperedInterval(math.Exp2(cents / 1200))
}

func (i TemperedInterval) Value() float64 {
	return float64(i)
}

func (i TemperedInterval) ToCents() float64 {
	return math.Round(math.Log10(i.ToFloat())/math.Log10(2)*120000) / 100
}

func (i TemperedInterval) ToFloat() float64 {
	return float64(i)
}

func (i TemperedInterval) String() string {
	return fmt.Sprintf("%.2f", i.ToFloat())
}

func (i TemperedInterval) IsUnison() bool {
	return i.Value() == 1.0
}
