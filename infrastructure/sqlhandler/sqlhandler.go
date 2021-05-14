package sqlhandler

type SQLHandleItf interface {
	Begin() (*Tx, error)
	Query(string, ...interface{}) (*Row, error)
	Exec(string, ...interface{}) (*Result, error)
	In(string, ...interface{})(string, []interface{}, error)
}

type TxItf interface {
	Commit() error
	Rollback() error
	Exec(string, ...interface{}) (*Result, error)
}

type ResultItf interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type RowItf interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
	Err() error
}
