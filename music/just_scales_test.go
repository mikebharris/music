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

func Test_ShouldReturn5LimitScaleBuildFromPythagoreanWithExpectedScaleDegrees(t *testing.T) {
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

func Test_shouldReturnThirteenLimitJustChromaticScaleBasedOnPureRatios(t *testing.T) {
	// Given
	// Mike read the page at https://en.wikipedia.org/wiki/Five-limit_tuning

	// When
	scale := NewJustIntonationChromaticScaleWithLimit(13)
	intervals := scale.Intervals()

	// Then
	assert.Equal(t, "13-limit Just Intonation", scale.System())
	assert.Equal(t, "Just Intonation chromatic scale based on 13-limit pure ratios.", scale.Description())

	assert.Equal(t, 13, len(intervals))
	assert.Equal(t, JustInterval{numerator: 1, denominator: 1}, intervals[0])
	assert.Equal(t, JustInterval{numerator: 13, denominator: 12}, intervals[1])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 7}, intervals[2])
	assert.Equal(t, JustInterval{numerator: 6, denominator: 5}, intervals[3])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 4}, intervals[4])
	assert.Equal(t, JustInterval{numerator: 4, denominator: 3}, intervals[5])
	assert.Equal(t, JustInterval{numerator: 7, denominator: 5}, intervals[6])
	assert.Equal(t, JustInterval{numerator: 3, denominator: 2}, intervals[7])
	assert.Equal(t, JustInterval{numerator: 8, denominator: 5}, intervals[8])
	assert.Equal(t, JustInterval{numerator: 5, denominator: 3}, intervals[9])
	assert.Equal(t, JustInterval{numerator: 7, denominator: 4}, intervals[10])
	assert.Equal(t, JustInterval{numerator: 13, denominator: 7}, intervals[11])
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
