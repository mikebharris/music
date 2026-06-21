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

// NewJustIntonationChromaticScaleCharacteristicOfLimit builds a chromatic scale that
// actively uses the characteristic intervals of the limit prime in each window. For
// each 100-cent window it prefers the lowest-Benedetti-height ratio whose numerator
// or denominator contains the limit prime as a factor, falling back to the simplest
// available ratio only when no prime-containing candidate exists. Perfect consonances
// (4/3 and 3/2) are always protected so that the scale remains tonally coherent.
func NewJustIntonationChromaticScaleCharacteristicOfLimit(limit int) JustScale {
	return JustScale{
		system:      fmt.Sprintf("%d-limit Just Intonation (characteristic)", limit),
		description: fmt.Sprintf("Just Intonation chromatic scale using characteristic %d-limit ratios.", limit),
		algorithm: func() []JustInterval {
			return computeJustScaleCharacteristicOfLimit(limit, func(interval JustInterval) bool {
				return false
			})
		},
	}
}

func computeJustScaleCharacteristicOfLimit(limit int, filter intervalFilterFunction) []JustInterval {
	var primeMultipliers [][][]uint
	for p := 2; p <= limit; p++ {
		if isPrime(p) {
			primeMultipliers = append(primeMultipliers, multipliers(uint(p)))
		}
	}
	pool := justIntervalsFromMultipliers(buildMultiplierTablesFrom(primeMultipliers...), filter)

	prime := uint(limit)
	var result = []JustInterval{Unison()}
	for r := 50.0; r <= 1200.0; r += 100 {
		var candidates []JustInterval
		for _, iv := range pool {
			cents := iv.ToCents()
			if cents >= r && cents < r+100 {
				candidates = append(candidates, iv)
			}
		}
		if len(candidates) == 0 {
			continue
		}

		simplest := candidates[0]
		for _, iv := range candidates[1:] {
			if iv.BenedettiHeight() < simplest.BenedettiHeight() {
				simplest = iv
			}
		}

		// Perfect consonances (4/3, 3/2) are protected: without this, those windows
		// would be displaced by exotic prime-containing ratios (e.g. 21/16, 16/11).
		if simplest.IsPerfect() {
			result = append(result, simplest)
			continue
		}

		var bestWithPrime JustInterval
		for _, iv := range candidates {
			if iv.Numerator()%prime == 0 || iv.Denominator()%prime == 0 {
				if bestWithPrime == (JustInterval{}) || iv.BenedettiHeight() < bestWithPrime.BenedettiHeight() {
					bestWithPrime = iv
				}
			}
		}

		if bestWithPrime != (JustInterval{}) {
			result = append(result, bestWithPrime)
		} else {
			result = append(result, simplest)
		}
	}
	return result
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

// NewPartch43ToneScale returns Harry Partch's 43-tone 11-limit scale as described
// in Genesis of a Music (1949/1974). The scale consists of 43 intervals per octave
// (44 entries including the octave) built from otonalities and utonalities on the
// 11-limit diamond, arranged in ascending pitch order.
func NewPartch43ToneScale() JustScale {
	return JustScale{
		system:      "Partch 43-tone",
		description: "Harry Partch's 43-tone 11-limit just intonation scale from Genesis of a Music.",
		algorithm:   computePartch43ToneScale,
	}
}

func computePartch43ToneScale() []JustInterval {
	return IntervalsFromIntegers([][]uint{
		{1, 1}, {81, 80}, {33, 32}, {21, 20}, {16, 15}, {12, 11}, {11, 10}, {10, 9},
		{9, 8}, {8, 7}, {7, 6}, {32, 27}, {6, 5}, {11, 9}, {5, 4}, {14, 11},
		{9, 7}, {21, 16}, {4, 3}, {27, 20}, {11, 8}, {7, 5}, {10, 7}, {16, 11},
		{40, 27}, {3, 2}, {32, 21}, {14, 9}, {11, 7}, {8, 5}, {18, 11}, {5, 3},
		{27, 16}, {12, 7}, {7, 4}, {16, 9}, {9, 5}, {20, 11}, {11, 6}, {15, 8},
		{40, 21}, {64, 33}, {160, 81}, {2, 1},
	})
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

// NewHarmonicSeriesScale returns a JustScale built from the first n partials of
// the harmonic series, expressed as ratios relative to the fundamental (partial 1).
// Each partial k gives the interval k/1, which is then octave-reduced to [1, 2).
// The octave (2/1) is always included as the final degree.
func NewHarmonicSeriesScale(partials uint) JustScale {
	return JustScale{
		system:      "Harmonic Series",
		description: fmt.Sprintf("First %d partials of the harmonic series, octave-reduced to a single octave.", partials),
		algorithm: func() []JustInterval {
			seen := make(map[JustInterval]bool)
			var intervals []JustInterval
			for k := uint(1); k <= partials; k++ {
				interval := JustInterval{numerator: k, denominator: 1}.OctaveReduce().Simplify()
				if seen[interval] {
					continue
				}
				seen[interval] = true
				intervals = append(intervals, interval)
			}
			intervals = append(intervals, Octave())
			SortIntervals(intervals)
			return selectJustestIntervals(intervals)
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

func (s JustScale) IntervalStrings() []string {
	var result []string
	for _, i := range s.Intervals() {
		result = append(result, i.String())
	}
	return result
}

func (s JustScale) ToFrequenciesForTonicOf(tonic uint, octaves uint) []float64 {
	var frequencies []float64
	for octave := uint(0); octave < octaves; octave++ {
		octaveMultiplier := math.Pow(2, float64(octave))
		for _, interval := range s.Intervals() {
			if interval.IsUnison() && octave > 0 {
				continue
			}
			f := float64(tonic) * octaveMultiplier * float64(interval.Numerator()) / float64(interval.Denominator())
			frequencies = append(frequencies, math.Round(f))
		}
	}
	return frequencies
}

type computeJustIntervalsFn func() []JustInterval

// ptolemyIntenseDiatonicSteps maps each diatonic mode to its sequence of successive
// step intervals. Defined at package level to avoid re-allocation on every call.
var ptolemyIntenseDiatonicSteps = map[MusicalMode][]JustInterval{
	LydianMode:     {GreaterMajorSecond(), LesserMajorSecond(), GreaterMajorSecond(), DiatonicSemitone(), LesserMajorSecond(), GreaterMajorSecond(), DiatonicSemitone()},
	IonianMode:     {GreaterMajorSecond(), LesserMajorSecond(), DiatonicSemitone(), GreaterMajorSecond(), LesserMajorSecond(), GreaterMajorSecond(), DiatonicSemitone()},
	MixolydianMode: {GreaterMajorSecond(), LesserMajorSecond(), DiatonicSemitone(), GreaterMajorSecond(), LesserMajorSecond(), DiatonicSemitone(), GreaterMajorSecond()},
	DorianMode:     {GreaterMajorSecond(), DiatonicSemitone(), LesserMajorSecond(), GreaterMajorSecond(), LesserMajorSecond(), DiatonicSemitone(), GreaterMajorSecond()},
	AeolianMode:    {GreaterMajorSecond(), DiatonicSemitone(), LesserMajorSecond(), GreaterMajorSecond(), DiatonicSemitone(), GreaterMajorSecond(), LesserMajorSecond()},
	PhrygianMode:   {DiatonicSemitone(), GreaterMajorSecond(), LesserMajorSecond(), GreaterMajorSecond(), DiatonicSemitone(), GreaterMajorSecond(), LesserMajorSecond()},
	LocrianMode:    {DiatonicSemitone(), GreaterMajorSecond(), LesserMajorSecond(), DiatonicSemitone(), GreaterMajorSecond(), LesserMajorSecond(), GreaterMajorSecond()},
}

func computePtolemyIntenseDiatonicScale(mode MusicalMode) []JustInterval {
	var interval = Unison()
	var intervals = []JustInterval{interval}

	for _, v := range ptolemyIntenseDiatonicSteps[mode] {
		interval = JustInterval{numerator: interval.numerator * v.numerator, denominator: interval.denominator * v.denominator}.Simplify()
		intervals = append(intervals, interval)
	}

	return intervals
}

func computePythagoreanIntervals() []JustInterval {
	var fifthsFromTonicToCompute = 6
	var intervals []JustInterval
	for i := -fifthsFromTonicToCompute; i <= fifthsFromTonicToCompute; i++ {
		intervals = append(intervals, PerfectFifth().ToPowerOf(i).OctaveReduce())
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

		acuteRatio := interval.Add(SyntonicComma())
		graveRatio := interval.Add(SyntonicComma().Reciprocal())

		if graveRatio.denominator < acuteRatio.denominator {
			intervals = append(intervals, graveRatio)
		} else {
			intervals = append(intervals, acuteRatio)
		}
	}
	return intervals
}

func computeJustScale(multipliers [][]uint, filter intervalFilterFunction) []JustInterval {
	return selectJustestIntervals(justIntervalsFromMultipliers(multipliers, filter))
}

func selectJustestIntervals(poolOfPotentialIntervals []JustInterval) []JustInterval {
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

		// chosen interval is the one with the lowest Benedetti height (numerator * denominator),
		// which reliably selects the simplest integer ratio.
		var chosenInterval JustInterval
		for i, interval := range intervalsInNoteRange {
			if i == 0 || interval.BenedettiHeight() < chosenInterval.BenedettiHeight() {
				chosenInterval = interval
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
