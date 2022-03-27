package store

import (
	"bytes"
	"encoding/gob"
	"errors"
	"time"

	"github.com/boltdb/bolt"
)

var (
	storeName = "store"

	ErrNotFound = errors.New("key not found")
	ErrBadValue = errors.New("bad value")
)

var DB *bolt.DB

func Open(path string) error {
	opts := &bolt.Options{
		Timeout: 50 * time.Millisecond,
	}
	if db, err := bolt.Open(path, 0o640, opts); err != nil {
		return err
	} else {
		err := db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(storeName))
			return err
		})
		if err != nil {
			return err
		} else {
			DB = db
			return nil
		}
	}
}

func Close() error {
	return DB.Close()
}

func Get(key string, value interface{}) error {
	return DB.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(storeName)).Cursor()
		if k, v := c.Seek([]byte(key)); k == nil || string(k) != key {
			return ErrNotFound
		} else if value == nil {
			return nil
		} else {
			d := gob.NewDecoder(bytes.NewReader(v))
			return d.Decode(value)
		}
	})
}

func Put(key string, value interface{}) error {
	if value == nil {
		return ErrBadValue
	}
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(value); err != nil {
		return err
	}
	return DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(storeName)).Put([]byte(key), buf.Bytes())
	})
}

func Delete(key string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		if key == "root" {
			return nil
		}

		c := tx.Bucket([]byte(storeName)).Cursor()
		if k, _ := c.Seek([]byte(key)); k == nil || string(k) != key {
			return ErrNotFound
		} else {
			return c.Delete()
		}
	})
}

func Clear() error {
	return DB.Update(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(storeName)).Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			if string(k) == "root" {
				continue
			}

			err := c.Delete()
			if err != nil {
				return err
			}
		}

		return nil
	})
}
