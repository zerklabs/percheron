package percheron

import (
	"bytes"
	"encoding/gob"
	// "github.com/garyburd/redigo/redis"
	"github.com/nu7hatch/gouuid"
	"github.com/zerklabs/auburn"
	"log"
	"net"
	"time"
)

type Bucket struct {
	Name    string
	Created time.Time
	Owner   *uuid.UUID
	ID      *uuid.UUID
}

type ObjMetadata struct {
	Name     string
	Size     int64
	Created  time.Time
	Modified time.Time
	Owner    *uuid.UUID
	ID       *uuid.UUID
	Checksum []byte
	HashType string
}

type PerchStore struct {
	Path      string
	RedisHost string
	RedisPort int
	Peers     []net.IPAddr
}

// TODO(rch): add directory creation
func (user *User) NewBucket(name string) (*Bucket, error) {
	bucket := new(Bucket)

	bucket.Name = name
	bucket.Created = time.Now()
	bucket.Owner = user.ID
	bucket.ID = auburn.GenUUID()

	return bucket, nil
}

// TODO(rch): add directory creation
func (bucket *Bucket) NewObject(name string) *ObjMetadata {
	obj := new(ObjMetadata)

	obj.Name = name
	obj.Created = time.Now()
	obj.Modified = obj.Created
	obj.Owner = bucket.Owner
	obj.ID = auburn.GenUUID()

	return obj
}

func NewPerchStore(folderPath string) *PerchStore {
	yes, err := DoesDirExist(folderPath)

	if err != nil {
		log.Fatal(err)
	}

	if yes {
		// no error means the stat returned successfully
		store := new(PerchStore)
		store.Path = folderPath

		return store
	}

	log.Fatalf("%s does not exist or cannot be accessed", folderPath)

	return new(PerchStore)
}

// Marshal the ObjMetadata struct into a byte array
func (self *ObjMetadata) Marshal() []byte {
	var bin bytes.Buffer

	// Create an encoder and send a value.
	enc := gob.NewEncoder(&bin)
	err := enc.Encode(self)

	if err != nil {
		log.Fatal("encode:", err)
	}

	return bin.Bytes()
}
