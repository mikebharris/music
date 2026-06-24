package scala

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mikebharris/music/music"
)

func NewScalaScaleFileFromScale(filename string, scale music.Scale) string {
	switch s := scale.(type) {
	case music.JustScale:
		return newScalaScaleFileFromJustScale(filename, s)
	case music.TemperedScale:
		return newScalaScaleFileFromTemperedScale(filename, s)
	default:
		return ""
	}
}

func newScalaScaleFileFromJustScale(filename string, scale music.JustScale) string {
	numberOfIntervalsInFile := 0
	intervalsContent := ""
	for _, interval := range scale.Intervals() {
		if interval.IsEqualTo(music.Unison()) {
			continue
		}
		intervalsContent += fmt.Sprintf("%d/%d\n", interval.Numerator(), interval.Denominator())
		numberOfIntervalsInFile++
	}
	return scaleFileHeader(filename, scale.System(), scale.Description(), numberOfIntervalsInFile) + intervalsContent
}

func newScalaScaleFileFromTemperedScale(filename string, scale music.TemperedScale) string {
	numberOfIntervalsInFile := 0
	intervalsContent := ""
	for _, interval := range scale.Intervals() {
		if interval == 1.0 {
			continue
		}
		intervalsContent += fmt.Sprintf("%.2f\n", interval.ToCents())
		numberOfIntervalsInFile++
	}
	return scaleFileHeader(filename, scale.System(), scale.Description(), numberOfIntervalsInFile) + intervalsContent
}

func scaleFileHeader(filename string, system string, description string, numberOfIntervalsInFile int) string {
	contents := "! " + filename + "\n"
	contents += "! generated using github.com/mikebharris/music/scala\n"
	contents += "!\n"
	contents += fmt.Sprintf("%s scale using %s\n", system, description)
	contents += fmt.Sprintf("%d\n", numberOfIntervalsInFile)
	return contents
}

func NewJustScaleFromScalaFile(scalaFile string) (music.JustScale, error) {
	if scalaFile == "" {
		return music.JustScale{}, errors.New("the Scala file is empty, nothing to convert")
	}
	fileContents := strings.Split(scalaFile, "\n")
	var numberOfIntervals = 0
	var intervals [][]uint
	for _, line := range fileContents {
		if lineIsCommentOrEmpty(line) {
			continue
		}
		if !lineIsJustInterval(line) && numberOfIntervals == 0 { // the first line we encounter like this we treat as the count of intervals specified in the file
			numberOfIntervals, _ = strconv.Atoi(line)
			continue
		}
		if lineIsJustInterval(line) {
			intervals = append(intervals, stringRatioToNumeratorAndDenominatorTuple(line))
		}
	}
	if len(intervals) == 0 && numberOfIntervals == 0 {
		return music.JustScale{}, errors.New("the Scala file does not contain any intervals")
	}
	if len(intervals) != numberOfIntervals {
		return music.JustScale{}, fmt.Errorf("number of intervals specified in Scala file does not number of intervals collected: %d vs %d", numberOfIntervals, len(intervals))
	}

	return music.NewJustIntonationChromaticScaleWith("", intervals), nil
}

func lineIsJustInterval(line string) bool {
	return strings.Contains(line, "/")
}

func lineIsCommentOrEmpty(line string) bool {
	return strings.HasPrefix(line, "!") || line == ""
}

func stringRatioToNumeratorAndDenominatorTuple(ratio string) []uint {
	v := strings.Split(ratio, "/")
	numerator, _ := strconv.Atoi(v[0])
	denominator, _ := strconv.Atoi(v[1])
	return []uint{uint(numerator), uint(denominator)}
}
