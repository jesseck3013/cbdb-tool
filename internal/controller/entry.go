package controller

import (
	"context"
	"strconv"

	"github.com/jesseck3013/cbdb-tool/internal/formatter"
	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/spf13/pflag"
)

type CLI struct {
	Store    *model.Service
	Renderer Renderer
}

func NewController(store model.Store, Renderer Renderer) *CLI {
	c := &CLI{
		Store:    model.NewSerice(store),
		Renderer: Renderer,
	}

	return c
}

func parsePersonByIDInput(id int64, flags *pflag.FlagSet) (model.PersonByIDInput, error) {
	input := model.PersonByIDInput{ID: id}

	for _, field := range model.ALLPersonFileds {
		v, err := flags.GetBool(string(field))
		if err != nil {
			return model.PersonByIDInput{}, err
		}

		if v {
			input.Fileds = append(input.Fileds, field)
		}
	}

	return input, nil
}

// 1. process the input
// 2. query data based on the input
// 3. render the input
func (c *CLI) Person(ctx context.Context, args []string, flags *pflag.FlagSet) error {
	if len(args) > 0 {
		idStr := args[0]
		id, err := strconv.Atoi(idStr)
		if err == nil {
			input, err := parsePersonByIDInput(int64(id), flags)

			if err != nil {
				return err
			}

			person, err := c.Store.FetchPersonByID(ctx, input)

			if err != nil {
				return err
			}
			formatter.PrintPerson(person)
		}
	}

	return nil
}
