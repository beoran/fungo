// drawing functions that work on a Drawable interface
// Builds on top of the fungo/SDL package
// this file contains the line, circle and ellipse drawing
// functions
package draw

import "fungo/sdl"
import "math"

type Surface struct { 
  * sdl.Surface
}

type Drawable interface {
  PutPixel(x, y int, color uint32)
  GetPixel(x, y int) (color uint32)
  BlendPixel(x, y int, color uint32, alpha uint8)
}

func FromSDL(s * sdl.Surface) (*Surface) {
  return &Surface{s}
}

func (s *Surface) toSDL() (*sdl.Surface) {
  return s.Surface
}

func (s * Surface) PutPixel(x1, y1 int, color uint32) {
  s.toSDL().PutPixel(x1, y1, color)
}

func (s * Surface) GetPixel(x1, y1 int) (color uint32) {
  return s.toSDL().GetPixel(x1, y1)
}

func (s * Surface) BlendPixel(x1, y1 int, color uint32, alpha uint8) {
  s.toSDL().BlendPixel(x1, y1, color, alpha)
}

// Draws a horizontal line
func (s * Surface) HLine(x1, y1, w int, color uint32) {
  s.toSDL().FillRectCoord(x1, y1, w, 1, color);
} 

// Draws a vertical line
func (s * Surface) VLine(x1, y1, h int, color uint32) {
  s.toSDL().FillRectCoord(x1, y1, 1, h, color);
} 

// Draws a box (open rectangle) 
func (s * Surface) Box(x1, y1, w, h int, color uint32) {
  s.HLine(x1	   , y1    , w, color)
  s.HLine(x1	   , y1 + h, w, color)
  s.VLine(x1	   , y1    , h, color)
  s.VLine(x1 + w , y1    , h, color)
}

// Draws a line from point x1 y1 to x2 y2
func (s * Surface) Line(x1, y1, x2, y2 int, color uint32) {
  // callback is a closure, saves us from having to pass 
  // explicitly the surface, color, etc
  // We don't use channels and goroutines since SDL 
  // may not support it during locking it's surfaces.
  cb := func(x, y int) {
    s.PutPixel(x, y, color)
  }
  DoLine(x1, y1, x2, y2, cb)
}

// Draws a line that starts at x1 and y1, and which 
// ends at x1 + w and y1 +h. Like the name suggests, 
// it could be the diagnoal of a rectangle.   
func (s * Surface) Diagonal(x1, y1, w, h int, color uint32) {
  s.Line(x1, y1, x1 + w, y1 + h, color)
}

// Draws a circle around the midpoint (x1, y1) with radius r
// and color color
func (s * Surface) Circle(x1, y1, r int, color uint32) {
  // callback is a closure, saves us from having to pass 
  // explicitly the surface, color, etc
  // We don't use channels and goroutines since SDL 
  // may not support it during locking it's surfaces.
  cb := func(x, y int) {
    s.PutPixel(x, y, color)
  }
  DoCircle(x1, y1, r, cb)
}

// Draws an arc between the two angles, expressed in radians
func (s * Surface) Arc(x1, y1, r int, ang1, ang2 float64, color uint32) {
  cb := func(x, y int) {
    s.PutPixel(x, y, color)
  }
  DoArc(x1, y1, r, ang1, ang2, cb)
}

// Draws an ellipse with the two give radii
func (s * Surface) Ellipse(x1, y1, rx, ry int, color uint32) {
  // callback is a closure, saves us from having to pass 
  // explicitly the surface, color, etc
  // We don't use channels and goroutines since SDL 
  // may not support it during locking it's surfaces.
  cb := func(x, y int) {
    s.PutPixel(x, y, color)
  }
  DoEllipse(x1, y1, rx, ry, cb)
}

// Callback for plotting a single value function
type PlotFunction func(x float64) (float64)

// Callback for plotting a multi-value relation
type MultiPlotFunction func(x float64) ([]float64)
  
// Callback for algorithmical drawing functions.
// Use a closure of this type to be able to draw to the drawable. 
type DrawCallback func(x, y int);

// Same, but can draw multiple Y values for one X point
type MultiDrawCallback func(x int, y []int);



// Plots the PlotFunction in the interval between and including 
// x1 and x2 using points
func (s * Surface) Plot(x1, x2 int, step float64,
  color uint32, calc PlotFunction) {
  // callback is a closure, saves us from having to pass 
  // explicitly the surface, color, etc
  // We don't use channels and goroutines since SDL 
  // may not support it during locking it's surfaces.
  cb := func(x, y int) {
    s.PutPixel(x, y, color)
  }
  DoPlot(x1, x2, step, calc, cb)
}

// Plots the PlotFunction in the interval between and including 
// x1 and x2 using lines
func (s * Surface) LinePlot(x1, x2 int, step float64,
  color uint32, calc PlotFunction) {
  // callback is a closure, saves us from having to pass 
  // explicitly the surface, color, etc
  // We don't use channels and goroutines since SDL 
  // may not support it during locking it's surfaces.
  oldx:= -1
  oldy:= -1
  fst := true   
  cb := func(x, y int) {
    if fst { 
      s.PutPixel(x, y, color)
      fst = false
      
    } else {
      s.Line(oldx, oldy, x, y, color)
    }
    oldx = x
    oldy = y
  }
  DoPlot(x1, x2, step, calc, cb)
}


// Plots the MultiPlotFunction in the interval between and including 
// x1 and x2. step determines the precision
func (s * Surface) MultiPlot(x1, x2 int, step float64, 
      color uint32, calc MultiPlotFunction) {
  // callback is a closure, saves us from having to pass 
  // explicitly the surface, color, etc
  // We don't use channels and goroutines since SDL 
  // may not support it during locking it's surfaces.
  cb := func(x int, ylist []int) {
    for _ , y := range ylist {
      s.PutPixel(x, y, color)
    }
  }
  DoMultiPlot(x1, x2, step, calc, cb)
}

// Plots the MultiPlotFunction in the interval between and including 
// x1 and x2. step determines the precision, using lines
// calc *must* alway return the same amount of pints, otherwise it 
// will crash this function
func (s * Surface) MultiLinePlot(x1, x2 int, step float64, 
      color uint32, calc MultiPlotFunction) {
  // callback is a closure 
  oldx:= -1
  oldylist:= make([]int, 0)
  fst := true   
  cb := func(x int,  ylist []int) {
    if fst {
      for _ , y := range ylist {    
        s.PutPixel(x, y, color)
      }
      fst = false      
    } else {
      for i , y := range ylist {
        oldy := oldylist[i]
        s.Line(oldx, oldy, x, y, color)
      }
    }
    oldx = x
    oldylist = ylist
  }
  DoMultiPlot(x1, x2, step, calc, cb)
}


// Example for use of Plot
func (s * Surface) PlotSin(x1, x2, cy int, color uint32) {
  calc := func(x float64) (float64) {
    x   = x / 100
    return (math.Sin(x) * 100) + float64(cy)
  }
  s.Plot(x1, x2, 1.0, color, calc)
}

// Example for use of MultiPlot. This is a slow way to plot draw 
// circles, so  better use Circle
func (s * Surface) PlotCircle(cx, cy, r int, color uint32) {
  rr   := float64(r) * float64(r)
  fcy  := float64(cy)
  fcx  := float64(cx)
  x1   := cx - r
  x2   := cx + r
  // This calculation function is a closure.
  // Ah, closures, how did I ever program without them? :)
  calc := func(x float64) ([]float64) {
    results := make([]float64, 2)
    mx      := x - fcx // Virtually center circle around 0,0
    y       := math.Sqrt(rr - mx*mx)
    results[0] =  fcy + y
    results[1] =  fcy - y
    return results
  }
  s.MultiLinePlot(x1, x2, 1, color, calc)
}


// Calls PlotFunction for all values between and including x1 and x2
// and plots the obtained values, rounded values with drawcallback
// Protected against NaN and Inf, so PlotFunction may return these 
// as well 
func DoPlot(x1, x2 int, step float64, 
           calc PlotFunction, draw DrawCallback){
  if (step <= 0.0) { step = 1.0 }           
  stop  := float64(x2) 
  for xf := float64(x1) ; xf <= stop ; xf += step {
    yf  := calc(xf) 
    if math.IsNaN(yf)    { continue } // skip NaN
    if math.IsInf(yf, 0) { continue } // skip Inf
    y   := round(yf)
    x   := round(xf)
    draw(x, y)
  }
}

// Calls MultiPlotFunction for all values between and including x1 and x2
// and plots the obtained values, rounded values with drawcallback
// Protected against NaN and Inf, so PlotFunction may return these 
// as well 
func DoMultiPlot(x1, x2 int, step float64, 
  calc MultiPlotFunction, draw MultiDrawCallback) {
  if (step <= 0.0) { step = 1.0 }
  stop  := float64(x2)
  for xf := float64(x1) ; xf <= stop ; xf += step {
    ylist  := calc(xf)
    ylen   := len(ylist)
    aid    := make([]int, ylen)
    j      := 0
    for i:=0 ; i < ylen ; i++ {
      yf := ylist[i]  
      if math.IsNaN(yf)    { continue } // skip NaN
      if math.IsInf(yf, 0) { continue } // skip Inf
      aid[j] = round(yf) 
      j++
    }
    cleanlist := make([]int, j)
    copy(cleanlist, aid) 
    draw(int(xf), cleanlist)
  }
}

// Calculates a Quadratic bezier
func QuadraticBezier(t, p0, p1, p2 float64) (float64) {
  return square(1 - t) * p0 + 2*(1-t)*t*p1 + square(t)*p2 
} 


func (s * Surface) PlotBezier(x1, y1, w, p0, p1, p2 int, color uint32) { 
  fx1 := float64(x1)
  fx2 := fx1 + float64(w)
  fy1 := float64(y1)
  fp0 := float64(p0)
  fp1 := float64(p1)
  fp2 := float64(p2)
  calc := func(x float64) (float64) {
    x   = (x - fx1) / (fx2 - fx1)
    y  := QuadraticBezier(x, fp0, fp1, fp2) 
    return y + fy1
  }
  s.LinePlot(x1, x1 + w, 1.0, color, calc)
}



// Helper functions for integer math 
func abs(v int) int { 
  if v >= 0 { return v }
  return -v
}

// Ternary operator for ints
func tern(cond bool, trueval int, falseval int) int {
  if cond { return trueval } 
  return falseval
}

// Ternary operator for float64
func ftern(cond bool, trueval float64, falseval float64) float64 {
  if cond { return trueval } 
  return falseval
}

// Rounds a float64 to the nearest integer
func round(val float64) (int) {
  return int(math.Floor(val + 0.5))
} 

// Returns the square of the value 
func square(value float64) (float64) {
  return value * value 
}

// All the D* actually come originally from Allegro, 
// so it's the Allegro giftware license which applies. 
// Calls the callback for every point on the line (x1 y1) -> (x2 y2)
func DoLine(x1, y1, x2, y2 int, callback DrawCallback) { 
  dx := x2 - x1
  dy := y2 - y1
  sdx := tern(dx < 0, -1 , 1) 
  sdy := tern(dy < 0, -1 , 1)
  dx = sdx * dx + 1
  dy = sdy * dy + 1
  x := 0
  y := 0
  px := x1
  py := y1

  if dx >= dy {
    for x = 0; x < dx; x++ {
      callback(px, py)  
      y += dy;
      if y >= dx {
        y -= dx;
        py += sdy;
      }
      px += sdx;
    }
  } else {
    for y = 0 ; y < dy; y++ {
      callback(px, py)
      x += dx;
      if x >= dy {
        x -= dy;
        px += sdx;
      }
      py += sdy;
    }
  }
}  
  
// Calls callback at each point of the circle 
func DoCircle(x, y, r int, callback DrawCallback) {
  cx    := 0;
  cy    := r;
  df    := 1 - r
  d_e   := 3;
  d_se  := -2 * r + 5

  for cx <= cy {
    callback(x+cx, y+cy)
    callback(x-cx, y+cy)
    callback(x+cx, y-cy)
    callback(x-cx, y-cy)
    callback(x+cy, y+cx)
    callback(x+cy, y-cx)
    callback(x-cy, y+cx)
    callback(x-cy, y-cx)
    if (df < 0)  {
      df += d_e
      d_e += 2
      d_se += 2
    } else {
      df += d_se
      d_e += 2
      d_se += 4
      cy--
    }
    cx++;
  }
}
  
// XXX: doesn't work for some reason. 
func DoEllipse(x, y, rx, ry int, callback DrawCallback) { 
  var ix, iy int;
  var h, i, j, k int;
  var oh, oi, oj, ok int;

  if (rx < 1) { rx = 1 } 

  if (ry < 1) { ry = 1 }

  h, i, j, k = 0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF

  if (rx > ry) {
    ix = 0;
    iy = rx * 64;

    for i > h {
      oh = h;
      oi = i;
      oj = j;
      ok = k;

      h = (ix + 32) >> 6;
      i = (iy + 32) >> 6;
      j = (h * ry) / rx;
      k = (i * ry) / rx;

      if (((h != oh) || (k != ok)) && (h < oi)) {
        callback(x+h, y+k)
          if (h != 0) { callback( x-h, y+k) }
          if (k != 0) {
            callback( x+h, y-k)
            if (h != 0) { callback( x-h, y-k) }
          }
      }

      if (((i != oi) || (j != oj)) && (h < i)) {
        callback( x+i, y+j)
          if (i != 0 ) { callback( x-i, y+j) }
          if (j != 0) {
            callback( x+i, y-j)
            if (i != 0) { callback( x-i, y-j) }
        }
      }

      ix = ix + iy / rx;
      iy = iy - ix / rx;
    }
  } else {
    ix = 0;
    iy = ry * 64;

      for i > h {
      oh = h;
      oi = i;
      oj = j;
      ok = k;

      h = (ix + 32) >> 6;
      i = (iy + 32) >> 6;
      j = (h * rx) / ry;
      k = (i * rx) / ry;

      if (((j != oj) || (i != oi)) && (h < i)) {
          callback( x+j, y+i)
          if (j!=0) { callback( x-j, y+i) }
          if (i!=0) { 
            callback( x+j, y-i) 
             if (j!=0) { callback( x-j, y-i) }
          }
      }

      if (((k != ok) || (h != oh)) && (h < oi)) {
          callback( x+k, y+h)
          if (k!=0) {  callback( x-k, y+h) }
          if (h!=0) {
              callback( x+k, y-h)
            if (k!=0) { callback( x-k, y-h) }
          }
      }

      ix = ix + iy / ry;
      iy = iy - ix / ry;

      } 
  } 
}

// Gets a point on the arc with radius r and angle a (expressed in radiants) 
func GetPointOnArc(r int,  a float64) (x int, y int) {
   var s, c float64;   
   s ,c  = math.Sincos(a)
   sl 	:= -s * float64(r);
   cl   := c  * float64(r);
   x     = (int)(ftern(c < 0, cl - 0.5, cl + 0.5))
   y     = (int)(ftern(s < 0, sl - 0.5, sl + 0.5))
   return x, y
}



// DoArc
//  Helper function for the arc function. Calculates the points in an arc
//  of radius r around point x, y, going anticlockwise from fixed point
//  binary angle ang1 to ang2, and calls the specified routine for each one. 
//  The output proc will be passed first a copy of the bmp parameter, then 
//  the x, y point, then a copy of the d parameter (so putpixel() can be 
//  used as the callback).
//
func DoArc(x1, y1, r int, ang1, ang2 float64, callback DrawCallback) {
  // A full circle can be drawn in a discrete number of steps, 
  // depending only on the radius.
  // I'm taking (r + 1) * 8 steps for a full circle, which is a safe 
  // if slower overestimation.
  csteps := (r + 1) * 8
  dang	 := ang2 - ang1
  steps  := (float64(csteps) * dang / (2 * math.Pi))
  isteps := int(steps) 
  angstep:= dang / steps 
  oldx   := x1
  oldy   := y1
  ang    := ang1
  for i := 0; i < isteps ; i++ {
    cx, cy := GetPointOnArc(r, ang);
    x , y  := cx + x1 , cy + y1    
    ang += angstep
    // Don't draw if not advanced enough
    if oldx == x && oldy == y { continue; }
    callback(x, y)
    oldx = x  
    oldy = y    
  }  
}


