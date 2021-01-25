package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bobek/bin-api-go/v1/binapi"
)

// http://requestbin.net/r/1io9ix61

func main() {
	println("Testing...")

	client := binapi.NewClient(nil, "client-test", os.Getenv("HMAC_TOKEN"))
	fmt.Println(client)

	//user, resp, _ := client.Users.Get(context.Background(), "bobek")
	//fmt.Println(user)
	//fmt.Println(resp)

	ups, resp, _ := client.UserProperties.List(context.Background())
	fmt.Println(ups)
	fmt.Println(resp)

	up, resp, err := client.UserProperties.Get(context.Background(), "str_property")
	fmt.Println(up)
	fmt.Printf("%+v\n", resp)
	fmt.Printf("%+v\n", err)

	upn, resp, err := client.UserProperties.Get(context.Background(), "non_existing")
	fmt.Println(upn)
	fmt.Printf("%+v\n", resp)
	fmt.Printf("%+v\n", err)
}
