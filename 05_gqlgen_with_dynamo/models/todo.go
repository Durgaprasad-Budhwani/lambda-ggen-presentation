package models

type Todo struct {
	ID     string `json:"id" dynamo:"id,hash"`
	UserID string `json:"userId" dynamo:"userId,range,required"`
	Text   string `json:"text" dynamo:"text,required"`
	Done   bool   `json:"done" dynamo:"done"`
}
