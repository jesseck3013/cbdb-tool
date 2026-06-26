package controller

import (
	"context"
	"os"
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

func parseCLIPersonByIDInput(id int64, flags *pflag.FlagSet) (model.PersonByIDInput, error) {
	input := model.PersonByIDInput{ID: id}

	all, err := flags.GetBool(string(model.All))
	if err != nil {
		return model.PersonByIDInput{}, err
	}

	if all {
		input.Fileds = model.GetAllPersonFieldsSlice()
		return input, nil
	}

	for field, _ := range model.ALLPersonFileds {
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

func (c *CLI) personByID(ctx context.Context, id int64, flags *pflag.FlagSet) error {
	input, err := parseCLIPersonByIDInput(id, flags)
	if err != nil {
		return err
	}

	return c.controller.personByID(ctx, os.Stdout, input)
}

func (c *CLI) personByName(ctx context.Context, name string) error {
	return c.controller.personByName(ctx, os.Stdout, name)
}

// entry point for the CLI's person command
func (c *CLI) Person(ctx context.Context, args []string, flags *pflag.FlagSet) error { //
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
