package test

import (
	TF "TwistedFate"
	"testing"
)

func TestMatchlistGetByAccountId(t *testing.T) {
	s, _ := TF.GetMatchlistByAccountId("szVA5tDUyU6RFarWzb9sW41A2CdE6TGitBMQN0LfK9vOFMU")
	if s.TotalGames == 0 {
		t.Errorf("Total games should not be 0.")
	}
}