package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Norzuiso/simplebank/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	arg := CreateTransfersParams{
		FromAccountID: sql.NullInt64{Int64: account1.ID, Valid: true},
		ToAccountID:   sql.NullInt64{Int64: account2.ID, Valid: true},
		Amount:        util.GenerateRandomAmountMoney(),
	}
	transfer, err := testQueries.CreateTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, account2.ID, transfer.ToAccountID.Int64)
	require.Equal(t, account1.ID, transfer.FromAccountID.Int64)
	require.Equal(t, arg.Amount, transfer.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)
	err := testQueries.DeleteTransfers(context.Background(), transfer.ID)

	require.NoError(t, err)

	trasnfer2, err := testQueries.GetTransfers(context.Background(), transfer.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, trasnfer2)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)
	transfer2, err := testQueries.GetTransfers(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.Amount, transfer2.Amount)
	require.Equal(t, transfer2.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transfer2.ToAccountID, transfer.ToAccountID)
	require.WithinDuration(t, transfer2.CreatedAt, transfer.CreatedAt, time.Second)

}

func TestUpdateTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)
	arg := UpdateTransfersParams{
		ID:     transfer.ID,
		Amount: util.GenerateRandomAmountMoney(),
	}
	err := testQueries.UpdateTransfers(context.Background(), arg)

	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfers(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, arg.Amount, transfer2.Amount)
	require.Equal(t, transfer2.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transfer2.ToAccountID, transfer.ToAccountID)
	require.WithinDuration(t, transfer2.CreatedAt, transfer.CreatedAt, time.Second)

}

func TestListOfTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer.ID)
	}
}
