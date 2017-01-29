package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"sort"

	"./khux"
)

var (
	slotNum    = flag.Int("slot", 1, "")
	keyblade   = flag.String("kb", "sl", "")
	enemyType  = flag.String("enemy", "", "")
	numResults = flag.Int("num", khux.DefaultNumResults, "")

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
	for k, v := range khux.Kbs {
		v.Power()
		v.Name = k
		khux.Kbs[k] = v
	}

	blades := []khux.KeyBlade{}
	khux.GetVals(khux.Kbs, &blades)

	sort.Sort(khux.ByPower(blades))

	for _, k := range blades {
		k.Print()
	}
}

func runMedalAnalysis() {
	data, err := ioutil.ReadFile(khux.Medalsfilename)
	khux.FatalOnErr(err)

	medals := []khux.Medal{}
	err = json.Unmarshal(data, &medals)
	khux.FatalOnErr(err)

	kb := khux.Kbs[*keyblade]
	kb.Buffs = khux.NewBuffStatus()

	// Apply iKairi to everything.
	kb.Buffs.ApplyBuff(&khux.Buff{
		Type:      khux.General,
		AtkDef:    khux.Attack,
		NumLevels: 3,
	})
	s := kb.Slots[*slotNum-1]

	for i, m := range medals {
		m.Clean()
		m.Calculate()
		m.ApplySlot(&s)
		m.ApplyBuffStatus(kb.Buffs)

		if *enemyType != "" {
			m.ApplyEnemyTypeAdvantage(khux.Type(*enemyType))
		}
		medals[i] = m
	}

	sort.Sort(khux.ByHighDmg(medals))

	results := make([]khux.PrintableMedal, 0, *numResults)
	for i, m := range medals {
		if i >= *numResults {
			break
		}
		results = append(results, m.GetPrintable())
		fmt.Printf("%v\n", m.GetPrintable())
	}
}
