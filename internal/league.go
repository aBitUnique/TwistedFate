package internal

import (
	"encoding/json"
	"fmt"
)

type LeagueEntryDto struct {
	LeagueId string
	SummonerId string
	SummonerName string
	QueueType    string
	Tier         string
	Rank         string
	LeaguePoints int
	Wins         int
	Losses       int
	HotStreak    bool
	Veteran      bool
	FreshBlood   bool
	Inactive     bool
	MiniSeries   MiniSeriesDto
}

type MiniSeriesDto struct {
	Losses int
	Progress string
	Target int
	Wins int
}

func (ledto LeagueEntryDto) String() string {
	return fmt.Sprintf("LeagueId: %s\nSummonerId: %s\nSummonerName: %s\nQueueType: %s\nTier: %s\nRank: %s\nLeaguePoints: %d\nWins: %d\nLosses: %d\nHotStreak: %t\nVeteran: %t\nFreshBlood: %t\nInactive: %t\n",
	ledto.LeagueId,
	ledto.SummonerId,
	ledto.SummonerName,
	ledto.QueueType,
	ledto.Tier,
	ledto.Rank,
	ledto.LeaguePoints,
	ledto.Wins,
	ledto.Losses,
	ledto.HotStreak,
	ledto.Veteran,
	ledto.FreshBlood,
	ledto.Inactive)
}

func GetLeagueEntryBySummonerId(summonerId string) ([]LeagueEntryDto, error) {
	var LeagueEntry []LeagueEntryDto
	resp, err := GetRequest(fmt.Sprintf("https://euw1.api.riotgames.com/lol/league/v4/entries/by-summoner/%s", summonerId))
	if err != nil {
		return LeagueEntry, err
	}
	err = json.NewDecoder(resp.Body).Decode(&LeagueEntry)
	if err != nil {
		return LeagueEntry, err
	}
	return LeagueEntry, nil
}