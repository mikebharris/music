# A music theory module for working with musical intervals, tuning systems and scales, both just and tempered, written in Go.

This Go module provides types and functions to represent and manipulate musical notes, intervals, and scales. It supports various tuning systems including just intonation, meantone temperament, and equal temperament.

## Features
- Representation of musical intervals for just intonation pure ratios and irrational tempered intervals
- Support for different tuning systems:
  - Just Intonation
  - Meantone Temperament
  - Bach's Well-Tempered Clavier
  - Equal Temperament
- Support for different just scales:
  - Pythagorean Scale
  - Turkish Saz Scale
  - Ptolemaic Scale
  - X-Limit Chromatic Scale
  - Harry Partch's 43-tone "Genesis" Scale
  - Derived from the Harmonic Series
  - User-defined scales
- Functions for manipulating intervals
  - Addition and subtraction of intervals
  - Greater than and less than comparisons
  - Conversion to cents
  - Octave reduction and lowest common denominator simplification

## Examples

Create a new scale for a Turkish Saz Tuning and print out its intervals:

```go
scale := NewSazScale()
intervals := scale.Intervals()

for i, interval := range intervals {
    fmt.Printf("Note %d has ratio of %d:%d, which in cents is %.2f cents\n", i, interval.Numerator(), interval.Denominator(), interval.ToCents())
}
```

Outputs:

```
Note 0 has ratio of 1:1, which in cents is 0.00 cents
Note 1 has ratio of 18:17, which in cents is 98.95 cents
Note 2 has ratio of 12:11, which in cents is 150.64 cents
Note 3 has ratio of 9:8, which in cents is 203.91 cents
Note 4 has ratio of 81:68, which in cents is 302.86 cents
Note 5 has ratio of 27:22, which in cents is 354.55 cents
Note 6 has ratio of 81:64, which in cents is 407.82 cents
Note 7 has ratio of 4:3, which in cents is 498.04 cents
Note 8 has ratio of 24:17, which in cents is 597.00 cents
Note 9 has ratio of 16:11, which in cents is 648.68 cents
Note 10 has ratio of 3:2, which in cents is 701.96 cents
Note 11 has ratio of 27:17, which in cents is 800.91 cents
Note 12 has ratio of 18:11, which in cents is 852.59 cents
Note 13 has ratio of 27:16, which in cents is 905.87 cents
Note 14 has ratio of 16:9, which in cents is 996.09 cents
Note 15 has ratio of 32:17, which in cents is 1095.04 cents
Note 16 has ratio of 64:33, which in cents is 1146.73 cents
Note 17 has ratio of 2:1, which in cents is 1200.00 cents
```

Subtract a synthonic comma from a Pythagorean major second to get a just major second:

```go
i := JustInterval{numerator: 9, denominator: 8}
j := SyntonicComma()
fmt.Println(i.Subtract(j))
```
Outputs:

```
10:9
```

Create a bespoke just scale:

```go
	scale := NewJustIntonationChromaticScaleWith("Bespoke scale based on provided ratios", [][]uint{{1, 1}, {14, 13}, {3, 2}, {16, 9}, {2, 1}})
	intervals := scale.Intervals()
    for i, interval := range intervals {
        fmt.Printf("Note %d has ratio of %d:%d, which in cents is %.2f cents\n", i, interval.Numerator(), interval.Denominator(), interval.ToCents())
    }
```
Outputs:

```
Note 0 has ratio of 1:1, which in cents is 0.00 cents
Note 1 has ratio of 14:13, which in cents is 128.30 cents
Note 2 has ratio of 3:2, which in cents is 701.96 cents
Note 3 has ratio of 16:9, which in cents is 996.09 cents
Note 4 has ratio of 2:1, which in cents is 1200.00 cents
```