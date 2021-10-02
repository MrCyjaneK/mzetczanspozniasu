package mzklib

type GetLinesJSON struct {
	Lines []Linie `json:"lines"`
}

type Linie struct {
	ID           int64  `json:"id"`
	LineType     string `json:"lineType"`
	LineValidity []struct {
		From  int64 `json:"from"`
		Until int64 `json:"until"`
	} `json:"lineValidity"`
	Name      string `json:"name"`
	NightLine bool   `json:"nightLine"`
}

func GetLines() (linie GetLinesJSON) {
	httpget("http://przystanki.zywiec.pl/getLines.json", &linie)
	return
}
