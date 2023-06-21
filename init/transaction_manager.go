package init

import (
	"context"
	"database/sql"
	"errors"
)

type TransactionResult struct {
	Rows   uint64
	Result interface{}
}

type TransactionManager struct {
	db *sql.DB
}

func NewTransactionManager(db *sql.DB) TransactionManager {
	return TransactionManager{db: db}
}

func (txw TransactionManager) ExecWithTransaction(txFunc func(*sql.Tx) error) error {
	tx, err := txw.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})

	if err != nil {
		return errors.Join(err, errors.New("error starting database transaction"))
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			_ = tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	return txFunc(tx)
}
