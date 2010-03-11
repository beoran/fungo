package main

import "fmt"
import "os"
import "time"
import "fungo/sdl"
// import "fungo/draw"
import "fungo/gui"
// import "rand"
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

func sleep(secs int) { 
  time.Sleep(int64(secs) * 1000000000);
}

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

//Tests the GUI
func TestGui() {  
  screen := sdl.OpenScreen(640, 480, 32, 0)
  /*
  font := sdl.LoadTTFont("data/sazanami-gothic.ttf", 24)
  defer font.Free()
  img := font.RenderBlended("Hello 世界!", white)
  defer img.Free()
  */
  // open the screen 
  gui    := gui.NewHanao(screen)
  for !gui.Done() {
    gui.Update()
      
    /*  
      gui.Draw(screen)
    */ 
    screen.Flip()
  }
}

func TestSetup() {
  sdl.Init(sdl.INIT_EVERYTHING)
}

func TestQuit() {
  sdl.Quit()    
} 

func main()	{
  TestSetup()
  TestGui()
  TestResults()
}



