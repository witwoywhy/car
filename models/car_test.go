package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCar(t *testing.T) {
	t.Run("can new car", func(t *testing.T) {
		want := &Car{
			Brand:    "honda",
			Name:     "brv",
			Model:    "e",
			SubModel: "taffeta",
			Color:    "white",
			Price:    915000,
		}

		car, err := NewCar("Honda", "BRV", "e", "Taffeta", "white", 915000)
		assert.Nil(t, err)
		assert.Equal(t, want, car)
	})

	type input struct {
		brand, name, model, subModel, color string
	}

	type tester struct {
		title string
		input input
		err   error
	}
	tests := []tester{
		{title: "brand", input: input{brand: ""}, err: errBrandInvalid},
		{title: "name", input: input{brand: "test", name: ""}, err: errorNameInvalid},
		{title: "model", input: input{brand: "test", name: "test", model: ""}, err: errorModelInvalid},
		{title: "color", input: input{brand: "test", name: "test", model: "test", color: ""}, err: errorColorInvalid},
	}

	for _, v := range tests {
		t.Run(fmt.Sprintf("%s invalid", v.title), func(t *testing.T) {
			_, err := NewCar(v.input.brand, v.input.name, v.input.model, "", v.input.color, 0)
			assert.Equal(t, v.err, err)
		})
	}
}
