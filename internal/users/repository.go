package users

type Repository interface {
	Create(user *UserModel) error
	FindOneUser(condition interface{}) (*UserModel, error)
}
