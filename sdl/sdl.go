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

/*
// free is a method on C int * pointers to method to free the associated memory 
func (self * C.int) free() {
  C.free(unsafe.Pointer(self))
}
*/
// cstring converts a string to a C string. This allocates memory, 
// so don't forget to add a "defer s.free()"
func cstr(self string) (*C.char) {
  return C.CString(self)
}

// Converts an int pointer to a C.int pointer
func cintptr(ptr * int)  (*C.int)  { 
  return (*C.int)(unsafe.Pointer(ptr))
}

/*
// cstring converts an int to a C int *. This allocates memory, 
// so don't forget to add a "defer s.free()"
func cintptrNew(self int) (* C.char) {
  return (*C.int) unsafe.Pointer(C.malloc(C.size_t())))
  return C.CString(self)
}
*/

type mystring string;

// Helper to convert strings to C strings 
func (self mystring) cstr() (* C.char) {
  return C.CString(string(self))
}

// converts ints to bools
func i2b(res int) (bool) {
  if res != 0 { return true } 
  return false
}

// converts bools to ints
func b2i(res bool) (int) {
  if res { return 1 } 
  return 0
}

// CPUINFO

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


// Error Handling

/* Does not compile due to CGO limitation.
// Sets the current SDL error message 
func SetError(fmt string) {
  cres  := cstr(fmt); defer cres.free()
  C.SDL_SetError(cres)
}
*/
// Gets the current error message of SDL
func GetError() (string) {
  res   := C.SDL_GetError()
  return C.GoString(res)
}

// Clears the error status of SDL
func ClearError() {
  C.SDL_ClearError();
}

// Error codes of SDL
type Errorcode int
const ( 
  ENOMEM = Errorcode(iota)
  EFREAD 
  EFWRITE
  EFSEEK
  UNSUPPORTED
  LASTERROR
)

// Raises an SDL error with the given code
func Error(code Errorcode) {
  C.SDL_Error(C.SDL_errorcode(code))
}

// Bindings to SDL_Image
// Load an image from an SDL data source.
// The 'type' may be one of: "BMP", "GIF", "PNG", etc.
// If the image format supports a transparent pixel, SDL will set the
// colorkey for the surface.  You can enable RLE acceleration on the
// surface afterwards by calling:
// SDL_SetColorKey(image, SDL_RLEACCEL, image->format->colorkey);
/*
For some reason, this function is not compile dproperly by cgo.
func ImgLoad(filename string) (* C.SDL_Surface) { 
  cfile := cstr(filename) ; defer cfile.free()
  return C.IMG_Load(cfile);
}
*/

// below not supported
// extern DECLSPEC SDL_Surface * SDLCALL IMG_LoadTyped_RW(SDL_RWops *src, int freesrc, char *type);
// Convenience functions 
//extern DECLSPEC SDL_Surface * SDLCALL IMG_Load_RW(SDL_RWops *src, int freesrc);
// extern DECLSPEC SDL_Surface * SDLCALL IMG_Load(const char *file);
//extern DECLSPEC SDL_Surface * SDLCALL IMG_ReadXPMFromArray(char **xpm);

// SDL RWOPS


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
  RW_SEEK_SET = 0 
  // Seek relative to current read point
  RW_SEEK_CUR = 1 
  // Seek relative to the end of data
  RW_SEEK_END = 2 
)

type seek_func * func (rwops * C.SDL_RWops, offset C.int, whence C.int) (C.int)

// I doubt these will work...
func RWSeek(rwops * C.SDL_RWops, offset, whence int) (int) {
  //return int(My_RWseek(rwops, C.int(offset), C.int(whence)))
  // int My_RWseek(SDL_RWops *ctx, int offset, int whence) { return (ctx)->seek(ctx, offset,whence); } 
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
  size    := buffer.cap
  n   := 1 
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
#define SDL_RWseek(ctx, offset, whence) (ctx)->seek(ctx, offset, whence)
#define SDL_RWtell(ctx)     (ctx)->seek(ctx, 0, RW_SEEK_CUR)
#define SDL_RWread(ctx, ptr, size, n) (ctx)->read(ctx, ptr, size, n)
#define SDL_RWwrite(ctx, ptr, size, n)  (ctx)->write(ctx, ptr, size, n)
#define SDL_RWclose(ctx)    (ctx)->close(ctx)
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
  if rw.rwops == nil { return nil }
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

// SDL mouse event handling

// Retrieve the current state of the mouse.
// The current button state is returned as a button bitmask, which can
// be tested using the sdl.BUTTON(X) functions, and x and y are set 
// to the current mouse cursor position.
func GetMouseState() (uint8, int, int) { 
  var x, y int
  px := (*C.int)(unsafe.Pointer(&x))
  py := (*C.int)(unsafe.Pointer(&y))
  but := uint8(C.SDL_GetMouseState(px, py))
  return but, x, y  
}


// Retrieve the current state of the mouse.
// The current button state is returned as a button bitmask, which can
// be tested using the SDL_BUTTON(X) macros, and x and y are set to the
// mouse deltas since the last call to SDL_GetRelativeMouseState().
func GetRelativeMouseState() (uint8, int, int) {
  var x, y int
  px := (*C.int)(unsafe.Pointer(&x))
  py := (*C.int)(unsafe.Pointer(&y))
  but := uint8(C.SDL_GetRelativeMouseState(px, py))
  return but, x, y  
}

// Set the position of the mouse cursor (generates a mouse motion event)
func WarpMouse(x, y uint16) { 
  C.SDL_WarpMouse(C.Uint16(x), C.Uint16(y))
}

//
// Create a cursor using the specified data and mask (in MSB format).
// The cursor width must be a multiple of 8 bits.
// The cursor is created in black and white according to the following:
// data  mask    resulting pixel on screen
//  0     1       White
//  1     1       Black
//  0     0       Transparent
//  1     0       Inverted color if possible, black if not.
//
// Cursors created with this function must be freed with
// SDL_FreeCursor().
/*
extern DECLSPEC SDL_Cursor * SDLCALL SDL_CreateCursor
    (Uint8 *data, Uint8 *mask, int w, int h, int hot_x, int hot_y);
*/

// Set the currently active cursor to the specified one.
// If the cursor is currently visible, the change will be immediately 
// represented on the display.

/*
extern DECLSPEC void SDLCALL SDL_SetCursor(SDL_Cursor *cursor);
*/

// Returns the currently active cursor.
/* 
extern DECLSPEC SDL_Cursor * SDLCALL SDL_GetCursor(void); 
*/

//
// Deallocates a cursor created with SDL_CreateCursor().
//
/*
extern DECLSPEC void SDLCALL SDL_FreeCursor(SDL_Cursor *cursor);
*/
// Toggle whether or not the cursor is shown on the screen.
// The cursor start off displayed, but can be turned off.
// ShowCursor() returns 1 if the cursor was being displayed
// before the call, or 0 if it was not.  You can query the current
// state by passing a 'toggle' value of -1.
/* 
extern DECLSPEC int SDLCALL SDL_ShowCursor(int toggle);
*/
// Used as a mask when testing buttons in buttonstate
// Button 1:  Left mouse button
// Button 2:  Middle mouse button
// Button 3:  Right mouse button
// Button 4:  Mouse wheel up   (may also be a real button)
// Button 5:  Mouse wheel down (may also be a real button)
func BUTTON(X uint8) (uint8) { 
  return (1 << ((X)-1))
}   

const BUTTON_LEFT       = 1
const BUTTON_MIDDLE     = 2
const BUTTON_RIGHT      = 3
const BUTTON_WHEELUP    = 4
const BUTTON_WHEELDOWN  = 5
/*
#define SDL_BUTTON_LMASK  SDL_BUTTON(SDL_BUTTON_LEFT)
#define SDL_BUTTON_MMASK  SDL_BUTTON(SDL_BUTTON_MIDDLE)
#define SDL_BUTTON_RMASK  SDL_BUTTON(SDL_BUTTON_RIGHT)
*/


// Time functions 
// This is the OS scheduler timeslice, in milliseconds
const TIMESLICE  = C.SDL_TIMESLICE

// This is the maximum resolution of the SDL timer on all platforms 
// Experimentally determined
const TIMER_RESOLUTION	= 10	

// Get the number of milliseconds since the SDL library initialization.
// Note that this value wraps if the program runs for more than ~49 days.
//
func GetTicks() (uint32) {  
  return uint32(C.SDL_GetTicks())
}

// Wait a specified number of milliseconds before returning 
func Delay(ms uint32) {
  C.SDL_Delay(C.Uint32(ms))
}

/* Function prototype for the timer callback function */
// Callback are not implemented yet.
// typedef Uint32 (SDLCALL *SDL_TimerCallback)(Uint32 interval);
/* Set a callback to run after the specified number of milliseconds has
 * elapsed. The callback function is passed the current timer interval
 * and returns the next timer interval.  If the returned value is the 
 * same as the one passed in, the periodic alarm continues, otherwise a
 * new alarm is scheduled.  If the callback returns 0, the periodic alarm
 * is cancelled.
 *
 * To cancel a currently running timer, call SDL_SetTimer(0, NULL);
 *
 * The timer callback function may run in a different thread than your
 * main code, and so shouldn't call any functions from within itself.
 *
 * The maximum resolution of this timer is 10 ms, which means that if
 * you request a 16 ms timer, your callback will run approximately 20 ms
 * later on an unloaded system.  If you wanted to set a flag signaling
 * a frame update at 30 frames per second (every 33 ms), you might set a 
 * timer for 30 ms:
 *   SDL_SetTimer((33/10)*10, flag_update);
 *
 * If you use this function, you need to pass SDL_INIT_TIMER to SDL_Init().
 *
 * Under UNIX, you should not use raise or use SIGALRM and this function
 * in the same program, as it is implemented using setitimer().  You also
 * should not use this function in multi-threaded applications as signals
 * to multi-threaded apps have undefined behavior in some implementations.
 *
 * This function returns 0 if successful, or -1 if there was an error.
 */
// extern DECLSPEC int SDLCALL SDL_SetTimer(Uint32 interval, SDL_TimerCallback callback);

/* New timer API, supports multiple timers
 * Written by Stephane Peter <megastep@lokigames.com>
 */

/* Function prototype for the new timer callback function.
 * The callback function is passed the current timer interval and returns
 * the next timer interval.  If the returned value is the same as the one
 * passed in, the periodic alarm continues, otherwise a new alarm is
 * scheduled.  If the callback returns 0, the periodic alarm is cancelled.
 */
//typedef Uint32 (SDLCALL *SDL_NewTimerCallback)(Uint32 interval, void *param);

/* Definition of the timer ID type */
//typedef struct _SDL_TimerID *SDL_TimerID;

/* Add a new timer to the pool of timers already running.
   Returns a timer ID, or NULL when an error occurs.
 */
//extern DECLSPEC SDL_TimerID SDLCALL SDL_AddTimer(Uint32 interval, SDL_NewTimerCallback callback, void *param);

/* Remove one of the multiple timers knowing its ID.
 * Returns a boolean value indicating success.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_RemoveTimer(SDL_TimerID t);

/* Ends C function definitions when using C++ */
