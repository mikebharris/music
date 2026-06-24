package scala

import (
	"errors"
	"strings"
	"testing"

	"github.com/mikebharris/music/music"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldGenerateScalaFileForPythagorean3LimitScale(t *testing.T) {
	// Given
	scale := music.NewPythagoreanScale()

	// When
	scalaFile := NewScalaScaleFileFromScale("pythagorean-3-limit.scl", scale)

	// Then
	contents := strings.Split(scalaFile, "\n")
	assert.Equal(t, "! pythagorean-3-limit.scl", contents[0])
	assert.Equal(t, "! generated using github.com/mikebharris/music/scala", contents[1])
	assert.Equal(t, "!", contents[2])
	assert.Equal(t, "Pythagorean scale using 3-limit Pythagorean ratios.", contents[3])
	assert.Equal(t, "13", contents[4])
	assert.Equal(t, "256/243", contents[5])
	assert.Equal(t, "9/8", contents[6])
	assert.Equal(t, "32/27", contents[7])
	assert.Equal(t, "81/64", contents[8])
	assert.Equal(t, "4/3", contents[9])
	assert.Equal(t, "1024/729", contents[10])
	assert.Equal(t, "729/512", contents[11])
	assert.Equal(t, "3/2", contents[12])
	assert.Equal(t, "128/81", contents[13])
	assert.Equal(t, "27/16", contents[14])
	assert.Equal(t, "16/9", contents[15])
	assert.Equal(t, "243/128", contents[16])
	assert.Equal(t, "2/1", contents[17])
}

func Test_ShouldGeneratePythagorean3LimitScaleFromScalaFile(t *testing.T) {
	// Given
	scalaFile :=
		"! pythagorean-3-limit.scl\n" +
			"!\n" +
			"13\n" +
			"256/243\n" +
			"9/8\n" +
			"32/27\n" +
			"81/64\n" +
			"4/3\n" +
			"1024/729\n" +
			"729/512\n" +
			"3/2\n" +
			"128/81\n" +
			"27/16\n" +
			"16/9\n" +
			"243/128\n" +
			"2/1\n"

	// When
	scale, err := NewJustScaleFromScalaFile(scalaFile)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, 13, len(scale.Intervals()))
	assert.Equal(t, music.NewInterval(256, 243), scale.Intervals()[0])
	assert.Equal(t, music.NewInterval(9, 8), scale.Intervals()[1])
	assert.Equal(t, music.NewInterval(32, 27), scale.Intervals()[2])
	assert.Equal(t, music.NewInterval(2, 1), scale.Intervals()[12])
}

func Test_ShouldReturnErrorIfSpecifiedNumberOfIntervalsDoesNotMatchCountOfConvertedIntervals(t *testing.T) {
	// Given
	scalaFile :=
		"! pythagorean-3-limit-missing-notes.scl\n" +
			"!\n" +
			"13\n" +
			"2/1\n"

	//When
	scale, err := NewJustScaleFromScalaFile(scalaFile)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("number of intervals specified in Scala file does not number of intervals collected: 13 vs 1"), err)
	assert.Empty(t, scale)
}

func Test_ShouldReturnErrorFileDoesNotSpecifyNumberOfIntervals(t *testing.T) {
	// Given
	scalaFile :=
		"! pythagorean-3-limit-missing-number-of-intervals.scl\n" +
			"!\n" +
			"2/1\n"

	//When
	scale, err := NewJustScaleFromScalaFile(scalaFile)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("number of intervals specified in Scala file does not number of intervals collected: 0 vs 1"), err)
	assert.Empty(t, scale)
}

func Test_ShouldReturnErrorIfScalaFileIsEmpty(t *testing.T) {
	// Given
	scalaFile := ""

	//When
	scale, err := NewJustScaleFromScalaFile(scalaFile)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("the Scala file is empty, nothing to convert"), err)
	assert.Empty(t, scale)
}

func Test_ShouldReturnErrorIfScalaFileContainsTemperedIntervals(t *testing.T) {
	// Given
	scalaFile :=
		"! scala-file-with-tempered-intervals.scl\n" +
			"!\n" +
			"1\n" +
			"1200.0\n"

	//When
	scale, err := NewJustScaleFromScalaFile(scalaFile)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("number of intervals specified in Scala file does not number of intervals collected: 1 vs 0"), err)
	assert.Empty(t, scale)
}

func Test_ShouldReturnErrorIfScalaFileContainsNoIntervals(t *testing.T) {
	// Given
	scalaFile :=
		"! scala-file-with-only-comments.scl\n"

	//When
	scale, err := NewJustScaleFromScalaFile(scalaFile)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("the Scala file does not contain any intervals"), err)
	assert.Empty(t, scale)
}

func Test_ShouldGenerateScalaFileForTwelveToneEqualTemperamentScale(t *testing.T) {
	// Given
	scale := music.NewEqualTemperamentScale(12)

	// When
	scalaFile := NewScalaScaleFileFromScale("12-tone-equal-temperament.scl", scale)

	// Then
	contents := strings.Split(scalaFile, "\n")
	assert.Equal(t, "! 12-tone-equal-temperament.scl", contents[0])
	assert.Equal(t, "! generated using github.com/mikebharris/music/scala", contents[1])
	assert.Equal(t, "!", contents[2])
	assert.Equal(t, "Equal Temperament scale using 12-tone equal temperament.", contents[3])
	assert.Equal(t, "12", contents[4])
	assert.Equal(t, "100.00", contents[5])
	assert.Equal(t, "200.00", contents[6])
	assert.Equal(t, "300.00", contents[7])
	assert.Equal(t, "400.00", contents[8])
	assert.Equal(t, "500.00", contents[9])
	assert.Equal(t, "600.00", contents[10])
	assert.Equal(t, "700.00", contents[11])
	assert.Equal(t, "800.00", contents[12])
	assert.Equal(t, "900.00", contents[13])
	assert.Equal(t, "1000.00", contents[14])
	assert.Equal(t, "1100.00", contents[15])
	assert.Equal(t, "1200.00", contents[16])
}

type AnotherScale struct {
	system      string
	description string
}

func (s AnotherScale) System() string {
	panic("not implemented")
}

func (s AnotherScale) Description() string {
	panic("not implemented")
}

func (s AnotherScale) IntervalStrings() []string {
	panic("not implemented")
}

func Test_ReturnsEmptyScalaFileWhenUnsupportedScaleProvided(t *testing.T) {
	// Given
	scale := AnotherScale{}

	// When
	scalaFile := NewScalaScaleFileFromScale("not-important.scl", scale)

	// Then
	assert.Empty(t, scalaFile)
}
