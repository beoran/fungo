//
// Go Language wrappers around SDL
//
package sdl

// #include <SDL.h>
import "C"
import "unsafe"

//
/*
import "fmt"
import "os"
import "runtime"
*/


// Helper functions 
// Allocates a string with the given byte length
// don't forget a call to defer s.free() ! 
func cstrNew(size int) (* C.char) {
  return (*C.char)(unsafe.Pointer(C.malloc(C.size_t(size))))  
}

// free is a method on C char * strings to method to free the associated memory 
func (self * C.char) free() {
  C.free(unsafe.Pointer(self))
}

// cstring converts a string to a C string. This allocates memory, 
// so don't forget to add a "defer s.free()"
func cstr(self string) (* C.char) {
  return C.CString(self)
}

type mystring string;

// Helper to convert strings to C strings 
func (self mystring) cstr() (* C.char) {
  return C.CString(string(self))
}


// The available application states
const (
  APPMOUSEFOCUS = C.SDL_APPMOUSEFOCUS	//0x01		// The app has mouse coverage */
  APPINPUTFOCUS = C.SDL_APPINPUTFOCUS	//0x02		// The app has input focus */
  APPACTIVE	= C.SDL_APPACTIVE		//0x04		// The application is active */
)

// 
// GetAppState returns the current state of the application, which is a
// bitwise combination of APPMOUSEFOCUS, APPINPUTFOCUS, and
// APPACTIVE.  If APPACTIVE is set, then the user is able to
// see your application, otherwise it has been iconified or disabled.
///
func GetAppState() (C.Uint8) { 
  return C.Uint8(C.SDL_GetAppState())
}



/* As of version 0.5, SDL is loaded dynamically into the application */

/* These are the flags which may be passed to SDL_Init() -- you should
   specify the subsystems which you will be using in your application.
*/
const( 
  INIT_TIMER 		= C.SDL_INIT_TIMER
  INIT_AUDIO 		= C.SDL_INIT_AUDIO
  INIT_VIDEO 		= C.SDL_INIT_VIDEO
  INIT_CDROM 		= C.SDL_INIT_CDROM
  INIT_JOYSTICK 	= C.SDL_INIT_JOYSTICK
  INIT_NOPARACHUTE	= C.SDL_INIT_NOPARACHUTE
  INIT_EVENTTHREAD 	= C.SDL_INIT_EVENTTHREAD
  INIT_EVERYTHING	= C.SDL_INIT_EVERYTHING
)


// This function loads the SDL dynamically linked library and initializes 
// the subsystems specified by 'flags' (and those satisfying dependencies)
// Unless the INIT_NOPARACHUTE flag is set, it will install cleanup
// signal handlers for some commonly ignored fatal signals (like SIGSEGV)
func Init(flags uint32) uint32 {  
  return uint32(C.SDL_Init(C.Uint32(flags)))
}

// This function initializes specific SDL subsystems 
func InitSubSystem(flags uint32) uint32 { 
   return uint32(C.SDL_InitSubSystem(C.Uint32(flags)))
}

// This function cleans up specific SDL subsystems 
func QuitSubSystem(flags uint32) { 
   C.SDL_QuitSubSystem(C.Uint32(flags))
}

// This function returns mask of the specified subsystems which have
// been initialized.
// If 'flags' is 0, it returns a mask of all initialized subsystems.
func WasInit(flags uint32) uint32 { 
   return uint32(C.SDL_WasInit(C.Uint32(flags)))
}

// This function returns true if the given subsystem was initialized, 
// and false if it wasn't.
func Initialized(flags uint32) bool {
  return (WasInit(flags) & flags > 0)
}

// This function initializes a subsystem if it wasn't initialized before
// Returns true if initialization was performed
func InitSubSystemOnce(flags uint32) bool {
  if Initialized(flags) { return false; } 
  InitSubSystem(flags)  
  return true
}

// This function cleans up all initialized subsystems and unloads the
// dynamically linked library.  You should call it upon all exit conditions.
func Quit() { 
   C.SDL_Quit()
}



  

