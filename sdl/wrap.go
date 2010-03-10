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
// import "fmt"
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

// Gets the amount of bits per pixel of the format of this surface
func (surface * Surface) BitsPerPixel() uint8 {
  return uint8(surface.surface.format.BitsPerPixel)
} 
 
// Gets the amount of bytes per pixel of the format of this surface 
func (surface * Surface) BytesPerPixel() uint8 {
  return uint8(surface.surface.format.BytesPerPixel)
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

// Gets the components of the given color
func (surface * Surface) GetRGB(color uint32) (r, g, b uint8)  {
  return getRGB(surface.surface.format, color)
}

// Gets the components of the given color
func (surface * Surface) GetRGBA(color uint32) (r, g, b, a uint8) {
  return getRGBA(surface.surface.format, color)
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

// Blits the entire src surface to the dst surface 
// using x, and y as the coordinates of the upper left corner
func (src * Surface) BlitTo(dst * Surface, x, y int) {
  rect    := Rect{ int16(x), int16(y), 0, 0}
  dstrect := rect.toSDL()  
  blitSurface(src.surface, dst.surface, nil, &dstrect)
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

// converts the event to the right type, or returns nil if it's
// not of this type 
func (event *Event) Keyboard() *KeyboardEvent {
  if event.Type==KEYUP || event.Type==KEYDOWN {
      return (*KeyboardEvent)(ptr(event))
  } 
  return nil;
}
 
func (event *Event) MouseButton() *MouseButtonEvent {
  if event.Type==MOUSEBUTTONDOWN || event.Type==MOUSEBUTTONUP {
      return (*MouseButtonEvent)(ptr(event));
  }
  return nil;
}
 
func (event *Event) MouseMotion() *MouseMotionEvent {
    if event.Type==MOUSEMOTION {
      return (*MouseMotionEvent)(ptr(event));
    } 
    return nil;
} 

func (event *Event) Active() *ActiveEvent {
    if event.Type==ACTIVEEVENT  {
      return (*ActiveEvent)(ptr(event));
    } 
    return nil;
} 

func (event *Event) JoyAxis() *JoyAxisEvent {
    if event.Type==JOYAXISMOTION  {
      return (*JoyAxisEvent)(ptr(event));
    } 
    return nil;
} 

func (event *Event) JoyBall() *JoyBallEvent {
    if event.Type==JOYBALLMOTION {
      return (*JoyBallEvent)(ptr(event));
    } 
    return nil;
} 

func (event *Event) JoyHat() *JoyHatEvent {
    if(event.Type==JOYHATMOTION) {
      return (*JoyHatEvent)(ptr(event));
    } 
    return nil;
} 

func (event *Event) JoyButton() *JoyButtonEvent {
    if event.Type==JOYBUTTONDOWN || event.Type == JOYBUTTONUP  {
      return (*JoyButtonEvent)(ptr(event));
    } 
    return nil;
} 

func (event *Event) Resize() *ResizeEvent {
    if event.Type==VIDEORESIZE {
      return (*ResizeEvent)(ptr(event));
    } 
    return nil;
} 

func (event *Event) Expose() *ExposeEvent {
    if event.Type==VIDEOEXPOSE {
      return (*ExposeEvent)(ptr(event));
    } 
    return nil;
} 

func (event *Event) Quit() *QuitEvent {
    if event.Type==QUIT {
      return (*QuitEvent)(ptr(event));
    } 
    return nil;
} 

func (event *Event) User() *UserEvent {
    if event.Type== USEREVENT {
      return (*UserEvent)(ptr(event));
    } 
    return nil;
} 

func (event *Event) SysWM() *SysWMEvent {
    if event.Type == SYSWMEVENT {
      return (*SysWMEvent)(ptr(event));
    } 
    return nil;
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
  fnt := TTFOpenFont(filename, ptsize)
  return wrapFont(fnt)
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

// Returns true if the pixel is outside the image boundaries
// and lay /not/ be drawn to or read from using the PutPixelXXX
// GetPixelXXX and BlendPixelXXX series of functions.
// Returns false if the pixel is safe to draw to.
// This may seeem conterintuitive, but normal use will be
// if surf.PixelOutside(x, y) { return }    
func (surface * Surface) PixelOutside(x, y int) (bool) {
  if x < 0 || y < 0    { return true } 
  if x >= surface.W()  { return true }
  if y >= surface.H()  { return true }
  return false
}

// Returns true if the pixel lies outside the clipping rectangle
// returns false if not so and it may be drawn
func (surface * Surface) PixelClip(x, y int) (bool) {
  cliprect := surface.surface.clip_rect
  minx     := int(cliprect.x)
  miny     := int(cliprect.y)
  maxx     := int(cliprect.x) + int(cliprect.w)
  maxy     := int(cliprect.y) + int(cliprect.h)
  if x < minx   { return true }
  if y < miny   { return true }
  if x >= maxx  { return true }
  if y >= maxy  { return true }
  return false
}

// Short for unsafe.Pointer
type ptr unsafe.Pointer

// Helpers for the drawing primitives. They return a pointer to the 
// location of the pixel

// Basically, the pixel is always located at  
// uintptr((y*(int(surface.pitch)) + x * bytesperpixel)) offset
// from the surface.pixels pointer
// only the case of 24 bits is more complicated 
// See http://www.libsdl.org/cgi/docwiki.cgi/Pixel_Access
// for more reference. 

// Returns a pointer that points to the location of the pixel 
// at x and y for a surface s with bbp8
// Does no checks for validity of x and y and no clipping!
func (s * Surface) pixelptr8(x, y int) (* uint8) {
  surface:= s.surface  
  pixels := uintptr(ptr(surface.pixels))
  offset := uintptr(y*(int(surface.pitch)) + x)
  return (* uint8)(ptr(pixels + offset))
}

// Returns a pointer that points to the location of the pixel 
// at x and y for a surface s with bbp16
// Does no checks of clipping!
func (s * Surface) pixelptr16(x, y int) (* uint16) {
  surface:= s.surface  
  pixels := uintptr(ptr(surface.pixels))  
  offset := uintptr((y*(int(surface.pitch)) + x<<1))
  return (* uint16)(ptr(pixels + offset))
}

// Returns four pointers that point to the location of the 
// r,g,b, and a channels of the pixel  at x and y for a surface s with bbp24,
// in that respective order. Does no checks of clipping!
func (s * Surface) pixelptr24(x, y int) (*byte, *byte, *byte, *byte) {
  surface:= s.surface
  format := surface.format  
  pixels := uintptr(ptr(surface.pixels))
  offset := uintptr(y*(int(surface.pitch)) + x*3)
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
func (s * Surface) pixelptr32(x, y int) (*uint32) {
  surface:= s.surface    
  pixels := uintptr(ptr(surface.pixels))  
  offset := uintptr((y*(int(surface.pitch)) + x<<2)) 
  // y * int(self.pitch) + x*4, so we use x << 2
  //fmt.Printf("pixelptr32: %p %d, %d %d, %d %d\n", pixels, offset,  
  //surface.pitch, x, y)
  return (* uint32)(ptr(pixels + offset))
}

// Putpixel drawing primitives
// Each of them is optimized for speed in a different ituation .They 
// should be called only after calling Lock() on the surface,
// They also do not do /any/ clipping or checking on x and y, so be 
// sure to check them with PixelOK.
// Puts a pixel to a surface with BPP 8   
func (s * Surface) RawPutPixel8(x, y int, color uint32) {
  ptr    := s.pixelptr8(x, y) 
  *ptr    = uint8(color)
}

// Puts a pixel to a surface with BPP 16
func (s * Surface) RawPutPixel16(x, y int, color uint32) {
  ptr    := s.pixelptr16(x, y)   
  *ptr    = uint16(color)
}

// Puts a pixel to a surface with BPP 24. Relatively slow!
func (s * Surface) RawPutPixel24(x, y int, color uint32) {
  format := s.surface.format
  rptr, gptr, bptr, aptr := s.pixelptr24(x, y)
  *rptr   = uint8(color >> uint32(format.Rshift))
  *gptr   = uint8(color >> uint32(format.Gshift))
  *bptr   = uint8(color >> uint32(format.Bshift))
  *aptr   = uint8(color >> uint32(format.Ashift))
}

// Puts a pixel to a surface with BPP 32
func (s * Surface) RawPutPixel32(x, y int, color uint32) {
  ptr    := s.pixelptr32(x, y)
  *ptr    = color
}

// Also allow put pixel with precalculated y pitch offset??? 

// Puts a pixel depending on the BytesPerPixel of the target surface
// format. Still doesn't check the x and y coordinates for validity. 
func (s * Surface) RawPutPixelBBP(x, y int, color uint32) {
  switch s.surface.format.BytesPerPixel {
    case 1:
      s.RawPutPixel8(x, y, color)
    case 2:
      s.RawPutPixel16(x, y, color)
    case 3:
      s.RawPutPixel24(x, y, color)
    case 4:  
      s.RawPutPixel32(x, y, color)
  }
}

// Get pixel from
// Gets a pixel from a surface with BPP 8
func (s * Surface) RawGetPixel8(x, y int) (color uint32) {
  ptr    := s.pixelptr8(x, y)
  return uint32(*ptr)
}

// Gets a pixel from a surface with BPP 16
func (s * Surface) RawGetPixel16(x, y int) (color uint32) {
  ptr    := s.pixelptr16(x, y)
  return uint32(*ptr)   
}

// Gets a pixel from a surface with BPP 24. Relatively slow!
func (s * Surface) RawGetPixel24(x, y int) (color uint32) {
  format := s.surface.format 
  rptr, gptr, bptr, aptr := s.pixelptr24(x, y)
  color   = uint32(*rptr) << uint32(format.Rshift)
  color  |= uint32(*gptr) << uint32(format.Gshift)
  color  |= uint32(*bptr) << uint32(format.Bshift)
  color  |= uint32(*aptr) << uint32(format.Ashift)
  return color
}

// Gets a pixel from a surface with BPP 32
func (s * Surface) RawGetPixel32(x, y int) (color uint32) {
  ptr    := s.pixelptr8(x, y)
  return uint32(*ptr)    
}

// Gets a pixel depending on the BytesPerPixel of the target surface
// format. Still doesn't check the x and y coordinates for validity. 
func (s * Surface) RawGetPixelBBP(x, y int) (color uint32) {
  switch s.surface.format.BytesPerPixel {
    case 1:
      return s.RawGetPixel8(x, y)
    case 2:
      return s.RawGetPixel16(x, y)
    case 3:
      return s.RawGetPixel24(x, y)
    case 4:  
      return s.RawGetPixel32(x, y)
    default: 
      return 0 
  }
  return 0
}

// Helps blending two colors 
func helpBlend(old, color, 
  rmask, gmask, bmask, amask, alpha uint32) (result uint32) {
  oldr := old & rmask
  oldg := old & gmask
  oldb := old & bmask
  olda := old & amask
  colr := color & rmask
  colg := color & gmask
  colb := color & bmask
  cola := color & amask
  // we add to every component
  // (new - old) * (alpha / 256) 
  // ((new - old) * alpha) >> 8)
  r    := (oldr + (((colr - oldr) * alpha) >> 8)) & rmask 
  g    := (oldg + (((colg - oldg) * alpha) >> 8)) & gmask
  b    := (oldb + (((colb - oldb) * alpha) >> 8)) & bmask
  a    := uint32(0)
  if amask > 0  {
    a    = (olda + (((cola - olda) * alpha) >> 8)) & amask
  }
  return  r | g | b | a
}

// Helps blending one component of a color
func helpBlendComponent(oldc, newc, alpha uint8) (uint8) {
  return oldc + (((newc-oldc)*alpha) >> 8)
}

// Returns the color masks of this surface's format in order r,g,b,a
func (s * Surface) ColorMasks() (uint32, uint32, uint32, uint32) {
  f      := s.surface.format
  return uint32(f.Rmask), uint32(f.Gmask), 
         uint32(f.Bmask), uint32(f.Amask)
}

// Blends a pixel with the one already there to a surface with BPP 8
// taking into account the value of alpha
func (s * Surface) RawBlendPixel8(x, y int, 
  color uint32, alpha uint8) {    
  ptr    := s.pixelptr8(x, y)
  old    := uint32(*ptr)
  oldr, oldg, oldb := s.GetRGB(old)
  // Messing with the palette would probably be faster, 
  // but it's cleaner to do it like this.
  colr, colg, colb := s.GetRGB(color)
  newr   := helpBlendComponent(oldr, colr, alpha)
  newg   := helpBlendComponent(oldg, colg, alpha)
  newb   := helpBlendComponent(oldb, colb, alpha)  
  newcol := s.MapRGB(newr, newg, newb)
  // And we have to do the same for the new color anyway, so... 
  *ptr    = uint8(newcol)
}


// Blends a pixel a pixel from a surface with BPP 24, taking into
// consideration the value of alpha. Relatively slow!
func (s * Surface) RawBlendPixel24(x, y int, color uint32, alpha uint8) {
  format := s.surface.format 
  rptr, gptr, bptr, aptr := s.pixelptr24(x, y)
  oldr   := *rptr
  oldg   := *gptr
  oldb   := *bptr
  olda   := *aptr
  colr   := uint8((color >> uint32(format.Rshift)) & 0xff)
  colg   := uint8((color >> uint32(format.Gshift)) & 0xff)
  colb   := uint8((color >> uint32(format.Bshift)) & 0xff)
  cola   := uint8((color >> uint32(format.Ashift)) & 0xff)
  newr   := helpBlendComponent(oldr, colr, alpha)
  newg   := helpBlendComponent(oldg, colg, alpha)
  newb   := helpBlendComponent(oldb, colb, alpha)
  newa   := helpBlendComponent(olda, cola, alpha)
  *rptr   = uint8(newr)
  *gptr   = uint8(newg)
  *bptr   = uint8(newb)
  *aptr   = uint8(newa)  
}

// Blends a pixel with the one already there to a surface with BPP 16
// taking into account the value of alpha
func (s * Surface) RawBlendPixel16(x, y int, 
  color uint32, alpha uint8) {
  rmask, gmask, bmask, amask := s.ColorMasks()  
  ptr    := s.pixelptr16(x, y)
  old    := uint32(*ptr)
  newcol := helpBlend(old, color, rmask,
            gmask, bmask, amask, uint32(alpha))
  *ptr    = uint16(newcol)
}


// Blends a pixel with the one already there to a surface with BPP 32
// taking into account the value of alpha
func (s * Surface) RawBlendPixel32(x, y int, 
  color uint32, alpha uint8) {
  rmask, gmask, bmask, amask := s.ColorMasks()  
  ptr    := s.pixelptr32(x, y)
  old    := *ptr
  newcol := helpBlend(old, color, rmask,
            gmask, bmask, amask, uint32(alpha))  
  *ptr    = newcol
}

// Puts a pixel depending on the BytesPerPixel of the target surface
// format. Still doesn't check the x and y coordinates for validity. 
func (s * Surface) RawBlendPixelBBP(x, y int, color uint32, 
    alpha uint8) {
    switch s.surface.format.BytesPerPixel {
    case 1:
      s.RawBlendPixel8(x, y, color, alpha)
    case 2:
      s.RawBlendPixel16(x, y, color, alpha)
    case 3:
      s.RawBlendPixel24(x, y, color, alpha)
    case 4:  
      s.RawBlendPixel32(x, y, color, alpha)
    default: 
      // Do nothing.
  }
}


// PutPixel for surfaces with 8 Bytes per pixel.
// Unsafe on the wrong BBP.
// It does protect against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.  
func (s * Surface) InPutPixel8(x, y int, color uint32) {
  if s.PixelOutside(x, y) { return }
  s.RawPutPixel8(x, y, color)
} 

// PutPixel for surfaces with 16 Bytes per pixel.
// Unsafe on the wrong BBP.
// It does protect against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.  
func (s * Surface) InPutPixel16(x, y int, color uint32) {
  if s.PixelOutside(x, y) { return }
  s.RawPutPixel16(x, y, color)
} 

// PutPixel for surfaces with 24 Bytes per pixel.
// Unsafe on the wrong BBP.
// It does protect against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.  
func (s * Surface) InPutPixel24(x, y int, color uint32) {
  if s.PixelOutside(x, y) { return }
  s.RawPutPixel24(x, y, color)
} 

// PutPixel for surfaces with 32 Bytes per pixel.
// Unsafe on the wrong BBP.
// It does protect against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.  
func (s * Surface) InPutPixel32(x, y int, color uint32) {
  if s.PixelOutside(x, y) { return }
  s.RawPutPixel32(x, y, color)
} 

// PutPixel for surfaces of all depth
// It protects against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.  
func (s * Surface) InPutPixelBBP(x, y int, color uint32) {
  if s.PixelOutside(x, y) { return }
  s.RawPutPixelBBP(x, y, color)
}  

// GetPixel for surfaces with 8 Bytes per pixel.
// Unsafe on the wrong BBP.
// It does protect against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.
// Returns 0 if x and y are out of bounds  
func (s * Surface) InGetPixel8(x, y int) (color uint32) {
  if s.PixelOutside(x, y) { return 0 }
  return s.RawGetPixel8(x, y)
} 

// GetPixel for surfaces with 16 Bytes per pixel.
// Unsafe on the wrong BBP.
// It does protect against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.
// Returns 0 if x and y are out of bounds  
func (s * Surface) InGetPixel16(x, y int) (color uint32) {
  if s.PixelOutside(x, y) { return 0 }
  return s.RawGetPixel16(x, y)
} 

// GetPixel for surfaces with 24 Bytes per pixel.
// Unsafe on the wrong BBP.
// It does protect against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.
// Returns 0 if x and y are out of bounds  
func (s * Surface) InGetPixel24(x, y int) (color uint32) {
  if s.PixelOutside(x, y) { return 0 }
  return s.RawGetPixel24(x, y)
} 

// GetPixel for surfaces with 32 Bytes per pixel.
// Unsafe on the wrong BBP.
// It does protect against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.
// Returns 0 if x and y are out of bounds  
func (s * Surface) InGetPixel32(x, y int) (color uint32) {
  if s.PixelOutside(x, y) { return 0 }
  return s.RawGetPixel32(x, y)
} 

// PutPixel for surfaces of all depth
// It protects against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.
// Returns 0 if x and y are out of bounds  
func (s * Surface) InGetPixelBBP(x, y int) (color uint32) {
  if s.PixelOutside(x, y) { return 0 }
  return s.RawGetPixelBBP(x, y)
}  


// BlendPixel for surfaces with 8 Bytes per pixel.
// Unsafe on the wrong BBP.
// It does protect against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.  
func (s * Surface) InBlendPixel8(x, y int, color uint32, alpha uint8) {
  if s.PixelOutside(x, y) { return }
  s.RawBlendPixel8(x, y, color, alpha)
} 

// BlendPixel for surfaces with 16 Bytes per pixel.
// Unsafe on the wrong BBP.
// It does protect against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.  
func (s * Surface) InBlendPixel16(x, y int, color uint32, alpha uint8) {
  if s.PixelOutside(x, y) { return }
  s.RawBlendPixel16(x, y, color, alpha)
} 

// BlendPixel for surfaces with 24 Bytes per pixel.
// Unsafe on the wrong BBP.
// It does protect against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.  
func (s * Surface) InBlendPixel24(x, y int, color uint32, alpha uint8) {
  if s.PixelOutside(x, y) { return }
  s.RawBlendPixel24(x, y, color, alpha)
} 

// BlendPixel for surfaces with 32 Bytes per pixel.
// Unsafe on the wrong BBP.
// It does protect against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.  
func (s * Surface) InBlendPixel32(x, y int, color uint32, alpha uint8) {
  if s.PixelOutside(x, y) { return }
  s.RawBlendPixel32(x, y, color, alpha)
} 

// PutPixel for surfaces of all depth
// It protects against out of bounds x and y values, 
// hence the name starts with "In". 
// However, it does not take clipping into consideration.  
func (s * Surface) InBlendPixelBBP(x, y int, color uint32, 
  alpha uint8) {
  if s.PixelOutside(x, y) { return }
  s.RawBlendPixelBBP(x, y, color, alpha)
}  


// Puts a pixel with the given color at the given coordinates
// Takes the clipping rectangle and surface bounds into consideration
// Locks and unlocks the surface if that is needed for drawing
func (s * Surface) PutPixel(x, y int, color uint32) {
  if s.PixelClip(x, y) { return }
  s.Lock()
  s.InPutPixelBBP(x, y, color)
  s.Unlock() 
}

// Gets the color of a pixel from this surface 
// Returns 0 if the pixel is outside of the clipping rectangle,
// or outside of the bounds of the surface
// Locks and unlocks the surface if that is needed
func (s * Surface) GetPixel(x, y int) (color uint32) {
  if s.PixelClip(x, y) { return 0 }
  s.Lock()
  res := s.InGetPixelBBP(x, y)
  s.Unlock()
  return res 
}

// Blends the color of a pixel from this surface, 
// taking alpha into consideration. 
// Returns 0 if the pixel is outside of the clipping rectangle,
// or outside of the bounds of the surface.
// Locks and unlocks the surface if that is needed for drawing
func (s * Surface) BlendPixel(x, y int, color uint32, alpha uint8) {
  if s.PixelClip(x, y) { return }
  s.Lock()
  s.InBlendPixelBBP(x, y, color, alpha)
  s.Unlock() 
}

