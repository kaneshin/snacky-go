package models

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	s "strconv"
)

type Toriko struct {
	*Model
}

func NewToriko() *Toriko {
	instance := &Toriko{}
	instance.InitModel()
	return instance
}

// == Snack

type TorikoSnackItem struct {
	XMLName  xml.Name         `xml:"item"`
	Id       int              `xml:"id"`
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
}

func (m *Toriko) getURL() *url.URL {
	u, err := url.Parse("http://www.sysbird.jp/toriko/api")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("apikey", "guest")
	q.Set("format", "xml")
	u.RawQuery = q.Encode()
	return u
}

func (m *Toriko) getSnack(key string, value int, limit int, offset int) (*[]TorikoSnackItem, error) {
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

func (m *Toriko) GetSnackByKind(kind int, limit int, offset int) (*[]TorikoSnackItem, error) {
	return m.getSnack("type", kind, limit, offset)
}

func (m *Toriko) GetSnackByYear(year int, limit int, offset int) (*[]TorikoSnackItem, error) {
	return m.getSnack("year", year, limit, offset)
}

// = Master

func (m *Toriko) getMaster(list string) ([]byte, error) {
	url := m.getURL()
	q := url.Query()
	q.Set("list", list)
	url.RawQuery = q.Encode()
	resp, err := http.Get(url.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// == Kind

type TorikoKindItem struct {
	XMLName xml.Name `xml:"type"`
	Id      int      `xml:"id"`
	Name    string   `xml:"name"`
}

type TorikoKindData struct {
	XMLName xml.Name         `xml:"okashinotoriko"`
	Items   []TorikoKindItem `xml:"type"`
}

func (m *Toriko) GetKindMaster() (*[]TorikoKindItem, error) {
	body, err := m.getMaster("type")
	var d TorikoKindData
	xml.Unmarshal(body, &d)
	return &d.Items, err
}

// == Area

type TorikoAreaItem struct {
	XMLName xml.Name `xml:"area"`
	Id      int      `xml:"id"`
	Name    string   `xml:"name"`
}

type TorikoAreaData struct {
	XMLName xml.Name         `xml:"okashinotoriko"`
	Items   []TorikoAreaItem `xml:"area"`
}

func (m *Toriko) GetAreaMaster() (*[]TorikoAreaItem, error) {
	body, err := m.getMaster("area")
	var d TorikoAreaData
	xml.Unmarshal(body, &d)
	return &d.Items, err
}

// == maker

type TorikoMakerItem struct {
	XMLName xml.Name `xml:"maker"`
	Id      int      `xml:"id"`
	Name    string   `xml:"name"`
}

type TorikoMakerData struct {
	XMLName xml.Name          `xml:"okashinotoriko"`
	Items   []TorikoMakerItem `xml:"maker"`
}

func (m *Toriko) GetMakerMaster() (*[]TorikoMakerItem, error) {
	body, err := m.getMaster("maker")
	var d TorikoMakerData
	xml.Unmarshal(body, &d)
	return &d.Items, err
}
