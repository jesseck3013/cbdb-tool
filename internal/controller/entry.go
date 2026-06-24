package controller

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jesseck3013/cbdb-tool/internal/model"
	"github.com/longbridgeapp/opencc"
	"github.com/spf13/pflag"
)

type CLI struct {
	Store    *model.Service
	Renderer Renderer
	cc       *opencc.OpenCC
}

func NewController(store model.Store, Renderer Renderer) *CLI {
	cc, err := opencc.New("s2t")
	if err != nil {
		msg := fmt.Sprintf("Failed to initialized opencc: %v", err)
		panic(msg)
	}
	c := &CLI{
		Store:    model.NewSerice(store),
		Renderer: Renderer,
		cc:       cc,
	}

	return c
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

	person, err := c.Store.FetchPersonByID(ctx, input)

	if err != nil {
		return err
	}
	return c.Renderer.PersonByID(os.Stdout, person, input.Fileds)
}

func simplifiedToTraditional(cc *opencc.OpenCC, in string) (string, error) {
	out, err := cc.Convert(in)
	if err != nil {
		return "", err
	}
	return out, nil
}

func (c *CLI) personByName(ctx context.Context, name string) error {
	name, err := simplifiedToTraditional(c.cc, name)
	if err != nil {
		return err
	}

	ps, err := c.Store.FetchPersonByName(ctx, name)
	if err != nil {
		return err
	}
	c.Renderer.PersonByName(os.Stdout, ps)
	return nil
}

// 1. process the input
// 2. query data based on the input
// 3. render the input
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
