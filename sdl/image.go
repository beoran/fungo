/*
* Bindings to SDL_Image
*/
package sdl

/* 
#include <SDL.h>
#include <SDL_image.h>
*/
import "C"
import "unsafe"

const IMAGE_MAJOR_VERSION = C.SDL_IMAGE_MAJOR_VERSION
const IMAGE_MINOR_VERSION = C.SDL_IMAGE_MINOR_VERSION
const IMAGE_PATCHLEVEL	  = C.SDL_IMAGE_PATCHLEVEL

// This macro can be used to fill a version structure with the compile-time
// version of the SDL_image library.
// #define SDL_IMAGE_VERSION(X)						\
// {									\
// 	(X)->major = SDL_IMAGE_MAJOR_VERSION;				\
// 	(X)->minor = SDL_IMAGE_MINOR_VERSION;				\
// 	(X)->patch = SDL_IMAGE_PATCHLEVEL;				\
// }

// This function gets the version of the dynamically linked SDL_image library.
// it should NOT be used to fill a version structure, instead you should
// use the SDL_IMAGE_VERSION() macro.
func IMG_LinkedVersion() (* C.SDL_version) {
  return C.IMG_Linked_Version()
}

// Load an image from an SDL data source.
// The 'type' may be one of: "BMP", "GIF", "PNG", etc.
// If the image format supports a transparent pixel, SDL will set the
// colorkey for the surface.  You can enable RLE acceleration on the
// surface afterwards by calling:
// SDL_SetColorKey(image, SDL_RLEACCEL, image->format->colorkey);

func IMG_Load(filename string) (* C.SDL_Surface) { 
  cfile := cstr(filename); defer cfile.free()
  return C.IMG_Load(cfile);
}

// below not supported
// extern DECLSPEC SDL_Surface * SDLCALL IMG_LoadTyped_RW(SDL_RWops *src, int freesrc, char *type);
/* Convenience functions */
//extern DECLSPEC SDL_Surface * SDLCALL IMG_Load_RW(SDL_RWops *src, int freesrc);

/* Invert the alpha of a surface for use with OpenGL
   This function is now a no-op, and only provided for backwards compatibility.
*/
//extern DECLSPEC int SDLCALL IMG_InvertAlpha(int on);

/* Functions to detect a file type, given a seekable source */
/*
extern DECLSPEC int SDLCALL IMG_isBMP(SDL_RWops *src);
extern DECLSPEC int SDLCALL IMG_isGIF(SDL_RWops *src);
extern DECLSPEC int SDLCALL IMG_isJPG(SDL_RWops *src);
extern DECLSPEC int SDLCALL IMG_isLBM(SDL_RWops *src);
extern DECLSPEC int SDLCALL IMG_isPCX(SDL_RWops *src);
extern DECLSPEC int SDLCALL IMG_isPNG(SDL_RWops *src);
extern DECLSPEC int SDLCALL IMG_isPNM(SDL_RWops *src);
extern DECLSPEC int SDLCALL IMG_isTIF(SDL_RWops *src);
extern DECLSPEC int SDLCALL IMG_isXCF(SDL_RWops *src);
extern DECLSPEC int SDLCALL IMG_isXPM(SDL_RWops *src);
extern DECLSPEC int SDLCALL IMG_isXV(SDL_RWops *src);
*/
/* Individual loading functions */
/*
extern DECLSPEC SDL_Surface * SDLCALL IMG_LoadBMP_RW(SDL_RWops *src);
extern DECLSPEC SDL_Surface * SDLCALL IMG_LoadGIF_RW(SDL_RWops *src);
extern DECLSPEC SDL_Surface * SDLCALL IMG_LoadJPG_RW(SDL_RWops *src);
extern DECLSPEC SDL_Surface * SDLCALL IMG_LoadLBM_RW(SDL_RWops *src);
extern DECLSPEC SDL_Surface * SDLCALL IMG_LoadPCX_RW(SDL_RWops *src);
extern DECLSPEC SDL_Surface * SDLCALL IMG_LoadPNG_RW(SDL_RWops *src);
extern DECLSPEC SDL_Surface * SDLCALL IMG_LoadPNM_RW(SDL_RWops *src);
extern DECLSPEC SDL_Surface * SDLCALL IMG_LoadTGA_RW(SDL_RWops *src);
extern DECLSPEC SDL_Surface * SDLCALL IMG_LoadTIF_RW(SDL_RWops *src);
extern DECLSPEC SDL_Surface * SDLCALL IMG_LoadXCF_RW(SDL_RWops *src);
extern DECLSPEC SDL_Surface * SDLCALL IMG_LoadXPM_RW(SDL_RWops *src);
extern DECLSPEC SDL_Surface * SDLCALL IMG_LoadXV_RW(SDL_RWops *src);
*/

//extern DECLSPEC SDL_Surface * SDLCALL IMG_ReadXPMFromArray(char **xpm);

/* We'll use SDL for reporting errors */
/*
#define IMG_SetError	SDL_SetError
#define IMG_GetError	SDL_GetError
*/