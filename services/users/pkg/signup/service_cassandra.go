package signup

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocql/gocql"
)

const (
	Keyspace = "users"
	Table    = "signups"
)

var (
	query = fmt.Sprintf("INSERT INTO %s.%s (email, password) VALUES (?, ?)", Keyspace, Table)
)

type cassandraSignup struct {
	session *gocql.Session
}

func newCassandraService() (Service, error) {
	cluster := gocql.NewCluster(strings.Split(os.Getenv("CASSANDRA_HOSTS"), ",")...)
	cluster.Keyspace = Keyspace
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("CASSANDRA_USERNAME"),
		Password: os.Getenv("CASSANDRA_PASSWORD"),
	}
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create Cassandra cluster session %w", err)
	}
	return &cassandraSignup{session: session}, nil
}

func (cs *cassandraSignup) Signup(user User) error {
	if err := cs.session.Query(query, string(user.Email), string(user.Password)).Exec(); err != nil {
		return fmt.Errorf("failed to inser user: %w", err)
	}
	return nil
}
