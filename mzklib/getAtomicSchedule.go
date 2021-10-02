package mzklib

import (
	"fmt"
	"sort"
)

type GetAtomicScheduleJSON struct {
	AllLinesInActiveStraps []struct {
		ID           int64         `json:"id"`
		LineType     interface{}   `json:"lineType"`
		LineValidity []interface{} `json:"lineValidity"`
		Name         string        `json:"name"`
		NightLine    bool          `json:"nightLine"`
	} `json:"allLinesInActiveStraps"`
	AutoLettersAssignment bool          `json:"autoLettersAssignment"`
	DayTypeSymbol         string        `json:"dayTypeSymbol"`
	GettingOutVariants    []interface{} `json:"gettingOutVariants"`
	LineSchedules         map[int]struct {
		Departures       []Departure `json:"departures"`
		Destination      string      `json:"destination"`
		IsGettingOutOnly bool        `json:"isGettingOutOnly"`
		LineName         string      `json:"lineName"`
		LineValidity     []struct {
			From  int64 `json:"from"`
			Until int64 `json:"until"`
		} `json:"lineValidity"`
		MainVariantID   int64 `json:"mainVariantId"`
		OverrideLetters struct {
			Two510 string `json:"2510"`
			Two512 string `json:"2512"`
			Two543 string `json:"2543"`
			Four   string `json:"4"`
			Four05 string `json:"405"`
		} `json:"overrideLetters"`
	} `json:"lineSchedules"`
	MultipleLettersAssignment bool        `json:"multipleLettersAssignment"`
	ScheduleDate              interface{} `json:"scheduleDate"`
}

type Departure struct {
	Brigade                  string `json:"brigade"`
	CourseID                 int64  `json:"courseId"`
	CourseStartSec           int64  `json:"courseStartSec"`
	CourseStartString        string
	ExpectedVehicleType      string        `json:"expectedVehicleType"`
	Letter                   string        `json:"letter"`
	MultipleLegends          []interface{} `json:"multipleLegends"`
	Operator                 string        `json:"operator"`
	OptionalDirection        string        `json:"optionalDirection"`
	OrderInCourse            int64         `json:"orderInCourse"`
	OverloadedID             interface{}   `json:"overloadedId"`
	ScheduledDepartureSec    int64         `json:"scheduledDepartureSec"`
	ScheduledDepartureString string
	TransportType            string      `json:"transportType"`
	VariantID                int64       `json:"variantId"`
	Visible                  bool        `json:"visible"`
	WorkaroundID             interface{} `json:"workaroundId"`
}

type BySec []Departure

func (a BySec) Len() int           { return len(a) }
func (a BySec) Less(i, j int) bool { return a[i].ScheduledDepartureSec < a[j].ScheduledDepartureSec }
func (a BySec) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func GetAtomicSchedule(symbol string) (sched GetAtomicScheduleJSON) {
	httpget("http://przystanki.zywiec.pl/getAtomicSchedule.json?symbol="+symbol, &sched)
	for i := range sched.LineSchedules {
		sort.Sort(BySec(sched.LineSchedules[i].Departures))
		for j := range sched.LineSchedules[i].Departures {
			hours := int64(sched.LineSchedules[i].Departures[j].ScheduledDepartureSec / 60 / 60)
			minutes := (sched.LineSchedules[i].Departures[j].ScheduledDepartureSec - (hours * 60 * 60)) / 60
			sched.LineSchedules[i].Departures[j].ScheduledDepartureString = fmt.Sprintf("%02d:%02d", hours, minutes)
		}
	}
	return
}
