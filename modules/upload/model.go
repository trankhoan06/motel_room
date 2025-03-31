package upload

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Url       string `json:"url" gorm:"column:url"`
	Width     int    `json:"width" gorm:"column:width"`
	Height    int    `json:"height" gorm:"column:height"`
	CloudName string `json:"cloud_name" gorm:"column:cloud_name"`
	Extension string `json:"extension" gorm:"column:extension"`
}

func (Image) TableName() string { return "image" }
func (i *Image) Fullfill(s string) {
	i.Url = fmt.Sprintf("%s/%s", s, i.Url)
}
func (i *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Can't unmarshal JSONB value")
	}
	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*i = img
	return nil
}
func (i *Image) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}
	return json.Marshal(i)
}
