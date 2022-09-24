package models

import "time"

type Caixinhas []Caixinha

type Caixinha struct {
	ID                  int       `json:"id"`
	Name                string    `json:"name"`
	TotalValue          float64   `json:"total_value"`
	CommunityTotalValue float64   `json:"community_total_value"`
	CommunityTax        float64   `json:"community_tax"`
	MarketTotalValue    float64   `json:"market_total_value"`
	MarketTax           float64   `json:"market_tax"`
	StartDate           string    `json:"start_date"`
	CreatedAtDate       time.Time `json:"created_at"`
	IsCommunity         bool      `json:"is_community"`
}
