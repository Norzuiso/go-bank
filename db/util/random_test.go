package util

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomInt(t *testing.T) {
	x := GenerateRandomInt(0, 10)
	var y int64 = GenerateRandomInt(0, 10)
	require.IsType(t, reflect.ValueOf(y), reflect.ValueOf(x))
}

func TestGenerateRandomString(t *testing.T) {
	var str string = "testing String"
	randStr := GenerateRandomString(6)

	require.IsType(t, reflect.ValueOf(str), reflect.ValueOf(randStr))
}

func TestGenerateRandomOwner(t *testing.T) {
	var str string = "123456"
	randStr := GenerateRandomOwner()

	require.Len(t, randStr, 6)
	require.IsType(t, reflect.ValueOf(str), reflect.ValueOf(randStr))
}

func TestGenerateRandomAmountMoney(t *testing.T) {
	x := GenerateRandomAmountMoney()
	var y int64 = GenerateRandomInt(0, 10)

	require.IsType(t, reflect.ValueOf(y), reflect.ValueOf(x))
}

func TestGenerateRandomCurrency(t *testing.T) {
	x := GenerateRandomCurrency()
	require.NotEmpty(t, x)
}
