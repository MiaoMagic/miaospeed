package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/moshaoli688/miaospeed/interfaces"
	"github.com/moshaoli688/miaospeed/service/macros/script"
	"github.com/moshaoli688/miaospeed/utils"
	"github.com/moshaoli688/miaospeed/vendors"
)

type ScriptTestCliParams struct {
	ScriptName string
}

func InitConfigScriptTest() *ScriptTestCliParams {
	stcp := &ScriptTestCliParams{}

	sflag := flag.NewFlagSet(cmdName+" script", flag.ExitOnError)
	sflag.StringVar(&stcp.ScriptName, "file", "", "specify a script file to perform a test.")

	parseFlag(sflag)

	return stcp
}

func RunCliScriptTest() {
	stcp := InitConfigScriptTest()

	if stcp.ScriptName == "" {
		utils.DErrorf("Script Test | File name cannot be empty.")
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(stcp.ScriptName)
	if err != nil {
		utils.DErrorf("Script Test | Cannot read the file, path=%s", stcp.ScriptName)
		os.Exit(1)
	}

	utils.VerboseLevel = utils.LTInfo
	utils.DWarnf("MiaoSpeed speedtesting client %s", utils.VERSION)

	vendor := vendors.Find(interfaces.VendorLocal)
	utils.DInfof("Script Test | Using vendor %s", vendor.Type())
	scriptResult := script.ExecScript(vendor, &interfaces.Script{
		Content: string(fileContent),
	})

	fmt.Println("\n" + utils.ToJSON(scriptResult))
}
