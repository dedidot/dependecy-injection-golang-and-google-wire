//go:build wireinject
//+build wireinject

package simple

import "github.com/google/wire"

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
			NewDatabaseMongoDB,
			NewDatabasePostgreSQL,
			NewDatabaseRepository,
		)
	return nil
}

var FooSet = wire.NewSet(NewFooRepository, NewFooService)
var BarSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FoobarService {
	wire.Build(FooSet, BarSet, NewFoobarService)
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(NewApplication, wire.FieldsOf(new(*Application), "Configuration"))
	return nil
}