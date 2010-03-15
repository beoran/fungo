package main

import "fmt"
import "os"
import "time"
import "fungo/sdl"
import "fungo/draw"
// import "fungo/gui"
import "rand"
import "fungo/tile"
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


func TestInit() {
  assert(sdl.Initialized(sdl.INIT_JOYSTICK & sdl.INIT_AUDIO), 
    "SDL Init everything initializeds Joystick aslo" )
  assert(sdl.Initialized(sdl.INIT_AUDIO), 
    "SDL Init everything initializes audio also" )
}

func TestCD() {
  fmt.Println("Number of CD Drives:", sdl.CDNumDrives())
  cd := sdl.OpenCD(0)
  assert(cd != nil, "Can open CD drive.")    
  if cd == nil { return }
  ntrack := cd.CountTracks()
  fmt.Printf("CD Status: %d, Tracks: %d. (%s)\n", int(cd.Status()), ntrack,
  cd.String())   
  for i := 0 ; i < ntrack; i++ {   
    track := cd.Track(i);
    if i == 0 {
      // Note: you won't hear music unles your CD drive has been connected 
      // with an analog audio cable to the motherboard. The chance is high, 
      // in contemporary PC's, that it hasnt been correctly conected, and
      // that no sound will be heard.
      // This functonality iseems to be slated to disappear in SDL1.3, 
      // but I'm including it because who knows who might need it.
      cd.PlayTracks(i, track.Offset(), 1, track.Length())
      // sleep(10)
      cd.Stop()
    }    
    fmt.Printf("Track %d: %d,  %d\n", i, track.Offset(), track.Length())
  }  
  // cd.Eject()    
  cd.Close()  
}

func TestJoystick() {
  joysticks := sdl.OpenAllJoysticks() 
  if len(joysticks) > 0 {  
    fmt.Println(joysticks)
    js0 := joysticks[0]
    //for { 
      sdl.JoystickUpdate()
      fmt.Println(js0.Button(0), js0.Button(1), js0.Button(2), js0.Button(3))
    // if js0.Button(0) != 0 { break }
    //}
  }
}

func TestKeyboard() {
  sdl.EnableUnicode(1)
  keys := sdl.GetKeyState()
  fmt.Println(keys)
}

func TestError() {
  sdl.Error(sdl.EFREAD)
  err := sdl.GetError()
  exp := "Error reading from datastream"
  assert(err == exp, "Can fetch correct error message") 
  // fmt.Println(err)
  sdl.ClearError()
  // err2 := sdl.GetError()
  // fmt.Println(err2)
}

//Tests some surface drawing and also some event handling
func TestSurface() {
  
  screen := sdl.OpenScreen(640, 480, 32, sdl.FULLSCREEN)
  white  := sdl.Color{255, 255, 255, 0};
  col    := screen.MapRGB(255, 255,   0);
  colb   := screen.MapRGB(0, 0, 0);
  // colw   := screen.MapRGB(255, 255, 255);    
  fmt.Println("bit depth", screen.Format().BitsPerPixel())
  screen.FillRect(nil, colb)
  screen.Flip()
  start := sdl.GetTicks()
  // screen.Lock()
  fmt.Println(screen.W(), screen.H())
  for x:= 0 ; x < screen.W(); x++ {
    for y:= 0 ; y < screen.H(); y++ {
      screen.BlendPixel(x, y, col, 8)
    }  
  }
  // screen.Unlock()
  stop := sdl.GetTicks()
  fmt.Println("1 screen of 400 px drawn in ", stop - start)
  frame := sdl.LoadSurface("data/tile_aqua.png")
  frame.Accellerate() // Very important, this one! ^_^
  defer frame.Free()
  frame.BlitTo(screen, 20, 20)
  screen.Flip()
  
  font := sdl.LoadTTFont("data/sazanami-gothic.ttf", 24)
  defer font.Free()
  img := font.RenderBlended("Hello 世界!", white)
  defer img.Free()
  ti  := tile.NewTile(tile.NORMAL, 8)
  ti.Add(frame)
  // ts   := tile.NewTileSet(32, 32)  
  // ts.Add(ti, -1)
  layer:= tile.NewLayer(100, 100, 32, 32)
  for x := 0; x< 100; x++ { 
    for y := 0; y< 100; y++ { 
      layer.Set(x, y, ti) 
    }
  }
  

  
  
  start1 := sdl.GetTicks()
  repeats := uint32(1000)
  for i:= uint32(0) ; i <  repeats ; i++ { 
    layer.Draw(screen, 35, 35)   
    layer.Draw(screen, 35, 35)   
    layer.Draw(screen, 35, 35)   
    layer.Draw(screen, 35, 35)
    screen.Flip()
  }    
  stop1 := sdl.GetTicks()
  delta := stop1 - start1
  fps   := repeats * 1000 / delta
  fmt.Println("Repeats delta fps:", repeats, delta, fps)
  img.BlitTo(screen, 20, 20)  
  rand.Seed(time.Nanoseconds())
  draw := draw.FromSDL(screen)
  draw.PutPixel(200, 200, rand.Uint32())
  draw.PutPixel(200, 250, rand.Uint32())
  draw.Line(200, 200, 200, 250, rand.Uint32())
  draw.PutPixel(300, 300, rand.Uint32())
  draw.PutPixel(350, 300, rand.Uint32())
  draw.Line(300, 300, 350, 300, rand.Uint32())
  draw.PutPixel(400, 400, rand.Uint32())
  draw.PutPixel(450, 450, rand.Uint32())
  draw.Line(400, 400, 410, 450, rand.Uint32())
  draw.Circle(400, 400, 50, rand.Uint32())
  draw.Ellipse(300, 300, 50, 100, rand.Uint32())
  draw.Arc(200, 200, 50, 0.0, 2.0, rand.Uint32())
  draw.PlotSin(100, 500, 200, rand.Uint32())
  draw.PlotCircle(100, 100, 50, rand.Uint32())
  draw.PlotBezier(100, 100, 100, 0, 500, 0, rand.Uint32())
  draw.Box(100, 100, 200, 100, rand.Uint32())
  
  
  
  screen.Flip()
  
  /*
  var ev * sdl.Event;
  for {
    ev = sdl.PollEvent()
    if ev == nil { continue }  
    // fmt.Println("Event type:", ev.Type)
    kb := ev.Keyboard()
    if kb != nil {  
      fmt.Println(kb); 
    }
    if ev.Quit() != nil { 
      break 
    }
  }
  */
  // Flush event queue iuntil quit requested
  // event := sdl.WaitEvent()
  // fmt.Println("Event type:", event.Type)
  // sdl.Delay(5000)
}

func TestSetup() {
  sdl.Init(sdl.INIT_EVERYTHING)
}

func TestQuit() {
  sdl.Quit()    
} 

func TestSound() {
  sdl.OpenMixerDefault()
  mus := sdl.LoadMusic("data/ufsjin.mid")
  snd := sdl.LoadSound("data/pugs.wav") 
  defer snd.Free()
  defer mus.Free() 
  fmt.Println(mus, snd)
  mus.Play()   
  snd.Play()
  sdl.Delay(500)
  snd.Play()
  sdl.Delay(2000)
  mus.Pause()
  sdl.CloseMixer()
}

func main()	{
  TestSetup()
  // TestSound()  
  TestJoystick()
  TestSurface()
  TestCpuinfo()
  TestKeyboard()  
  TestInit()
  TestCD()
  TestError()  
  TestQuit()
  TestResults()
}



