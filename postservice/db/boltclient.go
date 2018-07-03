package db

import (
	"github.com/perennial-microservices/blog/dbclient"
	"github.com/boltdb/bolt"
	"fmt"
	"strconv"
	"github.com/perennial-microservices/blog/postservice/model"
	"encoding/json"
	"github.com/manveru/faker"
	"strings"
)

var BucketName = "PostsBucket"

type PostsClient struct {
	dbclient.BoltClient
}

func (pc *PostsClient) Seed() {
	pc.initializeBucket()
	pc.seedAccounts()
}

func (pc *PostsClient) initializeBucket() {
	pc.Client().Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(BucketName))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

func (pc *PostsClient) seedAccounts() {
	f, _ := faker.New("en")
	for i := 0; i < 100; i++ {
		userId := strconv.Itoa(10000 + i)
		key := strconv.Itoa(20000 + i)
		// Create an instance of our Account struct
		acc := model.Post{
			Id: key,
			Title: f.Paragraph(1, false),
			Content: strings.Join(f.Paragraphs(2, false)," "),
			UserId: userId,
		}

		// Serialize the struct to JSON
		jsonBytes, _ := json.Marshal(acc)
		pc.Client().Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(BucketName))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
}

func (pc *PostsClient) GetOne(id string) (interface{}, error) {
	// Allocate an empty Account instance we'll let json.Unmarshal populate for us in a bit.
	account := model.Post{}

	// Read an object from the bucket using boltDB.View
	err := pc.Client().View(func(tx *bolt.Tx) error {
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
		return model.Post{}, err
	}
	// Return the Account struct and nil as error.
	return account, nil
}