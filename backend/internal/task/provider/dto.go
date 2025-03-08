package provider

import "time"

type Provider1Response struct {
	ID int `json:"id"`
	Value string `json:"value"`
	EstimatedDuration time.Duration `json:"estimated_duration"`
}

type Provider2Response struct {
	ID int `json:"id"`
	Value string `json:"zorluk"`
	EstimatedDuration time.Duration `json:"sure"`
}	