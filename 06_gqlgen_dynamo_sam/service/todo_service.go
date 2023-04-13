package service

import (
	"fmt"
	"github.com/lambda-gqgen-presentation/models"
	uuid "github.com/satori/go.uuid"

	"github.com/guregu/dynamo"
)

type TodoService struct {
	db *dynamo.DB
}

func NewTodoService(db *dynamo.DB) *TodoService {
	return &TodoService{db: db}
}

func (d *TodoService) GetTodosByUserId(userId string) ([]*models.Todo, error) {
	fmt.Printf("Getting item from DynamoDB... %s\n", userId)
	var result []*models.Todo
	err := d.db.Table("TODO").Get("userId", userId).
		Index("user_id").
		All(&result)

	return result, err
}

func (d *TodoService) AddTodo(todo models.Todo) (*models.Todo, error) {
	todo.ID = uuid.NewV4().String()
	err := d.db.Table("TODO").Put(todo).Run()
	return &todo, err
}
