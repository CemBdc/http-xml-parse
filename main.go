package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	resp, err := http.Get("http://realtime.paragaranti.com/asp/xml/icpiyasa.asp")

	if err != nil{
		log.Fatal("Error occured: %s", err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	var dm domesticMarket

	xml.Unmarshal(data, &dm)

	fmt.Println(dm.Stocks)

}

type stock struct{
	XMLName xml.Name `xml:"STOCK"`
	SYMBOL string `xml:"SYMBOL"`
	DESC string `xml:"DESC"`
	LAST string `xml:"LAST"`
	PERNC string `xml:"PERNC"`
	LAST_MOD string `xml:"LAST_MOD"`
}

// ICPIYASA Tag in xml
type domesticMarket struct{
	XMLName xml.Name `xml:"ICPIYASA"`
	Stocks []stock `xml:"STOCK"`
}

func (s stock) String() string{
	return fmt.Sprintf("\t Symbol: %s, Description: %s, Last Value: %s, Change: %s, Last Updated Date: %s \n", s.SYMBOL, s.DESC, s.LAST, s.PERNC, s.LAST_MOD)
}