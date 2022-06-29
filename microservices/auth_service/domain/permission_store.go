package domain

type PermissionStore interface {
	
	Get(id primitive.ObjectID) (*Permission, error)
	GetByRole(role Role) ([]*Permission, error)
	GetAll() ([]*Permission, error)
	Insert(user *Permission) (*Permission, error)
	Update(user *Permission) (*Permission, error)
	DeleteAll()
}
