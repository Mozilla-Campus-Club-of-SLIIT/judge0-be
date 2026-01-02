package types

import "time"

type User struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	FName     string    `json:"fName"`
	LName     string    `json:"lName"`
	Email     string    `json:"email"`
	RegNumber string    `json:"regNumber"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type Challenge struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	SampleInput  string    `json:"sample_input"`
	SampleOutput string    `json:"sample_output"`
	CreatedAt    time.Time `json:"created_at"`
}
