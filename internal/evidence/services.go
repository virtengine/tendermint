package evidence

import (
	"github.com/tendermint/tendermint/types"
)

//go:generate go run github.com/vektra/mockery/v2 --disable-version-string --case underscore --name BlockStore

type BlockStore interface {
	LoadBlockMeta(height int64) *types.BlockMeta
	LoadBlockCommit(height int64) *types.Commit
	Height() int64
}
