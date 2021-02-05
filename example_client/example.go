package main

import (
	"context"
	"fmt"
	"os"
	"reflect"

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

	getPrep, err := client.UserProperties.PrepareGet("str_property")
	fmt.Printf("%+v\n", getPrep)
	fmt.Printf("%+v\n", err)

	getResp, err := client.Send(context.Background(), getPrep)
	fmt.Println("type: ", reflect.TypeOf(getResp))
	fmt.Printf("%+v\n", getResp)
	fmt.Printf("%+v\n", err)
	fmt.Println("type: ", reflect.TypeOf(getPrep.Target))
	fmt.Printf("%+v\n", getPrep.Target)

	fmt.Println("listPrep:")
	listPrep, err := client.UserProperties.PrepareList()
	fmt.Printf("%+v\n", listPrep)
	fmt.Printf("%+v\n", err)
	fmt.Println("type: ", reflect.TypeOf(listPrep.Target))
	fmt.Printf("%+v\n", listPrep.Target)

	fmt.Println("listResp:")
	listResp, err := client.Send(context.Background(), listPrep)
	fmt.Println("type: ", reflect.TypeOf(listResp))
	fmt.Printf("%+v\n", listResp)
	fmt.Printf("%+v\n", err)
	fmt.Println("type: ", reflect.TypeOf(listPrep.Target))
	fmt.Printf("%+v\n", listPrep.Target)

	fmt.Println("POST DV:")
	dv := &binapi.DetailView{
		Timestamp:     "Fri 05 Feb 2021 02:42:24 PM CET",
		Duration:      120,
		CascadeCreate: true,
		UserId:        "bobek",
		ItemId:        "item001",
	}
	// resp, err = client.DetailViews.Post(context.Background(), dv)
	// fmt.Printf("%+v\n", resp)
	// fmt.Printf("%+v\n", err)

	dvPrep, err := client.DetailViews.PreparePost(dv)
	fmt.Printf("%+v\n", err)
	dvResp, err := client.Send(context.Background(), dvPrep)
	fmt.Println("type: ", reflect.TypeOf(dvResp))
	fmt.Printf("%+v\n", dvResp)
	fmt.Printf("%+v\n", err)

}
