package signup

type mySQLService struct {

}

func newMySQLService() (Service, error) {
	return &mySQLService{}, nil
}

func (m mySQLService) Signup(user User) error {
	return nil
}