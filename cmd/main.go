package main

import (
	"fmt"

	"github.com/bernylinville/gotop"
	"github.com/xxxserxxx/lingo/v2"
)

var conf gotop.Config

func main() {
	ec := run()
	// if ec > 0 {
	// 	if ec < 2 {
	// 		logpath := filepath.Join(conf.ConfigDir.QueryCacheFolder().Path)
	// 	}
	// }
	// os.Exit(ec)
}

func run() int {
	ling, err := lingo.New("en_US", ".", gotop.Dicts)
	if err != nil {
		fmt.Printf("failed to load language files: %s\n", err)
		return 2
	}
}
