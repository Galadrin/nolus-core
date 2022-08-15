package simapp

//
//import (
//	"github.com/cosmos/cosmos-sdk/simapp"
//	abci "github.com/tendermint/tendermint/abci/types"
//	"github.com/tendermint/tendermint/libs/log"
//	tmdb "github.com/tendermint/tm-db"
//	"gitlab-nomo.credissimo.net/nomo/cosmzone/app"
//)
//
////
//// New creates application instance with in-memory database and disabled logging.
//func New(dir string, withDefaultGenesisState bool) simapp.App {
//	db := tmdb.NewMemDB()
//	logger := log.NewNopLogger()
//
//	encoding := simapp.MakeEncodingConfig(app.ModuleBasics)
//
//	a := app.New(logger, db, nil, true, map[int64]bool{}, dir, 0, encoding,
//		simapp.EmptyAppOptions{})
//	// InitChain updates deliverState which is required when app.NewContext is called
//	genState := []byte("{}")
//	if withDefaultGenesisState {
//		genStateObj := NewDefaultGenesisState(encoding.Marshaler)
//		state, err := json.MarshalIndent(genStateObj, "", " ")
//		if err != nil {
//			panic(err)
//		}
//		genState = state
//	}
//	a.InitChain(abci.RequestInitChain{
//		ConsensusParams: defaultConsensusParams,
//		AppStateBytes:   genState,
//	})
//	return a
//}
//
////func TestSetup() (*app.App, error) {
////	rootApp := New(app.DefaultNodeHome, true)
////	nolusApp, ok := rootApp.(*app.App)
////	if !ok {
////		return nil, fmt.Errorf("invalid simapp created: %v", ok)
////	}
////	return nolusApp, nil
////}
////
////// NewDefaultGenesisState generates the default state for the application.
////func NewDefaultGenesisState(cdc codec.JSONCodec) app.GenesisState {
////	return app.ModuleBasics.DefaultGenesis(cdc)
////}
////
////var defaultConsensusParams = &abci.ConsensusParams{
////	Block: &abci.BlockParams{
////		MaxBytes: 200000,
////		MaxGas:   2000000,
////	},
////	Evidence: &tmproto.EvidenceParams{
////		MaxAgeNumBlocks: 302400,
////		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
////		MaxBytes:        10000,
////	},
////	Validator: &tmproto.ValidatorParams{
////		PubKeyTypes: []string{
////			tmtypes.ABCIPubKeyTypeEd25519,
////		},
////	},
////}
////
////// NewAppConstructor returns a new simapp AppConstructor
////func NewAppConstructor() network.AppConstructor {
////	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
////
////	return func(val network.Validator) servertypes.Application {
////		return app.New(val.Ctx.Logger, tmdb.NewMemDB(), nil, true, map[int64]bool{}, val.Ctx.Config.RootDir, 0, encoding,
////			simapp.EmptyAppOptions{},
////			baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
////			baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
////		)
////	}
////}
