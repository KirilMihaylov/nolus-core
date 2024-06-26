package v051

import (
	"github.com/Nolus-Protocol/nolus-core/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

const (
	// UpgradeName defines the on-chain upgrades name.
	UpgradeName = "v0.5.3"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{},
	},
}
