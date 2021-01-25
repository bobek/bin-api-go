package main

import (
	"context"
	"fmt"

	"github.com/bobek/bin-api-go/v1/binapi"
)

// http://requestbin.net/r/1io9ix61

func main() {
	println("Testing...")

	client := binapi.NewClient(nil)
	fmt.Println(client)

	user, resp, _ := client.Users.Get(context.Background(), "bobek")
	fmt.Println(user)
	fmt.Println(resp)
}
