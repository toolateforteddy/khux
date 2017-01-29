package khux

import "fmt"

const (
	Starlight     = "sl"
	TreasureTrove = "tt"
	LadyLuck      = "ll"
	ThreeWishes   = "tw"
	Olympia       = "ol"
	DivineRose    = "dr"
	MoogleOfGlory = "mog"
	Lionheart     = "lh"
)

type Slot struct {
	Type     Type
	Dir      Direction
	TypeMult float64
	DirMult  float64
}

type KeyBlade struct {
	Level int
	Slots []Slot

	CalculatedPower float64
	Buffs           BuffStatus

	Name string
}

var Kbs = map[string]KeyBlade{
	Starlight: {
		Level: 25,
		Slots: []Slot{
			{
				Type:     Magic,
				Dir:      Upright,
				TypeMult: 1.6,
				DirMult:  2.35,
			},
			{
				Type:     Speed,
				Dir:      Upright,
				TypeMult: 1.6,
				DirMult:  2.0,
			},
			{
				Type:     Power,
				Dir:      Reversed,
				TypeMult: 1.6,
				DirMult:  2.30,
			},
			{
				Type:     Speed,
				Dir:      Upright,
				TypeMult: 1.6,
				DirMult:  2.30,
			},
			{
				Type:     Magic,
				Dir:      Upright,
				TypeMult: 1.6,
				DirMult:  2.10,
			},
		}},
	TreasureTrove: {
		Level: 27,
		Slots: []Slot{
			{
				Type:     Power,
				Dir:      Upright,
				TypeMult: 1.4,
				DirMult:  1.9,
			},
			{
				Type:     Power,
				Dir:      Reversed,
				TypeMult: 1.3,
				DirMult:  1.9,
			},
			{
				Type:     Power,
				Dir:      Upright,
				TypeMult: 1.7,
				DirMult:  2.40,
			},
			{
				Type:     Power,
				Dir:      Upright,
				TypeMult: 1.5,
				DirMult:  2.00,
			},
			{
				Type:     Power,
				Dir:      Reversed,
				TypeMult: 1.8,
				DirMult:  2.40,
			},
		}},
	LadyLuck: {
		Level: 26,
		Slots: []Slot{
			{
				Type:     Speed,
				Dir:      Upright,
				TypeMult: 1.5,
				DirMult:  2.3,
			},
			{
				Type:     Speed,
				Dir:      Reversed,
				TypeMult: 1.7,
				DirMult:  1.9,
			},
			{
				Type:     Speed,
				Dir:      Upright,
				TypeMult: 1.3,
				DirMult:  1.9,
			},
			{
				Type:     Speed,
				Dir:      Upright,
				TypeMult: 1.3,
				DirMult:  1.90,
			},
			{
				Type:     Speed,
				Dir:      Reversed,
				TypeMult: 1.8,
				DirMult:  2.50,
			},
		}},
	ThreeWishes: {
		Level: 27,
		Slots: []Slot{
			{
				Type:     Magic,
				Dir:      Upright,
				TypeMult: 1.2,
				DirMult:  1.9,
			},
			{
				Type:     Magic,
				Dir:      Upright,
				TypeMult: 1.6,
				DirMult:  2.35,
			},
			{
				Type:     Magic,
				Dir:      Reversed,
				TypeMult: 1.5,
				DirMult:  2.0,
			},
			{
				Type:     Magic,
				Dir:      Upright,
				TypeMult: 1.8,
				DirMult:  2.00,
			},
			{
				Type:     Magic,
				Dir:      Reversed,
				TypeMult: 1.6,
				DirMult:  2.40,
			},
		}},
	Olympia: {
		Level: 30,
		Slots: []Slot{
			{
				Type:     Speed,
				Dir:      Reversed,
				TypeMult: 1.5,
				DirMult:  1.9,
			},
			{
				Type:     Power,
				Dir:      Reversed,
				TypeMult: 1.5,
				DirMult:  2.0,
			},
			{
				Type:     Speed,
				Dir:      Upright,
				TypeMult: 1.7,
				DirMult:  2.50,
			},
			{
				Type:     Power,
				Dir:      Upright,
				TypeMult: 2.0,
				DirMult:  2.80,
			},
			{
				Type:     Speed,
				Dir:      Upright,
				TypeMult: 1.6,
				DirMult:  2.10,
			},
		}},
	DivineRose: {
		Level: 26,
		Slots: []Slot{
			{
				Type:     Magic,
				Dir:      Upright,
				TypeMult: 1.5,
				DirMult:  1.9,
			},
			{
				Type:     Speed,
				Dir:      Reversed,
				TypeMult: 1.30,
				DirMult:  1.90,
			},
			{
				Type:     Magic,
				Dir:      Upright,
				TypeMult: 1.20,
				DirMult:  1.90,
			},
			{
				Type:     Magic,
				Dir:      Upright,
				TypeMult: 1.80,
				DirMult:  2.20,
			},
			{
				Type:     Speed,
				Dir:      Reversed,
				TypeMult: 2.10,
				DirMult:  2.70,
			},
		}},
	MoogleOfGlory: {
		Level: 27,
		Slots: []Slot{
			{
				Type:     Power,
				Dir:      Reversed,
				TypeMult: 1.70,
				DirMult:  1.90,
			},
			{
				Type:     Magic,
				Dir:      Upright,
				TypeMult: 1.20,
				DirMult:  2.00,
			},
			{
				Type:     Power,
				Dir:      Reversed,
				TypeMult: 1.30,
				DirMult:  2.10,
			},
			{
				Type:     Magic,
				Dir:      Reversed,
				TypeMult: 2.20,
				DirMult:  2.70,
			},
			{
				Type:     Power,
				Dir:      Upright,
				TypeMult: 1.6,
				DirMult:  2.00,
			},
		}},
	Lionheart: {
		Level: 20,
		Slots: []Slot{
			{
				Type:     Speed,
				Dir:      Upright,
				TypeMult: 1.1,
				DirMult:  1.8,
			},
			{
				Type:     Speed,
				Dir:      Upright,
				TypeMult: 1.1,
				DirMult:  1.9,
			},
			{
				Type:     Speed,
				Dir:      Reversed,
				TypeMult: 1.1,
				DirMult:  2.10,
			},
			{
				Type:     Speed,
				Dir:      Upright,
				TypeMult: 1.1,
				DirMult:  2.00,
			},
			{
				Type:     Speed,
				Dir:      Reversed,
				TypeMult: 1.1,
				DirMult:  2.60,
			},
		}},
}

var slotweights = []float64{
	0,
	1,
	2,
	5,
	5,
}

func (k *KeyBlade) Power() float64 {
	power := 0.0
	for i, slot := range k.Slots {
		power = power + slotweights[i]*slot.DirMult
	}
	k.CalculatedPower = power
	return power
}

func (k *KeyBlade) Print() {
	fmt.Printf("KB: %s@%v\n %v\n", k.Name, k.Level, k.CalculatedPower)
}

type ByPower []KeyBlade

func (a ByPower) Len() int           { return len(a) }
func (a ByPower) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPower) Less(i, j int) bool { return a[i].CalculatedPower > a[j].CalculatedPower }
