package kurs

import (
	"encoding/json"
	"github.com/wendylau87/xfers2021/entities"
)

func (u *usecase) CreateKurs(v entities.CreateKurs) (*entities.KursResponse, error) {
	return u.domain.CreateKurs(v)
}

func (u *usecase) GetKursByDate(startDate string, endDate string) ([]entities.KursResponses, error) {
	return u.domain.GetKursByDate(startDate, endDate)
}

func (u *usecase) GetKursByName(name string, startDate string, endDate string) (entities.KursResponses, error, int) {
	return u.domain.GetKursByName(name, startDate, endDate)
}

func (u *usecase) DeleteKurs(date string) error {
	return u.domain.DeleteKurs(date)
}

func (u *usecase) UpdateKurs(v entities.PutKurs) (*entities.KursResponse, error, int) {
	return u.domain.UpdateKurs(v)
}

func (u *usecase) GenerateKurs() (error) {
	createKurs, err := u.domain.GenerateKurs()
	if err != nil{
		return err
	}

	for _, v := range createKurs{
		jstr, _ := json.Marshal(v)
		u.logger.LogAccess(string(jstr))
		_, err = u.domain.CreateKurs(v)
		if err != nil{
			u.logger.LogError("Error when generate kurs because %s", err)
		}
	}
	return nil
}

