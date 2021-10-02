package mzklib

import "strconv"

type GetDirectionJSON struct {
	AltDestination interface{} `json:"altDestination"`
	AltStopPoints  []struct {
		City              interface{} `json:"city"`
		Distance          interface{} `json:"distance"`
		GettingOut        bool        `json:"gettingOut"`
		ID                int64       `json:"id"`
		Latitude          float64     `json:"latitude"`
		Longitude         float64     `json:"longitude"`
		Name              string      `json:"name"`
		OnRequest         bool        `json:"onRequest"`
		StopPointValidity []struct {
			From  int64 `json:"from"`
			Until int64 `json:"until"`
		} `json:"stopPointValidity"`
		Street  string `json:"street"`
		Symbol  string `json:"symbol"`
		Virtual bool   `json:"virtual"`
	} `json:"altStopPoints"`
	AltVariantLoid int64 `json:"altVariantLoid"`
	Connections    []struct {
		FromSymbol string `json:"fromSymbol"`
		Main       bool   `json:"main"`
		Nodes      []struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			OrderNo   int64   `json:"orderNo"`
		} `json:"nodes"`
		ToSymbol   string `json:"toSymbol"`
		ValidFrom  int64  `json:"validFrom"`
		ValidUntil int64  `json:"validUntil"`
	} `json:"connections"`
	Destination         interface{} `json:"destination"`
	HasAnotherDirection bool        `json:"hasAnotherDirection"`
	MainVariantLoid     int64       `json:"mainVariantLoid"`
	StopPoints          []struct {
		City              interface{} `json:"city"`
		Distance          interface{} `json:"distance"`
		GettingOut        bool        `json:"gettingOut"`
		ID                int64       `json:"id"`
		Latitude          float64     `json:"latitude"`
		Longitude         float64     `json:"longitude"`
		Name              string      `json:"name"`
		OnRequest         bool        `json:"onRequest"`
		StopPointValidity []struct {
			From  int64 `json:"from"`
			Until int64 `json:"until"`
		} `json:"stopPointValidity"`
		Street  string `json:"street"`
		Symbol  string `json:"symbol"`
		Virtual bool   `json:"virtual"`
	} `json:"stopPoints"`
}

func GetDirection(lineID int, isthere bool) (direction GetDirectionJSON) {
	httpget("http://przystanki.zywiec.pl/getDirection.json?lineId="+strconv.Itoa(lineID)+"&thereDirection="+strconv.FormatBool(isthere), &direction)
	return
}

func GetDirectionS(lineID string, isthere string) (direction GetDirectionJSON) {
	httpget("http://przystanki.zywiec.pl/getDirection.json?lineId="+lineID+"&thereDirection="+isthere, &direction)
	return
}
