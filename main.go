package main

import (
	"fmt"
	"my-go-test/defers"
	"my-go-test/exercise"
	"my-go-test/forcontinued"
	"my-go-test/functions"
	"my-go-test/importmath"
	"my-go-test/multipleresults"
	"my-go-test/namedresults"
	"my-go-test/numeric"
	"my-go-test/packages"
	"my-go-test/pointers"
	"my-go-test/structs"
	"my-go-test/switchs"
	"my-go-test/typeinference"
	"my-go-test/variables"
	"my-go-test/arrays"
	"my-go-test/maps"
	"my-go-test/methods"
	"my-go-test/interfaces"
	"my-go-test/errors"
)

func main() {
	fmt.Println("Welcome to the playgroud!")
	importmath.Main()
	packages.Main()
	functions.Main()
	multipleresults.Main()
	namedresults.Main()
	variables.Main()
	variables.MainInitializers()
	variables.MainVariableDeclarations()
	typeinference.MainTypeInference()
	numeric.MainNumericConstants()
	forcontinued.MainForContinued()
	exercise.MainLoopAndFunction()
	switchs.MainSwitchOS()
	switchs.MainSwitchWithNoCondition()
	defers.MainDefer()
	defers.MainDeferStack()
	pointers.MainPointer()
	structs.MainStruct()
	structs.MainStructPoint()
    arrays.MainArray()
    arrays.MainAppend()
    arrays.MainRangeContinue()
    exercise.MainPic()
    maps.MainMap()
    maps.MainMapMultate()
    exercise.MainMaps()
    functions.MainFunctionValue()
    functions.MainFunctionClosures()
    exercise.MainFibonacciClosure()
    methods.MainMethod()
    methods.MainPorinter()
    interfaces.MainInterface()
    interfaces.MainInterfaceValues()
    interfaces.MainTypeAssertion()
    interfaces.MainTypeSwitch()
    interfaces.MainStringer()
    errors.MainError()
}
