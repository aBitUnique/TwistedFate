package TwistedFate

import "encoding/json"

type Champion struct {
	Version string `json:"version"`
	ID      string `json:"id"`
	Key     string `json:"key"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Blurb   string `json:"blurb"`
	Info    struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
	} `json:"info"`
	Image struct {
		Full   string `json:"full"`
		Sprite string `json:"sprite"`
		Group  string `json:"group"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
	} `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   struct {
		Hp                   float64     `json:"hp"`
		Hpperlevel           float64     `json:"hpperlevel"`
		Mp                   float64     `json:"mp"`
		Mpperlevel           float64     `json:"mpperlevel"`
		Movespeed            int     `json:"movespeed"`
		Armor                int     `json:"armor"`
		Armorperlevel        float64 `json:"armorperlevel"`
		Spellblock           float64     `json:"spellblock"`
		Spellblockperlevel   float64 `json:"spellblockperlevel"`
		Attackrange          int     `json:"attackrange"`
		Hpregen              float64     `json:"hpregen"`
		Hpregenperlevel      float64     `json:"hpregenperlevel"`
		Mpregen              float64     `json:"mpregen"`
		Mpregenperlevel      float64     `json:"mpregenperlevel"`
		Crit                 int     `json:"crit"`
		Critperlevel         int     `json:"critperlevel"`
		Attackdamage         float64     `json:"attackdamage"`
		Attackdamageperlevel float64     `json:"attackdamageperlevel"`
		Attackspeedperlevel  float64 `json:"attackspeedperlevel"`
		Attackspeed          float64 `json:"attackspeed"`
	} `json:"stats"`
}

type ChampionRotationJson struct {
FreeChampionIds              []int `json:"freeChampionIds"`
FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

type ChampionRotation struct {
	FreeChampionIds              []string `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []string `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

func GetChampionRotation() (ChampionRotation, error) {
	cM := GetChampMap()
	var ChampionRotationJson ChampionRotationJson
	resp, err := GetRequest("https://euw1.api.riotgames.com/lol/platform/v3/champion-rotations")
	if err != nil {
		return ChampionRotation{}, err
	}
	err = json.NewDecoder(resp.Body).Decode(&ChampionRotationJson)
	if err != nil {
		return ChampionRotation{}, err
	}
	var FreeChampionIds []string
	for _, v := range ChampionRotationJson.FreeChampionIds {
		FreeChampionIds = append(FreeChampionIds, ChampionIntToString(v, cM))
	}
	var FreeChampionIdsForNewPlayers []string
	for _, v := range ChampionRotationJson.FreeChampionIds {
		FreeChampionIdsForNewPlayers = append(FreeChampionIdsForNewPlayers, ChampionIntToString(v, cM))
	}
	return ChampionRotation{FreeChampionIds: FreeChampionIds, FreeChampionIdsForNewPlayers: FreeChampionIdsForNewPlayers, MaxNewPlayerLevel: ChampionRotationJson.MaxNewPlayerLevel}, nil
}