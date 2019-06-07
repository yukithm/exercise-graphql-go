package exercise_graphql_go

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

const (
	rootDir = "."
)

type Resolver struct {
	todos []*Todo
}

func (r *Resolver) File() FileResolver {
	return &fileResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type fileResolver struct{ *Resolver }

func (r *fileResolver) Entries(ctx context.Context, obj *File) ([]*File, error) {
	if !obj.IsDir {
		return nil, nil
	}

	log.Printf("Resolve File.entries: %s", obj.Path)
	stats, err := ioutil.ReadDir(obj.Path)
	if err != nil {
		return nil, err
	}

	entries := make([]*File, 0, len(stats))
	for _, stat := range stats {
		entries = append(entries, statToFile(obj.Path, stat))
	}

	return entries, nil
}

func statToFile(dir string, stat os.FileInfo) *File {
	return &File{
		Path:  filepath.Join(dir, stat.Name()),
		Size:  stat.Size(),
		IsDir: stat.IsDir(),
	}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	todo := &Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	return r.todos, nil
}
func (r *queryResolver) Root(ctx context.Context) (*File, error) {
	stat, err := os.Stat(rootDir)
	if err != nil {
		return nil, err
	}

	return statToFile(rootDir, stat), nil
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *Todo) (*User, error) {
	return &User{
		ID:   obj.UserID,
		Name: "user " + obj.UserID,
	}, nil
}
