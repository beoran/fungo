package main

import "fmt"
import "os"
import "fungo/sdl"
// import "tamias"
// import "fungo/sdl"
/*
import "exp/draw"
import "exp/draw/x11"
*/

type testbool bool

type testresults struct {
  ok, failed, total int  
}

var suite = testresults{0,0,0}

func (cond testbool) should(err string, args ...) {
  assert(bool(cond), err, args)
}

func error(error string, args ...) {
  suite.failed++ 
  suite.total++   
  fmt.Fprintln(os.Stderr, "Failed assertion nr", suite.total, ":", error, args);
}

func no_error() {
  suite.ok++ 
  suite.total++ 
}

func assert(cond bool, err string,  args ...)  {
  if cond {
    no_error()
    return
  }
  error(err, args) 
}

func TestResults() {
  fmt.Fprintf(os.Stderr, "Test results: %d/%d test passed, %d/%d failed.\n",
    suite.ok, suite.total, suite.failed, suite.total)
}

// This test doesn't score anything, just prints results, which should print.
// Mysteriously stopped working (undefined symbol in shared lib)
/*
func TestCpuinfo() {
  fmt.Println("CPU Features:")
  fmt.Println("SSE:", sdl.HasSSE())
  fmt.Println("SSE2:", sdl.HasSSE2())
  fmt.Println("3DNow:", sdl.Has3DNow())
  fmt.Println("3DNowExt:", sdl.Has3DNowExt())
  fmt.Println("RDTRSC:", sdl.HasRDTRSC())
  fmt.Println("MMX:", sdl.HasMMX())
  fmt.Println("MMXExt:", sdl.HasMMXExt())
  fmt.Println("AltiVec:", sdl.HasAltiVec())
}
*/

func TestError() {
  sdl.Error(sdl.EFREAD)
  err := sdl.GetError()
  fmt.Println(err)
  sdl.ClearError()
  err2 := sdl.GetError()
  fmt.Println(err2)
}

func main()	{
  // TestCpuinfo()
  TestError()
  TestResults()
}



