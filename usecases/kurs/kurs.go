package kurs

import (
	"github.com/wendylau87/xfers2021/domain/kurs"
	"github.com/wendylau87/xfers2021/entities"
	"github.com/wendylau87/xfers2021/infrastructure/logger"
)

type UsecaseItf interface {
	CreateKurs(v entities.CreateKurs) (*entities.KursResponse, error)
	GetKursByDate(startDate string, endDate string)([]entities.KursResponses, error)
	GetKursByName(name string, startDate string, endDate string)(entities.KursResponses, error, int)
	DeleteKurs(date string)error
	UpdateKurs(v entities.PutKurs) (*entities.KursResponse, error, int)
	GenerateKurs()error
}

type usecase struct {
	logger logger.Logger
	domain kurs.DomainItf
}



func InitKursUsecase(logger logger.Logger, dom kurs.DomainItf) UsecaseItf {
	return &usecase{
		logger,
		dom,
	}
}