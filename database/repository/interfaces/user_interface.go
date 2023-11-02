package interfaces

type IUserRepo interface {
	CreateUser() error
	ReadUser() error
	UpdateUser() error
	DeleteUser() error
}
