// Wraps SDL_ttg functions
// Only wraps UTF-8 enabled functions. 
package sdl

//#include <SDL.h>
//#include <SDL_image.h>
//#include <SDL_mixer.h>
//#include <SDL_ttf.h>
import "C"
// import "unsafe"

// ZERO WIDTH NO-BREAKSPACE (Unicode byte order mark)
const UNICODE_BOM_NATIVE  = 0xFEFF
const UNICODE_BOM_SWAPPED = 0xFFFE
// #include <SDL_ttf.h>

// This function tells the library whether UNICODE text is generally
// byteswapped.  A UNICODE BOM character in a string will override
// this setting for the remainder of that string.
func TTF_ByteSwappedCode(swapped bool) { 
  C.TTF_ByteSwappedUNICODE(C.int(b2i(swapped)))
}

// Initialize the TTF engine - returns 0 if successful, -1 on error 
func TTF_Init() (int) { 
  return int(C.TTF_Init())
}

// Open a font file and create a font of the specified point size.
// Some .fon fonts will have several sizes embedded in the file, so the
// point size becomes the index of choosing which size.  If the value
// is too high, the last indexed size will be the default.
func TTF_OpenFont(file string, ptsize int) (* C.TTF_Font) {
  cfile  := cstr(file)
  defer cfile.free()
  return C.TTF_OpenFont(cfile, C.int(ptsize))
}

func TTF_OpenFontIndex(file string, ptsize int, index int32) (* C.TTF_Font) { 
  cfile := cstr(file) 
  defer cfile.free()
  return C.TTF_OpenFontIndex(cfile, C.int(ptsize), C.long(index));
}
// Not wrapped
//extern DECLSPEC TTF_Font * SDLCALL TTF_OpenFontRW(SDL_RWops *src, int freesrc, int ptsize);
//extern DECLSPEC TTF_Font * SDLCALL TTF_OpenFontIndexRW(SDL_RWops *src, int freesrc, int ptsize, long index);

// Font styles
// This font style is implemented by modifying the font glyphs, and
// doesn't reflect any inherent properties of the truetype font file.
const TTF_STYLE_NORMAL		=0x00
const TTF_STYLE_BOLD		=0x01
const TTF_STYLE_ITALIC		=0x02
const TTF_STYLE_UNDERLINE	=0x04

// Get the style of the font
func TTFGetFontStyle(font * C.TTF_Font) (int) { 
  return int(C.TTF_GetFontStyle(font))
}

// Set the style of the font
func TTFSetFontStyle(font * C.TTF_Font, style int) { 
  C.TTF_SetFontStyle(font, C.int(style))
}

// Get the total height of the font - usually equal to point size
func TTFFontHeight(font * C.TTF_Font) (int) { 
  return int(C.TTF_FontHeight(font))
}


// Get the offset from the baseline to the top of the font
// This is a positive value, relative to the baseline.
func TTFFontAscent(font * C.TTF_Font) (int) { 
  return int(C.TTF_FontAscent(font))
}

// Get the offset from the baseline to the bottom of the font
// This is a negative value, relative to the baseline.
func TTFFontDescent(font * C.TTF_Font) (int) { 
  return int(C.TTF_FontDescent(font))
}

// Get the recommended spacing between lines of text for this font 
func TTFFontLineSkip(font * C.TTF_Font) (int) { 
  return int(C.TTF_FontLineSkip(font))
}
// Get the number of faces of the font 
func TTFFontFaces(font * C.TTF_Font) (int32) {
  return int32(C.TTF_FontFaces(font))
}

// Get the font face attributes, if any 
func TTFFontFaceIsFixedWidth(font * C.TTF_Font) (bool) { 
  return i2b(int(C.TTF_FontFaceIsFixedWidth(font)))
}

// Font family name
func TTFFontFamilyName(font * C.TTF_Font) (string) { 
  return C.GoString(C.TTF_FontFaceFamilyName(font));
}

// Font face stylename
func TTFFontFaceStyleName(font * C.TTF_Font) (string) { 
  return C.GoString(C.TTF_FontFaceStyleName(font));
}

// Get the metrics (dimensions) of a glyph
// To understand what these metrics mean, here is a useful link:
// http://freetype.sourceforge.net/freetype2/docs/tutorial/step2.html
// retuns minx, maxx, miny, maxy, advance
func TTFGlyphMetrics(font * C.TTF_Font, ch uint16) (int, int, int, int, int) {
  var minx, maxx, miny, maxy, advance int
  pminx := cintptr(&minx)
  pmaxx := cintptr(&maxx)
  pminy := cintptr(&miny)
  pmaxy := cintptr(&maxy)
  padv  := cintptr(&advance)
  C.TTF_GlyphMetrics(font, C.Uint16(ch), pminx, pmaxx, pminy, pmaxy, padv)
  return minx, maxx, miny, maxy, advance
}

// Get the dimensions of a rendered string of text. Works for UTF-8 encoded text.
// returns w and h in that order
func TTFTextSize(font * C.TTF_Font, text string) (int, int) { 
  var w, h int
  ctext := cstr(text) ; defer ctext.free()
  pw := cintptr(&w)
  ph := cintptr(&h)
  C.TTF_SizeUTF8(font, ctext, pw, ph)
  return w, h
}

// Create an 8-bit palettized surface and render the given text at
// fast quality with the given font and color.  The 0 pixel is the
// colorkey, giving a transparent background, and the 1 pixel is set
// to the text color.
// This function returns the new surface, or NULL if there was an error.
// Works with UTF8 encoded strings.
func TTFRenderSolid(font * C.TTF_Font, text string, 
     color C.SDL_Color) (* C.SDL_Surface) { 
  ctext := cstr(text) ; defer ctext.free()
  return C.TTF_RenderUTF8_Solid(font, ctext, color)
}


// Create an 8-bit palettized surface and render the given glyph at
// fast quality with the given font and color.  The 0 pixel is the
// colorkey, giving a transparent background, and the 1 pixel is set
// to the text color.  The glyph is rendered without any padding or
// centering in the X direction, and aligned normally in the Y direction.
//   This function returns the new surface, or NULL if there was an error.
func TTFRenderGlyphSolid(font * C.TTF_Font, ch uint16, 
     color C.SDL_Color) (* C.SDL_Surface) { 
  return C.TTF_RenderGlyph_Solid(font, C.Uint16(ch), color)
}

// Create an 8-bit palettized surface and render the given text at
// high quality with the given font and colors.  The 0 pixel is background,
// while other pixels have varying degrees of the foreground color.
// This function returns the new surface, or NULL if there was an error.
func TTFRenderShaded(font * C.TTF_Font, text string, 
     fg, bg C.SDL_Color) (* C.SDL_Surface) { 
  ctext := cstr(text) ; defer ctext.free()
  return C.TTF_RenderUTF8_Shaded(font, ctext, fg, bg)
}


// Create an 8-bit palettized surface and render the given glyph at
// high quality with the given font and colors.  The 0 pixel is background,
// while other pixels have varying degrees of the foreground color.
// The glyph is rendered without any padding or centering in the X
// direction, and aligned normally in the Y direction.
// This function returns the new surface, or NULL if there was an error.
func TTFRenderGlyphShaded(font * C.TTF_Font, ch uint16, 
     fg, bg C.SDL_Color) (* C.SDL_Surface) { 
  return C.TTF_RenderGlyph_Shaded(font, C.Uint16(ch), fg, bg)
}

// Create a 32-bit ARGB surface and render the given text at high quality,
// using alpha blending to dither the font with the given color.
// This function returns the new surface, or NULL if there was an error.
func TTFRenderBlended(font * C.TTF_Font, text string, 
     color C.SDL_Color) (* C.SDL_Surface) { 
  ctext := cstr(text) ; defer ctext.free()
  return C.TTF_RenderUTF8_Blended(font, ctext, color)
}

// Create a 32-bit ARGB surface and render the given glyph at high quality,
// using alpha blending to dither the font with the given color.
// The glyph is rendered without any padding or centering in the X
// direction, and aligned normally in the Y direction.
// This function returns the new surface, or NULL if there was an error.
func TTFRenderGlyphBlended(font * C.TTF_Font, ch uint16, 
     color C.SDL_Color) (* C.SDL_Surface) { 
  return C.TTF_RenderGlyph_Blended(font, C.Uint16(ch), color)
}

// Close an opened font file
func TTFCloseFont(font * C.TTF_Font) { 
  C.TTF_CloseFont(font)
}

// De-initialize the TTF engine 
func TTFQuit() { 
  C.TTF_Quit()
}

// De-initialize the TTF engine 
func TTFWasInit() (bool) { 
  res := C.TTF_WasInit()
  return i2b(int(res))
}


