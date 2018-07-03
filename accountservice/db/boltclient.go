package db

import (
	"github.com/perennial-microservices/blog/dbclient"
	"github.com/boltdb/bolt"
	"fmt"
	"strconv"
	"github.com/perennial-microservices/blog/accountservice/model"
	"encoding/json"
)

var BucketName = "AccountsBucket"

type AccountsClient struct {
	dbclient.BoltClient
}

func (ac *AccountsClient) Seed() {
	ac.initializeBucket()
	ac.seedAccounts()
}

func (ac *AccountsClient) initializeBucket() {
	ac.Client().Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(BucketName))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

func (ac *AccountsClient) seedAccounts() {
	for i := 0; i < 100; i++ {
		key := strconv.Itoa(10000 + i)

		// Create an instance of our Account struct
		acc := model.Account{
			Id: key,
			Name: "Person_" + strconv.Itoa(i),
		}

		// Serialize the struct to JSON
		jsonBytes, _ := json.Marshal(acc)
		ac.Client().Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(BucketName))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
}

func (ac  *AccountsClient) GetOne(id string) (interface{}, error) {
	// Allocate an empty Account instance we'll let json.Unmarshal populate for us in a bit.
	account := model.Account{}

	// Read an object from the bucket using boltDB.View
	err := ac.Client().View(func(tx *bolt.Tx) error {
		// Read the bucket from the DB
		b := tx.Bucket([]byte(BucketName))

		// Read the value identified by our resourceId supplied as []byte
		accountBytes := b.Get([]byte(id))
		if accountBytes == nil {
			return fmt.Errorf("No account found for " + id)
		}
		// Unmarshal the returned bytes into the account struct we created at
		// the top of the function
		json.Unmarshal(accountBytes, &account)

		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	// If there were an error, return the error
	if err != nil {
		return model.Account{}, err
	}
	// Return the Account struct and nil as error.
	return account, nil
}