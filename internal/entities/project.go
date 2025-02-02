package entities

type Project struct {
	Id          string `db:"id"`
	ProjectName string `db:"project_name"`
	Description string `db:"description"`
}

type AssignUserDTO struct {
	UserId string `json:"userId"`
}

type CreateProjectDTO struct {
	ProjectName string `json:"projectName"`
	Description string `json:"description"`
}

type ProjectDTO struct {
	ProjectId string `json:"projectId"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
}
