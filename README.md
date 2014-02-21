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


### Storage Layout

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

### Object Storage
Each object is stored in 4MB chunks. If the object itself is less than or equal to 4MB then it is stored in a single file. Each chunk is named to include it's position in the actual file and the checksum of the chunk.

For perfect files (those less than or equal to 4MB), an example name would be:

0_f66c2834db5ad832ca3d31fdfae504ae07e9c95f4cf2e6beb8670b27961de45a

For larger files:

0_89dbf8cdd390aba7dcdb648553854ca7952d4bfe831044cb27a1161b4c6e5198
1_1115a3c7a7c6a2cc08ef3f3dfb429cd427bb9eb43ffe2c9efe51a938a3e249e5
[...]

### Metadata
  * [USERINFO.gob](docs/USERINFO.md)
  * [BUCKET.gob](docs/BUCKET.md)
  * [OBJMETADATA.gob](docs/OBJMETADATA.md)


## Constraints
  * [Follow same naming convention for buckets as S3](http://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html)
    * This means bucket names are unique throughout the entire system
