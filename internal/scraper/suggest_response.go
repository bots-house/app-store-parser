package scraper

import "encoding/xml"

type suggestXML struct {
	XMLName xml.Name    `xml:"plist"`
	Dict    suggestDict `xml:"dict"`
}

type suggestDict struct {
	Title string        `xml:"key"`
	Hints []suggestHint `xml:"array>dict"`
}

type suggestHint struct {
	Keys   []string `xml:"key"`
	Values []string `xml:"string"`
}
