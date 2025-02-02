package repository

type Repositories struct {
	User IUser
}

func NewRepositories(connection *Connection) *Repositories {
	return &Repositories{
		User: NewUser(connection.db),
	}
}
