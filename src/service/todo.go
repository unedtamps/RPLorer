package service

import (
	"context"

	r "github.com/unedtamps/go-backend/internal/repository"
	"github.com/unedtamps/go-backend/util"
)

type TodoService struct {
	*r.Store
}

type TodoParams struct {
	Title  string
	Desc   string
	UserId string
}

type TodoServiceI interface {
	CreateTodo(context.Context, TodoParams) (*r.CreateTodoRow, error)
	GetTodoByUserId(
		context.Context,
		string,
		int64,
		int64,
	) ([]*r.GetTodoByUserIdRow, *util.MetaData, error)
}

func newTodoService(store *r.Store) *TodoService {
	return &TodoService{
		Store: store,
	}
}

func (t *TodoService) CreateTodo(ctx context.Context, p TodoParams) (*r.CreateTodoRow, error) {
	todo, err := t.Queries.CreateTodo(ctx, r.CreateTodoParams{
		ID:          util.GenerateUUID(),
		UserID:      p.UserId,
		TodoName:    p.Title,
		Description: p.Desc,
	})
	return todo, err
}

func (t *TodoService) GetTodoByUserId(
	ctx context.Context,
	userID string,
	page int64,
	page_size int64,
) ([]*r.GetTodoByUserIdRow, *util.MetaData, error) {

	limit, offset := util.WithPagination(page, page_size)
	todo, err := t.Queries.GetTodoByUserId(ctx, r.GetTodoByUserIdParams{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	})
	metadata := util.WithMetadata(page, int64(len(todo)), nil)
	return todo, &metadata, err
}
