package percheron

import (
	"github.com/cabrel/auburn"
	"github.com/nu7hatch/gouuid"
	"log"
	"net"
	"time"
)

type User struct {
	Email    string
	Created  time.Time
	ID       *uuid.UUID
	AccessID string
}

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
	Path  string
	Peers []net.IPAddr
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
func (store *PerchStore) NewUserInfo(email string) (*User, error) {
	user := new(User)

	user.Email = email
	user.Created = time.Now()
	user.ID = auburn.GenUUID()

	return user, nil
}

// TODO(rch): add directory creation
func (bucket *Bucket) NewObject(name string) (*ObjMetadata, error) {
	obj := new(ObjMetadata)

	obj.Name = name
	obj.Created = time.Now()
	obj.Created = obj.Created
	obj.Owner = bucket.Owner
	obj.ID = auburn.GenUUID()

	return obj, nil
}

func NewPerchStore(folderPath string) *PerchStore {
	yes, err := DoesDirExist(folderPath)

	if err != nil {
		log.Fatal(err)
	}

	if yes {
		// no error means the stat returned successfully
		return &PerchStore{Path: folderPath}
	}

	log.Fatalf("%s does not exist or cannot be accessed", folderPath)

	return &PerchStore{}
}
