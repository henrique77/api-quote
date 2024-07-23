package model

type Metrics struct {
	ResultsPerCarrier      map[string]int     `json:"results_per_carrier,omitempty"`
	TotalFinalPrice        map[string]float64 `json:"total_final_price,omitempty"`
	AverageFinalPrice      map[string]float64 `json:"average_final_price,omitempty"`
	LeastExpensiveShipping float64            `json:"least_expensive_shipping,omitempty"`
	MostExpensiveShipping  float64            `json:"most_expensive_shipping,omitempty"`
}

type ResultsPerCarrier struct {
	Name     string
	Quantity int
}

type TotalFinalPrice struct {
	Name  string
	Total float64
}

type AverageFinalPrice struct {
	Name    string
	Average float64
}
