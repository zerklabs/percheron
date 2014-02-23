package percheron

import (
	"bytes"
	"encoding/gob"
	// "github.com/garyburd/redigo/redis"
	"github.com/zerklabs/auburn"
	"log"
	"path/filepath"
	"time"
)

type Object struct {
	Name         string
	Size         int64
	Created      time.Time
	Modified     time.Time
	BucketID     string
	OwnerID      string
	ID           string
	Checksum     []byte
	ChecksumType string
	Path         string
	Extra        map[string]string
	ACL          int
	Grants       map[string]string
}

// TODO(rch): add directory creation
func (bucket *Bucket) NewObject(name string) *Object {
	obj := new(Object)

	obj.Name = name
	obj.Created = time.Now()
	obj.Modified = obj.Created
	obj.OwnerID = bucket.OwnerID
	obj.BucketID = bucket.ID
	obj.ID = auburn.GenStrUUID()
	obj.Path = filepath.Join(bucket.Path, obj.ID)

	return obj
}

func (self *Object) Key() string {
	return fmt.Sprintf("p:o:%s:%s", self.BucketID, self.ID)
}

// Marshal the ObjMetadata struct into a byte array
func (self *Object) Marshal() []byte {
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
func (self *Object) Unmarshal(u []byte) *Object {
	dec := gob.NewDecoder(bytes.NewBuffer(u))

	err := dec.Decode(&self)

	if err != nil {
		log.Fatal("decode:", err)
	}

	return self
}
