package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab-nomo.credissimo.net/nomo/cosmzone/x/treasury/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

// GetParams todo split to multiple methods
func (k Keeper) GetParams(ctx sdk.Context) (state types.GenesisState) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GenesisStateKey)
	if b == nil {
		panic("treasury stored state must not have been nil")
	}

	k.cdc.MustUnmarshal(b, &state)
	return
}

func (k Keeper) AddProceeds(ctx sdk.Context, delta sdk.Coins) {
	genState := k.GetParams(ctx)
	if genState.FeeProceeds == nil {
		genState.FeeProceeds = sdk.NewCoins()
	}
	genState.FeeProceeds.Add(delta...)
	k.Logger(ctx).Info(fmt.Sprintf("New fee proceeds state: %s", genState.FeeProceeds))
	k.SetParams(ctx, genState)
}

// SetParams stores the genesis state. Needs a refactor to store parameters as separate values
func (k Keeper) SetParams(ctx sdk.Context, genState types.GenesisState) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&genState)
	store.Set(types.GenesisStateKey, b)
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
