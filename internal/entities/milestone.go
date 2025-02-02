package entities

import "time"

type Milestone struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
}

type CreateMilestoneDTO struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type MilestoneDTO struct {
	MilestoneId string    `json:"milestoneId"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}

type MilestoneUpdateStatusDTO struct {
	Status string `json:"status"`
}
