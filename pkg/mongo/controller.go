package mongo

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

var Repo *Repository

type Repository struct {
	*mgo.Session
	DatabaseName string
}

func NewStorage(dburl string, database string, username string, password string) (*Repository, error) {
	info := &mgo.DialInfo{
		Addrs:    []string{dburl},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}
	if Repo == nil {
		session, err := mgo.DialWithInfo(info)
		if err != nil {
			return nil, err
		}
		Repo = &Repository{
			Session:      session,
			DatabaseName: database,
		}
	}
	return Repo, nil
}

func (r *Repository) Close() error {
	r.Session.Close()
	return nil
}
