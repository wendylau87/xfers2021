package kurs

import (
	"errors"
	"github.com/wendylau87/xfers2021/entities"
)

func (d *domain) CreateKurs(v entities.CreateKurs) (*entities.KursResponse, error) {
	result := entities.KursResponse{}
	//NOTE: this is a transaction example.
	tx, err := d.SQLHandler.Begin()
	if err != nil {
		d.logger.LogError("%s", err)
		return nil, err
	}

	//Create KURS
	kurs, err := d.findKursByName(v.Name)
	if err != nil {
		_ = tx.Rollback()
		d.logger.LogError("Failed find kurs by Name : %s", err)
		return nil, err
	}
	if kurs.ID == 0 {
		kurs, err = d.createKurs(tx.Tx, entities.Kurs{
			Name: v.Name,
		})
		if err != nil {
			d.logger.LogError("Failed create new Kurs : %s", err)
			_ = tx.Rollback()
			return nil, err
		}
	}
	result.Name = kurs.Name


	//Create E-Rate
	erate, err := d.findERate(kurs.ID, v.ValidDate)
	if err != nil {
		_ = tx.Rollback()
		d.logger.LogError("Failed find E-Rate : %s", err)
		return nil, err
	}

	if erate.ID == 0 {
		erate, err = d.createERate(tx.Tx, entities.ERate{
			KursID:    kurs.ID,
			Buy:       v.ERate.Buy,
			Sell:      v.ERate.Sell,
			ValidDate: v.ValidDate,
		})
		if err != nil {
			d.logger.LogError("Failed create new E-Rate : %s", err)
			_ = tx.Rollback()
			return nil, err
		}
		result.ERate.Buy = erate.Buy
		result.ERate.Sell = erate.Sell
		result.ERate.ValidDate = erate.ValidDate
	}


	//Create TT Counter
	ttcounter, err := d.findTTCounter(kurs.ID, v.ValidDate)
	if err != nil {
		d.logger.LogError("Failed find TT Counter : %s", err)
		_ = tx.Rollback()
		return nil, err
	}
	if ttcounter.ID == 0 {
		ttcounter, err = d.createTTCounter(tx.Tx, entities.TTCounter{
			KursID:    kurs.ID,
			Buy:       v.TTCounter.Buy,
			Sell:      v.TTCounter.Sell,
			ValidDate: v.ValidDate,
		})
		if err != nil {
			d.logger.LogError("Failed create new TT Counter : %s", err)
			_ = tx.Rollback()
			return nil, err
		}

		result.TTCounter.Buy = ttcounter.Buy
		result.TTCounter.Sell = ttcounter.Sell
		result.TTCounter.ValidDate = ttcounter.ValidDate
	}

	//Create Bank Notes
	banknote, err := d.findBankNote(kurs.ID, v.ValidDate)
	if err != nil {
		d.logger.LogError("Failed find Bank Note : %s", err)
		_ = tx.Rollback()
		return nil, err
	}
	if banknote.ID == 0 {
		banknote, err = d.createBankNote(tx.Tx, entities.BankNote{
			KursID:    kurs.ID,
			Buy:       v.BankNote.Buy,
			Sell:      v.BankNote.Sell,
			ValidDate: v.ValidDate,
		})
		if err != nil {
			_ = tx.Rollback()
			d.logger.LogError("Failed create new Bank Note : %s", err)
			return nil, err
		}
		result.BankNote.ValidDate = banknote.ValidDate
		result.BankNote.Buy = banknote.Buy
		result.BankNote.Sell = banknote.Sell
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *domain) GetKursByDate(startDate string, endDate string) ([]entities.KursResponses, error) {
	results := []entities.KursResponses{}
	kurs, err := d.findAllKurs()
	if err != nil {
		return results, err
	}

	erates := []entities.ERate{}
	ttcounters := []entities.TTCounter{}
	banknotes := []entities.BankNote{}
	for _, v := range kurs {
		erate, err := d.findAllERateByDate(v.ID, startDate, endDate)
		if err != nil {
			return results, err
		}
		erates = append(erates, erate...)
		ttcounter, err := d.findAllTTCounterByDate(v.ID, startDate, endDate)
		if err != nil {
			return results, err
		}
		ttcounters = append(ttcounters, ttcounter...)
		banknote, err := d.findAllBankNoteByDate(v.ID, startDate, endDate)
		if err != nil {
			return results, err
		}
		banknotes = append(banknotes, banknote...)
	}

	for _, v := range kurs {
		var result entities.KursResponses
		result.Name = v.Name
		for _, obj := range erates {
			if obj.KursID == v.ID {
				result.ERates = append(result.ERates, entities.ERateResponse{
					Buy:       obj.Buy,
					Sell:      obj.Sell,
					ValidDate: obj.ValidDate,
				})
			}
		}

		for _, obj := range ttcounters {
			if obj.KursID == v.ID {
				result.TTCounters = append(result.TTCounters, entities.TTCounterResponse{
					Buy:       obj.Buy,
					Sell:      obj.Sell,
					ValidDate: obj.ValidDate,
				})
			}
		}

		for _, obj := range banknotes {
			if obj.KursID == v.ID {
				result.BankNotes = append(result.BankNotes, entities.BankNoteResponse{
					Buy:       obj.Buy,
					Sell:      obj.Sell,
					ValidDate: obj.ValidDate,
				})
			}
		}

		results = append(results, result)
	}

	return results, nil
}

func (d *domain) GetKursByName(name string, startDate string, endDate string) (entities.KursResponses, error, int) {
	var result entities.KursResponses
	kurs, err := d.findKursByName(name)
	if err != nil {
		return result, err, 500
	}

	if kurs.ID == 0 {
		return result, errors.New("Kurs tidak ditemukan."), 404
	}

	erates := []entities.ERate{}
	ttcounters := []entities.TTCounter{}
	banknotes := []entities.BankNote{}
	erate, err := d.findAllERateByDate(kurs.ID, startDate, endDate)
	if err != nil {
		return result, err, 500
	}
	erates = append(erates, erate...)
	ttcounter, err := d.findAllTTCounterByDate(kurs.ID, startDate, endDate)
	if err != nil {
		return result, err, 500
	}
	ttcounters = append(ttcounters, ttcounter...)
	banknote, err := d.findAllBankNoteByDate(kurs.ID, startDate, endDate)
	if err != nil {
		return result, err, 500
	}
	banknotes = append(banknotes, banknote...)

	result.Name = kurs.Name
	for _, obj := range erates {
		if obj.KursID == kurs.ID {
			result.ERates = append(result.ERates, entities.ERateResponse{
				Buy:       obj.Buy,
				Sell:      obj.Sell,
				ValidDate: obj.ValidDate,
			})
		}
	}

	for _, obj := range ttcounters {
		if obj.KursID == kurs.ID {
			result.TTCounters = append(result.TTCounters, entities.TTCounterResponse{
				Buy:       obj.Buy,
				Sell:      obj.Sell,
				ValidDate: obj.ValidDate,
			})
		}
	}

	for _, obj := range banknotes {
		if obj.KursID == kurs.ID {
			result.BankNotes = append(result.BankNotes, entities.BankNoteResponse{
				Buy:       obj.Buy,
				Sell:      obj.Sell,
				ValidDate: obj.ValidDate,
			})
		}
	}

	return result, nil, 500
}

func (d *domain) DeleteKurs(date string) error {
	//NOTE: this is a transaction example.
	tx, err := d.SQLHandler.Begin()
	if err != nil {
		return err
	}

	err = d.deleteERate(tx.Tx, date)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = d.deleteTTCounter(tx.Tx, date)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = d.deleteBankNote(tx.Tx, date)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (d *domain) UpdateKurs(v entities.PutKurs) (*entities.KursResponse, error, int) {
	result := entities.KursResponse{}
	kurs, err := d.findKursByName(v.Name)
	if err != nil {
		return nil, err, 500
	}

	if kurs.ID == 0 {
		return nil, errors.New("Kurs tidak ditemukan."), 404
	}
	result.Name = kurs.Name

	//NOTE: this is a transaction example.
	tx, err := d.SQLHandler.Begin()
	if err != nil {
		return nil, err, 500
	}

	//Update E-Rate
	erate, err := d.findERate(kurs.ID, v.ValidDate)
	if err != nil {
		_ = tx.Rollback()
		return nil, err, 500
	}
	if erate.ID == 0 {
		_ = tx.Rollback()
		return nil, errors.New("E-Rate tidak ditemukan."), 404
	}else{
		erate.Buy = v.ERate.Buy
		erate.Sell = v.ERate.Sell
		erate.ValidDate = v.ValidDate
		_, err = d.updateERate(tx.Tx, *erate)
		if err != nil {
			_ = tx.Rollback()
			return nil, err, 500
		}
		result.ERate.Buy = erate.Buy
		result.ERate.Sell = erate.Sell
		result.ERate.ValidDate = erate.ValidDate
	}

	//Update TT Counter
	ttcounter, err := d.findTTCounter(kurs.ID, v.ValidDate)
	if err != nil {
		_ = tx.Rollback()
		return nil, err, 500
	}
	if ttcounter.ID == 0 {
		_ = tx.Rollback()
		return nil, errors.New("TT Counter tidak ditemukan."), 404

	}else{
		ttcounter.Buy = v.TTCounter.Buy
		ttcounter.Sell = v.TTCounter.Sell
		ttcounter.ValidDate = v.ValidDate
		_, err = d.updateTTCounter(tx.Tx, *ttcounter)
		if err != nil {
			_ = tx.Rollback()
			return nil, err, 500
		}
		result.TTCounter.Buy = ttcounter.Buy
		result.TTCounter.Sell = ttcounter.Sell
		result.TTCounter.ValidDate = ttcounter.ValidDate
	}

	//Update Bank Notes
	banknote, err := d.findBankNote(kurs.ID, v.ValidDate)
	if err != nil {
		_ = tx.Rollback()
		return nil, err, 500
	}
	if banknote.ID == 0 {
		_ = tx.Rollback()
		return nil, errors.New("Bank Note tidak ditemukan."), 404
	}else{
		banknote.Buy = v.BankNote.Buy
		banknote.Sell = v.BankNote.Sell
		banknote.ValidDate = v.ValidDate
		banknote, err = d.updateBankNote(tx.Tx, *banknote)
		if err != nil {
			_ = tx.Rollback()
			return nil, err, 500
		}

		result.BankNote.ValidDate = banknote.ValidDate
		result.BankNote.Buy = banknote.Buy
		result.BankNote.Sell = banknote.Sell
	}

	if err = tx.Commit(); err != nil {
		return nil, err, 500
	}

	return &result, nil, 200
}

func (d *domain)GenerateKurs()([]entities.CreateKurs, error) {
	return d.callBCA()
}
