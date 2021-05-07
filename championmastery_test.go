package TwistedFate

import (
	"fmt"
	"testing"
)

func TestGetChampionMasteryBySummonerId(t *testing.T) {
	x, _ := GetChampionMasteryBySummonerId("7sn3aFF8iNezYkY19rGOCQUdZ1ZG0nMkg_NCueNuwrD3n24")
	for _, v := range x[:10] {
		fmt.Println(v)
	}
}

func TestGetChampionMasteryScoreBySummonerId(t *testing.T) {
	x, _ := GetChampionMasteryScoreBySummonerId("7sn3aFF8iNezYkY19rGOCQUdZ1ZG0nMkg_NCueNuwrD3n24")
	fmt.Println(x)
}

func TestGetChampionRotation(t *testing.T) {
	x, _ := GetChampionRotation()
	fmt.Println(x.FreeChampionIds)
}

func TestGetChampMap(t *testing.T) {
	GetChampMap()
}