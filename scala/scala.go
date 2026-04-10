package scala

import (
	"fmt"

	"github.com/mikebharris/music/music"
)

func NewScalaScaleFileFromScale[T music.JustScale | music.TemperedScale](filename string, scale T) string {
	switch s := any(scale).(type) {
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
