package main

import (
	"log"
	"os"

	userpb "github.com/anhbkpro/videos/2022/07/23/gen/go/user/v1"

	"google.golang.org/protobuf/proto"
)

func main() {
	u := userpb.User{
		Uuid:      "1-2-3-4",
		FullName:  "Anh",
		BirthYear: 2999,
	}

	b, err := proto.Marshal(&u)
	if err != nil {
		log.Fatalln("Marshaling error", err)
	}

	if err := os.WriteFile("user.bin", b, 0644); err != nil {
		log.Fatalln("Writing error", err)
	}
}
