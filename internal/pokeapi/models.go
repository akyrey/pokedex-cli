package pokeapi

import "fmt"

type Pagination[T any] struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []T
	Count    int `json:"count"`
}

type NameUrlStruct struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type EncounterDetail struct {
	ConditionValues []string `json:"condition_values"`
	Chance          int      `json:"chance"`
	MinLevel        int      `json:"min_level"`
	MaxLevel        int      `json:"max_level"`
}

type VersionDetail struct {
	Version          NameUrlStruct     `json:"version"`
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int               `json:"max_chance"`
}

type Encounter struct {
	Pokemon        NameUrlStruct   `json:"pokemon"`
	VersionDetails []VersionDetail `json:"version_details"`
}

type ExploreResponse struct {
	Location   NameUrlStruct `json:"location"`
	Name       string        `json:"name"`
	Encounters []Encounter   `json:"pokemon_encounters"`
	ID         int           `json:"id"`
	GameIndex  int           `json:"game_index"`
}

type Pokemon struct {
	Sprites struct {
		BackFemale       *string `json:"back_female"`
		BackShinyFemale  *string `json:"back_shiny_female"`
		FrontFemale      *string `json:"front_female"`
		FrontShinyFemale *string `json:"front_shiny_female"`
		BackDefault      string  `json:"back_default"`
		BackShiny        string  `json:"back_shiny"`
		FrontDefault     string  `json:"front_default"`
		FrontShiny       string  `json:"front_shiny"`
	} `json:"sprites"`
	Species                NameUrlStruct `json:"species"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Name                   string        `json:"name"`
	Moves                  []struct {
		Move                NameUrlStruct `json:"move"`
		VersionGroupDetails []struct {
			MoveLearnMethod NameUrlStruct `json:"move_learn_method"`
			VersionGroup    NameUrlStruct `json:"version_group"`
			LevelLearnedAt  int           `json:"level_learned_at"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Abilities []struct {
		Ability  NameUrlStruct `json:"ability"`
		Slot     int           `json:"slot"`
		IsHidden bool          `json:"is_hidden"`
	} `json:"abilities"`
	Forms       []NameUrlStruct `json:"forms"`
	GameIndices []struct {
		Version   NameUrlStruct `json:"version"`
		GameIndex int           `json:"game_index"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item           NameUrlStruct `json:"item"`
		VersionDetails []struct {
			Version NameUrlStruct `json:"version"`
			Rarity  int           `json:"rarity"`
		} `json:"version_details"`
	} `json:"held_items"`
	Stats []struct {
		Stat     NameUrlStruct `json:"stat"`
		BaseStat int           `json:"base_stat"`
		Effort   int           `json:"effort"`
	} `json:"stats"`
	Types []struct {
		Type NameUrlStruct `json:"type"`
		Slot int           `json:"slot"`
	} `json:"types"`
	Order          int  `json:"order"`
	Weight         int  `json:"weight"`
	ID             int  `json:"id"`
	Height         int  `json:"height"`
	BaseExperience int  `json:"base_experience"`
	IsDefault      bool `json:"is_default"`
}

func (p Pokemon) String() string {
	res := fmt.Sprintf("Name: %s\n", p.Name)
	res += fmt.Sprintf("Height: %d\n", p.Height)
	res += fmt.Sprintf("Weight: %d\n", p.Weight)
	res += "Stats:\n"
	for _, stat := range p.Stats {
		res += fmt.Sprintf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	res += "Types:\n"
	for _, t := range p.Types {
		res += fmt.Sprintf("  - %s\n", t.Type.Name)
	}

	return res
}
