Percheron
==========
# A Distributed Object [File] Store - in Go

## Design

```
             ------------
            |perch daemon|  --> distributed storage (GlusterFS, HDFS, etc..)
proxy -->   |perch daemon|  --> distributed storage (GlusterFS, HDFS, etc..)
            |perch daemon|  --> distributed storage (GlusterFS, HDFS, etc..)
             ------------
```

Where proxy is something on the order of (pick one):
  * [HAProxy](http://haproxy.1wt.eu/)
  * [nginx](http://nginx.org/)


Percheron is designed to provide a simple interface, similar to [S3](http://aws.amazon.com/s3/). The main design goals are:

  * RESTful HTTP interface
  * neutral to the backing file system (and expect neutrality in return)
  * impose a simple, straightforward storage [structure or heiarchy] design


### Identifiers
UUID v4 was chosen as a preventative measure against brute-force guessing of URL paths. Even if a person does not have access to download the file, knowing the path [if it was descriptive] can give someone insight into what kind of information the file contains.

This is however contradicted by the fact we use descriptive names for buckets. The reason for this was to not impose a URL that was completely 'unique', but at least 'readable' by a human.


### Storage Structure

```
  /storage (root of the backing storage)
  |
  |-- /<user id> (user root folder)
  |   |
  |   |-- USERINFO.gob
  |   |-- /<bucket> (bucket root folder)
  |   |   |
  |   |   |-- BUCKET.gob
  |   |   |-- /<object id> (object root folder)
  |   |   |   |
  |   |   |   |-- METADATA.gob
  |   |   |   |-- <object> (file)
  -------------

results in:

/storage/<UUIDv4>/<bucket>/<UUIDv4>/<object>
```

example filesystem paths:

```
/storage/d0bd8bb3-ae40-477e-af4e-7cb9d72e70e0/mybucket/eac590d4-2681-4947-9c1a-26c8e1765da2/myfile.zip
/storage/d0bd8bb3-ae40-477e-af4e-7cb9d72e70e0/mybucket/d9cfbc4e-49d8-4f8d-9973-8f3cecdfc857/myotherfile.zip
```

Where:
  * __d0bd8bb3-ae40-477e-af4e-7cb9d72e70e0__ is the user id
  * __mybucket__ is the bucket
  * __eac590d4-2681-4947-9c1a-26c8e1765da2__ is the object id of myfile.zip
  * __d9cfbc4e-49d8-4f8d-9973-8f3cecdfc857__ is the object id of myotherfile.zip

example URLs:
```
GET    http://storageproxy.example.org/mybucket/eac590d4-2681-4947-9c1a-26c8e1765da2
GET    http://storageproxy.example.org/mybucket/d9cfbc4e-49d8-4f8d-9973-8f3cecdfc857
```

### Metadata
  * [USERINFO.gob](docs/USERINFO.md)
  * [BUCKET.gob](docs/BUCKET.md)
  * [OBJMETADATA.gob](docs/OBJMETADATA.md)


## Constraints
  * [Follow same naming convention for buckets as S3](http://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html)
    * This means bucket names are unique throughout the entire system
