package benchmark_caller_stack

import (
	"runtime"
	"strings"
	"testing"
)

// RunFuncName
func RunFuncName() string {
	pc, _, _, ok := runtime.Caller(1)

	if !ok {
		return "can_not_get_function_name"
	}

	f := runtime.FuncForPC(pc)
	funcName := strings.Split(f.Name(), ".")
	return funcName[len(funcName)-1]
}

func CallerStack_target_deap(b *testing.B, deap int, target int) {
	if deap == target {
		b.StartTimer()
		RunFuncName()
		return
	} else {
		CallerStack_target_deap(b, deap+1, target)
	}
}
func CallerStack_do_nothing(b *testing.B, deap int, target int) {
	if deap == target {
		b.StartTimer()
		RunFuncName()
		return
	} else {
		CallerStack_target_deap(b, deap+1, target)
	}
}

func BenchmarkCallerStack_CallerStack_100_do_nothing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		CallerStack_do_nothing(b, 0, 100)
	}
}

func BenchmarkCallerStack_CallerStack_10000_do_nothing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		CallerStack_do_nothing(b, 0, 10000)
	}
}

func BenchmarkCallerStack_CallerStack_100_deap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		CallerStack_target_deap(b, 0, 100)
	}
}

func BenchmarkCallerStack_CallerStack_10000_deap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		CallerStack_target_deap(b, 0, 10000)
	}
}

func BenchmarkCallerStack_pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunFuncName()
	}
}
