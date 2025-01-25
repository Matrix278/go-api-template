package repository

type Repositories struct {
	User *User
}

func NewRepositories(connection *Connection) *Repositories {
	return &Repositories{
		User: NewUser(connection.db),
	}
}
