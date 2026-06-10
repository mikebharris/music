package music

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturnPythagorean3LimitScaleWithExpectedScaleDegrees(t *testing.T) {
	// Given
	scale := NewPythagoreanScale()

	// When
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "Pythagorean", scale.System())
	assert.Equal(t, "3-limit Pythagorean ratios.", scale.Description())
	assert.Equal(t, 14, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 256, denominator: 243}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 8}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 32, denominator: 27}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 81, denominator: 64}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 1024, denominator: 729}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 729, denominator: 512}, intervals[7])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[8])
	assert.Equal(t, JustInterval{numerator: 128, denominator: 81}, intervals[9])
	assert.Equal(t, JustInterval{numerator: 27, denominator: 16}, intervals[10])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 9}, intervals[11])
	assert.Equal(t, JustInterval{numerator: 243, denominator: 128}, intervals[12])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[13])
}

func Test_ShouldReturn5LimitScaleBuiltFromPythagoreanWithExpectedScaleDegrees(t *testing.T) {
	// Given
	scale := New5LimitPythagoreanScale()

	// When
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "5-limit Pythagorean", scale.System())
	assert.Equal(t, "5-limit just intonation pure ratios chromatic scale derived from applying syntonic comma to Pythagorean ratios.", scale.Description())
	assert.Equal(t, 14, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 15}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 10, denominator: 9}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 6, denominator: 5}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 4}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 64, denominator: 45}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 45, denominator: 32}, intervals[7])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[8])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 5}, intervals[9])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 3}, intervals[10])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 5}, intervals[11])
	assert.Equal(t, JustInterval{numerator: 15, denominator: 8}, intervals[12])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[13])
}

func Test_shouldReturnFiveLimitJustChromaticScaleBasedOnPureRatios(t *testing.T) {
	// Given
	// Mike read the page at https://en.wikipedia.org/wiki/Five-limit_tuning

	// When
	scale := NewJustIntonationChromaticScaleWithLimit(5)
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "5-limit Just Intonation", scale.System())
	assert.Equal(t, "Just Intonation chromatic scale based on 5-limit pure ratios.", scale.Description())

	assert.Equal(t, 13, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 15}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 8}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 6, denominator: 5}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 4}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 45, denominator: 32}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[7])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 5}, intervals[8])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 3}, intervals[9])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 5}, intervals[10])
	assert.Equal(t, JustInterval{numerator: 15, denominator: 8}, intervals[11])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[12])
}

func Test_shouldReturnSevenLimitJustChromaticScaleBasedOnPureRatios(t *testing.T) {
	// Given
	// https://en.wikipedia.org/wiki/Seven-limit_tuning

	// When
	scale := NewJustIntonationChromaticScaleWithLimit(7)
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "7-limit Just Intonation", scale.System())
	assert.Equal(t, "Just Intonation chromatic scale based on 7-limit pure ratios.", scale.Description())

	assert.Equal(t, 13, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 15, denominator: 14}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 7}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 6, denominator: 5}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 4}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 7, denominator: 5}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[7])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 5}, intervals[8])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 3}, intervals[9])
	assert.Equal(t, JustInterval{numerator: 7, denominator: 4}, intervals[10])
	assert.Equal(t, JustInterval{numerator: 15, denominator: 8}, intervals[11])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[12])
}

func Test_shouldReturnSevenLimitJustChromaticScaleBasedOnPureRatiosWhenElevenLimitScaleRequested(t *testing.T) {
	// Given
	// https://en.xen.wiki/w/11-limit

	// When
	scale := NewJustIntonationChromaticScaleWithLimit(11)
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "11-limit Just Intonation", scale.System())
	assert.Equal(t, "Just Intonation chromatic scale based on 11-limit pure ratios.", scale.Description())

	// The 11-limit introduces no new winners: every 11-based interval has a higher
	// Benedetti height than the 7-limit interval already occupying the same cent window,
	// so the scale is identical to 7-limit.
	assert.Equal(t, 13, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 15, denominator: 14}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 7}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 6, denominator: 5}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 4}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 7, denominator: 5}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[7])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 5}, intervals[8])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 3}, intervals[9])
	assert.Equal(t, JustInterval{numerator: 7, denominator: 4}, intervals[10])
	assert.Equal(t, JustInterval{numerator: 15, denominator: 8}, intervals[11])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[12])
}

func Test_shouldReturnGenuineElevenLimitChromaticScaleWithCharacteristicIntervals(t *testing.T) {
	// Given
	// NewJustIntonationChromaticScaleCharacteristicOfLimit uses a modified selection
	// algorithm: for each 100-cent window it prefers the lowest-Benedetti-height interval
	// whose numerator or denominator contains the limit prime, falling back to the overall
	// simplest only when no such candidate exists. Perfect consonances (4/3, 3/2) are
	// protected so the scale stays tonally coherent.

	// When
	scale := NewJustIntonationChromaticScaleCharacteristicOfLimit(11)
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "11-limit Just Intonation (characteristic)", scale.System())
	assert.Equal(t, "Just Intonation chromatic scale using characteristic 11-limit ratios.", scale.Description())

	assert.Equal(t, 13, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 22, denominator: 21}, intervals[1])  // undecimal-septimal semitone
	assert.Equal(t, JustInterval{numerator: 11, denominator: 10}, intervals[2])  // undecimal major second
	assert.Equal(t, JustInterval{numerator: 11, denominator: 9}, intervals[3])   // undecimal neutral third
	assert.Equal(t, JustInterval{numerator: 14, denominator: 11}, intervals[4])  // undecimal major third
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[5])    // perfect fourth (protected)
	assert.Equal(t, JustInterval{numerator: 11, denominator: 8}, intervals[6])   // undecimal tritone — the characteristic 11-limit interval
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[7])    // perfect fifth (protected)
	assert.Equal(t, JustInterval{numerator: 11, denominator: 7}, intervals[8])   // undecimal minor sixth
	assert.Equal(t, JustInterval{numerator: 18, denominator: 11}, intervals[9])  // undecimal major sixth
	assert.Equal(t, JustInterval{numerator: 11, denominator: 6}, intervals[10])  // undecimal minor seventh
	assert.Equal(t, JustInterval{numerator: 21, denominator: 11}, intervals[11]) // undecimal major seventh
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[12])
}

func Test_shouldReturnElevenLimitOvertoneScaleFromHarmonicSeries(t *testing.T) {
	// Given
	// The harmonic series naturally produces 11-limit intervals. The first 11 partials
	// octave-reduced give an otonal scale whose characteristic interval is 11/8 (the
	// undecimal tritone at 551.3 cents), which the Benedetti-height algorithm can never
	// produce via NewJustIntonationChromaticScaleWithLimit.
	// Using 11 partials (not 16) keeps the scale within 11-limit: the 16-partial version
	// introduces 13/8, crossing into 13-limit territory.

	// When
	scale := NewHarmonicSeriesScale(11)
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "Harmonic Series", scale.System())
	assert.Equal(t, "First 11 partials of the harmonic series, octave-reduced to a single octave.", scale.Description())

	assert.Equal(t, 7, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 8}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 4}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 11, denominator: 8}, intervals[3]) // undecimal tritone: the characteristic 11-limit interval
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 7, denominator: 4}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[6])
}

func Test_shouldReturnGenuineThirteenLimitChromaticScaleWithCharacteristicIntervals(t *testing.T) {
	// Given
	// Uses NewJustIntonationChromaticScaleCharacteristicOfLimit: for each window the
	// lowest-Benedetti-height interval containing prime 13 is preferred. Perfect
	// consonances (4/3, 3/2) are protected. Every other interval contains factor 13.

	// When
	scale := NewJustIntonationChromaticScaleCharacteristicOfLimit(13)
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "13-limit Just Intonation (characteristic)", scale.System())
	assert.Equal(t, "Just Intonation chromatic scale using characteristic 13-limit ratios.", scale.Description())

	assert.Equal(t, 13, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 13, denominator: 12}, intervals[1])  // tridecimal minor second
	assert.Equal(t, JustInterval{numerator: 15, denominator: 13}, intervals[2])  // tridecimal major second
	assert.Equal(t, JustInterval{numerator: 13, denominator: 11}, intervals[3])  // tridecimal minor third
	assert.Equal(t, JustInterval{numerator: 16, denominator: 13}, intervals[4])  // tridecimal major third
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[5])    // perfect fourth (protected)
	assert.Equal(t, JustInterval{numerator: 13, denominator: 9}, intervals[6])   // tridecimal tritone
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[7])    // perfect fifth (protected)
	assert.Equal(t, JustInterval{numerator: 13, denominator: 8}, intervals[8])   // tridecimal minor sixth
	assert.Equal(t, JustInterval{numerator: 22, denominator: 13}, intervals[9])  // tridecimal major sixth
	assert.Equal(t, JustInterval{numerator: 26, denominator: 15}, intervals[10]) // tridecimal minor seventh
	assert.Equal(t, JustInterval{numerator: 13, denominator: 7}, intervals[11])  // tridecimal major seventh
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[12])
}

func Test_shouldReturnGenuineSeventeenLimitChromaticScaleWithCharacteristicIntervals(t *testing.T) {
	// When
	scale := NewJustIntonationChromaticScaleCharacteristicOfLimit(17)
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "17-limit Just Intonation (characteristic)", scale.System())
	assert.Equal(t, "Just Intonation chromatic scale using characteristic 17-limit ratios.", scale.Description())

	assert.Equal(t, 13, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 17, denominator: 16}, intervals[1])  // septendecimal minor second
	assert.Equal(t, JustInterval{numerator: 17, denominator: 15}, intervals[2])  // septendecimal major second
	assert.Equal(t, JustInterval{numerator: 17, denominator: 14}, intervals[3])  // septendecimal minor third
	assert.Equal(t, JustInterval{numerator: 21, denominator: 17}, intervals[4])  // septendecimal major third
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[5])    // perfect fourth (protected)
	assert.Equal(t, JustInterval{numerator: 17, denominator: 12}, intervals[6])  // septendecimal tritone
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[7])    // perfect fifth (protected)
	assert.Equal(t, JustInterval{numerator: 17, denominator: 11}, intervals[8])  // septendecimal minor sixth
	assert.Equal(t, JustInterval{numerator: 17, denominator: 10}, intervals[9])  // septendecimal major sixth
	assert.Equal(t, JustInterval{numerator: 30, denominator: 17}, intervals[10]) // septendecimal minor seventh
	assert.Equal(t, JustInterval{numerator: 17, denominator: 9}, intervals[11])  // septendecimal major seventh
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[12])
}

func Test_shouldReturnGenuineNineteenLimitChromaticScaleWithCharacteristicIntervals(t *testing.T) {
	// When
	scale := NewJustIntonationChromaticScaleCharacteristicOfLimit(19)
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "19-limit Just Intonation (characteristic)", scale.System())
	assert.Equal(t, "Just Intonation chromatic scale using characteristic 19-limit ratios.", scale.Description())

	assert.Equal(t, 13, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 19, denominator: 18}, intervals[1])  // nonadecimal minor second
	assert.Equal(t, JustInterval{numerator: 19, denominator: 17}, intervals[2])  // nonadecimal major second
	assert.Equal(t, JustInterval{numerator: 19, denominator: 16}, intervals[3])  // nonadecimal minor third
	assert.Equal(t, JustInterval{numerator: 19, denominator: 15}, intervals[4])  // nonadecimal major third
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[5])    // perfect fourth (protected)
	assert.Equal(t, JustInterval{numerator: 55, denominator: 38}, intervals[6])  // tritone — no simple 19-containing ratio exists here (h=2090)
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[7])    // perfect fifth (protected)
	assert.Equal(t, JustInterval{numerator: 19, denominator: 12}, intervals[8])  // nonadecimal minor sixth
	assert.Equal(t, JustInterval{numerator: 19, denominator: 11}, intervals[9])  // nonadecimal major sixth
	assert.Equal(t, JustInterval{numerator: 33, denominator: 19}, intervals[10]) // nonadecimal minor seventh
	assert.Equal(t, JustInterval{numerator: 19, denominator: 10}, intervals[11]) // nonadecimal major seventh
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[12])
}

func Test_shouldReturnBenJohnstonInspired31LimitChromaticScaleWithCharacteristicIntervals(t *testing.T) {
	// Ben Johnston (1926–2019) used up to 31-limit JI in his late string quartets
	// (notably No. 9, 1988). His system extends 5-limit JI through successive primes,
	// adding new harmonic colours at each step. This scale applies the characteristic
	// algorithm — lowest-Benedetti-height ratio containing prime 31 per window — to
	// produce a chromatic scale in the spirit of Johnston's high-limit practice.
	// Note: all non-protected intervals are of the form 31/n, since 31 is large enough
	// that the simplest ratios containing it all have 31 in the numerator.

	// When
	scale := NewJustIntonationChromaticScaleCharacteristicOfLimit(31)
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "31-limit Just Intonation (characteristic)", scale.System())
	assert.Equal(t, "Just Intonation chromatic scale using characteristic 31-limit ratios.", scale.Description())

	assert.Equal(t, 13, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 31, denominator: 29}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 31, denominator: 28}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 31, denominator: 26}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 31, denominator: 24}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[5]) // perfect fourth (protected)
	assert.Equal(t, JustInterval{numerator: 31, denominator: 22}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[7]) // perfect fifth (protected)
	assert.Equal(t, JustInterval{numerator: 31, denominator: 19}, intervals[8])
	assert.Equal(t, JustInterval{numerator: 31, denominator: 18}, intervals[9])
	assert.Equal(t, JustInterval{numerator: 31, denominator: 17}, intervals[10])
	assert.Equal(t, JustInterval{numerator: 31, denominator: 16}, intervals[11])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[12])
}

func Test_buildAFiveLimitJustChromaticScaleWithLesserMinorSeventhBasedOnPureRatios(t *testing.T) {
	// Given
	// Mike read the page at https://en.wikipedia.org/wiki/Five-limit_tuning and therefore
	filterOutLesserMajorSecondAndGreaterMinorSeventh := func(interval JustInterval) bool {
		if interval.IsLesserMajorSecond() || interval.IsGreaterMinorSeventh() {
			return true
		}
		return false
	}

	// When
	scale := NewJustIntonationChromaticScaleWithLimitAndFilter(5, filterOutLesserMajorSecondAndGreaterMinorSeventh)
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "5-limit Just Intonation", scale.System())
	assert.Equal(t, "Just Intonation chromatic scale based on 5-limit pure ratios.", scale.Description())

	assert.Equal(t, 13, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 15}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 8}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 6, denominator: 5}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 4}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 45, denominator: 32}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[7])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 5}, intervals[8])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 3}, intervals[9])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 9}, intervals[10])
	assert.Equal(t, JustInterval{numerator: 15, denominator: 8}, intervals[11])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[12])
}

func Test_buildAFiveLimitJustChromaticScaleWithLesserMajorSecondBasedOnPureRatios(t *testing.T) {
	// Given
	// Mike read the page at https://en.wikipedia.org/wiki/Five-limit_tuning and
	filterOutGreaterMajorSecondAndLesserMinorSeventh := func(interval JustInterval) bool {
		if interval.IsGreaterMajorSecond() || interval.IsLesserMinorSeventh() {
			return true
		}
		return false
	}

	// When
	scale := NewJustIntonationChromaticScaleWithLimitAndFilter(5, filterOutGreaterMajorSecondAndLesserMinorSeventh)
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "5-limit Just Intonation", scale.System())
	assert.Equal(t, "Just Intonation chromatic scale based on 5-limit pure ratios.", scale.Description())

	assert.Equal(t, 13, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 15}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 10, denominator: 9}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 6, denominator: 5}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 4}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 45, denominator: 32}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[7])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 5}, intervals[8])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 3}, intervals[9])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 5}, intervals[10])
	assert.Equal(t, JustInterval{numerator: 15, denominator: 8}, intervals[11])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[12])
}

func Test_shouldReturnBespokeJustScaleBasedOnProvidedIntervals(t *testing.T) {
	// Given
	// Mike read the page at https://en.wikipedia.org/wiki/Five-limit_tuning

	// When
	scale := NewJustIntonationChromaticScaleWith("Bespoke scale based on provided ratios", [][]uint{{1, 1}, {14, 13}, {3, 2}, {16, 9}, {2, 1}})
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "Just Intonation", scale.System())
	assert.Equal(t, "Bespoke scale based on provided ratios", scale.Description())

	assert.Equal(t, 5, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 14, denominator: 13}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 9}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[4])
}

func Test_ShouldReturnIntenseDiatonicScaleInLydianModeWithScaleDegrees(t *testing.T) {
	// Given
	scale := NewIntenseDiatonicScale(LydianMode)

	// When
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "Ptolemy Intense Diatonic", scale.System())
	assert.Equal(t, "Ptolemy's 5-limit intense diatonic scale in Lydian mode.", scale.Description())
	assert.Equal(t, 8, len(intervals))

	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 8}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 4}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 45, denominator: 32}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 3}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 15, denominator: 8}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[7])
}

func Test_ShouldReturnIntenseDiatonicScaleInIonianModeWithScaleDegrees(t *testing.T) {
	// Given
	scale := NewIntenseDiatonicScale(IonianMode)

	// When
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "Ptolemy Intense Diatonic", scale.System())
	assert.Equal(t, "Ptolemy's 5-limit intense diatonic scale in Ionian mode.", scale.Description())
	assert.Equal(t, 8, len(intervals))

	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 8}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 4}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 3}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 15, denominator: 8}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[7])
}

func Test_ShouldReturnIntenseDiatonicScaleInMixolydianModeWithScaleDegrees(t *testing.T) {
	// Given
	scale := NewIntenseDiatonicScale(MixolydianMode)

	// When
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "Ptolemy Intense Diatonic", scale.System())
	assert.Equal(t, "Ptolemy's 5-limit intense diatonic scale in Mixolydian mode.", scale.Description())
	assert.Equal(t, 8, len(intervals))

	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 8}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 4}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 3}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 9}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[7])
}

func Test_ShouldReturnIntenseDiatonicScaleInDorianModeWithScaleDegrees(t *testing.T) {
	// Given
	scale := NewIntenseDiatonicScale(DorianMode)

	// When
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "Ptolemy Intense Diatonic", scale.System())
	assert.Equal(t, "Ptolemy's 5-limit intense diatonic scale in Dorian mode.", scale.Description())
	assert.Equal(t, 8, len(intervals))

	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 8}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 6, denominator: 5}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 3}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 9}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[7])
}

func Test_ShouldReturnIntenseDiatonicScaleInAeolianModeWithScaleDegrees(t *testing.T) {
	// Given
	scale := NewIntenseDiatonicScale(AeolianMode)

	// When
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "Ptolemy Intense Diatonic", scale.System())
	assert.Equal(t, "Ptolemy's 5-limit intense diatonic scale in Aeolian mode.", scale.Description())
	assert.Equal(t, 8, len(intervals))

	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 8}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 6, denominator: 5}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 5}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 5}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[7])
}

func Test_ShouldReturnIntenseDiatonicScaleInPhrygianModeWithScaleDegrees(t *testing.T) {
	// Given
	scale := NewIntenseDiatonicScale(PhrygianMode)

	// When
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "Ptolemy Intense Diatonic", scale.System())
	assert.Equal(t, "Ptolemy's 5-limit intense diatonic scale in Phrygian mode.", scale.Description())
	assert.Equal(t, 8, len(intervals))

	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 15}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 6, denominator: 5}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 5}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 5}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[7])
}

func Test_ShouldReturnIntenseDiatonicScaleInLocrianModeWithScaleDegrees(t *testing.T) {
	// Given
	scale := NewIntenseDiatonicScale(LocrianMode)

	// When
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "Ptolemy Intense Diatonic", scale.System())
	assert.Equal(t, "Ptolemy's 5-limit intense diatonic scale in Locrian mode.", scale.Description())
	assert.Equal(t, 8, len(intervals))

	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 15}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 6, denominator: 5}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 64, denominator: 45}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 5}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 9}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[7])
}

func Test_shouldReturnPartch43ToneScale(t *testing.T) {
	// Harry Partch (1901–1974) described this scale in Genesis of a Music (1949/1974).
	// It is an 11-limit 43-tone scale built from otonalities and utonalities on the
	// 11-limit tonality diamond, producing 43 distinct pitches per octave (44 entries
	// including the octave 2/1).

	// When
	scale := NewPartch43ToneScale()
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "Partch 43-tone", scale.System())
	assert.Equal(t, "Harry Partch's 43-tone 11-limit just intonation scale from Genesis of a Music.", scale.Description())
	assert.Equal(t, 44, len(intervals))

	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 81, denominator: 80}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 33, denominator: 32}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 21, denominator: 20}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 15}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 12, denominator: 11}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 11, denominator: 10}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 10, denominator: 9}, intervals[7])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 8}, intervals[8])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 7}, intervals[9])
	assert.Equal(t, JustInterval{numerator: 7, denominator: 6}, intervals[10])
	assert.Equal(t, JustInterval{numerator: 32, denominator: 27}, intervals[11])
	assert.Equal(t, JustInterval{numerator: 6, denominator: 5}, intervals[12])
	assert.Equal(t, JustInterval{numerator: 11, denominator: 9}, intervals[13])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 4}, intervals[14])
	assert.Equal(t, JustInterval{numerator: 14, denominator: 11}, intervals[15])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 7}, intervals[16])
	assert.Equal(t, JustInterval{numerator: 21, denominator: 16}, intervals[17])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[18])
	assert.Equal(t, JustInterval{numerator: 27, denominator: 20}, intervals[19])
	assert.Equal(t, JustInterval{numerator: 11, denominator: 8}, intervals[20])
	assert.Equal(t, JustInterval{numerator: 7, denominator: 5}, intervals[21])
	assert.Equal(t, JustInterval{numerator: 10, denominator: 7}, intervals[22])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 11}, intervals[23])
	assert.Equal(t, JustInterval{numerator: 40, denominator: 27}, intervals[24])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[25])
	assert.Equal(t, JustInterval{numerator: 32, denominator: 21}, intervals[26])
	assert.Equal(t, JustInterval{numerator: 14, denominator: 9}, intervals[27])
	assert.Equal(t, JustInterval{numerator: 11, denominator: 7}, intervals[28])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 5}, intervals[29])
	assert.Equal(t, JustInterval{numerator: 18, denominator: 11}, intervals[30])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 3}, intervals[31])
	assert.Equal(t, JustInterval{numerator: 27, denominator: 16}, intervals[32])
	assert.Equal(t, JustInterval{numerator: 12, denominator: 7}, intervals[33])
	assert.Equal(t, JustInterval{numerator: 7, denominator: 4}, intervals[34])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 9}, intervals[35])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 5}, intervals[36])
	assert.Equal(t, JustInterval{numerator: 20, denominator: 11}, intervals[37])
	assert.Equal(t, JustInterval{numerator: 11, denominator: 6}, intervals[38])
	assert.Equal(t, JustInterval{numerator: 15, denominator: 8}, intervals[39])
	assert.Equal(t, JustInterval{numerator: 40, denominator: 21}, intervals[40])
	assert.Equal(t, JustInterval{numerator: 64, denominator: 33}, intervals[41])
	assert.Equal(t, JustInterval{numerator: 160, denominator: 81}, intervals[42])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[43])
}

func Test_ShouldReturnSazScaleWithExpectedScaleDegrees(t *testing.T) {
	// Given
	scale := NewSazScale()

	// When
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "Saz", scale.System())
	assert.Equal(t, "Turkish Saz tuning ratios.", scale.Description())
	assert.Equal(t, 18, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 18, denominator: 17}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 12, denominator: 11}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 9, denominator: 8}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 81, denominator: 68}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 27, denominator: 22}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 81, denominator: 64}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[7])
	assert.Equal(t, JustInterval{numerator: 24, denominator: 17}, intervals[8])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 11}, intervals[9])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[10])
	assert.Equal(t, JustInterval{numerator: 27, denominator: 17}, intervals[11])
	assert.Equal(t, JustInterval{numerator: 18, denominator: 11}, intervals[12])
	assert.Equal(t, JustInterval{numerator: 27, denominator: 16}, intervals[13])
	assert.Equal(t, JustInterval{numerator: 16, denominator: 9}, intervals[14])
	assert.Equal(t, JustInterval{numerator: 32, denominator: 17}, intervals[15])
	assert.Equal(t, JustInterval{numerator: 64, denominator: 33}, intervals[16])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[17])
}

func Test_isPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return false for n less than or equal to 1",
			args: args{n: 1},
			want: false,
		},
		{
			name: "should return true for a prime number",
			args: args{n: 13},
			want: true,
		},
		{
			name: "should return false for a non-prime number",
			args: args{n: 15},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isPrime(tt.args.n), "isPrime(%v)", tt.args.n)
		})
	}
}

func Test_convertJustScaleToRelativeFrequencies(t *testing.T) {
	// Given
	scale := NewSazScale()

	// When
	frequencies := scale.ToFrequenciesForTonicOf(256.00, 2)

	// Then
	assert.Equal(t, []float64{256, 271, 279, 288, 305, 314, 324, 341, 361, 372, 384, 407, 419, 432, 455, 482, 496, 512, 542, 559, 576, 610, 628, 648, 683, 723, 745, 768, 813, 838, 864, 910, 964, 993, 1024}, frequencies)
}

func Test_ShouldReturnHarmonicSeriesScaleForFirstEightPartials(t *testing.T) {
	// Given
	scale := NewHarmonicSeriesScale(8)

	// When
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "Harmonic Series", scale.System())
	assert.Equal(t, 5, len(intervals))
	assert.Equal(t, Unison(), intervals[0])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 4}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 7, denominator: 4}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 2, denominator: 1}, intervals[4])
}

func Test_ASchismaIsTheDifferenceBetwixtAPythagoreanCommaAndASyntonicComma(t *testing.T) {
	// Given
	s := Schisma()
	pc := PythagoreanComma()
	sc := SyntonicComma()

	// Then
	assert.Equal(t, uint(32805), s.Numerator())
	assert.Equal(t, uint(32768), s.Denominator())
	assert.Equal(t, uint(531441), pc.Numerator())
	assert.Equal(t, uint(524288), pc.Denominator())
	assert.Equal(t, uint(81), sc.Numerator())
	assert.Equal(t, uint(80), sc.Denominator())
	assert.Equal(t, s, PythagoreanComma().Subtract(SyntonicComma()))
}
