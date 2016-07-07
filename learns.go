package main

import (
	"github.com/lrumler/learns/arrays"
	"github.com/lrumler/learns/conditionals"
	"github.com/lrumler/learns/functions"
	"github.com/lrumler/learns/hello"
	"github.com/lrumler/learns/loops"
	"github.com/lrumler/learns/maths"
	"github.com/lrumler/learns/pointers"
	"github.com/lrumler/learns/random"
	"github.com/lrumler/learns/slices"
	"github.com/lrumler/learns/structs"
	"github.com/lrumler/learns/switches"
	"github.com/lrumler/learns/types"
)

func main() {

	//
	arrays.WhatsAnArray()

	//
	conditionals.PrintSquares()

	//
	functions.DoAddition()
	functions.StringSwap()

	//
	hello.PrintHello()

	//
	loops.CountToAThousand()

	//
	maths.PrintPi()
	maths.TakeSquares()

	//
	pointers.DemonstratePointers()

	//
	random.FaveNum()

	//
	slices.HowToSlice()
	slices.RefsToArrays()
	slices.WhatsASlice()

	//
	structs.Literally()
	structs.LookAVertex()
	structs.NowWithPointers()

	//
	switches.TimeOfDayBasedGreeting()
	switches.WhenSaturday()
	switches.WhichOS()

	//
	types.BasicTypes()
	types.Conversions()
}
