package percheron

import (
	"bytes"
	"encoding/gob"
	"errors"
	// "github.com/garyburd/redigo/redis"
	"github.com/zerklabs/auburn"
	"log"
	"path/filepath"
	"time"
)

type Bucket struct {
	Name    string
	Created time.Time
	OwnerID string
	ID      string
	Path    string
}

// TODO(rch): add directory creation
func (store *PerchStore) NewBucket(name string, owner *User) (*Bucket, error) {
	bucket := new(Bucket)

	bucket.Name = name
	bucket.Created = time.Now()
	bucket.OwnerID = owner.ID
	bucket.ID = auburn.GenStrUUID()
	bucket.Path = filepath.Join(store.Path, bucket.Name)

	exists, err := DoesDirExist(bucket.Path)

	if err != nil {
		return &Bucket{}, err
	}

	if exists {
		return &Bucket{}, errors.New("Bucket already exists")
	}

	return bucket, nil
}

// Marshal the User struct into a byte array
func (self *Bucket) Marshal() []byte {
	var bin bytes.Buffer

	// Create an encoder and send a value.
	enc := gob.NewEncoder(&bin)
	err := enc.Encode(self)

	if err != nil {
		log.Fatal("encode:", err)
	}

	return bin.Bytes()
}

// Marshal the User struct into a byte array
func (self *Bucket) Unmarshal(u []byte) *Bucket {
	dec := gob.NewDecoder(bytes.NewBuffer(u))

	err := dec.Decode(&self)

	if err != nil {
		log.Fatal("decode:", err)
	}

	return self
}
