/*
* Bindings to SDL_RWops
*/
package sdl

/* 
#include <SDL.h>
#include <SDL_rwops.h>
*/
import "C"
import "unsafe"


// Functions to create SDL_RWops structures from various data sources
// Creates an RWOPS from a file
func RWFromFile(file, mode string)  (* C.SDL_RWops) {
  cfile := cstr(file) ; cfile.free()
  cmode := cstr(mode) ; cmode.free()  
  return C.SDL_RWFromFile(cfile, cmode);  
}

//Not supported
//extern DECLSPEC SDL_RWops * SDLCALL SDL_RWFromFP(FILE *fp, int autoclose);
//extern DECLSPEC SDL_RWops * SDLCALL SDL_RWFromMem(void *mem, int size);
//extern DECLSPEC SDL_RWops * SDLCALL SDL_RWFromConstMem(const void *mem, int size);
//extern DECLSPEC SDL_RWops * SDLCALL SDL_AllocRW(void);

func FreeRW(rwops * C.SDL_RWops) {
  C.SDL_FreeRW(rwops)
}

const (
  // Seek from the beginning of data 
  RW_SEEK_SET	= 0	
  // Seek relative to current read point
  RW_SEEK_CUR	= 1	
  // Seek relative to the end of data
  RW_SEEK_END	= 2	
)

type seek_func * func (rwops * C.SDL_RWops, offset C.int, whence C.int) (C.int)

// I doubt these will work...
func RWSeek(rwops * C.SDL_RWops, offset, whence int) (int) {
  tocalla := seek_func((unsafe.Pointer(rwops.seek)))
  tocall  := (*tocalla)
  return int(tocall(rwops, C.int(offset), C.int(whence)))
}

func RWTell(rwops * C.SDL_RWops) (int) {
  tocalla := seek_func((unsafe.Pointer(rwops.seek)))
  tocall  := (*tocalla)
  return int(tocall(rwops, C.int(0), C.int(RW_SEEK_CUR)))
}

/*
type read_func * func (rwops * C.SDL_RWops, offset C.int, whence C.int) (C.int)

func RWRead(rwops * C.SDL_RWops, buffer []byte) (int) {
  tocalla := read_func((unsafe.Pointer(rwops.read)))
  tocall  := (*tocalla)
  size	  := buffer.cap
  n	  := 1 
  return int(tocall(rwops, C.int(0), C.int(RW_SEEK_CUR)))
}
*/

/*
type close_func * func (rwops * C.SDL_RWops) (C.int)

func RWClose(rwops * C.SDL_RWops) (int) {
  tocalla := close_func(unsafe.Pointer(rwops.close))
  tocall  := (*tocalla)
  return int(tocall(rwops)))
}
*/

/* Macros to easily read and write from an SDL_RWops structure 
#define SDL_RWseek(ctx, offset, whence)	(ctx)->seek(ctx, offset, whence)
#define SDL_RWtell(ctx)			(ctx)->seek(ctx, 0, RW_SEEK_CUR)
#define SDL_RWread(ctx, ptr, size, n)	(ctx)->read(ctx, ptr, size, n)
#define SDL_RWwrite(ctx, ptr, size, n)	(ctx)->write(ctx, ptr, size, n)
#define SDL_RWclose(ctx)		(ctx)->close(ctx)
*/

/* Read an item of the specified endianness and return in native format 
extern DECLSPEC Uint16 SDLCALL SDL_ReadLE16(SDL_RWops *src);
extern DECLSPEC Uint16 SDLCALL SDL_ReadBE16(SDL_RWops *src);
extern DECLSPEC Uint32 SDLCALL SDL_ReadLE32(SDL_RWops *src);
extern DECLSPEC Uint32 SDLCALL SDL_ReadBE32(SDL_RWops *src);
extern DECLSPEC Uint64 SDLCALL SDL_ReadLE64(SDL_RWops *src);
extern DECLSPEC Uint64 SDLCALL SDL_ReadBE64(SDL_RWops *src);
*/

/* Write an item of native format to the specified endianness 
extern DECLSPEC int SDLCALL SDL_WriteLE16(SDL_RWops *dst, Uint16 value);
extern DECLSPEC int SDLCALL SDL_WriteBE16(SDL_RWops *dst, Uint16 value);
extern DECLSPEC int SDLCALL SDL_WriteLE32(SDL_RWops *dst, Uint32 value);
extern DECLSPEC int SDLCALL SDL_WriteBE32(SDL_RWops *dst, Uint32 value);
extern DECLSPEC int SDLCALL SDL_WriteLE64(SDL_RWops *dst, Uint64 value);
extern DECLSPEC int SDLCALL SDL_WriteBE64(SDL_RWops *dst, Uint64 value);
*/

//Wrappers

type RW struct { 
  rwops * C.SDL_RWops
  filename, mode string
}

func OpenRW(filename, mode string) (* RW) {
  rw := new(RW)
  rw.rwops = RWFromFile(filename, mode)
  if rw.rwops == nil { 
    return nil
  }
  rw.filename = filename
  rw.mode = mode
  return rw  
}

/*
func (rw * RW) Close() {
  if rw.rwops == nil { return } 
  RWClose(rw.rwops)  
}
*/
func (rw * RW) Free() {
  if rw.rwops == nil { return } 
  FreeRW(rw.rwops)
  rw.rwops = nil
}

