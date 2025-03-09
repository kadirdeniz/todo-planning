package provider

import "time"

type Provider1Response struct {
	ID int `json:"id"`
	Value int `json:"value"`
	EstimatedDuration time.Duration `json:"estimated_duration"`
}

type Provider2Response struct {
	ID int `json:"id"`
	Value int `json:"zorluk"`
	EstimatedDuration time.Duration `json:"sure"`
}	