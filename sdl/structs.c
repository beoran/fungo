


/* 
  We need a little trick to get keysyms right. 
  Keysym uses two enums in it's struct, but the godef tools
  misunderstands them as char types. The size of an enum
  with ggc is normally sizeof(int)0
*/
/*
typedef struct SDL_keysym {
	Uint8 scancode;			
	int sym;			
	int mod;			
	Uint16 unicode;			
} SDL_keysym;

#define SDL_keysym SDL_keysym_original
*/
#include <SDL/SDL.h>
/*
#undef SDL_keysym SDL_keysym_original
*/
/* 
  Generates the structs that are needed in go so they conform
  to what SDL and C expect
*/   

/*
typedef SDL_Surface $Surface;
typedef SDL_PixelFormat $PixelFormat;
typedef SDL_Palette $Palette;
typedef SDL_VideoInfo $VideoInfo;
typedef SDL_Overlay $Overlay;
*/


typedef SDL_Rect $Rect;
typedef SDL_Color $Color;
/*
typedef SDL_ActiveEvent $ActiveEvent;
typedef SDL_KeyboardEvent $KeyboardEvent;
typedef SDL_MouseMotionEvent $MouseMotionEvent;
typedef SDL_MouseButtonEvent $MouseButtonEvent;
typedef SDL_JoyAxisEvent $JoyAxisEvent;
typedef SDL_JoyBallEvent $JoyBallEvent;
typedef SDL_JoyHatEvent $JoyHatEvent;
typedef SDL_JoyButtonEvent $JoyButtonEvent;
typedef SDL_ResizeEvent $ResizeEvent;
typedef SDL_ExposeEvent $ExposeEvent;
typedef SDL_QuitEvent $QuitEvent;
typedef SDL_UserEvent $UserEvent;
typedef SDL_SysWMmsg $SysWMmsg;
typedef SDL_SysWMEvent $SysWMEvent;
typedef SDL_Event $Event;
typedef SDL_keysym $Keysym;
*/
