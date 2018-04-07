// this file was generated by gomacro command: import _b "runtime"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"runtime"
)

// reflection: allow interpreted code to import "runtime"
func init() {
	Packages["runtime"] = Package{
	Binds: map[string]Value{
		"BlockProfile":	ValueOf(runtime.BlockProfile),
		"Breakpoint":	ValueOf(runtime.Breakpoint),
		"CPUProfile":	ValueOf(runtime.CPUProfile),
		"Caller":	ValueOf(runtime.Caller),
		"Callers":	ValueOf(runtime.Callers),
		"CallersFrames":	ValueOf(runtime.CallersFrames),
		"Compiler":	ValueOf(runtime.Compiler),
		"FuncForPC":	ValueOf(runtime.FuncForPC),
		"GC":	ValueOf(runtime.GC),
		"GOARCH":	ValueOf(runtime.GOARCH),
		"GOMAXPROCS":	ValueOf(runtime.GOMAXPROCS),
		"GOOS":	ValueOf(runtime.GOOS),
		"GOROOT":	ValueOf(runtime.GOROOT),
		"Goexit":	ValueOf(runtime.Goexit),
		"GoroutineProfile":	ValueOf(runtime.GoroutineProfile),
		"Gosched":	ValueOf(runtime.Gosched),
		"KeepAlive":	ValueOf(runtime.KeepAlive),
		"LockOSThread":	ValueOf(runtime.LockOSThread),
		"MemProfile":	ValueOf(runtime.MemProfile),
		"MemProfileRate":	ValueOf(&runtime.MemProfileRate).Elem(),
		"MutexProfile":	ValueOf(runtime.MutexProfile),
		"NumCPU":	ValueOf(runtime.NumCPU),
		"NumCgoCall":	ValueOf(runtime.NumCgoCall),
		"NumGoroutine":	ValueOf(runtime.NumGoroutine),
		"ReadMemStats":	ValueOf(runtime.ReadMemStats),
		"ReadTrace":	ValueOf(runtime.ReadTrace),
		"SetBlockProfileRate":	ValueOf(runtime.SetBlockProfileRate),
		"SetCPUProfileRate":	ValueOf(runtime.SetCPUProfileRate),
		"SetCgoTraceback":	ValueOf(runtime.SetCgoTraceback),
		"SetFinalizer":	ValueOf(runtime.SetFinalizer),
		"SetMutexProfileFraction":	ValueOf(runtime.SetMutexProfileFraction),
		"Stack":	ValueOf(runtime.Stack),
		"StartTrace":	ValueOf(runtime.StartTrace),
		"StopTrace":	ValueOf(runtime.StopTrace),
		"ThreadCreateProfile":	ValueOf(runtime.ThreadCreateProfile),
		"UnlockOSThread":	ValueOf(runtime.UnlockOSThread),
		"Version":	ValueOf(runtime.Version),
	}, Types: map[string]Type{
		"BlockProfileRecord":	TypeOf((*runtime.BlockProfileRecord)(nil)).Elem(),
		"Error":	TypeOf((*runtime.Error)(nil)).Elem(),
		"Frame":	TypeOf((*runtime.Frame)(nil)).Elem(),
		"Frames":	TypeOf((*runtime.Frames)(nil)).Elem(),
		"Func":	TypeOf((*runtime.Func)(nil)).Elem(),
		"MemProfileRecord":	TypeOf((*runtime.MemProfileRecord)(nil)).Elem(),
		"MemStats":	TypeOf((*runtime.MemStats)(nil)).Elem(),
		"StackRecord":	TypeOf((*runtime.StackRecord)(nil)).Elem(),
		"TypeAssertionError":	TypeOf((*runtime.TypeAssertionError)(nil)).Elem(),
	}, Proxies: map[string]Type{
		"Error":	TypeOf((*P_runtime_Error)(nil)).Elem(),
	}, Untypeds: map[string]string{
		"Compiler":	"string:gc",
	}, Wrappers: map[string][]string{
		"BlockProfileRecord":	[]string{"Stack",},
	}, 
	}
}

// --------------- proxy for runtime.Error ---------------
type P_runtime_Error struct {
	Object	interface{}
	Error_	func(interface{}) string
	RuntimeError_	func(interface{}) 
}
func (P *P_runtime_Error) Error() string {
	return P.Error_(P.Object)
}
func (P *P_runtime_Error) RuntimeError()  {
	P.RuntimeError_(P.Object)
}
