package kurs

import (
	"database/sql"
	"encoding/json"
	"github.com/wendylau87/xfers2021/entities"
)

func(d *domain) createKurs(tx *sql.Tx, v entities.Kurs)(*entities.Kurs, error){
	result, err := tx.Exec(CreateKurs, v.Name)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	v.ID = int(id)

	return &v, nil
}


func(d *domain) findKursByName(name string)(*entities.Kurs, error){
	result := entities.Kurs{}
	rows, err := d.SQLHandler.Query(ReadKursByName, name)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&result.ID, &result.Name); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func(d *domain) findAllKurs()([]entities.Kurs, error){
	results := []entities.Kurs{}
	rows, err := d.SQLHandler.Query(ReadAllKurs)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		result := entities.Kurs{}
		if err = rows.Scan(&result.ID, &result.Name); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}



func(d *domain) createERate(tx *sql.Tx, v entities.ERate)(*entities.ERate, error){
	result, err := tx.Exec(CreateERate, v.KursID, v.Buy, v.Sell, v.ValidDate)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	v.ID = int(id)

	return &v, nil
}


func(d *domain) findERate(kursID int, date string)(*entities.ERate, error){
	result := entities.ERate{}
	rows, err := d.SQLHandler.Query(ReadERate, kursID, date)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&result.ID, &result.KursID, &result.Buy, &result.Sell, &result.ValidDate); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func(d *domain) findAllERateByDate(kursID int, startDate string, endDate string)([]entities.ERate, error){
	results := []entities.ERate{}
	rows, err := d.SQLHandler.Query(ReadAllERateByDate, kursID, startDate, endDate)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		result := entities.ERate{}
		if err = rows.Scan(&result.ID, &result.KursID, &result.Buy, &result.Sell, &result.ValidDate); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	jsonstr, _ := json.Marshal(results)
	d.logger.LogAccess(string(jsonstr))

	return results, nil
}

func(d *domain) updateERate(tx *sql.Tx, v entities.ERate)(*entities.ERate, error){
	_, err := tx.Exec(UpdateERate, v.Buy, v.Sell, v.KursID, v.ValidDate)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func(d *domain) deleteERate(tx *sql.Tx, date string)error{
	_, err := tx.Exec(DeleteERateByDate, date)
	if err != nil {
		return err
	}

	return nil
}

func(d *domain) createTTCounter(tx *sql.Tx, v entities.TTCounter)(*entities.TTCounter, error){
	result, err := tx.Exec(CreateTTCounter, v.KursID, v.Buy, v.Sell, v.ValidDate)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	v.ID = int(id)

	return &v, nil
}


func(d *domain) findTTCounter(kursID int, date string)(*entities.TTCounter, error){
	result := entities.TTCounter{}
	rows, err := d.SQLHandler.Query(ReadTTCounter, kursID, date)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&result.ID, &result.KursID, &result.Buy, &result.Sell, &result.ValidDate); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func(d *domain) findAllTTCounterByDate(kursID int, startDate string, endDate string)([]entities.TTCounter, error){
	results := []entities.TTCounter{}
	rows, err := d.SQLHandler.Query(ReadAllTTCounterByDate, kursID, startDate, endDate)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		result := entities.TTCounter{}
		if err = rows.Scan(&result.ID, &result.KursID, &result.Buy, &result.Sell, &result.ValidDate); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func(d *domain) updateTTCounter(tx *sql.Tx, v entities.TTCounter)(*entities.TTCounter, error){
	_, err := tx.Exec(UpdateTTCounter, v.Buy, v.Sell, v.KursID, v.ValidDate)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func(d *domain) deleteTTCounter(tx *sql.Tx, date string)error{
	_, err := tx.Exec(DeleteTTCounterByDate, date)
	if err != nil {
		return err
	}

	return nil
}

func(d *domain) createBankNote(tx *sql.Tx, v entities.BankNote)(*entities.BankNote, error){
	result, err := tx.Exec(CreateBankNote, v.KursID, v.Buy, v.Sell, v.ValidDate)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	v.ID = int(id)

	return &v, nil
}


func(d *domain) findBankNote(kursID int, date string)(*entities.BankNote, error){
	result := entities.BankNote{}
	rows, err := d.SQLHandler.Query(ReadBankNote, kursID, date)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&result.ID, &result.KursID, &result.Buy, &result.Sell, &result.ValidDate); err != nil {
			return nil, err
		}
	}

	return &result, nil
}


func(d *domain) findAllBankNoteByDate(kursID int, startDate string, endDate string)([]entities.BankNote, error){
	results := []entities.BankNote{}
	rows, err := d.SQLHandler.Query(ReadAllBankNoteByDate, kursID, startDate, endDate)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		result := entities.BankNote{}
		if err = rows.Scan(&result.ID, &result.KursID, &result.Buy, &result.Sell, &result.ValidDate); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func(d *domain) updateBankNote(tx *sql.Tx, v entities.BankNote)(*entities.BankNote, error){
	_, err := tx.Exec(UpdateBankNote, v.Buy, v.Sell, v.KursID, v.ValidDate)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func(d *domain) deleteBankNote(tx *sql.Tx, date string)error{
	_, err := tx.Exec(DeleteBankNoteByDate, date)
	if err != nil {
		return err
	}

	return nil
}