package mzklib

import "log"

type GetStopsJSON struct {
	StopPoints []StopPoint `json:"stopPoints"`
}

type StopPoint struct {
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
}

var Stops GetStopsJSON

func init() {
	log.Println("init(): Pobieranie stopów")
	Stops = GetStops()
	log.Println("init(): Liczba stopów:", len(Stops.StopPoints))
}

func GetStops() (stops GetStopsJSON) {
	httpget("http://przystanki.zywiec.pl/getStops.json", &stops)
	return stops
}

func GetStop(symbol string) StopPoint {
	for i := range Stops.StopPoints {
		if Stops.StopPoints[i].Symbol == symbol {
			return Stops.StopPoints[i]
		}
	}
	return StopPoint{}
}
