package main

import (
	"github.com/leerumler/learns/arrays"
	"github.com/leerumler/learns/conditionals"
	"github.com/leerumler/learns/functions"
	"github.com/leerumler/learns/hello"
	"github.com/leerumler/learns/loops"
	"github.com/leerumler/learns/maths"
	"github.com/leerumler/learns/pointers"
	"github.com/leerumler/learns/random"
	"github.com/leerumler/learns/slices"
	"github.com/leerumler/learns/structs"
	"github.com/leerumler/learns/switches"
	"github.com/leerumler/learns/types"
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
