package entities

type UserProject struct {
	Id        string `db:"id"`
	UserId    string `db:"user_id"`
	ProjectId string `db:"project_id"`
}

type UserProjectDTO struct {
	ProjectId   string `json:"projectId"`
	ProjectName string `json:"projectName"`
	Role        string `json:"role"`
}
