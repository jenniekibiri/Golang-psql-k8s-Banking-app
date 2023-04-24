package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

// func to create a new store object
func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

// func to exec a sql transaction
// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()

}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}	
type TransferResults struct {
	FromAccount Account `json:"from_account"`
	ToAccount   Account `json:"to_account"`
	Transfer    Transfer `json:"transfer"`
}

// func to transfer money from one account to another
func (store *Store) TransferTx( ctx context.Context, arg TransferTxParams) (TransferResults,error){
	var results TransferResults
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		results.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		// update the balance of the from account
		fromAccount, err := q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}
		results.FromAccount = fromAccount
		// update the balance of the to account
		toAccount, err := q.CreateEntry(ctx, CreateEntryParams{

			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})




		if err != nil {
			return err
		}

		// todo: update the balance of the to account

	
		return nil
	})
	return results, err
}


