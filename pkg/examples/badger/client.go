package badger

import (
	"fmt"

	badger "github.com/dgraph-io/badger/v3"
)

const (
	path = "./bin/badger/data"
)

type (
	Client interface {
		Update(key string, value string) error
		View(key string) (*badger.Item, error)
		Close()
	}
	client struct {
		db *badger.DB
	}
)

func NewClient() Client {
	fmt.Println("---------------New Badger Client-------------------")

	// open a Badger database
	opts := badger.DefaultOptions(path)
	db, err := badger.Open(opts)

	if err != nil {
		fmt.Println("Error opening Badger database:", err)
		return nil
	}

	return &client{
		db,
	}
}

func (c *client) Close() {
	fmt.Println("---------------Closing Badger Client---------------")
	c.db.Close()
}

func (c *client) Update(key string, value string) error {
	// write a key-value pair to the database
	err := c.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})

	if err != nil {
		fmt.Println("Error writing to database:", err)
		return err
	}
	return nil
}

func (c *client) View(key string) (*badger.Item, error) {
	// read a value from the database
	err := c.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			fmt.Printf("Value: %s\n", val)
			return nil
		})
		return err
	})

	if err != nil {
		fmt.Println("Error reading from database:", err)
		return nil, nil
	}
	return nil, nil
}
