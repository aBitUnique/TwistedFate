package TwistedFate

import (
	"encoding/json"
	"fmt"
)

type SummonerDto struct {
	Id string
	AccountId string
	Puuid string
	Name string
	ProfileIconId int
	RevisionDate float32
	SummonerLevel int
}

func (sDto SummonerDto) String() string {
	return fmt.Sprintf("Id: %s \nAccountId: %s \nPuuid: %s \nName: %s \nProfileIconId: %d \nRevisionDate: %g \nSummonerLevel: %d",
		sDto.Id,
		sDto.AccountId,
		sDto.Puuid,
		sDto.Name,
		sDto.ProfileIconId,
		sDto.RevisionDate,
		sDto.SummonerLevel)
}

func GetSummonerByAccountId(accountId string) (*SummonerDto, error) {
	Summoner := new(SummonerDto)
	resp, err := GetRequest(fmt.Sprintf("https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-account/%s", accountId))
	if err != nil {
		return Summoner, err
	}
	err = json.NewDecoder(resp.Body).Decode(Summoner)
	if err != nil {
		return Summoner, err
	}
	return Summoner, nil
}

func GetSummonerByName(name string) (*SummonerDto, error) {
	Summoner := new(SummonerDto)
	resp, err := GetRequest(fmt.Sprintf("https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s", name))
	if err != nil {
		return Summoner, err
	}
	err = json.NewDecoder(resp.Body).Decode(Summoner)
	if err != nil {
		return Summoner, err
	}
	return Summoner, nil
}

func GetSummonerByPuuid(puuid string) (*SummonerDto, error) {
	Summoner := new(SummonerDto)
	resp, err := GetRequest(fmt.Sprintf("https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/%s", puuid))
	if err != nil {
		return Summoner, err
	}
	err = json.NewDecoder(resp.Body).Decode(Summoner)
	if err != nil {
		return Summoner, err
	}
	return Summoner, nil
}

func GetSummonerBySummonerId(summonerId string) (*SummonerDto, error) {
	Summoner := new(SummonerDto)
	resp, err := GetRequest(fmt.Sprintf("https://euw1.api.riotgames.com/lol/summoner/v4/summoners/%s", summonerId))
	if err != nil {
		return Summoner, err
	}
	err = json.NewDecoder(resp.Body).Decode(Summoner)
	if err != nil {
		return Summoner, err
	}
	return Summoner, nil
}