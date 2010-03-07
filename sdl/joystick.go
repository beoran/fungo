/*
* Bindings to SDL_joystick
*/
package sdl

/* 
#include <SDL.h>
#include <SDL_joystick.h>
*/
import "C"
import "unsafe"

// In order to use these functions, SDL_Init() must have been called
// with the SDL_INIT_JOYSTICK flag.  This causes SDL to scan the system
// for joysticks, and load appropriate drivers.

// The joystick structure used to identify an SDL joystick 

// Count the number of joysticks attached to the system
func NumJoysticks() (int) {   
  return int(C.SDL_NumJoysticks())
}


// Get the implementation dependent name of a joystick.
// This can be called before any joysticks are opened.
// If no name can be found, this function returns NULL.
func JoystickName(index int) string { 
  return C.GoString(C.SDL_JoystickName(C.int(index)))
}

// Open a joystick for use - the index passed as an argument refers to
// the N'th joystick on the system.  This index is the value which will
// identify this joystick in future joystick events.
// This function returns a joystick identifier, or NULL if an error occurred.
func JoystickOpen(index int) (* C.SDL_Joystick) { 
  return C.SDL_JoystickOpen(C.int(index))
}


// Returns true if the joystick has been opened, or false if it has not.
func JoystickOpened(index int) (bool) {
  return i2b(int(C.SDL_JoystickOpened(C.int(index))))
}

// Get the device index of an opened joystick.
func JoystickIndex(joystick * C.SDL_Joystick) (int) { 
  return int(C.SDL_JoystickIndex(joystick))
}

// Get the number of general axis controls on a joystick
func JoystickNumAxes(joystick * C.SDL_Joystick) (int) { 
  return int(C.SDL_JoystickNumAxes(joystick))
}

// Get the number of trackballs on a joystick
// Joystick trackballs have only relative motion events associated
// with them and their state cannot be polled.
func JoystickNumBalls(joystick * C.SDL_Joystick) (int) { 
  return int(C.SDL_JoystickNumBalls(joystick))
}

// Get the number of POV hats on a joystick
func JoystickNumHats(joystick * C.SDL_Joystick) (int) { 
  return int(C.SDL_JoystickNumHats(joystick))
}

// Get the number of buttons on a joystick
func JoystickNumButtons(joystick * C.SDL_Joystick) (int) { 
  return int(C.SDL_JoystickNumButtons(joystick))
}

// Update the current state of the open joysticks.
// This is called automatically by the event loop if any joystick
// events are enabled.
func JoystickUpdate() {
  C.SDL_JoystickUpdate();
}

//
// Enable/disable joystick event polling.
// If joystick events are disabled, you must call SDL_JoystickUpdate()
// yourself and check the state of the joystick when you want joystick
// information.
// The state can be one of sdl.QUERY, sdl.ENABLE or sdl.IGNORE.
func JoystickEventState(state int) (int) { 
  return int(C.SDL_JoystickEventState(C.int(state)))
}

// Get the current state of an axis control on a joystick
// The state is a value ranging from -32768 to 32767.
// The axis indices start at index 0.
func JoystickGetAxis(joystick * C.SDL_Joystick, axis int) (uint16) { 
  return uint16(C.SDL_JoystickGetAxis(joystick, C.int(axis)))
}

//
// Get the current state of a POV hat on a joystick
// The return value is one of the following positions:
const HAT_CENTERED	=	0x00
const HAT_UP		=	0x01
const HAT_RIGHT		=	0x02
const HAT_DOWN		=	0x04
const HAT_LEFT		=	0x08
const HAT_RIGHTUP	= 	HAT_RIGHT | HAT_UP
const HAT_RIGHTDOWN	= 	HAT_RIGHT | HAT_DOWN
const HAT_LEFTUP	= 	HAT_LEFT  | HAT_UP
const HAT_LEFTDOWN	= 	HAT_LEFT  | HAT_DOWN

//The hat indices start at index 0.
func JoystickGetHat(joystick * C.SDL_Joystick, hat int) (uint8) { 
  return uint8(C.SDL_JoystickGetHat(joystick, C.int(hat)))
}

// Get the ball axis change since the last poll
// This returns the change in x and y, and true, or 0 ,0, false 
// or -1 if you passed it invalid parameters.
// The ball indices start at index 0.
func JoystickGetBall(joystick * C.SDL_Joystick, ball int) (int, int, bool) {
  x  := 0 		; y := 0
  dx := (*C.int)(unsafe.Pointer(&x))	; dy := (*C.int)(unsafe.Pointer(&y))
  res := i2b(int(C.SDL_JoystickGetBall(joystick, C.int(ball), dx, dy)));
  if !res { return 0, 0, false }
  return x, y, true
}

// Get the current state of a button on a joystick
// The button indices start at index 0.
func JoystickGetButton(joystick * C.SDL_Joystick, button int) (uint8) { 
  return uint8(C.SDL_JoystickGetButton(joystick, C.int(button)))
}


// Close a joystick previously opened with JoystickOpen()
func JoystickClose(joystick * C.SDL_Joystick) { 
  C.SDL_JoystickClose(joystick)
}
/**/
