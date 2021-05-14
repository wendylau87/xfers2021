package usecases

import (
	"github.com/wendylau87/xfers2021/domain"
	"github.com/wendylau87/xfers2021/infrastructure/logger"
	"github.com/wendylau87/xfers2021/usecases/kurs"
)

type Usecase struct {
	Kurs kurs.UsecaseItf
}

func Init(logger logger.Logger, dom *domain.Domain) *Usecase {
	return &Usecase{
		Kurs : kurs.InitKursUsecase(logger, dom.Kurs),
	}
}