package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
)

var (
	slotNum    = flag.Int("slot", 0, "")
	keyblade   = flag.String("kb", "sl", "")
	enemyType  = flag.String("enemy", "", "")
	numResults = flag.Int("num", defaultNumResults, "")

	kba = flag.Bool("kba", false, "")
)

func main() {
	flag.Parse()

	if *kba {
		runKbAnalysis()
	} else {
		runMedalAnalysis()
	}

}
func runKbAnalysis() {
	for k, v := range kbs {
		v.Power()
		v.name = k
		kbs[k] = v
	}

	blades := []KeyBlade{}
	GetVals(kbs, &blades)

	sort.Sort(byPower(blades))

	for _, k := range blades {
		k.Print()
	}
}

func runMedalAnalysis() {
	data, err := ioutil.ReadFile(medalsfilename)
	FatalOnErr(err)

	medals := []Medal{}
	err = json.Unmarshal(data, &medals)
	FatalOnErr(err)

	kb := kbs[*keyblade]
	kb.Buffs = NewBuffStatus()

	// Apply iKairi to everything.
	kb.Buffs.ApplyBuff(&Buff{
		Type:      General,
		AtkDef:    Attack,
		NumLevels: 3,
	})
	s := kb.Slots[*slotNum-1]

	for i, m := range medals {
		m.Clean()
		m.Calculate()
		m.ApplySlot(&s)
		m.ApplyBuffStatus(kb.Buffs)

		if *enemyType != "" {
			m.ApplyEnemyTypeAdvantage(Type(*enemyType))
		}
		medals[i] = m
	}

	sort.Sort(byHighDmg(medals))

	results := make([]PrintableMedal, 0, *numResults)
	for i, m := range medals {
		if i >= *numResults {
			break
		}
		results = append(results, m.GetPrintable())
		fmt.Printf("%v\n", m.GetPrintable())
	}
}
