package music

import (
	"fmt"
	"math"
	"slices"
)

type JustScale struct {
	system      string
	description string
	algorithm   computeJustIntervalsFn
}

func NewPythagoreanScale() JustScale {
	return JustScale{
		system:      "Pythagorean",
		description: "3-limit Pythagorean ratios.",
		algorithm:   computePythagoreanIntervals,
	}
}

func New5LimitPythagoreanScale() JustScale {
	return JustScale{
		system:      "5-limit Pythagorean",
		description: "5-limit just intonation pure ratios chromatic scale derived from applying syntonic comma to Pythagorean ratios.",
		algorithm:   compute5LimitPythagoreanIntervals,
	}
}

func NewJustIntonationChromaticScaleWithLimit(limit int) JustScale {
	return NewJustIntonationChromaticScaleWithLimitAndFilter(limit, func(interval JustInterval) bool {
		return false
	})
}

func NewJustIntonationChromaticScaleWithLimitAndFilter(limit int, filter func(interval JustInterval) bool) JustScale {
	return JustScale{
		system:      fmt.Sprintf("%d-limit Just Intonation", limit),
		description: fmt.Sprintf("Just Intonation chromatic scale based on %d-limit pure ratios.", limit),
		algorithm: func() []JustInterval {
			return computeJustScaleWithLimit(limit, filter)
		},
	}
}

func computeJustScaleWithLimit(limit int, filter func(interval JustInterval) bool) []JustInterval {
	var primeMultipliers [][][]uint
	for p := 2; p <= limit; p++ {
		if isPrime(p) {
			primeMultipliers = append(primeMultipliers, multipliers(uint(p)))
		}
	}
	return computeJustScale(buildMultiplierTablesFrom(primeMultipliers...), filter)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func NewIntenseDiatonicScale(mode MusicalMode) JustScale {
	return JustScale{
		system:      "Ptolemy Intense Diatonic",
		description: fmt.Sprintf("Ptolemy's 5-limit intense diatonic scale in %s mode.", mode),
		algorithm: func() []JustInterval {
			return computePtolemyIntenseDiatonicScale(mode)
		},
	}
}

func NewSazScale() JustScale {
	// as per https://en.wikipedia.org/wiki/Ba%C4%9Flama and the cura that I have
	return JustScale{
		system:      "Saz",
		description: "Turkish Saz tuning ratios.",
		algorithm:   computeSazScale,
	}
}

func NewJustIntonationChromaticScaleWith(description string, intervals [][]uint) JustScale {
	return JustScale{
		system:      "Just Intonation",
		description: description,
		algorithm: func() []JustInterval {
			return IntervalsFromIntegers(intervals)
		},
	}
}

func (s JustScale) System() string {
	return s.system
}

func (s JustScale) Description() string {
	return s.description
}

func (s JustScale) Intervals() []JustInterval {
	return s.algorithm()
}

type computeJustIntervalsFn func() []JustInterval

func computePtolemyIntenseDiatonicScale(mode MusicalMode) []JustInterval {
	greaterMajorSecond := GreaterMajorSecond()
	lesserMajorSecond := LesserMajorSecond()
	diatonicSemitone := DiatonicSemitone()

	var intervalMap = map[MusicalMode][]JustInterval{
		LydianMode:     {greaterMajorSecond, lesserMajorSecond, greaterMajorSecond, diatonicSemitone, lesserMajorSecond, greaterMajorSecond, diatonicSemitone},
		IonianMode:     {greaterMajorSecond, lesserMajorSecond, diatonicSemitone, greaterMajorSecond, lesserMajorSecond, greaterMajorSecond, diatonicSemitone},
		MixolydianMode: {greaterMajorSecond, lesserMajorSecond, diatonicSemitone, greaterMajorSecond, lesserMajorSecond, diatonicSemitone, greaterMajorSecond},
		DorianMode:     {greaterMajorSecond, diatonicSemitone, lesserMajorSecond, greaterMajorSecond, lesserMajorSecond, diatonicSemitone, greaterMajorSecond},
		AeolianMode:    {greaterMajorSecond, diatonicSemitone, lesserMajorSecond, greaterMajorSecond, diatonicSemitone, greaterMajorSecond, lesserMajorSecond},
		PhrygianMode:   {diatonicSemitone, greaterMajorSecond, lesserMajorSecond, greaterMajorSecond, diatonicSemitone, greaterMajorSecond, lesserMajorSecond},
		LocrianMode:    {diatonicSemitone, greaterMajorSecond, lesserMajorSecond, diatonicSemitone, greaterMajorSecond, lesserMajorSecond, greaterMajorSecond},
	}

	var interval = Unison()
	var intervals = []JustInterval{interval}

	for _, v := range intervalMap[mode] {
		interval = JustInterval{numerator: interval.numerator * v.numerator, denominator: interval.denominator * v.denominator}.Simplify()
		intervals = append(intervals, interval)
	}

	return intervals
}

func computePythagoreanIntervals() []JustInterval {
	var fifthsFromTonicToCompute = 6
	var intervals []JustInterval
	for i := -fifthsFromTonicToCompute; i <= fifthsFromTonicToCompute; i++ {
		if i < 0 {
			intervals = append(intervals, PerfectFifth().ToPowerOf(i).Reciprocal().OctaveReduce())
		} else {
			intervals = append(intervals, PerfectFifth().ToPowerOf(i).OctaveReduce())
		}
	}

	intervals = append(intervals, Octave())
	slices.SortFunc(intervals, func(i, j JustInterval) int {
		return i.sortWith(j)
	})
	return intervals
}

func compute5LimitPythagoreanIntervals() []JustInterval {
	var intervals []JustInterval
	for _, interval := range computePythagoreanIntervals() {
		if interval.IsPerfect() {
			intervals = append(intervals, interval)
			continue
		}

		graveRatio := interval.Add(AcuteUnison())
		acuteRatio := interval.Add(GraveUnison())

		if graveRatio.denominator < acuteRatio.denominator {
			intervals = append(intervals, graveRatio)
		} else {
			intervals = append(intervals, acuteRatio)
		}
	}
	return intervals
}

func computeJustScale(multipliers [][]uint, filter intervalFilterFunction) []JustInterval {
	poolOfPotentialIntervals := justIntervalsFromMultipliers(multipliers, filter)
	var preferredIntervals = []JustInterval{Unison()}
	centsInOctave := 1200.0
	for r := 50.0; r <= centsInOctave; r += 100 {
		var intervalsInNoteRange []JustInterval
		for _, interval := range poolOfPotentialIntervals {
			cents := interval.ToCents()
			if cents >= r && cents < r+100 {
				intervalsInNoteRange = append(intervalsInNoteRange, interval)
			}
		}

		//   chosen interval is the simplest integer ratio
		var chosenInterval JustInterval
		for i, interval := range intervalsInNoteRange {
			if i == 0 || (interval.numerator < chosenInterval.numerator && interval.denominator < chosenInterval.denominator) {
				chosenInterval = interval
				continue
			}
		}
		if chosenInterval == (JustInterval{}) {
			continue
		}
		preferredIntervals = append(preferredIntervals, chosenInterval)
	}

	return preferredIntervals
}

func buildMultiplierTablesFrom(multipliers ...[][]uint) [][]uint {
	if len(multipliers) == 1 {
		return multipliers[0]
	}
	return createMultiplierTableOf(multipliers[0], buildMultiplierTablesFrom(multipliers[1:]...))
}

func computeSazScale() []JustInterval {
	return IntervalsFromIntegers([][]uint{{1, 1}, {18, 17}, {12, 11}, {9, 8}, {81, 68}, {27, 22}, {81, 64}, {4, 3}, {24, 17}, {16, 11}, {3, 2}, {27, 17}, {18, 11}, {27, 16}, {16, 9}, {32, 17}, {64, 33}, {2, 1}})
}
