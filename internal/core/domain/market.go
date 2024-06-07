package domain

type Market struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Status   int    `json:"status"`
	Outcomes []MarketOutcome
}

type MarketOutcome struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}
