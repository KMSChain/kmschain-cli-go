package main

import (
	"cli/kmschain/cmd"
	//"cli/kmschain/logging"
	//"kmschain-sdk-go/kmschain"
)

func main() {
	//logging.SetupLogger()
	//configs.ParseConfig("config.json")
	//kmschain.CM_init()
	cmd.Execute()
}
