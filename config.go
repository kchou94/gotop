package gotop

import (
	"embed"

	"github.com/shibukawa/configdir"
)

// FIXME github action uses old(er) Go version that doesn't have embed
// go:embed "dicts/*.toml"
var Dicts embed.FS

type Config struct {
	ConfigDir configdir.ConfigDir
}
