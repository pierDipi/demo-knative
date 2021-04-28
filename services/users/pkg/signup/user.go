package signup

type Email string

type Password string

type User struct {
	Email    Email    `json:"email"`
	Password Password `json:"password"`
}

func (u User) Validate() error {
	// TODO validate
	return nil
}
