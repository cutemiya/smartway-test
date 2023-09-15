package user

type User struct {
	Name       string `json:"name" example:"Name"`
	Surname    string `json:"surname" example:"Surname"`
	Patronymic string `json:"patronymic" example:"Patronymic"`
}

type TimeDiapason struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
