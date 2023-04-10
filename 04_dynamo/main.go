package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	ID     string `json:"id" dynamo:"id,hash"`
	UserID string `json:"userId" dynamo:"userId,range,required"`
	Text   string `json:"text" dynamo:"text,required"`
	Done   bool   `json:"done" dynamo:"done"`
}

func main() {
	fmt.Println("Connecting to DynamoDB...")
	//sess := session.Must(session.NewSession())
	sess, err := session.NewSession()
	db := dynamo.New(sess)
	fmt.Println("Connected to DynamoDB...")
	table := db.Table("TODO")

	// put item
	todo := Todo{
		ID:     uuid.NewV4().String(),
		Text:   "Todo " + uuid.NewV4().String(),
		Done:   false,
		UserID: "1",
	}

	err = table.Put(todo).Run()
	fmt.Println("Put item into DynamoDB...")

	if err != nil {
		panic(err)
	}

	fmt.Println("Getting item from DynamoDB...")

	// get the same item
	var result []Todo
	err = table.Get("userId", todo.UserID).
		Index("user_id").
		All(&result)

	for _, item := range result {
		fmt.Println(item.Text)
	}

	fmt.Println("Got item from DynamoDB...")
	if err != nil {
		panic(err)
	}
}
