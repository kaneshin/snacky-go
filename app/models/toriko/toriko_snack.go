package toriko

import (
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/kaneshin/snacky-go/app"

	"io/ioutil"
	"net/http"
	s "strconv"

	e "github.com/kaneshin/snacky-go/app/entities"
)

type TorikoSnackItem struct {
	XMLName  xml.Name         `xml:"item"`
	Id       int32            `xml:"id"`
	Name     string           `xml:"name"`
	Kana     string           `xml:"kana"`
	Maker    string           `xml:"maker"`
	Price    string           `xml:"price"`
	Kind     string           `xml:"type"`
	MakedAt  string           `xml:"regist"`
	ImageUrl string           `xml:"image"`
	Comment  string           `xml:"comment"`
	Tags     []TorikoSnackTag `xml:"tags"`
}

type TorikoSnackTag struct {
	XMLName xml.Name `xml:"tags"`
	Tag     string   `xml:"tag"`
}

type TorikoSnackData struct {
	XMLName xml.Name          `xml:"okashinotoriko"`
	Items   []TorikoSnackItem `xml:"item"`
	Count   int32             `xml:"count"`
}

func (m *Toriko) UpdateSnacksByKind(kind int, limit int, offset int) error {
	items, err := m.GetSnacksByKind(kind, limit, offset)
	if err != nil {
		return err
	}
	if len(*items) == 0 {
		return errors.New("No items")
	}
	return m.insertSnacks(items)
}

func (m *Toriko) UpdateSnacksByYear(year int, limit int, offset int) error {
	items, err := m.GetSnacksByYear(year, limit, offset)
	if err != nil {
		return err
	}
	if len(*items) == 0 {
		return errors.New("No items")
	}
	return m.insertSnacks(items)
}

func (m *Toriko) GetSnacksByKind(kind int, limit int, offset int) (*[]TorikoSnackItem, error) {
	return m.getSnacks("type", kind, limit, offset)
}

func (m *Toriko) GetSnacksByYear(year int, limit int, offset int) (*[]TorikoSnackItem, error) {
	return m.getSnacks("year", year, limit, offset)
}

func (m *Toriko) getSnacks(key string, value int, limit int, offset int) (*[]TorikoSnackItem, error) {
	url := m.getURL()
	q := url.Query()
	q.Set(key, s.Itoa(value))
	q.Set("max", s.Itoa(limit))
	q.Set("offset", s.Itoa(offset))
	url.RawQuery = q.Encode()
	resp, err := http.Get(url.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var d TorikoSnackData
	xml.Unmarshal(body, &d)
	return &d.Items, err
}

func (m *Toriko) insertSnacks(items *[]TorikoSnackItem) error {
	for _, item := range *items {
		d := *e.NewSnack(item.Id, item.Name)
		if err := app.Dbd.Dbm.Insert(d.Scheme()); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
