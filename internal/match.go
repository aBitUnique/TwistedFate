package internal

import (
	"encoding/json"
	"fmt"
)

type MatchlistDto struct {
	StartIndex int
	TotalGames int
	EndIndex int
	Matches []MatchReferenceDto
}

type MatchReferenceDto struct {
	GameId float32
	Role string
	Season int
	PlatformId string
	Champion int
	Queue int
	Lane string
	Timestamp float32
}

func (mdto MatchlistDto) String() string {
	return fmt.Sprintf("Total matches: %d",
		mdto.TotalGames)
}

func GetMatchlistByAccountId(accountId string) (MatchlistDto, error) {
	var Matchlist MatchlistDto
	resp, err := GetRequest(fmt.Sprintf("https://euw1.api.riotgames.com/lol/match/v4/matchlists/by-account/%s", accountId))
	if err != nil {
		return Matchlist, err
	}
	err = json.NewDecoder(resp.Body).Decode(&Matchlist)
	if err != nil {
		return Matchlist, err
	}
	return Matchlist, nil
}