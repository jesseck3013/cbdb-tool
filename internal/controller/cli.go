package controller

import (
	"context"
	"os"
	"strconv"

	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/spf13/pflag"
)

type CLI struct {
	controller Controller
}

func NewCLI(c Controller) *CLI {
	return &CLI{
		controller: c,
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

	b, err := c.controller.personByID(ctx, input)
	if err != nil {
		return err
	}

	os.Stdout.Write(b)
	return nil
}

func (c *CLI) personByName(ctx context.Context, name string) error {
	b, err := c.controller.personByName(ctx, name)
	if err != nil {
		return err
	}
	os.Stdout.Write(b)
	return nil
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
