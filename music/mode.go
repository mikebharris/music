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

func (m MusicalMode) IsDiatonic() bool {
	switch m {
	case LydianMode, IonianMode, MixolydianMode, DorianMode, AeolianMode, PhrygianMode, LocrianMode:
		return true
	default:
		return false
	}
}
func (m MusicalMode) String() string {
	return string(m)
}
