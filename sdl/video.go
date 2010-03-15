//
// Go Language wrappers around SDL
//
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
//import "unsafe"
//import "runtime"

// Transparency definitions: These define alpha as the opacity of a surface
const ALPHA_OPAQUE	=255
const ALPHA_TRANSPARENT	=0

// Useful data types 
/*
typedef struct SDL_Rect {
	Sint16 x, y;
	Uint16 w, h;
} SDL_Rect;

typedef struct SDL_Color {
	Uint8 r;
	Uint8 g;
	Uint8 b;
	Uint8 unused;
} SDL_Color;
#define SDL_Colour SDL_Color

typedef struct SDL_Palette {
	int       ncolors;
	SDL_Color *colors;
} SDL_Palette;

// Everything in the pixel format structure is read-only 
typedef struct SDL_PixelFormat {
	SDL_Palette *palette;
	Uint8  BitsPerPixel;
	Uint8  BytesPerPixel;
	Uint8  Rloss;
	Uint8  Gloss;
	Uint8  Bloss;
	Uint8  Aloss;
	Uint8  Rshift;
	Uint8  Gshift;
	Uint8  Bshift;
	Uint8  Ashift;
	Uint32 Rmask;
	Uint32 Gmask;
	Uint32 Bmask;
	Uint32 Amask;

	// RGB color key information 
	Uint32 colorkey;
	// Alpha value information (per-surface alpha) 
	Uint8  alpha;
} SDL_PixelFormat;

// This structure should be treated as read-only, except for 'pixels',
// which, if not NULL, contains the raw pixel data for the surface.
typedef struct SDL_Surface {
	Uint32 flags;				// Read-only 
	SDL_PixelFormat *format;		// Read-only
	int w, h;				// Read-only
	Uint16 pitch;				// Read-only
	void *pixels;				// Read-write 
	int offset;				// Private 

	// Hardware-specific surface info
	struct private_hwdata *hwdata;

	// clipping information 
	SDL_Rect clip_rect;			// Read-only //
	Uint32 unused1;				// for binary compatibility //

	// Allow recursive locks 
	Uint32 locked;				// Private //

	// info for fast blit mapping to other surfaces 
	struct SDL_BlitMap *map;		// Private //

	// format version, bumped at every change to invalidate blit maps 
	unsigned int format_version;		// Private 

	// Reference count -- used when freeing surface 
	int refcount;				// Read-mostly 
} SDL_Surface;
*/
// These are the currently supported flags for the SDL_surface 
// Available for CreateRGBSurface() or SetVideoMode() 
// Surface is in system memory 
const  SWSURFACE =0x00000000	
// Surface is in video memory 
const  HWSURFACE =0x00000001	
// Use asynchronous blits if possible 
const  ASYNCBLIT =0x00000004	
// Available for SDL_SetVideoMode() 
// Allow any video depth/pixel-format 
const  ANYFORMAT =0x10000000	
// Surface has exclusive palette 
const  HWPALETTE =0x20000000	
// Set up double-buffered video mode 
const  DOUBLEBUF =0x40000000	
// Surface is a full screen display 
const  FULLSCREEN=0x80000000	
// Create an OpenGL rendering context 
const  OPENGL     =0x00000002      
// Create an OpenGL rendering context and use it for blitting 
const  OPENGLBLIT=0x0000000A	
// This video mode may be resized 
const  RESIZABLE =0x00000010	
// No window caption or edge frame 
const  NOFRAME   =0x00000020	
// Used internally (read-only) 
// Blit uses hardware acceleration 
const  HWACCEL    = 0x00000100	
// Blit uses a source color key 
const  SRCCOLORKEY=	0x00001000	
// Private flag 
const  RLEACCELOK =	0x00002000	
// Surface is RLE encoded 
const  RLEACCEL   =	0x00004000	
// Blit uses source alpha blending 
const  SRCALPHA   =	0x00010000	
// Surface uses preallocated memory 
const  PREALLOC   =	0x01000000	

// Evaluates to true if the surface needs to be locked before access 
func mustLock(surface * C.SDL_Surface) (bool) {
  aid := uint32(surface.flags) & (HWSURFACE|ASYNCBLIT|RLEACCEL)
  off := uint32(surface.offset) 
  res := (off!=0) || (aid!=0)
  return res
}


// Useful for determining the video hardware capabilities 
/*
typedef struct SDL_VideoInfo {
	Uint32 hw_available :1;	// Flag: Can you create hardware surfaces? 
	Uint32 wm_available :1;	// Flag: Can you talk to a window manager? 
	Uint32 UnusedBits1  :6;
	Uint32 UnusedBits2  :1;
	Uint32 blit_hw      :1;	// Flag: Accelerated blits HW --> HW 
	Uint32 blit_hw_CC   :1;	// Flag: Accelerated blits with Colorkey 
	Uint32 blit_hw_A    :1;	// Flag: Accelerated blits with Alpha 
	Uint32 blit_sw      :1;	// Flag: Accelerated blits SW --> HW 
	Uint32 blit_sw_CC   :1;	// Flag: Accelerated blits with Colorkey 
	Uint32 blit_sw_A    :1;	// Flag: Accelerated blits with Alpha 
	Uint32 blit_fill    :1;	// Flag: Accelerated color fill 
	Uint32 UnusedBits3  :16;
	Uint32 video_mem;	// The total amount of video memory (in K) 
	SDL_PixelFormat *vfmt;	// Value: The format of the video surface 
	int    current_w;	// Value: The current video mode width 
	int    current_h;	// Value: The current video mode height 
} SDL_VideoInfo;
*/

// The most common video overlay formats.
// For an explanation of these pixel formats, see:
// http://www.webartz.com/fourcc/indexyuv.htm
// For information on the relationship between color spaces, see:
// http://www.neuro.sfc.keio.ac.jp/~aly/polygon/info/color-space-faq.html

const  YV12_OVERLAY = 0x32315659	// Planar mode: Y + V + U  (3 planes) 
const  IYUV_OVERLAY = 0x56555949	// Planar mode: Y + U + V  (3 planes) 
const  YUY2_OVERLAY = 0x32595559	// Packed mode: Y0+U0+Y1+V0 (1 plane) 
const  UYVY_OVERLAY = 0x59565955	// Packed mode: U0+Y0+V0+Y1 (1 plane) 
const  YVYU_OVERLAY = 0x55595659	// Packed mode: Y0+V0+Y1+U0 (1 plane) 

// The YUV hardware video overlay 
/*
typedef struct SDL_Overlay {
	Uint32 format;				// Read-only 
	int w, h;				// Read-only 
	int planes;				// Read-only 
	Uint16 *pitches;			// Read-only 
	Uint8 **pixels;				// Read-write 

	// Hardware-specific surface info 
	struct private_yuvhwfuncs *hwfuncs;
	struct private_yuvhwdata *hwdata;

	// Special flags 
	Uint32 hw_overlay :1;	// Flag: This overlay hardware accelerated? 
	Uint32 UnusedBits :31;
} SDL_Overlay;
*/


// Public enumeration for setting the OpenGL window attributes. 
const (
    GL_RED_SIZE = iota
    GL_GREEN_SIZE
    GL_BLUE_SIZE
    GL_ALPHA_SIZE
    GL_BUFFER_SIZE
    GL_DOUBLEBUFFER
    GL_DEPTH_SIZE
    GL_STENCIL_SIZE
    GL_ACCUM_RED_SIZE
    GL_ACCUM_GREEN_SIZE
    GL_ACCUM_BLUE_SIZE
    GL_ACCUM_ALPHA_SIZE
    GL_STEREO
    GL_MULTISAMPLEBUFFERS
    GL_MULTISAMPLESAMPLES
    GL_ACCELERATED_VISUAL
    GL_SWAP_CONTROL
)

// flags for SDL_SetPalette() 
const  SDL_LOGPAL	=0x01
const  SDL_PHYSPAL	=0x02

// Function prototypes 
// These functions are used internally, and should not be used unless you
// have a specific need to specify the video driver you want to use.
// You should normally use SDL_Init() or SDL_InitSubSystem().
//
// SDL_VideoInit() initializes the video subsystem -- sets up a connection
// to the window manager, etc, and determines the current video mode and
// pixel format, but does not initialize a window or graphics mode.
// Note that event handling is activated by this routine.
//
// If you use both sound and video in your application, you need to call
// SDL_Init() before opening the sound device, otherwise under Win32 DirectX,
// you won't be able to set full-screen display modes.
func VideoInit(driver string, flags uint32) (int) { 
  cdrv := cstr(driver) 
  return int(C.SDL_VideoInit(cdrv, C.Uint32(flags)))
}

func VideoQuit() {
  C.SDL_VideoQuit()
}

// This function fills the given character buffer with the name of the
// video driver, and returns a pointer to it if the video driver has
// been initialized.  It returns NULL if no driver has been initialized.
func VideoDriverName() (string)  {
  maxlen := 255
  namebuf:= cstrNew(maxlen) ; defer namebuf.free()
  res := C.SDL_VideoDriverName(namebuf, C.int(maxlen))
  if res == nil { return "Video Not Initialized" }
  return C.GoString(res)
}

// This function returns a pointer to the current display surface.
// If SDL is doing format conversion on the display surface, this
// function returns the publicly visible surface, not the real video
// surface.
func getVideoSurface() (* C.SDL_Surface ) {
  return C.SDL_GetVideoSurface()
}

// This function returns a read-only pointer to information about the
// video hardware.  If this is called before SetVideoMode(), the 'vfmt'
// member of the returned structure will contain the pixel format of the
// "best" video mode.
func getVideoInfo() (* C.SDL_VideoInfo) { 
  return C.SDL_GetVideoInfo()
}

// Check to see if a particular video mode is supported.
// It returns 0 if the requested mode is not supported under any bit depth,
// or returns the bits-per-pixel of the closest available mode with the
// given width and height.  If this bits-per-pixel is different from the
// one used when setting the video mode, SetVideoMode() will succeed,
// but will emulate the requested bits-per-pixel with a shadow surface.
//
// The arguments to SDL_VideoModeOK() are the same ones you would pass to
// SDL_SetVideoMode()
func VideoModeOK(width, height, bpp int, flags uint32) bool { 
  return i2b(int(C.SDL_VideoModeOK(C.int(width),C.int(height),
	    C.int(bpp), C.Uint32(flags))))
}

// Return a pointer to an array of available screen dimensions for the
// given format and video flags, sorted largest to smallest.  Returns 
// NULL if there are no dimensions available for a particular format, 
// or (SDL_Rect **)-1 if any dimension is okay for the given format.
// If 'format' is NULL, the mode list will be for the format given 
// by SDL_GetVideoInfo()->vfmt
func ListModes() {
} 
// XXX: todo  
// extern DECLSPEC SDL_Rect ** SDLCALL SDL_ListModes(SDL_PixelFormat *format, Uint32 flags);

// Set up a video mode with the specified width, height and bits-per-pixel.
//
// If 'bpp' is 0, it is treated as the current display bits per pixel.
//
// If SDL_ANYFORMAT is set in 'flags', the SDL library will try to set the
// requested bits-per-pixel, but will return whatever video pixel format is
// available.  The default is to emulate the requested pixel format if it
// is not natively available.
//
// If SDL_HWSURFACE is set in 'flags', the video surface will be placed in
// video memory, if possible, and you may have to call SDL_LockSurface()
// in order to access the raw framebuffer.  Otherwise, the video surface
// will be created in system memory.
//
// If SDL_ASYNCBLIT is set in 'flags', SDL will try to perform rectangle
// updates asynchronously, but you must always lock before accessing pixels.
// SDL will wait for updates to complete before returning from the lock.
//
// If SDL_HWPALETTE is set in 'flags', the SDL library will guarantee
// that the colors set by SDL_SetColors() will be the colors you get.
// Otherwise, in 8-bit mode, SDL_SetColors() may not be able to set all
// of the colors exactly the way they are requested, and you should look
// at the video surface structure to determine the actual palette.
// If SDL cannot guarantee that the colors you request can be set, 
// i.e. if the colormap is shared, then the video surface may be created
// under emulation in system memory, overriding the SDL_HWSURFACE flag.
//
// If SDL_FULLSCREEN is set in 'flags', the SDL library will try to set
// a fullscreen video mode.  The default is to create a windowed mode
// if the current graphics system has a window manager.
// If the SDL library is able to set a fullscreen video mode, this flag 
// will be set in the surface that is returned.
//
// If SDL_DOUBLEBUF is set in 'flags', the SDL library will try to set up
// two surfaces in video memory and swap between them when you call 
// SDL_Flip().  This is usually slower than the normal single-buffering
// scheme, but prevents "tearing" artifacts caused by modifying video 
// memory while the monitor is refreshing.  It should only be used by 
// applications that redraw the entire screen on every update.
//
// If SDL_RESIZABLE is set in 'flags', the SDL library will allow the
// window manager, if any, to resize the window at runtime.  When this
// occurs, SDL will send a SDL_VIDEORESIZE event to you application,
// and you must respond to the event by re-calling SDL_SetVideoMode()
// with the requested size (or another size that suits the application).
//
// If SDL_NOFRAME is set in 'flags', the SDL library will create a window
// without any title bar or frame decoration.  Fullscreen video modes have
// this flag set automatically.
//
// This function returns the video framebuffer surface, or NULL if it fails.
//
// If you rely on functionality provided by certain video flags, check the
// flags of the returned surface to make sure that functionality is available.
// SDL will fall back to reduced functionality if the exact flags you wanted
// are not available.
func setVideoMode(width, height, bpp int, flags uint32) (* C.SDL_Surface) {
  return C.SDL_SetVideoMode(C.int(width), C.int(height), 
  C.int(bpp), C.Uint32(flags))
} 

// Makes sure the given list of rectangles is updated on the given screen.
// If 'x', 'y', 'w' and 'h' are all 0, SDL_UpdateRect will update the entire
// screen.
// These functions should not be called while 'screen' is locked.
// extern DECLSPEC void SDLCALL SDL_UpdateRects
//		(SDL_Surface *screen, int numrects, SDL_Rect *rects);
func updateRect(screen * C.SDL_Surface, x, y int32, w, h uint32) { 
  C.SDL_UpdateRect(screen, C.Sint32(x), C.Sint32(y), C.Uint32(w), C.Uint32(h))
}

// On hardware that supports double-buffering, this function sets up a flip
// and returns.  The hardware will wait for vertical retrace, and then swap
// video buffers before the next video surface blit or lock will return.
// On hardware that doesn not support double-buffering, this is equivalent
// to calling SDL_UpdateRect(screen, 0, 0, 0, 0);
// The SDL_DOUBLEBUF flag must have been passed to SDL_SetVideoMode() when
// setting the video mode for this function to perform hardware flipping.
// This function returns 0 if successful, or -1 if there was an error.
func flip(screen * C.SDL_Surface) { 
  C.SDL_Flip(screen)
}

// Set the gamma correction for each of the color channels.
// The gamma values range (approximately) between 0.1 and 10.0
// 
// If this function isn't supported directly by the hardware, it will
// be emulated using gamma ramps, if available.  If successful, this
// function returns 0, otherwise it returns -1.
func SetGamma(red, green, blue float) (int) { 
  return int(C.SDL_SetGamma(C.float(red), C.float(green), C.float(blue)))
}

// Set the gamma translation table for the red, green, and blue channels
// of the video hardware.  Each table is an array of 256 16-bit quantities,
// representing a mapping between the input and output for that channel.
// The input is the index into the array, and the output is the 16-bit
// gamma value at that index, scaled to the output color precision.
// 
// You may pass NULL for any of the channels to leave it unchanged.
// If the call succeeds, it will return 0.  If the display driver or
// hardware does not support gamma translation, or otherwise fails,
// this function will return -1.
// XXX: todo
// extern DECLSPEC int SDLCALL SDL_SetGammaRamp(const Uint16 *red, const Uint16 *green, const Uint16 *blue);
//

// Retrieve the current values of the gamma translation tables.
// 
// You must pass in valid pointers to arrays of 256 16-bit quantities.
// Any of the pointers may be NULL to ignore that channel.
// If the call succeeds, it will return 0.  If the display driver or
// hardware does not support gamma translation, or otherwise fails,
// this function will return -1.
// XXX: todo 
// extern DECLSPEC int SDLCALL SDL_GetGammaRamp(Uint16 *red, Uint16 *green, Uint16 *blue);

// Sets a portion of the colormap for the given 8-bit surface. If 'surface'
// is not a palettized surface, this function does nothing, returning 0.
// If all of the colors were set as passed to SDL_SetColors(), it will
// return 1. If not all the color entries were set exactly as given,
// it will return 0, and you should look at the surface palette to
// determine the actual color palette.
//
// When 'surface' is the surface associated with the current display, the
// display colormap will be updated with the requested colors.  If 
// SDL_HWPALETTE was set in SDL_SetVideoMode() flags, SDL_SetColors()
// will always return 1, and the palette is guaranteed to be set the way
// you desire, even if the window colormap has to be warped or run under
// emulation.
// extern DECLSPEC int SDLCALL SDL_SetColors(SDL_Surface *surface, 
//		SDL_Color *colors, int firstcolor, int ncolors);


// Sets a portion of the colormap for a given 8-bit surface.
// 'flags' is one or both of:
// SDL_LOGPAL  -- set logical palette, which controls how blits are mapped
//                to/from the surface,
// SDL_PHYSPAL -- set physical palette, which controls how pixels look on
//                the screen
// Only screens have physical palettes. Separate change of physical/logical
// palettes is only possible if the screen has SDL_HWPALETTE set.
//
// The return value is 1 if all colours could be set as requested, and 0
// otherwise.
//
// SDL_SetColors() is equivalent to calling this function with
//     flags = (SDL_LOGPAL|SDL_PHYSPAL).
// XXX: todo
// extern DECLSPEC int SDLCALL SDL_SetPalette(SDL_Surface *surface, int flags,
//				   SDL_Color *colors, int firstcolor,
//				   int ncolors);


// Maps an RGB triple to an opaque pixel value for a given pixel format
func mapRGB(format * C.SDL_PixelFormat, r, g, b uint8) uint32 {
  return uint32(C.SDL_MapRGB(format, C.Uint8(r), C.Uint8(g), C.Uint8(b)))
}

// Maps an RGBA quadruple to a pixel value for a given pixel format
func mapRGBA(format * C.SDL_PixelFormat, r, g, b, a uint8) uint32 {
  return uint32(C.SDL_MapRGBA(format, C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
}
 
// Maps a pixel value into the RGB components for a given pixel format
func getRGB(format * C.SDL_PixelFormat, pixel uint32) (byte, byte, byte) { 
  var r, b, g uint8 
  pr := cbyteptr(&r)
  pb := cbyteptr(&b)
  pg := cbyteptr(&g)
  C.SDL_GetRGB(C.Uint32(pixel), format, pr, pg, pb)
  return r, g, b
}

// Maps a pixel value into the RGBA components for a given pixel format
func getRGBA(format * C.SDL_PixelFormat, pixel uint32) (byte, byte, byte, byte) { 
  var r, b, g, a uint8 
  pr := cbyteptr(&r)
  pb := cbyteptr(&b)
  pg := cbyteptr(&g)
  pa := cbyteptr(&a)
  C.SDL_GetRGBA(C.Uint32(pixel), format, pr, pg, pb, pa)
  return r, g, b, a
}
 
// Allocate and free an RGB surface (must be called after SDL_SetVideoMode)
// If the depth is 4 or 8 bits, an empty palette is allocated for the surface.
// If the depth is greater than 8 bits, the pixel format is set using the
// flags '[RGB]mask'.
// If the function runs out of memory, it will return NULL.
//
// The 'flags' tell what kind of surface to create.
// SDL_SWSURFACE means that the surface should be created in system memory.
// SDL_HWSURFACE means that the surface should be created in video memory,
// with the same format as the display surface.  This is useful for surfaces
// that will not change much, to take advantage of hardware acceleration
// when being blitted to the display surface.
// SDL_ASYNCBLIT means that SDL will try to perform asynchronous blits with
// this surface, but you must always lock it before accessing the pixels.
// SDL will wait for current blits to finish before returning from the lock.
// SDL_SRCCOLORKEY indicates that the surface will be used for colorkey blits.
// If the hardware supports acceleration of colorkey blits between
// two surfaces in video memory, SDL will try to place the surface in
// video memory. If this isn't possible or if there is no hardware
// acceleration available, the surface will be placed in system memory.
// SDL_SRCALPHA means that the surface will be used for alpha blits and 
// if the hardware supports hardware acceleration of alpha blits between
// two surfaces in video memory, to place the surface in video memory
// if possible, otherwise it will be placed in system memory.
// If the surface is created in video memory, blits will be _much_ faster,
// but the surface format must be identical to the video surface format,
// and the only way to access the pixels member of the surface is to use
// the SDL_LockSurface() and SDL_UnlockSurface() calls.
// If the requested surface actually resides in video memory, SDL_HWSURFACE
// will be set in the flags member of the returned surface.  If for some
// reason the surface could not be placed in video memory, it will not have
// the SDL_HWSURFACE flag set, and will be created in system memory instead.
func createRGBSurface(flags uint32, width, height, depth int, 
  rmask, gmask, bmask, amask uint32) (* C.SDL_Surface) { 
 return C.SDL_CreateRGBSurface(C.Uint32(flags), C.int(width), C.int(height), 
	  C.int(depth),	C.Uint32(rmask), C.Uint32(gmask), 
	  C.Uint32(bmask), C.Uint32(amask));
}

func CreateRGBSurfaceFrom() {
}
// TODO
// extern DECLSPEC SDL_Surface * SDLCALL SDL_CreateRGBSurfaceFrom(void *pixels,
//			int width, int height, int depth, int pitch,
//			Uint32 Rmask, Uint32 Gmask, Uint32 Bmask, Uint32 Amask);

// Frees the memory associated with the surface
func freeSurface(surface * C.SDL_Surface) { 
  C.SDL_FreeSurface(surface)
}

//
// SDL_LockSurface() sets up a surface for directly accessing the pixels.
// Between calls to SDL_LockSurface()/SDL_UnlockSurface(), you can write
// to and read from 'surface->pixels', using the pixel format stored in 
// 'surface->format'.  Once you are done accessing the surface, you should 
// use SDL_UnlockSurface() to release it.
//
// Not all surfaces require locking.  If SDL_MUSTLOCK(surface) evaluates
// to 0, then you can read and write to the surface at any time, and the
// pixel format of the surface will not change.  In particular, if the
// SDL_HWSURFACE flag is not given when calling SDL_SetVideoMode(), you
// will not need to lock the display surface before accessing it.
// 
// No operating system or library calls should be made between lock/unlock
// pairs, as critical system locks may be held during this time.
//
// SDL_LockSurface() returns 0, or -1 if the surface couldn't be locked.
func lockSurface(surface * C.SDL_Surface) (int) { 
  return(int(C.SDL_LockSurface(surface)))
}

func unlockSurface(surface * C.SDL_Surface) { 
  C.SDL_UnlockSurface(surface)
}

// Load a surface from a seekable SDL data source (memory or file.)
// If 'freesrc' is non-zero, the source will be closed after being read.
// Returns the new surface, or NULL if there was an error.
// The new surface should be freed with SDL_FreeSurface().
// TODO
// extern DECLSPEC SDL_Surface * SDLCALL SDL_LoadBMP_RW(SDL_RWops *src, int freesrc);
// TODO
// Convenience macro -- load a surface from a file 
// #define SDL_LoadBMP(file)	SDL_LoadBMP_RW(SDL_RWFromFile(file, "rb"), 1)

//
// Save a surface to a seekable SDL data source (memory or file.)
// If 'freedst' is non-zero, the source will be closed after being written.
// Returns 0 if successful or -1 if there was an error.
// TODO
// extern DECLSPEC int SDLCALL SDL_SaveBMP_RW
//		(SDL_Surface *surface, SDL_RWops *dst, int freedst);

// Convenience macro -- save a surface to a file 
// #define SDL_SaveBMP(surface, file) \
//		SDL_SaveBMP_RW(surface, SDL_RWFromFile(file, "wb"), 1)
//

// Sets the color key (transparent pixel) in a blittable surface.
// If 'flag' is SDL_SRCCOLORKEY (optionally OR'd with SDL_RLEACCEL), 
// 'key' will be the transparent pixel in the source image of a blit.
// SDL_RLEACCEL requests RLE acceleration for the surface if present,
// and removes RLE acceleration if absent.
// If 'flag' is 0, this function clears any current color key.
// This function returns 0, or -1 if there was an error.
func setColorKey(surface * C.SDL_Surface, flag, key uint32) int { 
  return int(C.SDL_SetColorKey(surface, C.Uint32(flag), C.Uint32(key)))
}

// This function sets the alpha value for the entire surface, as opposed to
// using the alpha component of each pixel. This value measures the range
// of transparency of the surface, 0 being completely transparent to 255
// being completely opaque. An 'alpha' value of 255 causes blits to be
// opaque, the source pixels copied to the destination (the default). Note
// that per-surface alpha can be combined with colorkey transparency.
//
// If 'flag' is 0, alpha blending is disabled for the surface.
// If 'flag' is SDL_SRCALPHA, alpha blending is enabled for the surface.
// OR:ing the flag with SDL_RLEACCEL requests RLE acceleration for the
// surface; if SDL_RLEACCEL is not specified, the RLE accel will be removed.
//
// The 'alpha' parameter is ignored for surfaces that have an alpha channel.
func setAlpha(surface * C.SDL_Surface, flag uint32, alpha uint8) int { 
  return int(C.SDL_SetAlpha(surface, C.Uint32(flag), C.Uint8(alpha)))
}

// Sets the clipping rectangle for the destination surface in a blit.
//
// If the clip rectangle is NULL, clipping will be disabled.
// If the clip rectangle doesn't intersect the surface, the function will
// return SDL_FALSE and blits will be completely clipped.  Otherwise the
// function returns SDL_TRUE and blits to the surface will be clipped to
// the intersection of the surface area and the clipping rectangle.
//
// Note that blits are automatically clipped to the edges of the source
// and destination surfaces.
func setClipRect(surface * C.SDL_Surface, rect * C.SDL_Rect) (bool) { 
  return i2b(int(C.SDL_SetClipRect(surface, rect)))
}

// Gets the clipping rectangle for the destination surface in a blit.
// 'rect' must be a pointer to a valid rectangle which will be filled
// with the correct values.
func getClipRect(surface * C.SDL_Surface, rect * C.SDL_Rect) { 
  C.SDL_GetClipRect(surface, rect)
}


// Creates a new surface of the specified format, and then copies and maps 
// the given surface to it so the blit of the converted surface will be as 
// fast as possible.  If this function fails, it returns NULL.
//
// The 'flags' parameter is passed to SDL_CreateRGBSurface() and has those 
// semantics.  You can also pass SDL_RLEACCEL in the flags parameter and
// SDL will try to RLE accelerate colorkey and alpha blits in the resulting
// surface.
//
// This function is used internally by SDL_DisplayFormat().
func convertSurface(src * C.SDL_Surface, fmt * C.SDL_PixelFormat, 
  flags uint32) ( * C.SDL_Surface) {
  return C.SDL_ConvertSurface(src, fmt, C.Uint32(flags))
}

// This performs a fast blit from the source surface to the destination
// surface.  It assumes that the source and destination rectangles are
// the same size.  If either 'srcrect' or 'dstrect' are NULL, the entire
// surface (src or dst) is copied.  The final blit rectangles are saved
// in 'srcrect' and 'dstrect' after all clipping is performed.
// If the blit is successful, it returns 0, otherwise it returns -1.
//
// The blit function should not be called on a locked surface.
//
// The blit semantics for surfaces with and without alpha and colorkey
// are defined as follows:
//
// RGBA->RGB:
//     SDL_SRCALPHA set:
// 	alpha-blend (using alpha-channel).
// 	SDL_SRCCOLORKEY ignored.
//     SDL_SRCALPHA not set:
// 	copy RGB.
// 	if SDL_SRCCOLORKEY set, only copy the pixels matching the
// 	RGB values of the source colour key, ignoring alpha in the
// 	comparison.
// 
// RGB->RGBA:
//     SDL_SRCALPHA set:
// 	alpha-blend (using the source per-surface alpha value);
// 	set destination alpha to opaque.
//     SDL_SRCALPHA not set:
// 	copy RGB, set destination alpha to source per-surface alpha value.
//     both:
// 	if SDL_SRCCOLORKEY set, only copy the pixels matching the
// 	source colour key.
// 
// RGBA->RGBA:
//     SDL_SRCALPHA set:
// 	alpha-blend (using the source alpha channel) the RGB values;
// 	leave destination alpha untouched. [Note: is this correct?]
// 	SDL_SRCCOLORKEY ignored.
//     SDL_SRCALPHA not set:
// 	copy all of RGBA to the destination.
// 	if SDL_SRCCOLORKEY set, only copy the pixels matching the
// 	RGB values of the source colour key, ignoring alpha in the
// 	comparison.
// 
// RGB->RGB: 
//     SDL_SRCALPHA set:
// 	alpha-blend (using the source per-surface alpha value).
//     SDL_SRCALPHA not set:
// 	copy RGB.
//     both:
// 	if SDL_SRCCOLORKEY set, only copy the pixels matching the
// 	source colour key.
//
// If either of the surfaces were in video memory, and the blit returns -2,
// the video memory was lost, so it should be reloaded with artwork and 
// re-blitted:
//	while ( SDL_BlitSurface(image, imgrect, screen, dstrect) == -2 ) {
//		while ( SDL_LockSurface(image) < 0 )
//			Sleep(10);
//		-- Write image pixels to image->pixels --
//		SDL_UnlockSurface(image);
//	}
// This happens under DirectX 5.0 when the system switches away from your
// fullscreen application.  The lock will also fail until you have access
// to the video memory again.
 
// You should call SDL_BlitSurface() unless you know exactly how SDL
// blitting works internally and how to use the other blit functions.
func blitSurface(src, dst * C.SDL_Surface, srcrect, dstrect *C.SDL_Rect) (int) {
  return int(C.SDL_UpperBlit(src, srcrect, dst, dstrect))
}

// This function performs a fast fill of the given rectangle with 'color'
// The given rectangle is clipped to the destination surface clip area
// and the final fill rectangle is saved in the passed in pointer.
// If 'dstrect' is NULL, the whole surface will be filled with 'color'
// The color should be a pixel of the format used by the surface, and 
// can be generated by the SDL_MapRGB() function.
// This function returns 0 on success, or -1 on error.
func fillRect(dst * C.SDL_Surface, dstrect * C.SDL_Rect, color uint32) (int) { 
  return int(C.SDL_FillRect(dst, dstrect, C.Uint32(color)))
}

// This function takes a surface and copies it to a new surface of the
// pixel format and colors of the video framebuffer, suitable for fast
// blitting onto the display surface.
//
// If you want to take advantage of hardware colorkey or alpha blit
// acceleration, you should set the colorkey and alpha value before
// calling this function.
//
// If the conversion fails or runs out of memory, it returns NULL
func displayFormat(surface * C.SDL_Surface) (* C.SDL_Surface) { 
  return C.SDL_DisplayFormat(surface)
}

// This function takes a surface and copies it to a new surface of the
// pixel format and colors of the video framebuffer (if possible),
// suitable for fast alpha blitting onto the display surface.
// The new surface will always have an alpha channel.
//
// If you want to take advantage of hardware colorkey or alpha blit
// acceleration, you should set the colorkey and alpha value before
// calling this function.
//
// If the conversion fails or runs out of memory, it returns NULL
func displayFormatAlpha(surface * C.SDL_Surface) (* C.SDL_Surface) { 
  return C.SDL_DisplayFormatAlpha(surface)
}
 

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 
// YUV video surface overlay functions , for video playback                                      
// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 

// This function creates a video output overlay
// Calling the returned surface an overlay is something of a misnomer because
// the contents of the display surface underneath the area where the overlay
// is shown is undefined - it may be overwritten with the converted YUV data.
func CreateYUVOverlay(width, height int, format uint32, 
  display * C.SDL_Surface ) (* C.SDL_Overlay) { 
  return C.SDL_CreateYUVOverlay(C.int(width), C.int(height),
				C.Uint32(format), display)
}

// Lock an overlay for direct access, and unlock it when you are done 
func LockYUVOverlay(overlay * C.SDL_Overlay) (int) { 
  return int(C.SDL_LockYUVOverlay(overlay))
}

func UnlockYUVOverlay(overlay * C.SDL_Overlay) {
  C.SDL_UnlockYUVOverlay(overlay)
}

// Blit a video overlay to the display surface.
// The contents of the video surface underneath the blit destination are
// not defined.  
// The width and height of the destination rectangle may be different from
// that of the overlay, but currently only 2x scaling is supported.
func DisplayYUVOverlay(overlay * C.SDL_Overlay, dstrect * C.SDL_Rect) (int) { 
  return int(C.SDL_DisplayYUVOverlay(overlay, dstrect))
}

// Free a video overlay 
func FreeYUVOverlay(overlay * C.SDL_Overlay) { 
  C.SDL_FreeYUVOverlay(overlay)
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 
// OpenGL support functions.                                                 
// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 

//
// Dynamically load an OpenGL library, or the default one if path is NULL
//
// If you do this, you need to retrieve all of the GL functions used in
// your program from the dynamic library using SDL_GL_GetProcAddress().
// extern DECLSPEC int SDLCALL SDL_GL_LoadLibrary(const char *path);

//
// Get the address of a GL function
// extern DECLSPEC void * SDLCALL SDL_GL_GetProcAddress(const char* proc);

//
// Set an attribute of the OpenGL subsystem before intialization.
// extern DECLSPEC int SDLCALL SDL_GL_SetAttribute(SDL_GLattr attr, int value);
func GLSetAttribute(attr, value int) (int) {
  return int(C.SDL_GL_SetAttribute(C.SDL_GLattr(attr), C.int(value)))
} 

// Get an attribute of the OpenGL subsystem from the windowing
// interface, such as glX. This is of course different from getting
// the values from SDL's internal OpenGL subsystem, which only
// stores the values you request before initialization.
//
// Developers should track the values they pass into SDL_GL_SetAttribute
// themselves if they want to retrieve these values.
// extern DECLSPEC int SDLCALL SDL_GL_GetAttribute(SDL_GLattr attr, int* value);
func GLGetAttribute(attr int) (int, int) {
  var value C.int  
  res := int(C.SDL_GL_GetAttribute(C.SDL_GLattr(attr), &value))
  return res, int(value)
} 


//
// Swap the OpenGL buffers, if double-buffering is supported.
// extern DECLSPEC void SDLCALL SDL_GL_SwapBuffers(void);
func GLSwapBuffers() {
  C.SDL_GL_SwapBuffers() 
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 
// These functions allow interaction with the window manager, if any.        
// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 

//
// Sets/Gets the title and icon text of the display window (UTF-8 encoded)
func WM_SetCaption(title, icon string) {
  ctitle := cstr(title)
  cicon  := cstr(icon)  
  C.SDL_WM_SetCaption(ctitle, cicon)
}

// extern DECLSPEC void SDLCALL SDL_WM_GetCaption(char **title, char **icon);


// Sets the icon for the display window.
// This function must be called before the first call to SDL_SetVideoMode().
// It takes an icon surface, and a mask in MSB format.
// If 'mask' is NULL, the entire icon surface will be used as the icon.
// TODO: support mask 
func wm_SetIcon(icon * C.SDL_Surface) { 
  C.SDL_WM_SetIcon(icon, nil);
}

// This function iconifies the window, and returns 1 if it succeeded.
// If the function succeeds, it generates an SDL_APPACTIVE loss event.
// This function is a noop and returns 0 in non-windowed environments.
func WM_Iconify() (int) { 
  return int(C.SDL_WM_IconifyWindow())
}

// Toggle fullscreen mode without changing the contents of the screen.
// If the display surface does not require locking before accessing
// the pixel information, then the memory pointers will not change.
//
// If this function was able to toggle fullscreen mode (change from 
// running in a window to fullscreen, or vice-versa), it will return 1.
// If it is not implemented, or fails, it returns 0.
//
// The next call to SDL_SetVideoMode() will set the mode fullscreen
// attribute based on the flags parameter - if SDL_FULLSCREEN is not
// set, then the display will be windowed by default where supported.
//
// This is currently only implemented in the X11 video driver.
func WM_ToggleFullScreen(surface * C.SDL_Surface) (int) { 
  return int(C.SDL_WM_ToggleFullScreen(surface))
}

// This function allows you to set and query the input grab state of
// the application.  It returns the new input grab state.
type SDL_GrabMode int
const (
	GRAB_QUERY 	= SDL_GrabMode(-1)
	GRAB_OFF	= SDL_GrabMode(0)
	GRAB_ON 	= SDL_GrabMode(1)
)

// Grabbing means that the mouse is confined to the application window,
// and nearly all keyboard input is passed directly to the application,
// and not interpreted by a window manager, if any.
func WM_GrabInput(mode SDL_GrabMode) (SDL_GrabMode) { 
  return SDL_GrabMode(C.SDL_WM_GrabInput(C.SDL_GrabMode(mode)))
}


// Wrappers
 



