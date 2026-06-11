package music

// Interval is a type constraint satisfied by both JustInterval and TemperedInterval,
// used by generic functions that operate on either kind of interval.
type Interval interface {
	JustInterval | TemperedInterval
}

// Scale is implemented by both JustScale and TemperedScale.
type Scale interface {
	System() string
	Description() string
	IntervalStrings() []string
}

// NoteName represents a Western note letter name (A–H, where H is used in German notation).
// These constants are defined for future use in note-naming functionality.
type NoteName rune

const (
	A NoteName = 'A'
	B NoteName = 'B'
	C NoteName = 'C'
	D NoteName = 'D'
	E NoteName = 'E'
	F NoteName = 'F'
	G NoteName = 'G'
	H NoteName = 'H' // H is used in German/Central European notation (= B♮ in English)
)
