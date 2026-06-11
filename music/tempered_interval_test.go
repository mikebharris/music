package music

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemperedInterval_String(t *testing.T) {
	tests := []struct {
		name string
		i    TemperedInterval
		want string
	}{
		{name: "Test String() method", i: TemperedInterval(1.5), want: "1.5"},
		{name: "Full precision for irrational ratio", i: FromCents(100), want: "1.0594630943592953"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.i.String(), "String()")
		})
	}
}

func TestFromCents(t *testing.T) {
	tests := []struct {
		name      string
		cents     float64
		wantCents float64
	}{
		{name: "700 cents round-trips back to 700 cents", cents: 700.0, wantCents: 700.0},
		{name: "0 cents is unison", cents: 0.0, wantCents: 0.0},
		{name: "1200 cents is an octave", cents: 1200.0, wantCents: 1200.0},
		{name: "100 cents is one equal-tempered semitone", cents: 100.0, wantCents: 100.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.InDelta(t, tt.wantCents, FromCents(tt.cents).ToCents(), 0.001)
		})
	}
}
