package sdl

// #include <SDL.h>
// #include <SDL_error.h>
import "C"

/* Does not compile due to CGO limitation. 
func SetError(fmt string) {
  cres  := cstr(fmt); defer cres.free()
  C.SDL_SetError(cres)
}
*/

func GetError() (string) {
  res   := C.SDL_GetError()
  return C.GoString(res)
}

func ClearError() {
  C.SDL_GetError();
}

type Errorcode int

const ( 
  ENOMEM = Errorcode(iota)
  EFREAD 
  EFWRITE
  EFSEEK
  UNSUPPORTED
  LASTERROR
)

func Error(code Errorcode) {
  C.SDL_Error(C.SDL_errorcode(code))
}


