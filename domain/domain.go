package domain

import (
	"github.com/wendylau87/xfers2021/domain/kurs"
	"github.com/wendylau87/xfers2021/infrastructure/logger"
	"github.com/wendylau87/xfers2021/infrastructure/sqlhandler"
)

type Domain struct {
	Kurs kurs.DomainItf
}

func Init(
	logger logger.Logger,
	sql sqlhandler.SQLHandler,
) *Domain {
	return &Domain{
		Kurs : kurs.InitKursDomain(logger, sql),
	}
}