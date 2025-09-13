package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

// SELECT * from users;
func GetAllUsers() *[]User {
	var users []User
	// DB.Find(&users)
	DB.Raw("select * from users").Scan(&users)
	return &users
}

// SELECT * from users WHERE id = ?;
func FindUserByID(id int) *User {
	var user User
	// DB.First(&user, id)
	DB.Raw("select * from users where id = ?", id).Scan(&user)
	return &user
}

func (u *User) Create() {
	DB.Create(&u)
}

// INSERT INTO users(name, age) VALUES(?, ?);
func (u *User) Add(name string, age int) {
	// user := User{
	// 	Name: name,
	// 	Age:  age,
	// }
	// DB.Create(&user)
	DB.Raw("insert into users(name, age) values(?, ?)", name, age).Scan(&u)
}

// UPDATE users SET name = ?, age = ?;
func (u *User) Update() {
	user := User{
		Name: u.Name,
		Age:  u.Age,
	}
	DB.Model(&u).Updates(user)

	// DB.Raw("update users set name = ?, age = ?", name, age).Scan(&u)
}

// DELETE FROM users;
func (u *User) Delete() {
	DB.Delete(&u)

}
