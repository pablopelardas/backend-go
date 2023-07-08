package db

import (
	"context"
	"testing"
	"time"

	"github.com/pablopelardas/backend-go/utils"
	"github.com/stretchr/testify/require"
)

func resetEnvironment(t *testing.T) {
	testQueries.ResetTransfers(context.Background())
	testQueries.ResetEntries(context.Background())
	testQueries.ResetAccounts(context.Background())
}

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	resetEnvironment(t)
	account := createRandomAccount(t)
	createRandomEntry(t, account)
	resetEnvironment(t)
}

func TestGetEntry(t *testing.T) {
	resetEnvironment(t)
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
	resetEnvironment(t)
}

func TestListEntries(t *testing.T) {
	resetEnvironment(t)
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}
	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, account.ID, entry.AccountID)
	}
	resetEnvironment(t)
}

func TestDeleteEntry(t *testing.T) {
	resetEnvironment(t)
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.Empty(t, entry2)
	resetEnvironment(t)
}

func TestUpdateEntry(t *testing.T) {
	resetEnvironment(t)
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)
	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: utils.RandomMoney(),
	}
	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.NotEqual(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
	resetEnvironment(t)
}