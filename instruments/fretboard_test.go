package instruments

import (
	"reflect"
	"testing"

	"github.com/mikebharris/music/music"
	"github.com/stretchr/testify/assert"
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
				octaves: 3,
				scale:   music.NewJustIntonationChromaticScaleWithLimit(2),
			},
			want: Fretboard{
				System:      "2-limit Just Intonation",
				Description: "Fret positions based on Just Intonation chromatic scale based on 2-limit pure ratios.",
				ScaleLength: 650.0,
				Frets: []Fret{
					{Position: 0, Comment: "Perfect Unison", Interval: "1:1", Cents: 0.0},
					{Position: 325.0, Comment: "Perfect Octave", Interval: "2:1", Cents: 1200.0},
					{Position: 487.5, Comment: "Perfect Octave", Interval: "4:1", Cents: 2400.00},
					{Position: 568.75, Comment: "Perfect Octave", Interval: "8:1", Cents: 3600.00},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFretboardFromJustScale(tt.args.length, tt.args.octaves, tt.args.scale); !reflect.DeepEqual(got, tt.want) {
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
		mode    music.MusicalMode
	}
	tests := []struct {
		name string
		args args
		want Fretboard
	}{
		{
			name: "Test Tempered Chromatic Scale Fretboard",
			args: args{
				length:  650.0,
				octaves: 1,
				scale:   music.NewEqualTemperamentScale(12),
			},
			want: Fretboard{
				System:      "Equal Temperament",
				Description: "Chromatic fret positions based on 12-tone equal temperament.",
				ScaleLength: 650.0,
				Frets: []Fret{
					{Position: 0, Interval: "1"},
					{Position: 36.48, Interval: "1.0594630943592953", Cents: 100},
					{Position: 70.92, Interval: "1.122462048309373", Cents: 200},
					{Position: 103.42, Interval: "1.189207115002721", Cents: 300},
					{Position: 134.09, Interval: "1.2599210498948732", Cents: 400},
					{Position: 163.05, Interval: "1.3348398541700344", Cents: 500},
					{Position: 190.38, Interval: "1.414213562373095", Cents: 600},
					{Position: 216.18, Interval: "1.4983070768766815", Cents: 700},
					{Position: 240.53, Interval: "1.5874010519681994", Cents: 800},
					{Position: 263.51, Interval: "1.6817928305074292", Cents: 900},
					{Position: 285.2, Interval: "1.7817974362806788", Cents: 1000},
					{Position: 305.67, Interval: "1.887748625363387", Cents: 1100},
					{Position: 325, Interval: "2", Cents: 1200},
				},
			},
		},
		{
			name: "Test Tempered Chromatic Scale Fretboard with 2 Octaves",
			args: args{
				length:  650.0,
				octaves: 2,
				scale:   music.NewEqualTemperamentScale(2),
			},
			want: Fretboard{
				System:      "Equal Temperament",
				Description: "Chromatic fret positions based on 2-tone equal temperament.",
				ScaleLength: 650.0,
				Frets: []Fret{
					{Position: 0, Interval: "1"},
					{Position: 190.38, Interval: "1.414213562373095", Cents: 600},
					{Position: 325, Interval: "2", Cents: 1200},
					{Position: 420.19, Interval: "1.414213562373095", Cents: 1800},
					{Position: 487.5, Interval: "2", Cents: 2400},
				},
			},
		},
		{
			name: "Test Tempered Diatonic Lydian Scale Fretboard",
			args: args{
				length:  650.0,
				octaves: 1,
				scale:   music.NewEqualTemperamentScale(12),
				mode:    music.LydianMode,
			},
			want: Fretboard{
				System:      "Equal Temperament",
				Description: "Diatonic (Lydian) fret positions based on 12-tone equal temperament.",
				ScaleLength: 650.0,
				Frets: []Fret{
					{Position: 0, Interval: "1"},
					{Position: 70.92, Interval: "1.122462048309373", Cents: 200},
					{Position: 134.09, Interval: "1.2599210498948732", Cents: 400}, // augmented 4th
					{Position: 190.38, Interval: "1.414213562373095", Cents: 600},
					{Position: 216.18, Interval: "1.4983070768766815", Cents: 700},
					{Position: 263.51, Interval: "1.6817928305074292", Cents: 900},
					{Position: 305.67, Interval: "1.887748625363387", Cents: 1100},
					{Position: 325, Interval: "2", Cents: 1200},
				},
			},
		},
		{
			name: "Test Tempered Diatonic Ionian Scale Fretboard",
			args: args{
				length:  650.0,
				octaves: 1,
				scale:   music.NewEqualTemperamentScale(12),
				mode:    music.IonianMode,
			},
			want: Fretboard{
				System:      "Equal Temperament",
				Description: "Diatonic (Ionian) fret positions based on 12-tone equal temperament.",
				ScaleLength: 650.0,
				Frets: []Fret{
					{Position: 0, Interval: "1"},
					{Position: 70.92, Interval: "1.122462048309373", Cents: 200},
					{Position: 134.09, Interval: "1.2599210498948732", Cents: 400},
					{Position: 163.05, Interval: "1.3348398541700344", Cents: 500}, // perfect 4th
					{Position: 216.18, Interval: "1.4983070768766815", Cents: 700},
					{Position: 263.51, Interval: "1.6817928305074292", Cents: 900},
					{Position: 305.67, Interval: "1.887748625363387", Cents: 1100},
					{Position: 325, Interval: "2", Cents: 1200},
				},
			},
		},
		{
			name: "Test Tempered Diatonic Mixolydian Scale Fretboard",
			args: args{
				length:  650.0,
				octaves: 1,
				scale:   music.NewEqualTemperamentScale(12),
				mode:    music.MixolydianMode,
			},
			want: Fretboard{
				System:      "Equal Temperament",
				Description: "Diatonic (Mixolydian) fret positions based on 12-tone equal temperament.",
				ScaleLength: 650.0,
				Frets: []Fret{
					{Position: 0, Interval: "1"},
					{Position: 70.92, Interval: "1.122462048309373", Cents: 200},
					{Position: 134.09, Interval: "1.2599210498948732", Cents: 400},
					{Position: 163.05, Interval: "1.3348398541700344", Cents: 500}, // perfect 4th
					{Position: 216.18, Interval: "1.4983070768766815", Cents: 700},
					{Position: 263.51, Interval: "1.6817928305074292", Cents: 900},
					{Position: 285.2, Interval: "1.7817974362806788", Cents: 1000}, // flattened 7th
					{Position: 325, Interval: "2", Cents: 1200},
				},
			},
		},
		{
			name: "Test Tempered Diatonic Dorian Scale Fretboard",
			args: args{
				length:  650.0,
				octaves: 1,
				scale:   music.NewEqualTemperamentScale(12),
				mode:    music.DorianMode,
			},
			want: Fretboard{
				System:      "Equal Temperament",
				Description: "Diatonic (Dorian) fret positions based on 12-tone equal temperament.",
				ScaleLength: 650.0,
				Frets: []Fret{
					{Position: 0, Interval: "1"},
					{Position: 70.92, Interval: "1.122462048309373", Cents: 200},
					{Position: 103.42, Interval: "1.189207115002721", Cents: 300},  // flattened third
					{Position: 163.05, Interval: "1.3348398541700344", Cents: 500}, // perfect 4th
					{Position: 216.18, Interval: "1.4983070768766815", Cents: 700},
					{Position: 263.51, Interval: "1.6817928305074292", Cents: 900},
					{Position: 285.2, Interval: "1.7817974362806788", Cents: 1000}, // flattened 7th
					{Position: 325, Interval: "2", Cents: 1200},
				},
			},
		},
		{
			name: "Test Tempered Diatonic Aeolian Scale Fretboard",
			args: args{
				length:  650.0,
				octaves: 1,
				scale:   music.NewEqualTemperamentScale(12),
				mode:    music.AeolianMode,
			},
			want: Fretboard{
				System:      "Equal Temperament",
				Description: "Diatonic (Aeolian) fret positions based on 12-tone equal temperament.",
				ScaleLength: 650.0,
				Frets: []Fret{
					{Position: 0, Interval: "1"},
					{Position: 70.92, Interval: "1.122462048309373", Cents: 200},
					{Position: 103.42, Interval: "1.189207115002721", Cents: 300},  // flattened 3rd
					{Position: 163.05, Interval: "1.3348398541700344", Cents: 500}, // perfect 4th
					{Position: 216.18, Interval: "1.4983070768766815", Cents: 700},
					{Position: 240.53, Interval: "1.5874010519681994", Cents: 800}, // flattened 6th
					{Position: 285.2, Interval: "1.7817974362806788", Cents: 1000}, // flattened 7th
					{Position: 325, Interval: "2", Cents: 1200},
				},
			},
		},
		{
			name: "Test Tempered Diatonic Phrygian Scale Fretboard",
			args: args{
				length:  650.0,
				octaves: 1,
				scale:   music.NewEqualTemperamentScale(12),
				mode:    music.PhrygianMode,
			},
			want: Fretboard{
				System:      "Equal Temperament",
				Description: "Diatonic (Phrygian) fret positions based on 12-tone equal temperament.",
				ScaleLength: 650.0,
				Frets: []Fret{
					{Position: 0, Interval: "1"},
					{Position: 36.48, Interval: "1.0594630943592953", Cents: 100},  // flattened 2nd
					{Position: 103.42, Interval: "1.189207115002721", Cents: 300},  // flattened 3rd
					{Position: 163.05, Interval: "1.3348398541700344", Cents: 500}, // perfect 4th
					{Position: 216.18, Interval: "1.4983070768766815", Cents: 700},
					{Position: 240.53, Interval: "1.5874010519681994", Cents: 800}, // flattened 6th
					{Position: 285.2, Interval: "1.7817974362806788", Cents: 1000}, // flattened 7th
					{Position: 325, Interval: "2", Cents: 1200},
				},
			},
		},
		{
			name: "Test Tempered Diatonic Locrian Scale Fretboard",
			args: args{
				length:  650.0,
				octaves: 1,
				scale:   music.NewEqualTemperamentScale(12),
				mode:    music.LocrianMode,
			},
			want: Fretboard{
				System:      "Equal Temperament",
				Description: "Diatonic (Locrian) fret positions based on 12-tone equal temperament.",
				ScaleLength: 650.0,
				Frets: []Fret{
					{Position: 0, Interval: "1"},
					{Position: 36.48, Interval: "1.0594630943592953", Cents: 100},  // flattened 2nd
					{Position: 103.42, Interval: "1.189207115002721", Cents: 300},  // flattened 3rd
					{Position: 163.05, Interval: "1.3348398541700344", Cents: 500}, // perfect 4th
					{Position: 190.38, Interval: "1.414213562373095", Cents: 600},  // diminished 5th
					{Position: 240.53, Interval: "1.5874010519681994", Cents: 800}, // flattened 6th
					{Position: 285.2, Interval: "1.7817974362806788", Cents: 1000}, // flattened 7th
					{Position: 325, Interval: "2", Cents: 1200},
				},
			},
		},
		{
			name: "Test Quarter-Comma Meantone Diatonic Aeolian Scale Fretboard",
			args: args{
				length:  650.0,
				octaves: 1,
				scale:   music.NewQuarterCommaMeantoneScale(),
				mode:    music.AeolianMode,
			},
			want: Fretboard{
				System:      "Quarter-Comma Meantone",
				Description: "Diatonic (Aeolian) fret positions based on Meantone temperament achieved by narrowing of fifths by 0.25 of a syntonic comma (81/80).",
				ScaleLength: 650.0,
				Frets: []Fret{
					{Position: 0, Interval: "1"},
					{Position: 68.6, Interval: "1.118", Cents: 193.1},     // major 2nd
					{Position: 106.52, Interval: "1.196", Cents: 309.86},  // minor 3rd
					{Position: 163.84, Interval: "1.337", Cents: 502.8},   // perfect 4th
					{Position: 215.22, Interval: "1.495", Cents: 696.17},  // perfect 5th
					{Position: 243.75, Interval: "1.6", Cents: 813.69},    // minor 6th
					{Position: 286.67, Interval: "1.789", Cents: 1006.98}, // minor 7th
					{Position: 325, Interval: "2", Cents: 1200},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFretboardFromTemperedScale(tt.args.length, tt.args.octaves, tt.args.scale, tt.args.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFretboardFromTemperedScale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringReturnsFretboardInJsonFormat(t *testing.T) {

	// Given
	f := Fretboard{
		System:      "Equal Temperament",
		Description: "Fret positions based on 12-tone equal temperament.",
		ScaleLength: 650.0,
		Frets: []Fret{
			{Label: "0.00 cents", Position: 0},
			{Label: "100.00 cents", Position: 36.48},
		},
	}

	// When
	returnedJson := f.String()

	// Then
	expectedJson := `{"system": "Equal Temperament", "description": "Fret positions based on 12-tone equal temperament.","scaleLength": 650, "frets": [ {"label": "0.00 cents", "position": 0}, {"label": "100.00 cents", "position": 36.48} ] }`
	assert.JSONEq(t, expectedJson, returnedJson)
}
