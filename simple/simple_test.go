package simple

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWireSucces(t *testing.T) {
	injector, err := ServiceInjector(false)
	if err != nil {
		panic(err)
	}
	fmt.Println(injector.SimpleRepository)
	fmt.Println(err)
	assert.Equal(t, nil, err)
	assert.Equal(t, false, injector.Error)
}

func TestWireFailed(t *testing.T) {
	injector, errInjector := ServiceInjector(true)
	// if errInjector != nil {
	// 	panic(errInjector)
	// }
	// Assert
	fmt.Println(injector)
	fmt.Println(errInjector)
	assert.Equal(t, errors.New("failed NewSimpleService"), errInjector)
	assert.Nil(t, injector)
}
