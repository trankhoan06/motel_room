package model

import "main.go/modules/rent/model"

type SortSearch struct {
	Province  string          `json:"province"`
	LowPrice  int             `json:"low_price"`
	HighPrice int             `json:"high_price"`
	LowArea   int             `json:"low_area"`
	HighArea  int             `json:"high_area"`
	TypeRoom  *model.TypeRoom `json:"type_room"`
}
