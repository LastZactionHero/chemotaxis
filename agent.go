package main

import "github.com/jinzhu/gorm"

// MaxPastValues maximum number of prior valuations to remember
const MaxPastValues = 7

// Agent is a unit making stock decisions
type Agent struct {
	CashOnHand float64
	Holding    Holding
	PastValues [MaxPastValues]float64
}

// Value of the agent's holdings
func (a *Agent) Value(db *gorm.DB) float64 {
	return a.Holding.Value(db)
}

// agent
// cash on hand
// holding
// past values
// +die
// +purchase
// +sell
// +decide
