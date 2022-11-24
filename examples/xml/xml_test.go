package xml

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type urlset struct {
	Xmlns1 xml.Attr `xml:"xmlns,attr"`
	Xmlns2 xml.Attr `xml:"xmlns:xsi,attr"`
	Xmlns3 xml.Attr `xml:"xsi:schemaLocation,attr"`
	URL    []url    `xml:"url"`
}

type url struct {
	Loc      string `xml:"loc"`
	LastMod  string `xml:"lastmod"`
	Priority string `xml:"priority"`
}

func TestXml(t *testing.T) {
	o := urlset{
		Xmlns1: xml.Attr{
			Name:  xml.Name{Local: "xmlns"},
			Value: "http://www.sitemaps.org/schemas/sitemap/0.9",
		},
		Xmlns2: xml.Attr{
			Name:  xml.Name{Local: "xmlns:xsi"},
			Value: "http://www.w3.org/2001/XMLSchema-instance",
		},
		Xmlns3: xml.Attr{
			Name:  xml.Name{Local: "xsi:schemaLocation"},
			Value: "http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd",
		},
		URL: []url{
			{Loc: "http://sadovnichii75.mech.math.msu.su/", LastMod: "2022-11-22T08:46:35+00:00", Priority: "1.00"},
			{Loc: "http://sadovnichii75.mech.math.msu.su/2", LastMod: "2022-11-22T08:46:35+00:00", Priority: "1.00"},
		},
	}
	buf, err := xml.Marshal(&o)
	assert.Nil(t, err)
	f, err := os.Create("test.xml")
	assert.Nil(t, err)
	_, err = f.Write(buf)
	assert.Nil(t, err)
	fmt.Println(string(buf))
	//assert.Equal(t, string(buf), "<Object><Loc>12</Loc><Value>test</Value></Object>")
}
