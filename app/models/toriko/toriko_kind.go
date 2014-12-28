package toriko

import (
	"encoding/xml"
	"fmt"

	"github.com/kaneshin/snacky-go/app"
	e "github.com/kaneshin/snacky-go/app/entities"
)

type TorikoKindItem struct {
	XMLName xml.Name `xml:"type"`
	Id      int32    `xml:"id"`
	Name    string   `xml:"name"`
}

type TorikoKindData struct {
	XMLName xml.Name         `xml:"okashinotoriko"`
	Items   []TorikoKindItem `xml:"type"`
}

func (m *Toriko) UpdateKindMaster() error {
	items, err := m.getKindMaster()
	if err != nil {
		return err
	}
	return m.insertKinds(items)
}

func (m *Toriko) getKindMaster() (*[]TorikoKindItem, error) {
	body, err := m.getMaster("type")
	var d TorikoKindData
	xml.Unmarshal(body, &d)
	return &d.Items, err
}

func (m *Toriko) insertKinds(items *[]TorikoKindItem) error {
	for _, item := range *items {
		d := *e.NewKind(item.Id, item.Name)
		if err := app.Dbd.Dbm.Insert(d.Scheme()); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
