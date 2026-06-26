package controller

import (
	"context"
	"fmt"
	"io"

	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/longbridgeapp/opencc"
)

type controller struct {
	Store    *model.Service
	Renderer Renderer
	opencc   *opencc.OpenCC
}

func newController(store model.Store, Renderer Renderer) *controller {
	cc, err := opencc.New("s2t")
	if err != nil {
		msg := fmt.Sprintf("Failed to initialized opencc: %v", err)
		panic(msg)
	}
	c := &controller{
		Store:    model.NewSerice(store),
		Renderer: Renderer,
		opencc:   cc,
	}

	return c
}

func (c *controller) personByID(ctx context.Context, w io.Writer, input model.PersonByIDInput) error {
	person, err := c.Store.FetchPersonByID(ctx, input)

	if err != nil {
		return err
	}
	return c.Renderer.PersonByID(w, person, input.Fileds)
}

func (c *controller) simplifiedToTraditional(in string) (string, error) {
	out, err := c.opencc.Convert(in)
	if err != nil {
		return "", err
	}
	return out, nil
}

func (c *controller) personByName(ctx context.Context, w io.Writer, name string) error {
	name, err := c.simplifiedToTraditional(name)
	if err != nil {
		return err
	}

	ps, err := c.Store.FetchPersonByName(ctx, name)
	if err != nil {
		return err
	}
	c.Renderer.PersonByName(w, ps)
	return nil
}
