package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Norzuiso/simplebank/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntrie(t *testing.T) Entry {
	account := CreateRandomAccount(t)
	arg := CreateEntriesParams{
		AccountID: sql.NullInt64{account.ID, true},
		Amount:    util.GenerateRandomAmountMoney(),
	}

	entrie, err := testQueries.CreateEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entrie)

	require.Equal(t, arg.AccountID, entrie.AccountID)
	require.Equal(t, arg.Amount, entrie.Amount)
	require.NotEmpty(t, entrie.ID)

	return entrie
}

func TestGetEntry(t *testing.T) {
	entry := createRandomEntrie(t)

	entry2, err := testQueries.GetEntries(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotNil(t, entry2)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry.ID, entry2.ID)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, entry.Amount, entry2.Amount)
	require.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestCreateEntries(t *testing.T) {
	createRandomEntrie(t)
}

func TestDeleteEntrie(t *testing.T) {
	entry := createRandomEntrie(t)
	err := testQueries.DeleteEntries(context.Background(), entry.ID)

	require.NoError(t, err)
	entry2, err := testQueries.GetEntries(context.Background(), entry.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)

}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntrie(t)
	}
	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotNil(t, entry.ID)
		require.NotEmpty(t, entry.ID)
	}
}

func TestUpdateEntriy(t *testing.T) {
	entry1 := createRandomEntrie(t)
	arg := UpdateEntriesParams{
		ID:     entry1.ID,
		Amount: util.GenerateRandomAmountMoney(),
	}
	err := testQueries.UpdateEntries(context.Background(), arg)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntries(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, arg.Amount, entry2.Amount)
	require.Equal(t, entry1.ID, entry2.ID)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)

}
