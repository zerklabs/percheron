package main

import (
	"flag"
	"github.com/zerklabs/auburn"
	"github.com/zerklabs/percheron"
	"log"
	"runtime"
)

var (
	listenIP        = flag.String("host", "", "IP to run the webserver on")
	listenOn        = flag.Int("port", 8080, "Port to run the webserver on")
	certificatePath = flag.String("cert", "", "Certificate file for TLS (.pem) (Optional)")
	keyPath         = flag.String("key", "", "Private key for certificate (Required if cert given)")
	storePath       = flag.String("path", "", "Root of storage path")

	store   *percheron.PerchStore
	users   []percheron.User
	objects []percheron.Object
	buckets []percheron.Bucket
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// bind the command line flags
	flag.Parse()

	server := auburn.New(*listenIP, *listenOn, *certificatePath, *keyPath)
	store = percheron.NewPerchStore(*storePath)

	go fetchUsers()
	go fetchBuckets()
	go fetchObjects()

	server.Handle("/lookup/users", lookupUsers)
	server.Handle("/lookup/objects", lookupObjects)
	server.Handle("/lookup/buckets", lookupBuckets)

	server.Start()
}

func fetchUsers() {
}

func fetchBuckets() {
}

func fetchObjects() {
}

func lookupUsers(req *auburn.AuburnHttpRequest) {

}

func lookupObjects(req *auburn.AuburnHttpRequest) {

}

func lookupBuckets(req *auburn.AuburnHttpRequest) {

}
