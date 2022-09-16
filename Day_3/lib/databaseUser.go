package database

import (
	mysql "Day2Assignment/config"
	"Day2Assignment/models"
)

func CreateUser(user models.User) error {
	result := mysql.DB.Create(&user)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(id string, user models.User) error {
	if err := mysql.DB.Model(user).Where("id = ?", id).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id string) error{
	var user models.User
	if err := mysql.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(id string) (models.User, error) {
	var user models.User
	result := mysql.DB.Where("id = ?", id).First(&user)
	return user, result.Error
}


func GetAllUser() ([]models.User, error) {
	var users []models.User
	result := mysql.DB.Find(&users)

	return users, result.Error
}