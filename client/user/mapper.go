package user

import "smartway-test/model"

func MapToUserServiceModel(userModel User) model.User {
	return model.User{
		Name:       userModel.Name,
		Surname:    userModel.Surname,
		Patronymic: userModel.Patronymic,
	}
}
