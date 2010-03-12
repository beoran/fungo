package sdl

// #include <SDL.h>
import "C"
// import "unsafe"

// General keyboard/mouse state definitions 
const (
  RELEASED 	= C.SDL_RELEASED
  PRESSED	= C.SDL_PRESSED
)

// Event enumerations 
type SDL_Events uint8

const (
  // Unused (do not remove)
  NOEVENT		= iota
  // Application loses/gains visibility 
  ACTIVEEVENT   	
  // Keys pressed 
  KEYDOWN		
  // Keys released
  KEYUP			
  // Mouse moved 
  MOUSEMOTION		
  // Mouse button pressed 
  MOUSEBUTTONDOWN	
  // Mouse button released
  MOUSEBUTTONUP		
  // Joystick axis motion
  JOYAXISMOTION		
  // Joystick trackball motion 
  JOYBALLMOTION		
  // Joystick hat position change 
  JOYHATMOTION		
  // Joystick button pressed 
  JOYBUTTONDOWN		
  // Joystick button released        
  JOYBUTTONUP		
  // User-requested quit 
  QUIT			
  // System specific event 
  SYSWMEVENT		
  // User resized video mode   
  VIDEORESIZE		
  // Screen needs to be redrawn 
  VIDEOEXPOSE		
  // Events SDL_USEREVENT through SDL_MAXEVENTS-1 are for your use 
  USEREVENT		= 24
  // This last event is only for bounding internal arrays
  // It is the number of bits in the event mask datatype -- Uint3    
  NUMEVENTS		= 32
)

type SDL_EventMasks uint32

func EVENTMASK(x SDL_Events) (SDL_EventMasks) { 
  return SDL_EventMasks((1<<(x)))
}  

var (
  ACTIVEEVENTMASK	= EVENTMASK(ACTIVEEVENT)
  KEYDOWNMASK		= EVENTMASK(KEYDOWN)
  KEYUPMASK		= EVENTMASK(KEYUP)
  KEYEVENTMASK		= EVENTMASK(KEYDOWN) | EVENTMASK(KEYUP)
  MOUSEMOTIONMASK	= EVENTMASK(MOUSEMOTION)
  MOUSEBUTTONDOWNMASK	= EVENTMASK(MOUSEBUTTONDOWN)
  MOUSEBUTTONUPMASK	= EVENTMASK(MOUSEBUTTONUP)
  MOUSEEVENTMASK	= EVENTMASK(MOUSEMOTION)|
			  EVENTMASK(MOUSEBUTTONDOWN)|
	                  EVENTMASK(MOUSEBUTTONUP)
  JOYAXISMOTIONMASK	= EVENTMASK(JOYAXISMOTION)
  JOYBALLMOTIONMASK	= EVENTMASK(JOYBALLMOTION)
  JOYHATMOTIONMASK	= EVENTMASK(JOYHATMOTION)
  JOYBUTTONDOWNMASK	= EVENTMASK(JOYBUTTONDOWN)
  JOYBUTTONUPMASK	= EVENTMASK(JOYBUTTONUP)
  JOYEVENTMASK		= EVENTMASK(JOYAXISMOTION)|
	                  EVENTMASK(JOYBALLMOTION)|
	                  EVENTMASK(JOYHATMOTION)|
	                  EVENTMASK(JOYBUTTONDOWN)|
	                  EVENTMASK(JOYBUTTONUP)
  VIDEORESIZEMASK	= EVENTMASK(VIDEORESIZE)
  VIDEOEXPOSEMASK	= EVENTMASK(VIDEOEXPOSE)
  QUITMASK		= EVENTMASK(QUIT)
  SYSWMEVENTMASK	= EVENTMASK(SYSWMEVENT)
)
const ALLEVENTS		= SDL_EventMasks(0xFFFFFFFF)

type ActiveEvent 	C.SDL_ActiveEvent 
type KeyboardEvent 	C.SDL_KeyboardEvent
type MouseMotionEvent 	C.SDL_MouseMotionEvent
type MouseButtonEvent 	C.SDL_MouseButtonEvent
type JoyAxisEvent	C.SDL_JoyAxisEvent
type JoyBallEvent 	C.SDL_JoyBallEvent
type JoyHatEvent 	C.SDL_JoyHatEvent
type JoyButtonEvent 	C.SDL_JoyButtonEvent
type ResizeEvent	C.SDL_ResizeEvent
type ExposeEvent	C.SDL_ExposeEvent
type QuitEvent		C.SDL_QuitEvent
type UserEvent 		C.SDL_UserEvent
type SysWMmsg 		C.SDL_SysWMmsg
type SysWMEvent		C.SDL_SysWMEvent


type EventUnion 	C.SDL_Event
// Little trick we use here to get a type field in there
type Event struct {
	Type uint8
	Pad0 [31]byte 
	// sizeof(SDL_Event) is 20 on a 32 bits platform 
	// but to be sure on 64 bit platforms, I made it a bit bigger
}

type Keysym		C.SDL_keysym 

// Predefined event masks 
func (e * ActiveEvent) Type() (byte) {
  return GetType(ptr(e))  
}

// Whether given states were gained or lost (1/0) 
func (e * ActiveEvent) Gain() (bool) {
  return i2b(int(e.gain))
}

// A mask of the focus states 
func (e * ActiveEvent) State() (byte) {
  return byte(e.state)
}

// A mask of the focus states 
func (e * KeyboardEvent) State() (byte) {
  return byte(e.state)
}

// Which keyboard generated the event
func (e * KeyboardEvent) Which() (byte) {
  return byte(e.which)
}

// Type of the event
func (e * MouseMotionEvent) Type() (byte) {
  return GetType(ptr(e))  
}

// Which mouse generated the event
func (e * MouseMotionEvent) Which() (byte) {
  return byte(e.which)
}

// Current mouse button state
func (e * MouseMotionEvent) State() (byte) {
  return byte(e.state)
}
// Current X coordinates of cursor
func (e * MouseMotionEvent) X() (int) {
  return int(e.x)
}

// Current Y coordinates of cursor
func (e * MouseMotionEvent) Y() (int) {
  return int(e.y)
}

// Relative motion along X axis
func (e * MouseMotionEvent) XRel() (int) {
  return int(e.xrel)
}
// Relatve motion along Y axis
func (e * MouseMotionEvent) YRel() (int) {
  return int(e.yrel)
}

// Type of the event
func (e * MouseButtonEvent) Type() (byte) {
  return GetType(ptr(e))  
}

// Which mouse generated the event
func (e * MouseButtonEvent) Which() (byte) {
  return byte(e.which)
}

// The button that was pressed or released
func (e * MouseButtonEvent) Button() (byte) {
  return byte(e.button)
}

// PRESSED or RELEASED
func (e * MouseButtonEvent) State() (byte) {
  return byte(e.state)
}

// X coordinates of click or release
func (e * MouseButtonEvent) X() (int) {
  return int(e.x)
}

// Y coordinates of click or release
func (e * MouseButtonEvent) Y() (int) {
  return int(e.y)
}

// Type of the event
func (e * JoyAxisEvent) Type() (byte) {
  return GetType(ptr(e))  
}

// Which joystick generated the event
func (e * JoyAxisEvent) Which() (byte) {
  return byte(e.which)
}

// Which axis moved
func (e * JoyAxisEvent) Axis() (byte) {
  return byte(e.which)
}

// Joystick axis current value
func (e * JoyAxisEvent) Value() (int) {
  return int(e.value)
}

// Type of the event
func (e * JoyBallEvent) Type() (byte) {
  return GetType(ptr(e))  
}

// Which joystick generated the event
func (e * JoyBallEvent) Which() (byte) {
  return byte(e.which)
}

// Which ball on the joystick generated the event
func (e * JoyBallEvent) Ball() (byte) {
  return byte(e.which)
}

// Relative motion along X axis
func (e * JoyBallEvent) XRel() (int) {
  return int(e.xrel)
}

// Relatve motion along Y axis
func (e * JoyBallEvent) YRel() (int) {
  return int(e.yrel)
}

// Type of the event
func (e * JoyHatEvent) Type() (byte) {
  return GetType(ptr(e))  
}

// Which joystick generated the event
func (e * JoyHatEvent) Which() (byte) {
  return byte(e.which)
}

// Which hat moved
func (e * JoyHatEvent) Hat() (byte) {
  return byte(e.which)
}

// Joystick hat current value
func (e * JoyHatEvent) Value() (byte) {
  return byte(e.value)
}


// Type of the event
func (e * JoyButtonEvent) Type() (byte) {
  return GetType(ptr(e))  
}

// Which joystick generated the event
func (e * JoyButtonEvent) Which() (byte) {
  return byte(e.which)
}

// The button that was pressed or released
func (e * JoyButtonEvent) Button() (byte) {
  return byte(e.button)
}

// PRESSED or RELEASED
func (e * JoyButtonEvent) State() (byte) {
  return byte(e.state)
}

func (e * ResizeEvent) Type() (byte) {
  return GetType(ptr(e))  
}

func (e * ResizeEvent) W() (int) {
  return int(e.w)
}

func (e * ResizeEvent) H() (int) {
  return int(e.h)
}

func (e * ExposeEvent) Type() (byte) {
  return GetType(ptr(e))  
}

func (e * QuitEvent) Type() (byte) {
  return GetType(ptr(e))  
}

// TODO: make UserEvent actually useful.
// data1, data2
func (e * UserEvent) Type() (byte) {
  return GetType(ptr(e))  
}

func (e * UserEvent) Code() (int) {
  return int(e.code)  
}


// Pumps the event loop, gathering events from the input devices.
// This function updates the event queue and internal input device state.
// This should only be run in the thread that sets the video mode.
func PumpEvents() {
  C.SDL_PumpEvents()
}

type SDL_eventaction int

const ( 
  ADDEVENT 	= SDL_eventaction(C.SDL_ADDEVENT)
  PEEKEVENT 	= SDL_eventaction(C.SDL_PEEKEVENT)
  GETEVENT  	= SDL_eventaction(C.SDL_GETEVENT)
)
// Checks the event queue for messages and optionally returns them.
// If 'action' is ADDEVENT, up to 'numevents' events will be added to
// the back of the event queue.
// If 'action' is PEEKEVENT, up to 'numevents' events at the front
// of the event queue, matching 'mask', will be returned and will not
// be removed from the queue.
// If 'action' is GETEVENT, up to 'numevents' events at the front 
// of the event queue, matching 'mask', will be returned and will be
// removed from the queue.
// This function returns the number of events actually stored, or -1
// if there was an error.  This function is thread-safe.
func peepEvents( events * C.SDL_Event, numevents int, 
		 action SDL_eventaction, mask uint32) (int) {
  return int(C.SDL_PeepEvents(events, C.int(numevents), 
    C.SDL_eventaction(action), C.Uint32(mask)))
}


// Polls for currently pending events, and returns 1 if there are any pending
// events, or 0 if there are none available.  If 'event' is not NULL, the next
// event is removed from the queue and stored in that area.
func pollEvent(event * C.SDL_Event) (int) { 
  return int(C.SDL_PollEvent(event))
}
  

// Waits indefinitely for the next available event, returning 1, or 0 if there
// was an error while waiting for events.  If 'event' is not NULL, the next
// event is removed from the queue and stored in that area.
func waitEvent(event * C.SDL_Event) (int) { 
  return int(C.SDL_WaitEvent(event))
}
 
// Add an event to the event queue.
// This function returns 0 on success, or -1 if the event queue was full
// or there was some other error.
func pushEvent(event * C.SDL_Event) (int) { 
  return int(C.SDL_PushEvent(event))
}

// Event filtereing is not supported yet due to cg not supporting callbacks 
// well.
//
// This function sets up a filter to process all events before they
// change internal state and are posted to the internal event queue.
// The filter is protypted as:
// typedef int (SDLCALL *SDL_EventFilter)(const SDL_Event *event);
//
//   If the filter returns 1, then the event will be added to the internal queue.
//   If it returns 0, then the event will be dropped from the queue, but the 
//   internal state will still be updated.  This allows selective filtering of
//   dynamically arriving events.
// 
//   WARNING:  Be very careful of what you do in the event filter function, as 
//             it may run in a different thread!
// 
//   There is one caveat when dealing with the SDL_QUITEVENT event type.  The
//   event filter is only called when the window manager desires to close the
//   application window.  If the event filter returns 1, then the window will
//   be closed, otherwise the window will remain open if possible.
//   If the quit event is generated by an interrupt signal, it will bypass the
//   internal queue and be delivered to the application at the next event poll.
// 
// extern DECLSPEC void SDLCALL SDL_SetEventFilter(SDL_EventFilter filter);
//
//   Return the current event filter - can be used to "chain" filters.
//   If there is no event filter set, this function returns NULL.
// 
// extern DECLSPEC SDL_EventFilter SDLCALL SDL_GetEventFilter(void);


const (
  QUERY		= -1
  IGNORE	= 0
  DISABLE	= 0
  ENABLE	= 1
)
//
//   This function allows you to set the state of processing certain events.
//   If 'state' is set to SDL_IGNORE, that event will be automatically dropped
//   from the event queue and will not event be filtered.
//   If 'state' is set to SDL_ENABLE, that event will be processed normally.
//   If 'state' is set to SDL_QUERY, SDL_EventState() will return the 
//   current processing state of the specified event.
func EventState(kind uint8, state int) (uint8) {
  return uint8(C.SDL_EventState(C.Uint8(kind), C.int(state)))
}

// Returns true if a quit was requested.
func QuitRequested() (bool) { 
  C.SDL_PumpEvents()
  return i2b(int(C.SDL_PeepEvents(nil, 0, C.SDL_eventaction(PEEKEVENT), 		
	  C.Uint32(QUITMASK))))
}	

func GetType(e ptr) (uint8) { 
  return ((*Event)(e)).Type
}

func (e * KeyboardEvent) Type() (uint8) {
  return GetType(ptr(e))
}


func (e * KeyboardEvent) Keysym() (*Keysym) {
  return (*Keysym)(&e.keysym)
}


func (k * Keysym) Unicode() (int) {
  return int(k.unicode)
}

func (k * Keysym) Mod() (byte) {
  return byte(k.mod)
}

func (k * Keysym) Sym() (int) {
  return int(k.sym)
}

func (k * Keysym) Scancode() (int) {
  return int(k.scancode)
}
