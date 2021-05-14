package kurs

import (
	"github.com/wendylau87/xfers2021/entities"
	"github.com/wendylau87/xfers2021/infrastructure/logger"
	"github.com/wendylau87/xfers2021/infrastructure/sqlhandler"
)

type DomainItf interface {
	CreateKurs(v entities.CreateKurs) (*entities.KursResponse, error)
	GetKursByDate(startDate string, endDate string)([]entities.KursResponses, error)
	GetKursByName(name string, startDate string, endDate string)(entities.KursResponses, error, int)
	DeleteKurs(date string)error
	UpdateKurs(v entities.PutKurs) (*entities.KursResponse, error, int)
	GenerateKurs()([]entities.CreateKurs, error)
}

type domain struct {
	logger logger.Logger
	SQLHandler sqlhandler.SQLHandler
}

func InitKursDomain(logger logger.Logger, sql sqlhandler.SQLHandler) DomainItf {
	return &domain{
		logger,
		sql,
	}
}
