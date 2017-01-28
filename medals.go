package main

import (
	"math"
	"strings"
)

const medalsfilename = "./medals.json"
const chips = 1000
const boosted = 1000
const defaultNumResults = 7

type Direction string

const (
	Reversed = Direction("reversed")
	Upright  = Direction("upright")
)

type Type string

const (
	Magic = Type("magic")
	Power = Type("power")
	Speed = Type("speed")

	General = Type("general")
)

type Targetting string

const (
	AoE    = Targetting("AoE")
	Single = Targetting("Single")
	Random = Targetting("Random")
)

type Medal struct {
	BaseAtk  float64
	Boosted  bool `json:",omitempty"`
	Dir      Direction
	Guilt    float64
	HighMult float64 `json:",omitempty"`
	LowMult  float64 `json:",omitempty"`
	Mult     float64 `json:",omitempty"`
	Name     string
	SpCost   int        `json:"cost,omitempty"`
	Target   Targetting `json:",omitempty"`
	Type     Type

	CalculatedHighDmg float64 `json:",omitempty"`
	CalculatedLowDmg  float64 `json:",omitempty"`

	SlotApplied bool
}

type PrintableMedal struct {
	Name    string
	HighDmg float64
	LowDmg  float64
	Target  Targetting
	Cost    int
}

type byHighDmg []Medal

func (a byHighDmg) Len() int           { return len(a) }
func (a byHighDmg) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byHighDmg) Less(i, j int) bool { return a[i].CalculatedHighDmg > a[j].CalculatedHighDmg }

func (m *Medal) Clean() {
	if m.Mult != 0 {
		m.LowMult = m.Mult
		m.HighMult = m.Mult
	}
}

func (m *Medal) Calculate() {
	if m.CalculatedHighDmg != 0 {
		return
	}
	dmg := m.BaseAtk + chips
	if m.Boosted {
		dmg = dmg + boosted
	}

	dmg = dmg * (m.Guilt + 1)

	m.CalculatedHighDmg = dmg * m.HighMult
	m.CalculatedLowDmg = dmg * m.LowMult
}

func (m *Medal) ApplySlot(s *Slot) {
	if m.SlotApplied {
		return
	}
	if m.CalculatedHighDmg == 0 {
		m.Calculate()
	}
	if m.Type != s.Type {
		return
	}
	mult := s.TypeMult
	if m.Dir == s.Dir {
		mult = s.DirMult
	}
	m.applyMultiplier(mult)
	m.SlotApplied = true
}

func (m *Medal) ApplyBuffStatus(bs BuffStatus) {
	buffPercentage := 1.0

	netGeneral := bs[General].NetMultiplier()
	if netGeneral != 0 {
		buffPercentage = buffPercentage * netGeneral
	}

	netTyped := bs[m.Type].NetMultiplier()
	if netTyped != 0 {
		buffPercentage = buffPercentage * netTyped
	}

	m.applyMultiplier(buffPercentage)
}

// Medal Type, Enemy Type
var typeAdvantageMap = map[Type]map[Type]float64{
	Power: {
		Power: 1,
		Speed: 2,
		Magic: 0.5,
	},
	Speed: {
		Power: 0.5,
		Speed: 1,
		Magic: 2,
	},
	Magic: {
		Power: 2,
		Speed: 0.5,
		Magic: 1,
	},
}

func (m *Medal) ApplyEnemyTypeAdvantage(enemyType Type) {
	advMult := typeAdvantageMap[m.Type][enemyType]

	// WoFF medals ignore enemy type.
	if strings.HasPrefix(m.Name, "WoFF") {
		advMult = 2.0
	}

	m.applyMultiplier(advMult)
}

func (m *Medal) applyMultiplier(mult float64) {
	m.CalculatedHighDmg = m.CalculatedHighDmg * mult
	m.CalculatedLowDmg = m.CalculatedLowDmg * mult
}

func (m *Medal) GetPrintable() PrintableMedal {
	return PrintableMedal{
		Name:    m.Name,
		HighDmg: math.Floor(m.CalculatedHighDmg),
		LowDmg:  math.Floor(m.CalculatedLowDmg),
		Target:  m.Target,
		Cost:    m.SpCost,
	}
}
