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
import "fmt"
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
func (rect *Rect) ptoSDL() (* C.SDL_Rect) {
  return (* C.SDL_Rect)(unsafe.Pointer(rect))
}  
  
func (rect Rect) toSDL() (C.SDL_Rect) {
  return *(*C.SDL_Rect)(unsafe.Pointer(&rect))
}

// wraps a rectangle
func wrapRect(rect * C.SDL_Rect) (* Rect) {
  return (* Rect)(unsafe.Pointer(rect))
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
func (color Color) toSDL() (C.SDL_Color) {
  return *(* C.SDL_Color)(unsafe.Pointer(&color))
}

// Converts struct to SDL struct pointer  of same layout  
func (color *Color) ptoSDL() (* C.SDL_Color) {
  return (* C.SDL_Color)(unsafe.Pointer(color))
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
  clean  := func(surf * Surface) { surf.Free(); }
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

// Returns the width of the surface (as a int16) 
func (surface * Surface) W16() int16 {
  return int16(surface.surface.w)
}

// Returns the height of the surface (as a int16) 
func (surface * Surface) H16() int16 {
  return int16(surface.surface.w)
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
  if  rect != nil { sdlrect = rect.ptoSDL() }
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

// Pushes an event to the event queue
func PushEvent(event * Event) {
  pushEvent((*C.SDL_Event)(unsafe.Pointer(event)))
}

// Waits until an event is received and returns it
func WaitEvent() (* Event) {
  var event * Event
  for { // ever 
    event     = PollEvent()
    if event != nil { break }
  }
  return event
}

// TTFont is a TrueType Font
type TTFont struct {
  font * C.TTF_Font
} 

// wraps the C SDL font into a TTFONT
func wrapFont(cfont * C.TTF_Font) (* TTFont) {
  if cfont == nil { return nil }
  result := &TTFont{cfont}
  clean  := func(f * TTFont) { f.Free() } 
  runtime.SetFinalizer(result, clean)
  return result
}

// Returns true if the font is in an unusable state
func (font * TTFont) Destroyed() (bool) {
  return font.font == nil 
}

// Releases the memory associated to this font
func (font * TTFont) Free() {
  if font.Destroyed() { return } 
  TTFCloseFont(font.font)  
}

// Initializes the TTF engine once. Returns true if all is AOK.
func initTTFOnce() (bool) {
  if TTFWasInit() { return true; } 
  res := TTFInit()
  if res < 0 { return false } 
  return true
}

// Loads a truetype font from the named file with the given point size. 
// May return nil on failiure  
func LoadTTFont(filename string, ptsize int) (* TTFont) {
  initTTFOnce() // ensure ttf engine is initialized
  return wrapFont(TTFOpenFont(filename, ptsize))
}

// Loads a truetype font from the named file with the given point size.
// Can load a font from a multi font ttf file by passing the index 
// May return nil on failiure  
func LoadTTFontIndex(filename string, ptsize int, 
      index int32) (* TTFont) {
  initTTFOnce() // ensure ttf engine is initialized
  return wrapFont(TTFOpenFontIndex(filename, ptsize, index))
}


// Get the style of the font
func (font * TTFont) Style() (int) {
  return TTFGetFontStyle(font.font)
}

// Set the style of the font
func (font * TTFont) SetStyle(style int) (int) {
  TTFSetFontStyle(font.font, style)
  return font.Style()
}

// Get the total height of the font - usually equal to point size
func (font * TTFont) Height() (int) {
  return TTFFontHeight(font.font)
}

// Get the offset from the baseline to the top of the font
// This is a positive value, relative to the baseline.
func (font * TTFont) Ascent() (int) {
  return TTFFontAscent(font.font)
}

// Get the offset from the baseline to the bottom of the font
// This is a negative value, relative to the baseline.
func (font * TTFont) Descent() (int) {
  return TTFFontDescent(font.font)
}

// Get the recommended spacing between lines of text for this font
func (font * TTFont) LineSkip() (int) {
  return TTFFontLineSkip(font.font)
}
 
// Get the number of faces of the font
func (font * TTFont) Faces() (int32) {
  return TTFFontFaces(font.font)
}

// Returns true if the font is a monospaced fot, false if not.
func (font * TTFont) IsFixedWidth() (bool) {
  return TTFFontFaceIsFixedWidth(font.font)
}

// Returns the name of the font as a string 
func (font * TTFont) String() (string) {
  return TTFFontFamilyName(font.font)
}

// Returns the name of the font style as a string 
func (font * TTFont) StyleName() (string) {
  return TTFFontFaceStyleName(font.font)
}

// Get the metrics (dimensions) of a glyph
// Retunrs minx, maxx, miny, maxy, advance in that order
func (font * TTFont) Metrics(ch uint16) (int, int, int, int, int) {
  return TTFGlyphMetrics(font.font, ch)
}

// Get the dimensions an UTF-8 encoded text would get when 
// rendered with this font.
// Returns width and height in that order.  
func (font * TTFont) Size(text string) (int, int) {
  return TTFTextSize(font.font, text)
}

// Create an 8-bit palettized surface and render the given text at
// fast quality with the given font and color.  The 0 pixel is the
// colorkey, giving a transparent background, and the 1 pixel is set
// to the text color.
// This function returns the new surface, or NULL if there was an error.
// Works with UTF8 encoded strings.
func (font * TTFont) RenderSolid(text string, color Color) (* Surface) { 
  surf := TTFRenderSolid(font.font, text, color.toSDL()) 
  return wrapSurface(surf)
}

// Create an 8-bit palettized surface and render the given glyph at
// fast quality with the given font and color.  The 0 pixel is the
// colorkey, giving a transparent background, and the 1 pixel is set
// to the text color.  The glyph is rendered without any padding or
// centering in the X direction, and aligned normally in the Y direction.
// This function returns the new surface, or NULL if there was an error.
func (font * TTFont) RenderGlyphSolid(ch uint16, 
      color Color) (* Surface) { 
  surf := TTFRenderGlyphSolid(font.font, ch, color.toSDL()) 
  return wrapSurface(surf)
}

// Create an 8-bit palettized surface and render the given text at
// high quality with the given font and colors.  The 0 pixel is background,
// while other pixels have varying degrees of the foreground color.
// This function returns the new surface, or NULL if there was an error.
func (font * TTFont) RenderShaded(text string, 
  fg, bg Color) (* Surface) { 
  surf := TTFRenderShaded(font.font, text, fg.toSDL(), bg.toSDL()) 
  return wrapSurface(surf)
}

// Create an 8-bit palettized surface and render the given glyph at
// high quality with the given font and colors.  The 0 pixel is background,
// while other pixels have varying degrees of the foreground color.
// The glyph is rendered without any padding or centering in the X
// direction, and aligned normally in the Y direction.
// This function returns the new surface, or NULL if there was an error.
func (font * TTFont) RenderGlyphShaded(ch uint16, 
      fg, bg Color) (* Surface) { 
  surf := TTFRenderGlyphShaded(font.font, ch, fg.toSDL(), bg.toSDL()) 
  return wrapSurface(surf)
}

// Create a 32-bit ARGB surface and render the given text at high quality,
// using alpha blending to dither the font with the given color.
// This function returns the new surface, or NULL if there was an error.
func (font * TTFont) RenderBlended(text string, 
  color Color) (* Surface) { 
  surf := TTFRenderBlended(font.font, text, color.toSDL()) 
  return wrapSurface(surf)
}

// Create a 32-bit ARGB surface and render the given glyph at high quality,
// using alpha blending to dither the font with the given color.
// The glyph is rendered without any padding or centering in the X
// direction, and aligned normally in the Y direction.
// This function returns the new surface, or NULL if there was an error.
func (font * TTFont) RenderGlyphBlended(ch uint16, 
      color Color) (* Surface) { 
  surf := TTFRenderGlyphBlended(font.font, ch, color.toSDL()) 
  return wrapSurface(surf)
}

// Returns whether it's OK to put or get a pixel from these
// coordinates or not 
func (surface * Surface) PixelOK(x, y int16) (bool) {
  if x < 0 || y < 0      { return false } 
  if x >= surface.W16()  { return false }
  if y >= surface.H16()  { return false }
  return true
}

// Short for unsafe.Pointer
type ptr unsafe.Pointer

// Helpers for the drawing primitives. They return a pointer to the 
// location of the pixel

// Returns a pointer that points to the location of the pixel 
// at x and y for a surface s with bbp8
// Does no checks of clipping!
func (s * Surface) pixelptr8(x, y int16) (* uint8) {
  surface:= s.surface  
  pixels := uintptr(ptr(surface.pixels))
  offset := uintptr(y*(int16(surface.pitch)) + x)
  return (* uint8)(ptr(pixels + offset))
}

// Returns a pointer that points to the location of the pixel 
// at x and y for a surface s with bbp16
// Does no checks of clipping!
func (s * Surface) pixelptr16(x, y int16) (* uint16) {
  surface:= s.surface  
  pixels := uintptr(ptr(surface.pixels))
  offset := uintptr(y*(int16(surface.pitch) << 1) + x)
  return (* uint16)(ptr(pixels + offset))
}

// Returns four pointers that point to the location of the 
// r,g,b, and a channels of the pixel  at x and y for a surface s with bbp24,
// in that respective order. Does no checks of clipping!
func (s * Surface) pixelptr24(x, y int16) (*byte, *byte, *byte, *byte) {
  surface:= s.surface
  format := surface.format  
  pixels := uintptr(ptr(surface.pixels))
  offset := uintptr(y*(int16(surface.pitch)) + x*3)
  ptrbase:= pixels + offset
  rptr   := (*uint8)(ptr(ptrbase + uintptr(format.Rshift >> 3)))  
  gptr   := (*uint8)(ptr(ptrbase + uintptr(format.Gshift >> 3)))
  bptr   := (*uint8)(ptr(ptrbase + uintptr(format.Bshift >> 3)))
  aptr   := (*uint8)(ptr(ptrbase + uintptr(format.Ashift >> 3)))
  return rptr, gptr, bptr, aptr
}

// Returns a pointer that points to the location of the pixel
// at x and y for a surface s with bbp32
// Does no checks of clipping
func (s * Surface) pixelptr32(x, y int16) (*uint32) {
  surface:= s.surface    
  pixels := uintptr(ptr(surface.pixels))  
  off    := int32(y*(int16(surface.pitch) << 2) + x)
  offset := uintptr(off)
  fmt.Printf("pixelptr32: %p %p, %d %d, %d %d\n", pixels, offset, off, 
   surface.pitch, x, y)
  return (* uint32)(ptr(pixels + offset))
}

// Putpixel drawing primitives
// Each of them is optimized for speed in a different ituation .They 
// should be called only after calling Lock() on the surface,
// They also do not do /any/ clipping or checking on x and y, so be 
// sure to check them with PixelOK.
// Puts a pixel to a surface with BPP 8   
func (s * Surface) PutPixel8(x, y int16, color uint32) {
  ptr    := s.pixelptr8(x, y) 
  *ptr    = uint8(color)
}

// Puts a pixel to a surface with BPP 16
func (s * Surface) PutPixel16(x, y int16, color uint32) {
  ptr    := s.pixelptr16(x, y)   
  *ptr    = uint16(color)
}

// Puts a pixel to a surface with BPP 24. Relatively slow!
func (s * Surface) PutPixel24(x, y int16, color uint32) {
  format := s.surface.format
  rptr, gptr, bptr, aptr := s.pixelptr24(x, y)
  *rptr   = uint8(color >> uint32(format.Rshift))
  *gptr   = uint8(color >> uint32(format.Gshift))
  *bptr   = uint8(color >> uint32(format.Bshift))
  *aptr   = uint8(color >> uint32(format.Ashift))
}

// Puts a pixel to a surface with BPP 32
func (s * Surface) PutPixel32(x, y int16, color uint32) {
  ptr    := s.pixelptr32(x, y)
  *ptr    = color
}

// Also allow put pixel with precalculated y pitch offset??? 

// Puts a pixel depending on the BytesPerPixel of the target surface
// format. Still doesn't check the x and y coordinates for validity. 
func (s * Surface) PutPixelBBP(x, y int16, color uint32) {
  switch s.surface.format.BytesPerPixel {
    case 1:
      s.PutPixel8(x, y, color)
    case 2:
      s.PutPixel16(x, y, color)
    case 3:
      s.PutPixel24(x, y, color)
    case 4:  
      s.PutPixel32(x, y, color)
  }
}

// Get pixel from
// Gets a pixel from a surface with BPP 8
func (s * Surface) GetPixel8(x, y int16) (color uint32) {
  surface:= s.surface  
  pixels := uintptr(ptr(surface.pixels))
  offset := uintptr(y*(int16(surface.pitch)) + x)
  ptr    := (* uint8)(ptr(pixels + offset))
  return uint32(*ptr)
}

// Gets a pixel from a surface with BPP 16
func (s * Surface) GetPixel16(x, y int16) (color uint32) {
  surface:= s.surface  
  pixels := uintptr(ptr(surface.pixels))
  offset := uintptr(y*(int16(surface.pitch) << 1) + x)
  ptr    := (* uint16)(ptr(pixels + offset))
  return uint32(*ptr)   
}

// Gets a pixel from a surface with BPP 24. Relatively slow!
func (s * Surface) GetPixel24(x, y int16) (color uint32) {
  surface:= s.surface
  format := surface.format  
  pixels := uintptr(ptr(surface.pixels))
  offset := uintptr(y*(int16(surface.pitch)) + x*3)
  ptrbase:= pixels + offset
  rptr   := (*uint8)(ptr(ptrbase + uintptr(format.Rshift >> 3)))  
  gptr   := (*uint8)(ptr(ptrbase + uintptr(format.Gshift >> 3)))
  bptr   := (*uint8)(ptr(ptrbase + uintptr(format.Bshift >> 3)))
  aptr   := (*uint8)(ptr(ptrbase + uintptr(format.Ashift >> 3)))
  color   = uint32(*rptr) << uint32(format.Rshift)
  color  |= uint32(*gptr) << uint32(format.Gshift)
  color  |= uint32(*bptr) << uint32(format.Bshift)
  color  |= uint32(*aptr) << uint32(format.Ashift)
  return color
}

// Gets a pixel from a surface with BPP 32
func (s * Surface) GetPixel32(x, y int16) (color uint32) {
  surface:= s.surface  
  pixels := uintptr(ptr(surface.pixels))
  offset := uintptr(y*(int16(surface.pitch) << 2) + x)
  ptr    := (* uint32)(ptr(pixels + offset))
  return uint32(*ptr)    
}

// Gets a pixel depending on the BytesPerPixel of the target surface
// format. Still doesn't check the x and y coordinates for validity. 
func (s * Surface) GetPixelBBP(x, y int16) (color uint32) {
  switch s.surface.format.BytesPerPixel {
    case 1:
      return s.GetPixel8(x, y)
    case 2:
      return s.GetPixel16(x, y)
    case 3:
      return s.GetPixel24(x, y)
    case 4:  
      return s.GetPixel32(x, y)
    default: 
      return 0 
  }
  return 0
}

// Blends the pixel with the one already there using alpha as a gradation
/*
func (s * Surface) BlendPixel8(x, y int16, color uint32, alpha uint8) {
  surface:= s.surface  
  pixels := uintptr(ptr(surface.pixels))
  offset := uintptr(y*(int16(surface.pitch)) + x)
  ptr    := (* uint8)(ptr(pixels + offset))
  return uint32(*ptr)
}
*/

/*
//==================================================================================
// Put pixel with alpha blending
//==================================================================================
void _PutPixelAlpha(SDL_Surface *surface, Sint16 x, Sint16 y, Uint32 color, Uint8 alpha)
{
  if(x>=sge_clip_xmin(surface) && x<=sge_clip_xmax(surface) && y>=sge_clip_ymin(surface) && y<=sge_clip_ymax(surface)){
    Uint32 Rmask = surface->format->Rmask, Gmask = surface->format->Gmask, Bmask = surface->format->Bmask, Amask = surface->format->Amask;
    Uint32 R,G,B,A=0;
  
    switch (surface->format->BytesPerPixel) {
      case 1: { // Assuming 8-bpp 
        if( alpha == 255 ){
          *((Uint8 *)surface->pixels + y*surface->pitch + x) = color;
        }else{
          Uint8 *pixel = (Uint8 *)surface->pixels + y*surface->pitch + x;
          
          Uint8 dR = surface->format->palette->colors[*pixel].r;
          Uint8 dG = surface->format->palette->colors[*pixel].g;
          Uint8 dB = surface->format->palette->colors[*pixel].b;
          Uint8 sR = surface->format->palette->colors[color].r;
          Uint8 sG = surface->format->palette->colors[color].g;
          Uint8 sB = surface->format->palette->colors[color].b;
          
          dR = dR + ((sR-dR)*alpha >> 8);
          dG = dG + ((sG-dG)*alpha >> 8);
          dB = dB + ((sB-dB)*alpha >> 8);
        
          *pixel = SDL_MapRGB(surface->format, dR, dG, dB);
        }
      }
      break;

      case 2: { // Probably 15-bpp or 16-bpp    
        if( alpha == 255 ){
          *((Uint16 *)surface->pixels + y*surface->pitch/2 + x) = color;
        }else{
          Uint16 *pixel = (Uint16 *)surface->pixels + y*surface->pitch/2 + x;
          Uint32 dc = *pixel;
        
          R = ((dc & Rmask) + (( (color & Rmask) - (dc & Rmask) ) * alpha >> 8)) & Rmask;
          G = ((dc & Gmask) + (( (color & Gmask) - (dc & Gmask) ) * alpha >> 8)) & Gmask;
          B = ((dc & Bmask) + (( (color & Bmask) - (dc & Bmask) ) * alpha >> 8)) & Bmask;
          if( Amask )
            A = ((dc & Amask) + (( (color & Amask) - (dc & Amask) ) * alpha >> 8)) & Amask;

          *pixel= R | G | B | A;
        }
      }
      break;

      case 3: { // Slow 24-bpp mode, usually not used 
        Uint8 *pix = (Uint8 *)surface->pixels + y * surface->pitch + x*3;
        Uint8 rshift8=surface->format->Rshift/8;
        Uint8 gshift8=surface->format->Gshift/8;
        Uint8 bshift8=surface->format->Bshift/8;
        Uint8 ashift8=surface->format->Ashift/8;
        
        
        if( alpha == 255 ){
            *(pix+rshift8) = color>>surface->format->Rshift;
            *(pix+gshift8) = color>>surface->format->Gshift;
            *(pix+bshift8) = color>>surface->format->Bshift;
          *(pix+ashift8) = color>>surface->format->Ashift;
        }else{
          Uint8 dR, dG, dB, dA=0;
          Uint8 sR, sG, sB, sA=0;
          
          pix = (Uint8 *)surface->pixels + y * surface->pitch + x*3;
          
          dR = *((pix)+rshift8); 
                dG = *((pix)+gshift8);
                dB = *((pix)+bshift8);
          dA = *((pix)+ashift8);
          
          sR = (color>>surface->format->Rshift)&0xff;
          sG = (color>>surface->format->Gshift)&0xff;
          sB = (color>>surface->format->Bshift)&0xff;
          sA = (color>>surface->format->Ashift)&0xff;
          
          dR = dR + ((sR-dR)*alpha >> 8);
          dG = dG + ((sG-dG)*alpha >> 8);
          dB = dB + ((sB-dB)*alpha >> 8);
          dA = dA + ((sA-dA)*alpha >> 8);

          *((pix)+rshift8) = dR; 
                *((pix)+gshift8) = dG;
                *((pix)+bshift8) = dB;
          *((pix)+ashift8) = dA;
        }
      }
      break;

      case 4: { // Probably 32-bpp 
        if( alpha == 255 ){
          *((Uint32 *)surface->pixels + y*surface->pitch/4 + x) = color;
        }else{
          Uint32 *pixel = (Uint32 *)surface->pixels + y*surface->pitch/4 + x;
          Uint32 dc = *pixel;
      
          R = ((dc & Rmask) + (( (color & Rmask) - (dc & Rmask) ) * alpha >> 8)) & Rmask;
          G = ((dc & Gmask) + (( (color & Gmask) - (dc & Gmask) ) * alpha >> 8)) & Gmask;
          B = ((dc & Bmask) + (( (color & Bmask) - (dc & Bmask) ) * alpha >> 8)) & Bmask;
          if( Amask )
            A = ((dc & Amask) + (( (color & Amask) - (dc & Amask) ) * alpha >> 8)) & Amask;
          
          *pixel = R | G | B | A;
        }
      }
      break;
    }
  }
}
*/

 


