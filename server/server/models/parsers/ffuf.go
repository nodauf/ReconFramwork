package modelsParsers

import "time"

type Ffuf struct {
	Commandline string    `json:"commandline"`
	Time        time.Time `json:"time"`
	Results     []struct {
		Input struct {
			FUZZ string `json:"FUZZ"`
		} `json:"input"`
		Position         int    `json:"position"`
		Status           int    `json:"status"`
		Length           int    `json:"length"`
		Words            int    `json:"words"`
		Lines            int    `json:"lines"`
		ContentType      string `json:"content-type"`
		Redirectlocation string `json:"redirectlocation"`
		Resultfile       string `json:"resultfile"`
		URL              string `json:"url"`
		Host             string `json:"host"`
	} `json:"results"`
}
