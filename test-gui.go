package main

import "fmt"
import "os"
import "time"
import "fungo/sdl"
// import "fungo/draw"
import "fungo/gui"
import "reflect" 
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
  sdl.EnableUnicode(1)
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

type Foo struct {
}

type Any interface{}
type Message string
type Method reflect.Method

type Object interface {
  Send(Message , ...Any) 
  DefineMethod(Message, Any)
  GetMethod(Message, Any) 
}

func Hello() {
  fmt.Println("Hello Dynamic Call")
}

func Hello2(extra string) {
  fmt.Println("Hello Dynamic Call", extra)
}

func Square(val int) (int) {
  return val * val
}

// Can call any function with any arguments
// Slow but flexible.
// TODO: Think about what to do with the return values.
// A raw []reflect.Value is not so interesting
func DynamicCall(fun Any, args ...Any) (Any) {
  fval, ok    := reflect.NewValue(fun).(*reflect.FuncValue)  
  if ! ok { error("Not a function") ; return nil }
  // fmt.Println(fval)
  funkind, ok2 := fval.Type().(*reflect.FuncType)
  if ! ok2 { error("Not a function") ; return nil } 
  // fmt.Println(funkind, funkind.Name(), funkind.String(), funkind.NumIn())
  nargs := funkind.NumIn()
  vargs := make([]reflect.Value, nargs)
  if len(args) < nargs { error("Too few arguments") ; return nil }
  for i:= 0 ; i < nargs ; i ++ {
    val := reflect.NewValue(args[i])
    if val.Type() != funkind.In(i) { 
      error("Wrong argument type: ", val.Type(), 
      "expected", funkind.In(i) ) 
      return nil
    }
    vargs[i] = val
  }
  results := fval.Call(vargs)
  return results
}

func TimeFunction(repeats int, totime func()) (int64) {
  start := time.Nanoseconds()
  for i:= 0 ; i < repeats; i++ {
    totime()
  } 
  stop := time.Nanoseconds()
  return stop - start
}


// Tests the object system
func TestObject() {
  DynamicCall(Hello, "Ignore me", "Ignore me too")
  
  DynamicCall(Hello2, 1)
  Hello2("World")
  tf1 := func() {
    DynamicCall(Square, 12345)
  }
  tf2 := func() {
    Square(12345)
  }
  repeats := 100000
  t1 := TimeFunction(repeats, tf1)
  t2 := TimeFunction(repeats, tf2)
  fmt.Println(repeats, t1, t2)
  
}

func TestSetup() {
  sdl.InitDefault()
}

func TestQuit() {
  sdl.Quit()    
} 

func main()	{
  TestSetup()
  // TestObject()
  TestGui()
  TestResults()
}



