package controller

import (
	"context"
	"fmt"

	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/longbridgeapp/opencc"
)

type PersonService struct {
	Store    *model.Service
	Renderer Renderer
	opencc   *opencc.OpenCC
}

func NewPersonService(store model.Store, Renderer Renderer) *PersonService {
	cc, err := opencc.New("s2t")
	if err != nil {
		msg := fmt.Sprintf("Failed to initialized opencc: %v", err)
		panic(msg)
	}
	c := &PersonService{
		Store:    model.NewSerice(store),
		Renderer: Renderer,
		opencc:   cc,
	}

	return c
}

func (c *PersonService) PersonByID(ctx context.Context, input model.PersonByIDInput) ([]byte, error) {
	person, err := c.Store.FetchPersonByID(ctx, input)

	if err != nil {
		return []byte{}, err
	}
	return c.Renderer.PersonByID(person, input.Fileds)
}

func (c *PersonService) simplifiedToTraditional(in string) (string, error) {
	out, err := c.opencc.Convert(in)
	if err != nil {
		return "", err
	}
	return out, nil
}

func (c *PersonService) PersonByName(ctx context.Context, name string) ([]byte, error) {
	name, err := c.simplifiedToTraditional(name)
	if err != nil {
		return []byte{}, err
	}

	ps, err := c.Store.FetchPersonByName(ctx, name)
	if err != nil {
		return []byte{}, err
	}
	return c.Renderer.PersonByName(ps)
}
