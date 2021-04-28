package signup

import (
	"errors"
	"fmt"
	"strings"
)

var (
	InvalidUser = errors.New("invalid user")
)

type Service interface {
	Signup(user User) error
}

type ServiceType string

const (
	ServiceTypeCassandra ServiceType = "cassandra"
	ServiceTypeMySQL     ServiceType = "mysql"
)

func NewService(t ServiceType) (Service, error) {
	t = ServiceType(strings.ToLower(string(t)))
	if t == ServiceTypeCassandra {
		return newCassandraService()
	} else if t == ServiceTypeMySQL {
		return newMySQLService()
	}
	return nil, UnsupportedServiceType{Type: t}
}

type UnsupportedServiceType struct {
	Type ServiceType
}

func (t UnsupportedServiceType) Error() string {
	return fmt.Sprint("unsupported service type", t.Type)
}
