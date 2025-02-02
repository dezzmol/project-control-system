package entities

import "time"

type BugReport struct {
	Id          string    `db:"id"`
	Description string    `db:"description"`
	DateCreated time.Time `db:"date_created"`
	Status      string    `db:"status"`
}

type BugReportDTO struct {
	BugReportId string    `json:"bugReportId"`
	Description string    `json:"description"`
	ReportDate  time.Time `json:"reportDate"`
}

type CreateBugReportDTO struct {
	Description string `json:"description"`
}

type UpdateBugReportDTO struct {
	Status string `json:"status"`
}
