package tInterface

import (
	"fmt"
	"sort"
)

type Hero struct {
	Name                string
	Age                 int
	CombatEffectiveness int //战斗力
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].CombatEffectiveness < hs[j].CombatEffectiveness;
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func TSort() {
	fmt.Println("====")
	var heroList HeroSlice= []Hero{
		{Name: "男刀", Age: 3, CombatEffectiveness: 300},
		{Name: "盖伦", Age: 5, CombatEffectiveness: 100},
		{Name: "女枪", Age: 1, CombatEffectiveness: 400},
		{Name: "火男", Age: 2, CombatEffectiveness: 200},
	}
	heroList = append(heroList, Hero{Name: "猫咪", Age: 6, CombatEffectiveness: 30})
	fmt.Printf("%T,%v\n", heroList,heroList)
	sort.Sort(heroList)
	fmt.Printf("%T,%v\n", heroList,heroList)
}