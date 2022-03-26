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

type Manager struct {
	store *bolt.DB
}

func New(path string) (*Manager, error) {
	opts := &bolt.Options{
		Timeout: 50 * time.Millisecond,
	}
	if db, err := bolt.Open(path, 0640, opts); err != nil {
		return nil, err
	} else {
		err := db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(storeName))
			return err
		})
		if err != nil {
			return nil, err
		} else {
			return &Manager{store: db}, nil
		}
	}
}

func (m *Manager) Close() error {
	return m.store.Close()
}

func (m *Manager) get(key string, value interface{}) error {
	return m.store.View(func(tx *bolt.Tx) error {
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

func (m *Manager) put(key string, value interface{}) error {
	if value == nil {
		return ErrBadValue
	}
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(value); err != nil {
		return err
	}
	return m.store.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(storeName)).Put([]byte(key), buf.Bytes())
	})
}

func (m *Manager) delete(key string) error {
	return m.store.Update(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(storeName)).Cursor()
		if k, _ := c.Seek([]byte(key)); k == nil || string(k) != key {
			return ErrNotFound
		} else {
			return c.Delete()
		}
	})
}
