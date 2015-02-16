package helper

import (
	"log"
	"os"

	"github.com/jonasnick/go-ethereum/ethutil"
	logpkg "github.com/jonasnick/go-ethereum/logger"
)

var Logger logpkg.LogSystem
var Log = logpkg.NewLogger("TEST")

func init() {
	Logger = logpkg.NewStdLogSystem(os.Stdout, log.LstdFlags, logpkg.InfoLevel)
	logpkg.AddLogSystem(Logger)

	ethutil.ReadConfig(".ethtest", "/tmp/ethtest", "")
}
