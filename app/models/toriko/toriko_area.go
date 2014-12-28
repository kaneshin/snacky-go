package toriko

import (
	"encoding/xml"
	"fmt"

	"github.com/kaneshin/snacky-go/app"
	e "github.com/kaneshin/snacky-go/app/entities"
)

type TorikoAreaItem struct {
	XMLName xml.Name `xml:"area"`
	Id      int32    `xml:"id"`
	Name    string   `xml:"name"`
}

type TorikoAreaData struct {
	XMLName xml.Name         `xml:"okashinotoriko"`
	Items   []TorikoAreaItem `xml:"area"`
}

func (m *Toriko) UpdateAreaMaster() error {
	items, err := m.getAreaMaster()
	if err != nil {
		return err
	}
	return m.insertAreas(items)
}

func (m *Toriko) getAreaMaster() (*[]TorikoAreaItem, error) {
	body, err := m.getMaster("area")
	var d TorikoAreaData
	xml.Unmarshal(body, &d)
	return &d.Items, err
}

func (m *Toriko) insertAreas(items *[]TorikoAreaItem) error {
	for _, item := range *items {
		d := *e.NewArea(item.Id, item.Name)
		if err := app.Dbd.Dbm.Insert(d.Scheme()); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
