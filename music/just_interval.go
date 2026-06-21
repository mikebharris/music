package music

import (
	"cmp"
	"fmt"
	"math"
	"slices"
)

type JustInterval struct {
	numerator   uint
	denominator uint
	name        string
}

func NewInterval(numerator, denominator uint) JustInterval {
	return JustInterval{numerator: numerator, denominator: denominator}.Simplify()
}

func (i JustInterval) Numerator() uint {
	return i.numerator
}

func (i JustInterval) Denominator() uint {
	return i.denominator
}

func (i JustInterval) IsUnison() bool {
	return i.numerator == 1 && i.denominator == 1
}

func (i JustInterval) IsEqualTo(other JustInterval) bool {
	return i.numerator == other.numerator && i.denominator == other.denominator
}

func (i JustInterval) IsDiminishedFifth() bool {
	return i.numerator == 64 && i.denominator == 45
}

func (i JustInterval) IsLesserMajorSecond() bool {
	return i.numerator == 10 && i.denominator == 9
}

func (i JustInterval) IsGreaterMajorSecond() bool {
	return i.numerator == 9 && i.denominator == 8
}

func (i JustInterval) IsLesserMinorSeventh() bool {
	return i.numerator == 16 && i.denominator == 9
}

func (i JustInterval) IsGreaterMinorSeventh() bool {
	return i.numerator == 9 && i.denominator == 5
}

func (i JustInterval) Add(other JustInterval) JustInterval {
	interval := JustInterval{
		numerator:   i.numerator * other.numerator,
		denominator: i.denominator * other.denominator,
	}.Simplify()

	return interval
}

func (i JustInterval) IsPerfectFourth() bool {
	return i.numerator == 4 && i.denominator == 3
}

func (i JustInterval) IsPerfect() bool {
	simplestForm := i.OctaveReduce().Simplify()
	return simplestForm.IsUnison() || simplestForm.IsPerfectFourth() || simplestForm.IsPerfectFifth() || simplestForm.IsOctave()
}

func (i JustInterval) IsPerfectFifth() bool {
	return i.numerator == 3 && i.denominator == 2
}

func (i JustInterval) IsOctave() bool {
	return i.numerator == 2 && i.denominator == 1
}

func (i JustInterval) Simplify() JustInterval {
	if i.denominator == 0 || i.numerator == 0 {
		return JustInterval{numerator: 1, denominator: 1}
	}
	gcd := func(a, b uint) uint {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}(i.numerator, i.denominator)
	i.numerator = i.numerator / gcd
	i.denominator = i.denominator / gcd
	return i
}

func (i JustInterval) OctaveReduce() JustInterval {
	if i.numerator == 0 || i.denominator == 0 {
		return Unison()
	}
	for i.ToFloat() >= 2.0 || i.ToFloat() < 1.0 {
		if i.ToFloat() < 1.0 {
			i.numerator *= 2
		}
		if i.ToFloat() >= 2.0 {
			i.denominator *= 2
		}
	}
	return i
}

func (i JustInterval) LessThan(other JustInterval) bool {
	return i.numerator*other.denominator < other.numerator*i.denominator
}

func (i JustInterval) GreaterThan(other JustInterval) bool {
	return !i.LessThan(other) && !i.IsEqualTo(other)
}

func (i JustInterval) Subtract(other JustInterval) JustInterval {
	if i.LessThan(other) {
		return JustInterval{numerator: i.denominator * other.numerator, denominator: i.numerator * other.denominator}.Simplify()
	} else if i.GreaterThan(other) {
		return JustInterval{numerator: i.numerator * other.denominator, denominator: i.denominator * other.numerator}.Simplify()
	}
	return JustInterval{numerator: 1, denominator: 1}
}

func (i JustInterval) Name() string {
	for _, n := range intervalNames {
		if n.numerator == i.numerator && n.denominator == i.denominator {
			return n.name
		}
	}
	return ""
}

func (i JustInterval) ToFloat() float64 {
	return float64(i.numerator) / float64(i.denominator)
}

func (i JustInterval) ToTemperedInterval() TemperedInterval {
	if i.denominator == 0 || i.numerator == 0 {
		return TemperedInterval(0)
	}
	return TemperedInterval(float64(i.numerator) / float64(i.denominator))
}

func (i JustInterval) ToPowerOf(p int) JustInterval {
	if p < 0 {
		return i.Reciprocal().ToPowerOf(-p)
	}
	if p == 0 {
		return Unison()
	}
	return JustInterval{
		numerator:   uint(math.Pow(float64(i.numerator), float64(p))),
		denominator: uint(math.Pow(float64(i.denominator), float64(p))),
	}.Simplify()
}

func (i JustInterval) Reciprocal() JustInterval {
	interval := JustInterval{denominator: i.numerator, numerator: i.denominator}
	return interval
}

// https://en.xen.wiki/w/Benedetti_height
func (i JustInterval) BenedettiHeight() uint {
	return i.numerator * i.denominator
}

// https://en.xen.wiki/w/Tenney_norm
func (i JustInterval) TenneyNorm() float64 {
	return math.Log2(float64(i.BenedettiHeight()))
}

func (i JustInterval) HarmonicDistance() float64 {
	return i.TenneyNorm()
}

func PerfectFourth() JustInterval {
	return JustInterval{numerator: 4, denominator: 3}
}
func Unison() JustInterval {
	return JustInterval{numerator: 1, denominator: 1}
}

func GraveUnison() JustInterval {
	return JustInterval{numerator: 80, denominator: 81}
}

func PythagoreanComma() JustInterval {
	return JustInterval{numerator: 531441, denominator: 524288}
}

func AcuteUnison() JustInterval {
	return SyntonicComma()
}

func SyntonicComma() JustInterval {
	return JustInterval{numerator: 81, denominator: 80}
}

func Dieses() JustInterval {
	return JustInterval{numerator: 128, denominator: 125}
}

func Schisma() JustInterval {
	return JustInterval{numerator: 32805, denominator: 32768}
}

func JustChromaticSemitone() JustInterval {
	return JustInterval{numerator: 25, denominator: 24}
}

func LesserMajorSecond() JustInterval {
	return JustInterval{numerator: 10, denominator: 9}
}
func GreaterMajorSecond() JustInterval {
	return JustInterval{numerator: 9, denominator: 8}
}
func DiatonicSemitone() JustInterval {
	return JustInterval{numerator: 16, denominator: 15}
}
func PerfectFifth() JustInterval {
	return JustInterval{numerator: 3, denominator: 2}
}
func Octave() JustInterval {
	return JustInterval{numerator: 2, denominator: 1}
}

// Invert returns the octave complement of the interval (2/1 ÷ interval).
// The interval must be octave-reduced (i.e. in [1/1, 2/1]); compound intervals
// such as a twelfth (3/1) will produce a musically incorrect result.
func (i JustInterval) Invert() JustInterval {
	return Octave().Subtract(i.OctaveReduce())
}

// DeviationFromEqualTemperament returns the deviation in cents from the nearest
// 12-tone equal temperament pitch.
func (i JustInterval) DeviationFromEqualTemperament() float64 {
	cents := i.ToCents()
	nearest := math.Round(cents/100) * 100
	return cents - nearest
}

var intervalNames = []JustInterval{
	{1, 1, "Perfect Unison"},
	{531441, 524288, "Pythagorean Comma"},
	{32805, 32768, "Schisma"},
	{225, 224, "Septimal Kleisma"},
	{81, 80, "Syntonic Comma"},
	{80, 81, "Grave Unison"},
	{128, 125, "Dieses (Diminished Second)"},
	{25, 24, "Just (Lesser) Chromatic Semitone"},
	{256, 243, "Pythagorean Minor Second"},
	{135, 128, "Greater Chromatic Semitone"},
	{27, 25, "Acute Minor Second"},
	{17, 16, "Large Septendecimal Semitone"},
	{16, 15, "Minor Second"},
	{13, 12, "Tridecimal Minor Second (Avicenna)"},
	{12, 11, "Undecimal Minor Second"},
	{15, 14, "Septimal Minor Second"},
	{10, 9, "Just (Lesser) Major Second"},
	{9, 8, "Pythagorean (Greater) Major Second"},
	{8, 7, "Septimal Major Second"},
	{7, 6, "Septimal Minor Third"},
	{19, 16, "Otonol Minor Third"},
	{6, 5, "Minor Third"},
	{5, 4, "Major Third"},
	{9, 7, "Septimal Major Third"},
	{32, 27, "Pythagorean Minor Third"},
	{81, 64, "Pythagorean Major Third"},
	{4, 3, "Perfect Fourth"},
	{12, 16, "Septimal Sub Forth"},
	{11, 8, "Undecimal Tritone"},
	{45, 32, "Augmented Fourth"},
	{7, 5, "Septimal Augmented Fourth"},
	{1024, 729, "Pythagorean Diminished Fifth"},
	{729, 512, "Pythagorean Augmented Fourth"},
	{64, 45, "Diminished Fifth"},
	{10, 7, "Septimal Diminished Fifth"},
	{40, 27, "Grave Fifth"},
	{3, 2, "Perfect Fifth"},
	{8, 5, "Just Minor Sixth"},
	{128, 81, "Pythagorean Minor Sixth"},
	{13, 8, "Lesser Tridecimal Neutral Sixth"},
	{5, 3, "Major Sixth"},
	{225, 128, "Acute Augmented Sixth"},
	{27, 16, "Pythagorean Major Sixth"},
	{16, 9, "Pythagorean (Lesser) Minor Seventh"},
	{9, 5, "Just (Greater) Minor Seventh"},
	{7, 4, "Septimal (Harmonic) Minor Seventh"},
	{15, 8, "Just Major Seventh"},
	{13, 7, "Tridecimal Major Seventh"},
	{243, 128, "Pythagorean Major Seventh"},
	{48, 25, "Diminished Octave"},
	{125, 64, "Just Augmented Seventh"},
	{160, 81, "Grave Octave"},
	{2, 1, "Perfect Octave"},
}

func (i JustInterval) sortWith(j JustInterval) int {
	return cmp.Compare(i.numerator*j.denominator, j.numerator*i.denominator)
}

func (i JustInterval) String() string {
	return fmt.Sprintf("%d:%d", i.numerator, i.denominator)
}

func (i JustInterval) ToCents() float64 {
	return math.Log2(float64(i.numerator)/float64(i.denominator)) * 1200
}

func (i JustInterval) Diff(other JustInterval) JustInterval {
	if i.LessThan(other) {
		return other.Subtract(i)
	}
	return i.Subtract(other)
}

func IntervalsFromIntegers(integers [][]uint) []JustInterval {
	var intervals []JustInterval
	for _, pair := range integers {
		intervals = append(intervals, FromIntArray(pair))
	}
	return intervals
}

func FromIntArray(i []uint) JustInterval {
	return JustInterval{numerator: i[0], denominator: i[1]}.Simplify()
}

func SortIntervals(intervals []JustInterval) {
	slices.SortFunc(intervals, func(i, j JustInterval) int {
		return i.sortWith(j)
	})
}

// intervalFilterFunction defines a function type for excluding certain ratios based on scale symmetry.
type intervalFilterFunction func(ratio JustInterval) bool

// multipliers generates ratio pairs {base^n, 1} and {1, base^n} for n=1..power,
// plus the unison {1, 1}. For base 3 a power of 2 is used so the Pythagorean
// chromatic scale has enough distinct pitches; all other bases use power 1.
func multipliers(base uint) [][]uint {
	power := 1
	if base == 3 {
		power = 2
	}
	result := make([][]uint, 0, 2*power+1)
	for n := power; n >= 1; n-- {
		exp := uint(math.Pow(float64(base), float64(n)))
		result = append(result, []uint{exp, 1})
	}
	result = append(result, []uint{1, 1})
	for n := 1; n <= power; n++ {
		exp := uint(math.Pow(float64(base), float64(n)))
		result = append(result, []uint{1, exp})
	}
	return result
}

func justIntervalsFromMultipliers(multiplierList [][]uint, filter intervalFilterFunction) []JustInterval {
	seen := make(map[JustInterval]bool)
	var intervals []JustInterval
	for _, multiplier := range multiplierList {
		interval := JustInterval{numerator: multiplier[0], denominator: multiplier[1]}.OctaveReduce().Simplify()
		if interval.IsDiminishedFifth() {
			continue
		}
		if filter(interval) {
			continue
		}
		if seen[interval] {
			continue
		}
		seen[interval] = true
		intervals = append(intervals, interval)
	}
	octave := Octave()
	if !seen[octave] {
		intervals = append(intervals, octave)
	}
	SortIntervals(intervals)
	return intervals
}

func createMultiplierTableOf(multiplierListA, multiplierListB [][]uint) [][]uint {
	var multiplierTable [][]uint
	for _, multiplierA := range multiplierListA {
		for _, multiplierB := range multiplierListB {
			multiplierTable = append(multiplierTable, []uint{multiplierA[0] * multiplierB[0], multiplierA[1] * multiplierB[1]})
		}
	}
	return multiplierTable
}
