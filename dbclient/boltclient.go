package dbclient

import (
	"log"
	"github.com/boltdb/bolt"
)

type IBoltClient interface {
	OpenBoltDb(databaseName string)
	Seed()
	Client() (*bolt.DB)
	Check() (bool)
	GetOne(id string) (interface{}, error)
}

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb(databaseName string) {
	var err error
	bc.boltDB, err = bolt.Open(databaseName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BoltClient) Client() (*bolt.DB) {
	return bc.boltDB
}

func (bc *BoltClient) Check() (bool) {
	return bc.boltDB != nil
}