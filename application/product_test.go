package application_test

import (
	"testing"

	"github.com/danvieira97/arq-hexagonal-projeto/application"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{
		Name:   "Product 1",
		Status: application.ENABLED,
		Price:  0,
	}
	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		ID:     uuid.New().String(),
		Name:   "Product 1",
		Status: application.ENABLED,
		Price:  10,
	}

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be DISABLED or ENABLED", err.Error())

	product.Status = application.DISABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater than or equal to zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewString()

	err := product.GetID()
	require.NotNil(t, err)
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Product1"

	err := product.GetName()
	require.NotNil(t, err)
}
