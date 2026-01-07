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
	if i.denominator == 0 {
		return i
	}
	if i.numerator == 0 {
		return JustInterval{numerator: 0, denominator: 1}
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
	return TemperedInterval(float64(i.numerator) / float64(i.denominator))
}

func (i JustInterval) ToPowerOf(p int) JustInterval {
	return JustInterval{numerator: uint(math.Pow(float64(i.numerator), math.Abs(float64(p)))), denominator: uint(math.Pow(float64(i.denominator), math.Abs(float64(p))))}.Simplify()
}

func (i JustInterval) Reciprocal() JustInterval {
	interval := JustInterval{denominator: i.numerator, numerator: i.denominator}
	return interval
}

func perfectFourth() JustInterval {
	return JustInterval{4, 3, "Perfect Fourth"}
}

var Unison = JustInterval{numerator: 1, denominator: 1}
var AcuteUnison = JustInterval{numerator: 81, denominator: 80}
var SyntonicComma = JustInterval{numerator: 81, denominator: 80}
var Dieses = JustInterval{numerator: 128, denominator: 125}
var JustChromaticSemitone = JustInterval{numerator: 25, denominator: 24}
var GraveUnison = JustInterval{numerator: 80, denominator: 81}
var LesserMajorSecond = JustInterval{numerator: 10, denominator: 9}
var GreaterMajorSecond = JustInterval{numerator: 9, denominator: 8}
var DiatonicSemitone = JustInterval{numerator: 16, denominator: 15}
var PerfectFifth = JustInterval{numerator: 3, denominator: 2}
var Octave = JustInterval{numerator: 2, denominator: 1}

var intervalNames = []JustInterval{
	{1, 1, "Perfect Unison"},
	{225, 224, "Septimal Kleisma"},
	{80, 81, "Syntonic Comma"},
	{81, 80, "Grave Unison"},
	{128, 125, "Dieses (Diminished Second)"},
	{25, 24, "Just (Lesser) Chromatic Semitone"},
	{256, 243, "Pythagorean Minor Second"},
	{135, 128, "Greater Chromatic Semitone"},
	{27, 25, "Acute Minor Second"},
	{16, 15, "Minor Second"},
	{13, 12, "Tridecimal Minor Second (Avicenna)"},
	{12, 11, "Undecimal Minor Second"},
	{15, 14, "Septimal Minor Second"},
	{10, 9, "Just (Lesser) Major Second"},
	{9, 8, "Pythagorean (Greater) Major Second"},
	{8, 7, "Septimal Major Second"},
	{6, 5, "Minor Third"},
	{5, 4, "Major Third"},
	{32, 27, "Pythagorean Minor Third"},
	{81, 64, "Pythagorean Major Third"},
	{4, 3, "Perfect Fourth"},
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
	{5, 3, "Major Sixth"},
	{27, 16, "Pythagorean Major Sixth"},
	{16, 9, "Pythagorean (Lesser) Minor Seventh"},
	{9, 5, "Just (Greater) Minor Seventh"},
	{7, 4, "Septimal (Harmonic) Minor Seventh"},
	{15, 8, "Just Major Seventh"},
	{13, 7, "Tridecimal Major Seventh"},
	{243, 128, "Pythagorean Major Seventh"},
	{2, 1, "Perfect Octave"},
}

func (i JustInterval) sortWith(j JustInterval) int {
	return cmp.Compare(i.numerator*j.denominator, j.numerator*i.denominator)
}

func (i JustInterval) String() string {
	return fmt.Sprintf("%d:%d", i.numerator, i.denominator)
}

func (i JustInterval) ToCents() float64 {
	return math.Log10(float64(i.numerator)/float64(i.denominator)) / math.Log10(2) * 1200
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

func multipliers(base uint) [][]uint {
	return [][]uint{{base, 1}, {1, 1}, {1, base}}
}

func justIntervalsFromMultipliers(multiplierList [][]uint, filter intervalFilterFunction) []JustInterval {
	var intervals []JustInterval
	for _, multiplier := range multiplierList {
		interval := JustInterval{numerator: multiplier[0], denominator: multiplier[1]}.OctaveReduce()
		if interval.IsDiminishedFifth() {
			continue
		}
		if filter(interval) {
			continue
		}
		intervals = append(intervals, interval)
	}
	intervals = append(intervals, Octave)
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
