package mzklib

type GetRealtimeJSON struct {
	Departures []struct {
		CourseID              int64  `json:"courseId"`
		DirectionName         string `json:"directionName"`
		Lack                  bool   `json:"lack"`
		LineName              string `json:"lineName"`
		OnStopPoint           bool   `json:"onStopPoint"`
		OrderInCourse         int64  `json:"orderInCourse"`
		Passed                bool   `json:"passed"`
		RealDeparture         int64  `json:"realDeparture"`
		ScheduledDeparture    int64  `json:"scheduledDeparture"`
		ScheduledDepartureSec int64  `json:"scheduledDepartureSec"`
		VariantID             int64  `json:"variantId"`
		VehicleID             string `json:"vehicleId"`
	} `json:"departures"`
	Error           interface{} `json:"error"`
	ResponseDate    int64       `json:"responseDate"`
	StopPointID     int64       `json:"stopPointId"`
	StopPointName   string      `json:"stopPointName"`
	StopPointSymbol string      `json:"stopPointSymbol"`
}

func GetRealtime(stopPointSymbol string) (linie GetRealtimeJSON) {
	httpget("http://przystanki.zywiec.pl/getRealtime.json?stopPointSymbol="+stopPointSymbol, &linie)
	return
}
