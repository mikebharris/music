package music

const (
	LydianMode     MusicalMode = "Lydian"
	IonianMode     MusicalMode = "Ionian"
	MixolydianMode MusicalMode = "Mixolydian"
	DorianMode     MusicalMode = "Dorian"
	AeolianMode    MusicalMode = "Aeolian"
	PhrygianMode   MusicalMode = "Phrygian"
	LocrianMode    MusicalMode = "Locrian"
)

type MusicalMode string

var modeBrightness = map[MusicalMode]int{
	LydianMode:     7,
	IonianMode:     6,
	MixolydianMode: 5,
	DorianMode:     4,
	AeolianMode:    3,
	PhrygianMode:   2,
	LocrianMode:    1,
}

func (m MusicalMode) IsDiatonic() bool {
	switch m {
	case LydianMode, IonianMode, MixolydianMode, DorianMode, AeolianMode, PhrygianMode, LocrianMode:
		return true
	default:
		return false
	}
}

func (m MusicalMode) BrightnessOrder() int {
	return modeBrightness[m]
}

func (m MusicalMode) IsBrighterThan(other MusicalMode) bool {
	return m.BrightnessOrder() > other.BrightnessOrder()
}

func (m MusicalMode) String() string {
	return string(m)
}
