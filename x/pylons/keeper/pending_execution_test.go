package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/Pylons-tech/pylons/x/pylons/types"
)

func createNPendingExecution(k *Keeper, ctx sdk.Context, n int) []types.Execution {
	items := make([]types.Execution, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].ID = k.AppendPendingExecution(ctx, items[i])
	}
	return items
}

func TestPendingExecutionGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPendingExecution(&keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetPendingExecution(ctx, item.ID))
	}
}

func TestPendingExecutionExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPendingExecution(&keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasPendingExecution(ctx, item.ID))
	}
}

func TestPendingExecutionGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPendingExecution(&keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllPendingExecution(ctx))
}

func TestPendingExecutionCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPendingExecution(&keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetPendingExecutionCount(ctx))
}
