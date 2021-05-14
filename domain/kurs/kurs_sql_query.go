package kurs

const(
	CreateKurs = `INSERT INTO kurs(name) VALUES(?)`;
	ReadAllKurs = `SELECT id, name FROM kurs`
	ReadKursByName = `SELECT id, name FROM kurs where name = ?`

	CreateERate = `INSERT INTO e_rate(kurs_id, buy, sell, valid_date) VALUES(?, ?, ?, ?)`;
	ReadERate = `SELECT id, kurs_id, buy, sell, valid_date FROM e_rate where kurs_id=? and valid_date = ?`
	ReadAllERateByDate = `SELECT id, kurs_id, buy, sell, valid_date FROM e_rate where kurs_id IN (?) and (valid_date >=? and valid_date <= ?)`
	UpdateERate = `UPDATE e_rate SET buy=?, sell =? WHERE kurs_id =? and valid_date =?`
	DeleteERateByDate = `DELETE FROM e_rate WHERE valid_date = ?`

	CreateTTCounter = `INSERT INTO tt_counter(kurs_id, buy, sell, valid_date) VALUES(?, ?, ?, ?)`;
	ReadTTCounter = `SELECT id, kurs_id, buy, sell, valid_date FROM tt_counter where kurs_id=? and valid_date =?`
	ReadAllTTCounterByDate = `SELECT id, kurs_id, buy, sell, valid_date FROM tt_counter where kurs_id IN (?) and (valid_date >=? and valid_date <= ?)`
	UpdateTTCounter = `UPDATE tt_counter SET buy=?, sell =? WHERE kurs_id =? and valid_date =?`
	DeleteTTCounterByDate = `DELETE FROM tt_counter WHERE valid_date = ?`

	CreateBankNote = `INSERT INTO bank_notes(kurs_id, buy, sell, valid_date) VALUES(?, ?, ?, ?)`;
	ReadBankNote = `SELECT id, kurs_id, buy, sell, valid_date FROM bank_notes where kurs_id=? and valid_date =?`
	ReadAllBankNoteByDate = `SELECT id, kurs_id, buy, sell, valid_date FROM bank_notes where kurs_id IN (?) and (valid_date >=? and valid_date <= ?)`
	UpdateBankNote = `UPDATE bank_notes SET buy=?, sell =? WHERE kurs_id =? and valid_date =?`
	DeleteBankNoteByDate = `DELETE FROM bank_notes WHERE valid_date = ?`

)
