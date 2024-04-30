package user_domain

type getUser struct {
	Name string
}

func (u getUser) Handle() string {
	return u.Name
}

func GetUserFactory() getUser {
	return getUser{
		Name: "User's name",
	}
}
