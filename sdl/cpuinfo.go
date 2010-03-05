package sdl

// #include <SDL.h>
import "C"

// converts ints to bools
func i2b(res int) (bool) {
  if res != 0 { return true } 
  return false
}

// This function returns true if the CPU has the RDTSC instruction
func HasRDTRSC() (bool) {
  return i2b(int(C.SDL_HasRDTSC()))
}

// This function returns true if the CPU has MMX features
func HasMMX() (bool) {
  return i2b(int(C.SDL_HasMMX()))
}
 

// This function returns true if the CPU has MMX Ext. features
func HasMMXExt() (bool) {
  return i2b(int(C.SDL_HasMMXExt()))
}
 
// This function returns true if the CPU has 3DNow features
func Has3DNow() (bool) {
  return i2b(int(C.SDL_Has3DNow()))
}

// This function returns true if the CPU has 3DNow Ext. features
func Has3DNowExt() (bool) {
  return i2b(int(C.SDL_Has3DNowExt()))
}

// This function returns true if the CPU has SSE features
func HasSSE() (bool) {
  return i2b(int(C.SDL_HasSSE()))
}
 
// This function returns true if the CPU has SSE2 features
func HasSSE2() (bool) {
  return i2b(int(C.SDL_HasSSE2()))
}

// This function returns true if the CPU has AltiVec features
func HasAltiVec() (bool) {
  return i2b(int(C.SDL_HasAltiVec()))
}
 