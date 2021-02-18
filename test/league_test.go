package test

import (
	TF "TwistedFate"
	"fmt"
	"testing"
)

func TestGetLeagueEntryByAccountId(t *testing.T) {
	s, _ := TF.GetLeagueEntryBySummonerId("97Vtefw0C2x_UtDXe94THK1kwGivO8ZhBZsDmXxIQXS676I")
	fmt.Println((*s)[0])
	if (*s)[0].SummonerName != "Gwal66" {
		t.Errorf("Should be user Gwal66 but got %s", (*s)[0].SummonerName)
	}
}