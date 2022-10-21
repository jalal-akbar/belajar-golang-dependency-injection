//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
)

func ServiceInjector(isError bool /*paramter injection*/) (*SimpleService, error) { // Error Handling
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

// Multiple Binding
func DatabaseInjector() *DatabaseRepository {
	wire.Build(NewMongoDB, NewPostgreeSQL, NewDatabaseRepository)
	return nil
}

// Provider Set
// Group Foo and Bar
var fooSet = wire.NewSet(NewFooService, NewFoo)
var barSet = wire.NewSet(NewBarService, NewBarRepository)

func FooBarInjector() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBar)
	return nil
}

// Binding Interface
// Fail
// Default Wire Looking a struct not interface
// func SayHelloInjector() *SayHelloService {
// 	wire.Build(NewSayHelloImpl, NewSayHelloService)
// 	return nil
// }

// Success
var helloSet = wire.NewSet(NewSayHelloImpl, wire.Bind(new(SayHello), new(*SayHelloImpl))) // Bind SayHello interface to NewSayHelloImpl

func SayHelloInjector() *SayHelloService {
	wire.Build(helloSet, NewSayHelloService)
	return nil
}

// Struct Provider
var aDanb = wire.NewSet(NewA, NewB)

func StructProviderInjector() *AdanB {
	wire.Build(aDanb, wire.Struct(new(AdanB), "A", "B"))
	return nil
}

// Binding Values
func BindingValuesInjector() *AdanB {
	wire.Build(wire.Value(Avalue), wire.Value(Bvalue), wire.Struct(new(AdanB), "*"))
	return nil
}

// Struct Field Provider
func ConfigurationProvider() *Configuration {
	wire.Build(NewApplication, wire.FieldsOf(new(*Application), "Configuration"))
	return nil
}

// Cleanup Function
func FileInjector(name string) (*Connention, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
