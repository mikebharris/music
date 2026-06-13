package instruments

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/mikebharris/music/music"
)

type Fret struct {
	Label    string  `json:"label"`
	Position float64 `json:"position"`
	Comment  string  `json:"comment,omitempty"`
	Interval string  `json:"interval,omitempty"`
	Cents    float64 `json:"cents,omitempty"`
}

type Fretboard struct {
	System      string  `json:"system"`
	Description string  `json:"description,omitempty"`
	ScaleLength float64 `json:"scaleLength"`
	Frets       []Fret  `json:"frets"`
}

// String returns the Fretboard as a JSON string.
func (f Fretboard) String() string {
	b, _ := json.MarshalIndent(f, "", "  ")
	return string(b)
}

func NewFretboardFromJustScale(length float64, octaves int, scale music.JustScale, mode music.MusicalMode) Fretboard {
	fretboard := Fretboard{
		System:      scale.System(),
		Description: fmt.Sprintf("%s fret positions based on %s", scaleTypeFromMode(mode), scale.Description()),
		ScaleLength: length,
	}
	fretboard.makeJustFrets(scale.Intervals(), mode, octaves)
	return fretboard
}

func NewFretboardFromTemperedScale(length float64, octaves int, scale music.TemperedScale, mode music.MusicalMode) Fretboard {
	fretboard := Fretboard{
		System:      scale.System(),
		Description: fmt.Sprintf("%s fret positions based on %s", scaleTypeFromMode(mode), scale.Description()),
		ScaleLength: length,
	}
	fretboard.makeTemperedFrets(scale.Intervals(), mode, octaves)
	return fretboard
}

func scaleTypeFromMode(mode music.MusicalMode) string {
	var f string
	if mode.IsDiatonic() {
		f = "Diatonic (" + mode.String() + ")"
	} else {
		f = "Chromatic"
	}
	return f
}

func (f *Fretboard) makeTemperedFrets(intervals []music.TemperedInterval, mode music.MusicalMode, octaves int) {
	for octave := 0; octave < octaves; octave++ {
		for i, interval := range intervals {
			if octave > 0 && i == 0 {
				continue // skip unison on subsequent octaves
			}
			semitone := int(math.Round(interval.ToCents() / 100))
			if mode.IsDiatonic() && i > 0 && !diatonicFretMap(semitone, mode) {
				continue // skip the note if it's not in the diatonic scale
			}
			cents := interval.ToCents() + float64(octave)*1200
			f.Frets = append(f.Frets, Fret{
				Position: math.Round((f.ScaleLength-(f.ScaleLength/interval.Value())/math.Pow(2, float64(octave)))*100) / 100,
				Interval: interval.String(),
				Cents:    cents,
			})
		}
	}
}

func diatonicFretMap(intervalPosition int, mode music.MusicalMode) bool {
	var diatonicModes = map[music.MusicalMode][]bool{
		music.LydianMode:     {true, false, true, false, true, false, true, true, false, true, false, true, true},
		music.IonianMode:     {true, false, true, false, true, true, false, true, false, true, false, true, true},
		music.MixolydianMode: {true, false, true, false, true, true, false, true, false, true, true, false, true},
		music.DorianMode:     {true, false, true, true, false, true, false, true, false, true, true, false, true},
		music.AeolianMode:    {true, false, true, true, false, true, false, true, true, false, true, false, true},
		music.PhrygianMode:   {true, true, false, true, false, true, false, true, true, false, true, false, true},
		music.LocrianMode:    {true, true, false, true, false, true, true, false, true, false, true, false, true},
	}
	if intervalPosition > 0 && intervalPosition < 12 {
		return diatonicModes[mode][intervalPosition]
	}
	return true
}

func (f *Fretboard) makeJustFrets(intervals []music.JustInterval, mode music.MusicalMode, octaves int) {
	for octave := 0; octave < octaves; octave++ {
		for _, interval := range intervals {
			if octave > 0 && interval == music.Unison() {
				continue // skip unison on subsequent octaves
			}
			semitone := int(math.Round(interval.ToCents() / 100))
			if mode.IsDiatonic() && semitone > 0 && !diatonicFretMap(semitone, mode) {
				continue
			}
			// Fret position: distance from nut = ScaleLength * (1 - denom / (numer * 2^octave))
			position := f.ScaleLength * (1 - float64(interval.Denominator())/(float64(interval.Numerator())*math.Pow(2, float64(octave))))
			f.Frets = append(f.Frets, Fret{
				Position: math.Round(position*100) / 100,
				Comment:  interval.Name(),
				Interval: music.Octave().ToPowerOf(octave).Add(interval).String(),
				Cents:    (float64(octave) * music.Octave().ToCents()) + interval.ToCents(),
			})
		}
	}
}
