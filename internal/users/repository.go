package users

type Repository interface {
	Create(user *UserModel) error
}
