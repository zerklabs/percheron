package main

import (
	"github.com/zerklabs/auburn"
	"github.com/zerklabs/percheron"
)

func main() {
	server := auburn.New("127.0.0.1", 8080, "", "", false)

	server.AddRoute("/upload", uploadHandler)
	server.Start()
}

func uploadHandler(req *auburn.HttpTransaction) {
  req.Request.FormFile.
}
