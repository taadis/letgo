package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

//
type Article struct {
	XMLName xml.Name `xml:"article"`
	Id      string   `xml:"id,attr"`
	Title   string   `xml:"Title"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"`
}

//
type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

//
func main() {
	log.Println("read_xml_file.go")
	xmlFile, err := os.Open("article.xml")
	if err != nil {
		log.Panicln(err)
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Panicln(err)
	}
	article := Article{}
	err = xml.Unmarshal(xmlData, &article)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(article)
}
