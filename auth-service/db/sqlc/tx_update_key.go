package db

import (
	"context"
)

type UpdateKeysTxParams struct {
	UpdateKeysParams
}

type UpdateKeysTxResult struct {
	KeyPair KeyPair
}

func (store *SQLStore) UpdateKeysTx(ctx context.Context, arg UpdateKeysTxParams) (UpdateKeysTxResult, error) {
	var result UpdateKeysTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.KeyPair, err = q.UpdateKeys(ctx, arg.UpdateKeysParams)
		return err
	})

	return result, err
}
