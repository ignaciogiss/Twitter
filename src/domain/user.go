package domain

// TODO de todo y password wrap
type User struct {
	Name string
	Mail string
	Nick string
	password string
}

func NewUser(name string, mail string, nick string, password string) *User {

	return &User{name, mail, nick, password  }
}


func (user User)PasswordOk ( pass string ) bool {
	if user.password == pass {
		return true // TODO
	}
	return false
}