package main

type AtkDef string

const (
	Attack  = AtkDef("attack")
	Defense = AtkDef("defense")
)

type BuffStatus map[Type]BuffTypeStatus

type BuffTypeStatus map[AtkDef]int

func NewBuffStatus() BuffStatus {
	return map[Type]BuffTypeStatus{
		General: BuffTypeStatus{},
		Power:   BuffTypeStatus{},
		Speed:   BuffTypeStatus{},
		Magic:   BuffTypeStatus{},
	}
}

func (b BuffTypeStatus) NetLevels() int {
	return b[Attack] + b[Defense]
}
func (b BuffTypeStatus) NetMultiplier() float64 {
	return 1.0 + (0.2 * float64(b.NetLevels()))
}

type Buff struct {
	Type      Type
	AtkDef    AtkDef
	NumLevels int
}

func (bs BuffStatus) ApplyBuff(b *Buff) {
	currentBuf := bs[b.Type][b.AtkDef]
	currentBuf = currentBuf + b.NumLevels
	if currentBuf >= 3 {
		currentBuf = 3
	}

	bs[b.Type][b.AtkDef] = currentBuf
}
