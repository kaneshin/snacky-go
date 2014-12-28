package toriko

import (
	"encoding/xml"
	"fmt"

	"github.com/kaneshin/snacky-go/app"
	e "github.com/kaneshin/snacky-go/app/entities"
)

// == maker

type TorikoMakerItem struct {
	XMLName xml.Name `xml:"maker"`
	Id      int32    `xml:"id"`
	Name    string   `xml:"name"`
}

type TorikoMakerData struct {
	XMLName xml.Name          `xml:"okashinotoriko"`
	Items   []TorikoMakerItem `xml:"maker"`
}

func (m *Toriko) UpdateMakerMaster() error {
	items, err := m.getMakerMaster()
	if err != nil {
		return err
	}
	return m.insertMakers(items)
}

func (m *Toriko) getMakerMaster() (*[]TorikoMakerItem, error) {
	body, err := m.getMaster("maker")
	var d TorikoMakerData
	xml.Unmarshal(body, &d)
	return &d.Items, err
}

func (m *Toriko) insertMakers(items *[]TorikoMakerItem) error {
	for _, item := range *items {
		d := *e.NewMaker(item.Id, item.Name)
		if err := app.Dbd.Dbm.Insert(d.Scheme()); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
