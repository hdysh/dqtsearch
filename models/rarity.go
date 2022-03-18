package models

type Rarity struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Frame      string `json:"frame"`
	Background string `json:"background"`
}
