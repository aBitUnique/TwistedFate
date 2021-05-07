package test

import (
	TF "TwistedFate"
	"testing"
)

func TestSummonerGetByName(t *testing.T) {
	s, _ := TF.GetSummonerByName("aBitUnique")
	if s.Name != "aBitUnique" {
		t.Errorf("Should have returned user aBitUnique, instead returned %s", s.Name)
	}
}

func TestSummonerGetByAccount(t *testing.T) {
	s, _ := TF.GetSummonerByAccountId("szVA5tDUyU6RFarWzb9sW41A2CdE6TGitBMQN0LfK9vOFMU")
	if s.Name != "aBitUnique" {
		t.Errorf("Should have returned user aBitUnique, instead returned %s", s.Name)
	}
}

func TestSummonerGetByPuuid(t *testing.T) {
	s, _ := TF.GetSummonerByPuuid("MGvEZ-y0RUzHpqE6CdeFex-FAmnmyfoZbK9CYc_dCjtXHn49MIVxmFcV_reolXNTWqzHnza5cJY80g")
	if s.Name != "aBitUnique" {
		t.Errorf("Should have returned user aBitUnique, instead returned %s", s.Name)
	}
}

func TestSummonerGetBySummonerId(t *testing.T) {
	s, _ := TF.GetSummonerBySummonerId("7sn3aFF8iNezYkY19rGOCQUdZ1ZG0nMkg_NCueNuwrD3n24")
	if s.Name != "aBitUnique" {
		t.Errorf("Should have returned user aBitUnique, instead returned %s", s.Name)
	}
}
