package mzklib

type GetRunningVehiclesJSON struct {
	Offline  bool `json:"offline"`
	Vehicles []struct {
		Angle                      int64       `json:"angle"`
		CourseLoid                 int64       `json:"courseLoid"`
		DayCourseLoid              int64       `json:"dayCourseLoid"`
		DelaySec                   interface{} `json:"delaySec"`
		DistanceToNearestStopPoint int64       `json:"distanceToNearestStopPoint"`
		LastPingDate               int64       `json:"lastPingDate"`
		Latitude                   float64     `json:"latitude"`
		LineName                   string      `json:"lineName"`
		Longitude                  float64     `json:"longitude"`
		NearestSymbol              string      `json:"nearestSymbol"`
		OnStopPoint                string      `json:"onStopPoint"`
		Operator                   string      `json:"operator"`
		OptionalDirection          string      `json:"optionalDirection"`
		OrderInCourse              int64       `json:"orderInCourse"`
		ReachedMeters              int64       `json:"reachedMeters"`
		VariantLoid                int64       `json:"variantLoid"`
		VehicleID                  string      `json:"vehicleId"`
	} `json:"vehicles"`
}

func GetRunningVehicles() (linie GetRunningVehiclesJSON) {
	httpget("http://przystanki.zywiec.pl/getRunningVehicles.json", &linie)
	return
}
