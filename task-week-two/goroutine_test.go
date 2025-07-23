package main

import "testing"

func TestPrintNums(t *testing.T) {
	PrintNums()
}

func TestTriggerTask(t *testing.T) {
	funs := []func(int, int) int{InitFuncs, InitFuncs, InitFuncs, InitFuncs, InitFuncs, InitFuncs}
	TriggerTask(funs)
}

func TestTriggerTask2(t *testing.T) {
	funs := []func(int, int) int{InitFuncs, InitFuncs, InitFuncs, InitFuncs, InitFuncs, InitFuncs}
	TriggerTask2(funs)
}
