package percheron

import (
	"bytes"
	"encoding/gob"
	// "github.com/garyburd/redigo/redis"
	"github.com/nu7hatch/gouuid"
	"github.com/zerklabs/auburn"
	"log"
	"time"
)

type User struct {
	Email    string
	Created  time.Time
	ID       *uuid.UUID
	AccessID string
}

// TODO(rch): add directory creation
func (store *PerchStore) NewUser(email string) *User {
	user := new(User)

	user.Email = email
	user.Created = time.Now()
	user.ID = auburn.GenUUID()

	return user
}

// Marshal the User struct into a byte array
func (self *User) Marshal() []byte {
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
func (self *User) Unmarshal(u []byte) *User {
	dec := gob.NewDecoder(bytes.NewBuffer(u))

	err := dec.Decode(&self)

	if err != nil {
		log.Fatal("decode:", err)
	}

	return self
}
