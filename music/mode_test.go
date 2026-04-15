package music

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMusicalMode_IsDiatonic(t *testing.T) {
	tests := []struct {
		name string
		m    MusicalMode
		want bool
	}{
		{name: "Lydian is diatonic", m: LydianMode, want: true},
		{name: "Athenian is not a diatonic mode", m: "Athenian", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.IsDiatonic(); got != tt.want {
				t.Errorf("IsDiatonic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMusicalMode_String(t *testing.T) {
	tests := []struct {
		name string
		m    MusicalMode
		want string
	}{
		{name: "ModeLydian string representation", m: LydianMode, want: "Lydian"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMusicalMode_BrightnessOrder(t *testing.T) {
	assert.Equal(t, 7, LydianMode.BrightnessOrder())
	assert.Equal(t, 6, IonianMode.BrightnessOrder())
	assert.Equal(t, 5, MixolydianMode.BrightnessOrder())
	assert.Equal(t, 4, DorianMode.BrightnessOrder())
	assert.Equal(t, 3, AeolianMode.BrightnessOrder())
	assert.Equal(t, 2, PhrygianMode.BrightnessOrder())
	assert.Equal(t, 1, LocrianMode.BrightnessOrder())
	assert.Equal(t, 0, MusicalMode("Athenian").BrightnessOrder(), "non-diatonic mode returns 0")
}

func TestMusicalMode_IsBrighterThan(t *testing.T) {
	assert.True(t, LydianMode.IsBrighterThan(IonianMode), "Lydian should be brighter than Ionian")
	assert.True(t, IonianMode.IsBrighterThan(LocrianMode), "Ionian should be brighter than Locrian")
	assert.False(t, LocrianMode.IsBrighterThan(LydianMode), "Locrian should not be brighter than Lydian")
	assert.False(t, DorianMode.IsBrighterThan(DorianMode), "a mode is not brighter than itself")
}
