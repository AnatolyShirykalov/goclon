package parser

import (
	"bytes"
	"compress/gzip"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
	"gopkg.in/resty.v1"
	yaml "gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net/url"
)

func gUnzipData(data []byte) (resData []byte, err error) {
	b := bytes.NewBuffer(data)

	var r io.Reader
	r, err = gzip.NewReader(b)
	if err != nil {
		return
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return
	}

	resData = resB.Bytes()

	return
}

var Headers map[string]string

func init() {
	dat, err := ioutil.ReadFile("./config/headers.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(dat), &Headers)
	if err != nil {
		panic(err)
	}
}

func GetPage(client *resty.Client, urlOrPath string) []byte {
	u, err := url.Parse(urlOrPath)
	if err != nil {
		panic(err)
	}
	base, err := url.Parse("http://classic-online.ru")
	if err != nil {
		panic(err)
	}
	uri := base.ResolveReference(u).String()
	resp, err := client.R().Get(uri)
	if resp.StatusCode() == 302 {
		return []byte{}
	}
	gz := resp.Body()
	bytes, err := gUnzipData(gz)
	if err != nil {
		panic(err)
	}
	dec := charmap.Windows1251.NewDecoder()
	u8bs, err := dec.Bytes(bytes)
	return u8bs
}

func GetDoc(client *resty.Client, url string) (ret *goquery.Document) {
	data := GetPage(client, url)
	if len(data) == 0 {
		ret = nil
		return
	}
	ret, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	return
}

func NewClient() *resty.Client {
	client := resty.New()
	client.SetHeaders(Headers)
	return client
}

func Page(client *resty.Client, uri string) {
	doc := GetDoc(client, uri)
	if doc == nil {
		return
	}
	c := composer(doc)
	piece := piece(doc, c)
	group := group(doc)
	perform := perform(doc, piece, group)
	comments(client, doc, perform)
}
