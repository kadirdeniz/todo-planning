package model

import "time"

type Task struct {
	ID        int  
	Value     string
	EstimatedDuration time.Duration
}
