package TwistedFate

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type championMasteryDtoJson struct {
	ChampionId                   int    `json:"championId"`
	ChampionLevel                int    `json:"championLevel"`
	ChampionPoints               int    `json:"championPoints"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	TokensEarned                 int    `json:"tokensEarned"`
	SummonerId                   string `json:"summonerId"`
}

type ChampionMasteryDto struct {
	ChampionId                   string    `json:"championId"`
	ChampionLevel                int    `json:"championLevel"`
	ChampionPoints               int    `json:"championPoints"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	TokensEarned                 int    `json:"tokensEarned"`
	SummonerId                   string `json:"summonerId"`
}

func GetChampMap() []Champion {
	x, _ := GetRequest("https://ddragon.leagueoflegends.com/cdn/11.8.1/data/en_US/champion.json")
	var y map[string]interface{}
	json.NewDecoder(x.Body).Decode(&y)
	w := y["data"]
	t, _ := json.Marshal(w)
	m := make(map[string]Champion)
	var ChampionList []Champion
	err := json.Unmarshal(t, &m)
	for _, v := range m {
		ChampionList = append(ChampionList, v)
	}
	if err != nil {
		fmt.Println(err)
	}
	return ChampionList
}

func ChampionIntToString(ChampionId int, cM []Champion) string {
	for _, v := range cM {
		i, _ := strconv.Atoi(v.Key)
		if i == ChampionId {
			return v.Name
		}
	}
	return "Not Found."
}

func (x championMasteryDtoJson) ConvertToChampionMastery(cM []Champion) ChampionMasteryDto {

	y := ChampionMasteryDto{
		ChampionId: ChampionIntToString(x.ChampionId, cM),
		ChampionLevel: x.ChampionLevel,
		ChampionPoints: x.ChampionPoints,
		LastPlayTime: x.LastPlayTime,
		ChampionPointsSinceLastLevel: x.ChampionPointsSinceLastLevel,
		ChampionPointsUntilNextLevel: x.ChampionPointsUntilNextLevel,
		ChestGranted: x.ChestGranted,
		TokensEarned: x.TokensEarned,
		SummonerId: x.SummonerId,
	}

	return y
}


func GetChampionMasteryBySummonerId(summonerId string) ([]ChampionMasteryDto, error) {
	var ChampionMastery []championMasteryDtoJson
	resp, err := GetRequest(fmt.Sprintf("https://euw1.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-summoner/%s", summonerId))
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(resp.Body).Decode(&ChampionMastery)
	if err != nil {
		return nil, err
	}
	cM := GetChampMap()
	var ChampionMasteryList []ChampionMasteryDto
	for _, v := range ChampionMastery {
		ChampionMasteryList = append(ChampionMasteryList, v.ConvertToChampionMastery(cM))
	}
	return ChampionMasteryList, nil
}

func GetChampionMasteryScoreBySummonerId(summonerId string) (int, error) {
	var ChampionMasteryScore int
	resp, err := GetRequest(fmt.Sprintf("https://euw1.api.riotgames.com/lol/champion-mastery/v4/scores/by-summoner/%s", summonerId))
	if err != nil {
		return 0, err
	}
	err = json.NewDecoder(resp.Body).Decode(&ChampionMasteryScore)
	return ChampionMasteryScore, nil
}