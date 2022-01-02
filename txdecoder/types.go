package txdecoder

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibctransfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	ibctypes "github.com/cosmos/ibc-go/modules/core/types"
	nfttypes "github.com/crypto-org-chain/chain-main/v3/x/nft/types"

	gaialiquiditytypes "github.com/gravity-devs/liquidity/x/liquidity/types"

	osmoepochtypes "github.com/osmosis-labs/osmosis/x/epochs/types"
	osmogammtypes "github.com/osmosis-labs/osmosis/x/gamm/types"
	osmoincentivestypes "github.com/osmosis-labs/osmosis/x/incentives/types"
	osmolockuptypes "github.com/osmosis-labs/osmosis/x/lockup/types"
	osmopoolincentivestypes "github.com/osmosis-labs/osmosis/x/pool-incentives/types"

	// osmotxfeestypes "github.com/osmosis-labs/osmosis/x/txfees/types"

	terramarkettypes "github.com/terra-money/core/x/market/types"
	terraoracletypes "github.com/terra-money/core/x/oracle/types"
	terravestingtypes "github.com/terra-money/core/x/vesting/types"
	terrawasmtypes "github.com/terra-money/core/x/wasm/types"
)

func RegisterDefaultInterfaces(interfaceRegistry types.InterfaceRegistry) {
	std.RegisterInterfaces(interfaceRegistry)

	authtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	distributiontypes.RegisterInterfaces(interfaceRegistry)
	evidencetypes.RegisterInterfaces(interfaceRegistry)
	govtypes.RegisterInterfaces(interfaceRegistry)
	proposal.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	vestingtypes.RegisterInterfaces(interfaceRegistry)
	ibctypes.RegisterInterfaces(interfaceRegistry)
	ibctransfertypes.RegisterInterfaces(interfaceRegistry)
	nfttypes.RegisterInterfaces(interfaceRegistry)
	authztypes.RegisterInterfaces(interfaceRegistry)
	feegranttypes.RegisterInterfaces(interfaceRegistry)

	gaialiquiditytypes.RegisterInterfaces(interfaceRegistry)

	osmoepochtypes.RegisterInterfaces(interfaceRegistry)
	osmogammtypes.RegisterInterfaces(interfaceRegistry)
	osmoincentivestypes.RegisterInterfaces(interfaceRegistry)
	osmolockuptypes.RegisterInterfaces(interfaceRegistry)
	osmopoolincentivestypes.RegisterInterfaces(interfaceRegistry)
	// osmotxfeestypes.RegisterInterfaces(interfaceRegistry)

	terramarkettypes.RegisterInterfaces(interfaceRegistry)
	terraoracletypes.RegisterInterfaces(interfaceRegistry)
	terravestingtypes.RegisterInterfaces(interfaceRegistry)
	terrawasmtypes.RegisterInterfaces(interfaceRegistry)
}
