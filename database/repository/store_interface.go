package repository

type Store interface {
	CreateUser() error
	ReadUser() error
	UpdateUser() error
	DeleteUser() error
}
