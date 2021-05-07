package TwistedFate

import (
	"fmt"
	"testing"
)

func TestMatchlistGetByAccountId(t *testing.T) {
	s, _ := GetMatchlistByAccountId("szVA5tDUyU6RFarWzb9sW41A2CdE6TGitBMQN0LfK9vOFMU")
	for _, v := range s.Matches {
		fmt.Printf("%f ", v.GameId)
	}
	if s.TotalGames == 0 {
		t.Errorf("Total games should not be 0.")
	}
}