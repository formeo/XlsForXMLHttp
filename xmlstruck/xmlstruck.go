package xmlstruck

import (
	"encoding/xml"
)

//Servers Шапка XML
type Servers struct {
	XMLName xml.Name `xml:"xslfiles"`
	Version string   `xml:"version,attr"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`
	Svs     []Files  `xml:"files"`
}

//Files тело XML
type Files struct {
	Number   string `xml:"number"`
	Date     string `xml:"date"`
	Summ     string `xml:"summ"`
	Filename string `xml:"filename"`
}
