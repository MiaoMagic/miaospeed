package main

import (
	"github.com/moshaoli688/miaospeed/utils"
)

var COMPILATIONTIME string
var BUILDCOUNT string
var COMMIT string
var BRAND string

func main() {
	utils.COMPILATIONTIME = COMPILATIONTIME
	utils.BUILDCOUNT = BUILDCOUNT
	utils.COMMIT = COMMIT
	utils.BRAND = BRAND

	RunCli()
}
