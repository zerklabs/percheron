OBJECT.gob
===============

```
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

```

* `BucketID`: UUID v4
* `OwnerID`: UUID v4
* `ID`: UUID v4
* `ChecksumType`: (One of) SHA1, SHA256, SHA512
* `Extra`: Map for storing additional information related to the object
* `ACL`: (One of) ACL_PRIVATE (default), ACL_PUBLIC_READ, ACL_PUBLIC_READ_WRITE, ACL_AUTH_READ
* `Grants`: Key: x-perch-grant-`type`, Value: email address
