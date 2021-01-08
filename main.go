package main

import (
	"context"
	"fmt"
	"log"

	g "google.golang.org/grpc"

	"github.com/WeFolks/search_service/grpc"
)

func main() {
	var conn *g.ClientConn
	conn, err := g.Dial(":9000", g.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	defer conn.Close()

	c := grpc.NewSearchServiceClient(conn)

	request := grpc.SearchRequest{
		Name:     "folks",
		Category: "",
	}

	response, err := c.GetItems(context.Background(), &request)

	// item := grpc.Item{
	// 	Id:          "12345",
	// 	Name:        "folks",
	// 	Owner:       "",
	// 	Category:    "",
	// 	Description: "",
	// 	Type:        1,
	// }

	// _, err = c.AddItem(context.Background(), &item)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	items := response.Items

	for _, element := range items {
		fmt.Println(element.Name)
		fmt.Println(element.Description)
	}
}
