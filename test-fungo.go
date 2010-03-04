package main

import "fmt"
import "os"
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

func main()	{
  TestResults();
}
