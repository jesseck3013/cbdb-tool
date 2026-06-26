package controller

import (
	"context"
	"strconv"

	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/spf13/pflag"
)

type CLI struct {
	controller *controller
}

func NewCLI(store model.Store, renderer Renderer) *CLI {
	ctrl := newController(store, renderer)

	return &CLI{
		controller: ctrl,
	}
}

func parsePersonByIDInput(id int64, flags *pflag.FlagSet) (model.PersonByIDInput, error) {
	input := model.PersonByIDInput{ID: id}

	all, err := flags.GetBool(string(model.All))
	if err != nil {
		return model.PersonByIDInput{}, err
	}

	if all {
		input.Fileds = model.ALLPersonFileds
		return input, nil
	}

	for _, field := range model.ALLPersonFileds {
		v, err := flags.GetBool(string(field))
		if err != nil {
			return model.PersonByIDInput{}, err
		}

		if v {
			if field == model.All {
				input.Fileds = model.ALLPersonFileds
				break
			}
			input.Fileds = append(input.Fileds, field)
		}
	}

	return input, nil
}

func (c *CLI) personByID(ctx context.Context, id int64, flags *pflag.FlagSet) error {
	input, err := parsePersonByIDInput(id, flags)
	if err != nil {
		return err
	}

	return c.controller.personByID(ctx, input)
}

func (c *CLI) personByName(ctx context.Context, name string) error {
	return c.controller.personByName(ctx, name)
}

// entry point for the CLI's person command
func (c *CLI) Person(ctx context.Context, args []string, flags *pflag.FlagSet) error {
	if len(args) > 0 {
		arg := args[0]
		id, err := strconv.Atoi(arg)
		if err == nil {
			return c.personByID(ctx, int64(id), flags)
		} else {
			return c.personByName(ctx, arg)
		}
	}

	return nil
}
