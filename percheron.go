package percheron

import (
	"github.com/nu7hatch/gouuid"
	"time"
)

type User struct {
	Email   string
	Created time.Time
	ID      *uuid.UUID
}

type Bucket struct {
	Name    string
	Created time.Time
	Owner   *uuid.UUID
	ID      *uuid.UUID
}

type ObjMetadata struct {
	Name         string
	Size         int64
	Created      time.Time
	Modified     time.Time
	Owner        uuid.UUID
	ID           uuid.UUID
	Checksum     string
	ChecksumHash string
}

// TODO(rch): add directory creation
func (user *User) NewBucket(name string) (*Bucket, error) {
	id, err := uuid.NewV4()

	if err != nil {
		return &Bucket{}, err
	}

	bucket := new(Bucket)

	bucket.Name = name
	bucket.Created = time.Now()
	bucket.Owner = user.ID
	bucket.ID = id

	return bucket, nil
}

// TODO(rch): add directory creation
func NewUserInfo(email string) (*User, error) {
	id, err := uuid.NewV4()

	if err != nil {
		return &User{}, err
	}

	user := new(User)

	user.Email = email
	user.Created = time.Now()
	user.ID = id

	return user, nil
}

// TODO(rch): add directory creation
func (bucket *Bucket) NewObject(name string) (*ObjMetadata, error) {
	id, err := uuid.NewV4()

	if err != nil {
		return &ObjMetadata{}, err
	}

	obj := new(ObjMetadata)

	obj.Name = name
	obj.Created = time.Now()
	obj.Created = obj.Created
	obj.Owner = bucket.Owner
	obj.ID = id

	return obj, nil
}
