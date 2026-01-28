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
		{
			name: "Test String() method",
			i:    TemperedInterval(1.5),
			want: "1.50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.i.String(), "String()")
		})
	}
}
