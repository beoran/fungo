package sdl

//struct private_hwdata{};
//struct SDL_BlitMap{};
//#define map _map
//
//#include <SDL.h>
//#include <SDL_image.h>
//#include <SDL_mixer.h>
//#include <SDL_ttf.h>
import "C"
import "unsafe"
import "runtime"

// Wrap.go contains most high-level wrappers of the low level
// C-interface functions
 
// Rect is a rectangle 
/*
type Rect struct { 
  X int16
  Y int16
  W uint16
  H uint16
} 
*/
  
func NewRect(x, y int16, w, h uint16) (* Rect) {
  rect      := &Rect{x, y, w, h}
  return rect
} 
  
// To prevent problemw with memory management, 
// we let use Go structures for Rect and Color
// This converts the pointer from Go to SDL's C pointer
func (rect *Rect) toSDL() (* C.SDL_Rect) {
  return (* C.SDL_Rect)(unsafe.Pointer(rect))
}  
  
/*  
type Color struct { 
  R uint8
  G uint8
  B uint8
  unused uint8 
} 
*/

func NewColor(r, g, b uint8) (* Color) {
  color      := &Color{r,g,b,0}
  return color
}   
   
// Converts struct to SDL struct of same layout  
func (color *Color) toSDL() (* C.SDL_Color) {
  return (* C.SDL_Color)(unsafe.Pointer(color))
}

func wrapRect(rect * C.SDL_Rect) (* Rect) {
  return (* Rect)(unsafe.Pointer(rect))
}


// Have to think a bit about how to wrap Palette
/*
typedef struct SDL_Palette {
  int       ncolors;
  SDL_Color *colors;
} SDL_Palette;
*/

type PixelFormat struct { 
  format * C.SDL_PixelFormat
}

func wrapPixelFormat(format * C.SDL_PixelFormat) (* PixelFormat) {
  return &PixelFormat{format}
}


func (format * PixelFormat) BitsPerPixel() uint8 {
  return uint8(format.format.BitsPerPixel)
} 
 
func (format * PixelFormat) BytesPerPixel() uint8 {
  return uint8(format.format.BytesPerPixel)
} 

func (format * PixelFormat) Rloss() uint8 {
  return uint8(format.format.Rloss)
} 

func (format * PixelFormat) Gloss() uint8 {
  return uint8(format.format.Gloss)
} 

func (format * PixelFormat) Bloss() uint8 {
  return uint8(format.format.Bloss)
} 

func (format * PixelFormat) Aloss() uint8 {
  return uint8(format.format.Aloss)
} 

func (format * PixelFormat) Rshift() uint8 {
  return uint8(format.format.Rshift)
} 

func (format * PixelFormat) Gshift() uint8 {
  return uint8(format.format.Gshift)
} 

func (format * PixelFormat) Bshift() uint8 {
  return uint8(format.format.Bshift)
} 

func (format * PixelFormat) Ashift() uint8 {
  return uint8(format.format.Ashift)
} 

func (format * PixelFormat) Rmask() uint32 {
  return uint32(format.format.Rmask)
} 

func (format * PixelFormat) Gmask() uint32 {
  return uint32(format.format.Gmask)
} 

func (format * PixelFormat) Bmask() uint32 {
  return uint32(format.format.Bmask)
} 

func (format * PixelFormat) Amask() uint32 {
  return uint32(format.format.Amask)
} 

func (format * PixelFormat) Colorkey() uint32 {
  return uint32(format.format.colorkey)
} 

func (format * PixelFormat) Alpha() uint8 {
  return uint8(format.format.alpha)
} 

// Maps the color with the given r, g, b components to the color format
func (format * PixelFormat) MapRGB(r, g, b uint8) uint32 {
  return mapRGB(format.format, r, g, b)
}

// Maps the color with the given r, g, b, a components to 
// the color format
func (format * PixelFormat) MapRGBA(r, g, b, a uint8) uint32 {
  return mapRGBA(format.format, r, g, b, a)
}

// Gets the components of the given color
func (format * PixelFormat) GetRGB(color uint32) (r, g, b uint8)  {
  return getRGB(format.format, color)
}

// Gets the components of the given color
func (format * PixelFormat) GetRGBA(color uint32) (r, g, b, a uint8) {
  return getRGBA(format.format, color)
}


// A surface is a bitmap
// Note that you cannot construct this yourself using new
// or &Surface 
type Surface struct { 
  surface * C.SDL_Surface
}

// Wraps an SDL surface into a Surface
// Returns nil if the original surface was also nil 
func wrapSurface(surface * C.SDL_Surface) (* Surface) {
  if surface == nil { return nil }  
  result := &Surface{surface}
  clean  := func(surf Surface) { surf.Free(); }
  runtime.SetFinalizer(result, clean)
  return result
}  

// Returns the fags of the Surface
func (surface * Surface) Flags() uint32 {
  return uint32(surface.surface.flags)
} 

// Returns the PixelFormat associated with the Surface
func (surface * Surface) Format() (*PixelFormat) {
  return wrapPixelFormat(surface.surface.format)
} 

// Returns the width of the surface
func (surface * Surface) W() int {
  return int(surface.surface.w)
} 

// Returns the height of the surface
func (surface * Surface) H() int {
  return int(surface.surface.h)
} 

// Returns the pitch of the surface
func (surface * Surface) Pitch() uint16 {
  return uint16(surface.surface.pitch)
} 

// Returns the offset of the surface
func (surface * Surface) Offset() int {
  return int(surface.surface.offset)
} 
  
// Think about how to wrap pixels if at all? 
// Perhaps it's better to provide several putpixels
// for every color depth

// Returns the Clipping rectangle f the surface
func (surface * Surface) ClipRect() (*Rect) {
  return wrapRect(&surface.surface.clip_rect)
} 

// Returns the current active video surface or screen
func GetScreen() (*Surface) {
  return wrapSurface(getVideoSurface())
}

// Opens the screen with the given width, height, color depth 
// and flags
func OpenScreen(width, height, bpp int, flags uint32) (* Surface) {   
  screen := setVideoMode(width, height, bpp,flags)
  return wrapSurface(screen)
} 

// Makes a new surface with the given flags, size, depth, and 
// masks 
func NewSurfaceRGB(flags uint32, width, height, depth int, 
  rmask, gmask, bmask, amask uint32) (* Surface) {
  surf := createRGBSurface(flags, width, height, depth, rmask, gmask,
  bmask, amask)
  return wrapSurface(surf)
} 

// Returns the masks for this bpp without alpha, and also the 
// new bpp to use in case the given bpp is unavailable 
func masksForDepthNoAlpha(bpp uint8) (rmask uint32, 
      gmask uint32, bmask uint32, amask uint32, bppres uint8) { 
  amask = 0       
  switch bpp { 
    case 8:
      rmask = 0xFF >> 6 << 5
      gmask = 0xFF >> 5 << 2
      bmask = 0xFF >> 6
    case 12:
      rmask = 0xFF >> 4 << 8
      gmask = 0xFF >> 4 << 4
      bmask = 0xFF >> 4
    case 15:
      rmask = 0xFF >> 3 << 10
      gmask = 0xFF >> 3 << 5
      bmask = 0xFF >> 3
    case 16:
      rmask = 0xFF >> 3 << 11
      gmask = 0xFF >> 2 << 5
      bmask = 0xFF >> 3
    case 24, 32:
      rmask = 0xFF << 16
      gmask = 0xFF << 8
      bmask = 0xFF
    default: // set bbp to 32 bits if nothing better is found 
      rmask = 0xFF << 16
      gmask = 0xFF << 8
      bmask = 0xFF
      bpp = 32
  } 
  return rmask, gmask, bmask, amask, bpp
}  
      
// Returns the masks for this bpp with alpha, and also the 
// new bpp to use in case the given bpp is unavailable or unusable  
func masksForDepthAlpha(bpp uint8) (rmask uint32, 
  gmask uint32, bmask uint32, amask uint32, bppres uint8) {
  switch bpp { 
    case 16:
      rmask = 0xF << 8
      gmask = 0xF << 4
      bmask = 0xF
      amask = 0xF << 12
      // 4444 format
    case 32:
      rmask = 0xFF << 16
      gmask = 0xFF << 8
      bmask = 0xFF
      amask = 0xFF << 24
      // 8888 format    
    default: // set bbp to 32 bits if nothing better is found
      rmask = 0xFF << 16
      gmask = 0xFF << 8
      bmask = 0xFF
      amask = 0xFF << 24
      bpp = 32
  } 
  return rmask, gmask, bmask, amask, bpp
}  
      


// Returns true if the surface has already been deallocated,
// not allocated yet or is otherwise unusable, false if not
func (surface * Surface) Destroyed() (bool) {
  return surface.surface == nil
}

// Cleans up the memory associated with the surface
func (surface * Surface) Free() {
  if surface.Destroyed() { return }
  freeSurface(surface.surface)
} 

// Loads a surface from a .png, .jpg, .bmp or .xcf file. 
// Remeber to call Display or DisplayAlpha on it to speed up 
// it's blitting speed! 
func LoadSurface(filename string) (*Surface) {
  img := imgLoad(filename)
  return wrapSurface(img)
}

// Makes sure that the given rectangle of the screen is updated
// Don't call on Surfaces that are not screens. 
func (screen * Surface) UpdateRect(x, y int32, w, h uint32) { 
  updateRect(screen.surface, x, y, w, h)  
}

// Makes sure that the entire screen is updated
func (screen * Surface) Flip() {
  flip(screen.surface)
} 

// Maps the color with the given r, g, b components to the color format
// of this surface
func (surface * Surface) MapRGB(r, g, b uint8) uint32 {
  return mapRGB(surface.surface.format, r, g, b)
}

// Maps the color with the given r, g, b, a components to 
// the color format of this surface
func (surface * Surface) MapRGBA(r, g, b, a uint8) uint32 {
  return mapRGBA(surface.surface.format, r, g, b, a)
}

// Locks the surface for duirect drawuing to it's pixels 
// Does nothing if it's not needed to lock this surface
// Important note: according to SDL docs, you may not do any OS calls 
// or threading during the lock on a surface.
func (surface * Surface) Lock() (bool) {
  if !mustLock(surface.surface) { return false }
  lockSurface(surface.surface)
  return true
}

// Unlocks the surface after direct drawing
func (surface * Surface) Unlock() (bool) {
  if !mustLock(surface.surface) { return false }
  unlockSurface(surface.surface)
  return true
}

// Sets the color key for this surface. If flag contains RLEACCEL
// the bitmap will be rle acellerated as well 
func (surface * Surface) SetColorKey(flag, key uint32) int { 
  return setColorKey(surface.surface, flag, key)
}

// Sets the global alpha level for this surface. If flag 
// contains RLEACCEL the bitmap will be rle acellerated as well. 
func (surface * Surface) SetAlpha(flag uint32, alpha uint8) int {
  return setAlpha(surface.surface, flag, alpha)
}

//Converts a Rect to an SDL rect or to nil depending on 
//rect being nil itself or not
// useful for APIs that allow a nil rect to mean "all"  
func nilOrRect(rect * Rect) (* C.SDL_Rect) {
  var sdlrect * C.SDL_Rect = nil
  if  rect != nil { sdlrect = rect.toSDL() }
  return sdlrect
}

// Sets the clipping rectangle for the destination surface in a blit.
// If the clip rectangle is nil, clipping will be disabled.
func (surface * Surface) SetClipRect(rect * Rect) (bool) { 
  sdlrect := nilOrRect(rect)
  return setClipRect(surface.surface, sdlrect) 
}

// Blits this surface to the dst surface
// if srcrect or dstrect are nil, the whole surface is copied 
func (src * Surface) BlitRect(dst * Surface, srcrect, dstrect *Rect) (int) {
  sdlsrcrect := nilOrRect(srcrect)
  sdldstrect := nilOrRect(dstrect)
  return blitSurface(src.surface, dst.surface, sdlsrcrect, sdldstrect)
}

// This function performs a fast fill of the given rectangle with 'color'
// If 'dstrect' is nil, the whole surface will be filled with 'color'
func (dst * Surface) FillRect(dstrect * Rect, color uint32) (int) {
  sdldstrect := nilOrRect(dstrect)
  return fillRect(dst.surface, sdldstrect, color)
}

// This function takes a surface and copies it to a new surface of the
// pixel format and colors of the video framebuffer, suitable for fast
// blitting onto the display surface.
func (surface * Surface) DisplayFormat()  (* Surface) {
  newsurf := displayFormat(surface.surface) 
  return wrapSurface(newsurf) 
}

// This function takes a surface and copies it to a new surface of the
// pixel format and colors of the video framebuffer, suitable for fast
// alpha blitting onto the display surface.
func (surface * Surface) DisplayFormatAlpha()  (* Surface) {
  newsurf := displayFormat(surface.surface) 
  return wrapSurface(newsurf) 
}

// This function makes the bitmap faster to blit to the screen.
// It will automatically take care of reeing the unaccellerated 
// surface.  
func (surface * Surface) Accellerate()  (* Surface) {
  newsurf := displayFormat(surface.surface) 
  freeSurface(surface.surface) 
  // free the old surface 
  surface.surface = newsurf   
  return surface
}

// This function makes the bitmap faster to Alpha blit to the screen.
// It will automatically take care of reeing the unaccellerated 
// surface.  
func (surface * Surface) AccellerateAlpha()  (* Surface) {
  newsurf := displayFormatAlpha(surface.surface) 
  freeSurface(surface.surface) 
  // free the old surface 
  surface.surface = newsurf   
  return surface
}

// Sets this surface as the surface that the window manager or OS
// should use for this application.
func (icon * Surface) WMSetIcon() {
  wm_SetIcon(icon.surface)  
}

/*
func (event * C.SDL_Event) Type() (uint8) {
  uint8(event.type)
} 
*/

// Event interface wrappers
 
 // What kind (type) of event is it 
func (event * Event) Kind() (uint8) {
  return event.Type
} 

// Polls the event queue, returning nil if no events are available
// or if the event gets caught  
func PollEvent() (* Event) {
  event := &Event{}
  ok    := pollEvent((*C.SDL_Event)(unsafe.Pointer(event)))
  if ok == 0 { return nil; } 
  return event 
}


/*
// Cgo can't access the "type" field of the SDL_Event union type directly.
// This function casts the union into a minimal struct that contains only the
// leading type byte.
/*
func eventType(evt *C.SDL_Event) byte {
  return byte(((*struct {
  _type C.Uint8
  })(unsafe.Pointer(evt)))._type)
}
*/
 
