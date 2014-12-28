package toriko

import (
	"encoding/xml"

	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	s "strconv"
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
