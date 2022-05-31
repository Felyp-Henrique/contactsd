//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
)

type Injection struct {
	DataSource        MongoDataSource
	ContactRepository ContactRepository
}

func NewInjection(datasource MongoDataSource, repository ContactRepository) Injection {
	return Injection{
		DataSource:        datasource,
		ContactRepository: repository,
	}
}

func GetInjection() Injection {
	wire.Build(NewInjection, NewMongoDataSource, NewContactRepositoryMongo)
	return Injection{}
}
