package toriko

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	m "github.com/kaneshin/snacky-go/app/models"
)

type Toriko struct {
	*m.Model
}

func NewToriko() *Toriko {
	instance := &Toriko{}
	instance.InitModel()
	return instance
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
