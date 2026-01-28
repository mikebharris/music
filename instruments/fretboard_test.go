package instruments

import (
	"reflect"
	"testing"

	"github.com/mikebharris/music"
)

func Test_newFretboardFromJustScale(t *testing.T) {
	type args struct {
		length  float64
		octaves int
		scale   music.JustScale
	}
	tests := []struct {
		name string
		args args
		want Fretboard
	}{
		{
			name: "Test Just Scale Fretboard",
			args: args{
				length:  650.0,
				octaves: 2,
				scale:   music.NewJustIntonationChromaticScaleWithLimit(2),
			},
			want: Fretboard{
				System:      "2-limit Just Intonation",
				Description: "Fret positions based on Just Intonation chromatic scale based on 2-limit pure ratios.",
				ScaleLength: 650.0,
				Frets: []Fret{
					{Label: "1:1", Position: 0, Comment: "Perfect Unison", Interval: "1:1"},
					{Label: "2:1", Position: 325.0, Comment: "Perfect Octave", Interval: "2:1"},
					{Label: "2:1", Position: 487.5, Comment: "Perfect Octave", Interval: "2:1"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newFretboardFromJustScale(tt.args.length, tt.args.octaves, tt.args.scale); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFretboardFromJustScale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newFretboardFromTemperedScale(t *testing.T) {
	type args struct {
		length  float64
		octaves int
		scale   music.TemperedScale
	}
	tests := []struct {
		name string
		args args
		want Fretboard
	}{
		{
			name: "Test Tempered Scale Fretboard",
			args: args{
				length:  650.0,
				octaves: 1,
				scale:   music.NewEqualTemperamentScale(12),
			},
			want: Fretboard{
				System:      "Equal Temperament",
				Description: "Fret positions based on 12-tone equal temperament.",
				ScaleLength: 650.0,
				Frets: []Fret{
					{Label: "0.00 cents", Position: 0},
					{Label: "100.00 cents", Position: 36.48},
					{Label: "200.00 cents", Position: 70.92},
					{Label: "300.00 cents", Position: 103.42},
					{Label: "400.00 cents", Position: 134.09},
					{Label: "500.00 cents", Position: 163.05},
					{Label: "600.00 cents", Position: 190.38},
					{Label: "700.00 cents", Position: 216.18},
					{Label: "800.00 cents", Position: 240.53},
					{Label: "900.00 cents", Position: 263.51},
					{Label: "1000.00 cents", Position: 285.2},
					{Label: "1100.00 cents", Position: 305.67},
					{Label: "1200.00 cents", Position: 325},
				},
			},
		},
		{
			name: "Test Tempered Scale Fretboard with 2 Octaves",
			args: args{
				length:  650.0,
				octaves: 2,
				scale:   music.NewEqualTemperamentScale(2),
			},
			want: Fretboard{
				System:      "Equal Temperament",
				Description: "Fret positions based on 2-tone equal temperament.",
				ScaleLength: 650.0,
				Frets: []Fret{
					{Label: "0.00 cents", Position: 0},
					{Label: "600.00 cents", Position: 190.38},
					{Label: "1200.00 cents", Position: 325},
					{Label: "1800.00 cents", Position: 420.19},
					{Label: "2400.00 cents", Position: 487.5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newFretboardFromTemperedScale(tt.args.length, tt.args.octaves, tt.args.scale); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFretboardFromTemperedScale() = %v, want %v", got, tt.want)
			}
		})
	}
}
