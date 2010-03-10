// drawing functions that work on a Drawable interface
// Builds on top of the fungo/SDL package
// this file contains the line, circle and ellipse drawing
// functions
package draw

import "fungo/sdl"

type Surface sdl.Surface

type Drawable interface {
  PutPixel(x, y int, color uint32)
  GetPixel(x, y int) (color uint32)
  BlendPixel(x, y int, color uint32, alpha uint8)
}

func FromSDL(s * sdl.Surface) (*Surface) {
  return (*Surface)(s)
}

func (s *Surface) toSDL() (*sdl.Surface) {
  return (*sdl.Surface)(s)
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

func (s * Surface) Line(x1, y1, x2, y2 int, color uint32) {
  // callback is a closure, saves us from having to pass 
  // explicitly the surface, color, etc
  // We don't use channels and goroutines since SDL 
  // may not support it during locking it's surfaces.
  cb := func(x, y int) {
    s.PutPixel(x, y, color)
  }
  BresenhamLine(x1, y1, x2, y2, cb)
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
  BresenhamCircle(x1, y1, r, cb)
}

// Drawss an ellipse with the two give radii
func (s * Surface) Ellipse(x1, y1, rx, ry int, color uint32) {
  // callback is a closure, saves us from having to pass 
  // explicitly the surface, color, etc
  // We don't use channels and goroutines since SDL 
  // may not support it during locking it's surfaces.
  cb := func(x, y int) {
    s.PutPixel(x, y, color)
  }
  BresenhamEllipse(x1, y1, rx, ry, cb)
}


// Callback for algorythmicall drawing functions.
// Use a closure to be able to draw to the drawable. 
type DrawCallback func(x, y int);

// Helper functions for integer math 
func abs(v int) int { 
  if v >= 0 { return v }
  return -v
}

// Ternary operator
func tern(cond bool, trueval int, falseval int) int {
  if cond { return trueval } 
  return falseval
}

// All the BresenHam* are based on algorithms from SGE 
// Calls the callback for every point on the line (x1 y1) -> (x2 y2)
func BresenhamLine(x1, y1, x2, y2 int, callback DrawCallback) { 
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
func BresenhamCircle(x, y, r int, callback DrawCallback) {
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
func BresenhamEllipse(x, y, rx, ry int, callback DrawCallback) { 
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
        callback( x+h, y+k)
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
/*
#define DO_BEZIER(function)\
  
  *  Note: I don't think there is any great performance win in translating this to fixed-point integer math,
  *  most of the time is spent in the line drawing routine.
  \
  float x = float(x1), y = float(y1);\
  float xp = x, yp = y;\
  float delta;\
  float dx, d2x, d3x;\
  float dy, d2y, d3y;\
  float a, b, c;\
  int i;\
  int n = 1;\
  Sint16 xmax=x1, ymax=y1, xmin=x1, ymin=y1;\
  \
  // compute number of iterations \
  if(level < 1)\
    level=1;\
  if(level >= 15)\
    level=15; \
  while (level-- > 0)\
    n*= 2;\
  delta = float( 1.0 / float(n) );\
  \
  // compute finite differences \
  // a, b, c are the coefficient of the polynom in t defining the parametric curve \
  // The computation is done independently for x and y \
  a = float(-x1 + 3*x2 - 3*x3 + x4);\
  b = float(3*x1 - 6*x2 + 3*x3);\
  c = float(-3*x1 + 3*x2);\
  \
  d3x = 6 * a * delta*delta*delta;\
  d2x = d3x + 2 * b * delta*delta;\
  dx = a * delta*delta*delta + b * delta*delta + c * delta;\
  \
  a = float(-y1 + 3*y2 - 3*y3 + y4);\
  b = float(3*y1 - 6*y2 + 3*y3);\
  c = float(-3*y1 + 3*y2);\
  \
  d3y = 6 * a * delta*delta*delta;\
  d2y = d3y + 2 * b * delta*delta;\
  dy = a * delta*delta*delta + b * delta*delta + c * delta;\
  \
  if (SDL_MUSTLOCK(surface) && _sge_lock) {\
    if (SDL_LockSurface(surface) < 0)\
      return;\
  }\
  \
  // iterate \
  for (i = 0; i < n; i++) {\
    x += dx; dx += d2x; d2x += d3x;\
    y += dy; dy += d2y; d2y += d3y;\
    if(Sint16(xp) != Sint16(x) || Sint16(yp) != Sint16(y)){\
      function;\
      if(_sge_update==1){\
        xmax= (xmax>Sint16(xp))? xmax : Sint16(xp);  ymax= (ymax>Sint16(yp))? ymax : Sint16(yp);\
        xmin= (xmin<Sint16(xp))? xmin : Sint16(xp);  ymin= (ymin<Sint16(yp))? ymin : Sint16(yp);\
        xmax= (xmax>Sint16(x))? xmax : Sint16(x);    ymax= (ymax>Sint16(y))? ymax : Sint16(y);\
        xmin= (xmin<Sint16(x))? xmin : Sint16(x);    ymin= (ymin<Sint16(y))? ymin : Sint16(y);\
      }\
    }\
    xp = x; yp = y;\
  }\
  \
  // unlock the display \
  if (SDL_MUSTLOCK(surface) && _sge_lock) {\
    SDL_UnlockSurface(surface);\
  }\
  \
  // Update the area \
  sge_UpdateRect(surface, xmin, ymin, xmax-xmin+1, ymax-ymin+1);
*/  
