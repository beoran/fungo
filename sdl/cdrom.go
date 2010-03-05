package sdl

// #include <SDL.h>
import "C"
import "unsafe" 



/* In order to use these functions, SDL_Init() must have been called
   with the SDL_INIT_CDROM flag.  This causes SDL to scan the system
   for CD-ROM drives, and load appropriate drivers.
*/

/* The maximum number of CD-ROM tracks on a disk */
const MAX_TRACKS = C.SDL_MAX_TRACKS

/* The types of CD-ROM track possible */
const AUDIO_TRACK = C.SDL_AUDIO_TRACK
const DATA_TRACK  = C.SDL_DATA_TRACK

type CDstatus int;
const ( 
  CD_TRAYEMPTY  = CDstatus(iota)
  CD_STOPPED
  CD_PLAYING
  CD_PAUSED
  CD_ERROR      = CDstatus(-1)

)


/* Given a status, returns true if there's a disk in the drive */
func CD_INDRIVE(status) (bool) 
  return ((int)(status) > 0)
}

/*
typedef struct SDL_CDtrack {
  Uint8 id;   
  Uint8 type;   
  Uint16 unused;
  Uint32 length;  
  Uint32 offset;  
} SDL_CDtrack;
*/

/* This structure is only current as of the last call to SDL_CDStatus() */
/*
typedef struct SDL_CD {
  int id;     
  CDstatus status;  
  int numtracks;    
  int cur_track;    
  int cur_frame;    
  SDL_CDtrack track[SDL_MAX_TRACKS+1];
} SDL_CD;
*/

/* Conversion functions from frames to Minute/Second/Frames and vice versa */
const CD_FPS=75
func FRAMES_TO_MSF(value int) (m, s, f int) {
  f     = value % CD_FPS
  value = value / CD_FPS
  s     = value % 60
  value = value / 60
  m     = value   
  return 
}  
  
func MSF_TO_FRAMES(m, s, f int) {
  return ((m)*60*CD_FPS+(s)*CD_FPS+(f))
}

// CD-audio API functions
// Returns the number of CD-ROM drives on the system, or -1 if
// SDL_Init() has not been called with the SDL_INIT_CDROM flag.
func CDNumDrives() int { 
  int(C.SDL_CDNumDrives(void))
}

// Returns a human-readable, system-dependent identifier for the CD-ROM.
// Example: "/dev/cdrom"  "E:"  "/dev/disk/ide/1/master"
func CDName(drive int) string { 
  return C.GoString(C.SDL_CDName(C.int(drive)))
}

// Opens a CD-ROM drive for access.  It returns a drive handle on success,
// or NULL if the drive was invalid or busy.  This newly opened CD-ROM
// becomes the default CD used when other CD functions are passed a NULL
// CD-ROM handle.
// Drives are numbered starting with 0.  Drive 0 is the system default CD-ROM.
func CDOpen(drive int) (* C.SDL_CD) { 
  return C.SDL_CDOpen(C.int(drive))
}

// This function returns the current status of the given drive.
// If the drive has a CD in it, the table of contents of the CD and current
// play position of the CD will be stored in the SDL_CD structure.
func CDStatus(cdrom * C.SDL_CD)  (CDstatus) { 
  return CDstatus(C.SDL_CDStatus(cdrom))
}

// Play the given CD starting at 'start_track' and 'start_frame' for 'ntracks'
// tracks and 'nframes' frames.  If both 'ntrack' and 'nframe' are 0, play 
// until the end of the CD.  This function will skip data tracks.
// This function should only be called after calling SDL_CDStatus() to 
// get track information about the CD.
// For example:
// Play entire CD:
//  if ( CD_INDRIVE(SDL_CDStatus(cdrom)) ) { CDPlayTracks(cdrom, 0, 0, 0, 0); }
// Play last track:
//  if ( CD_INDRIVE(SDL_CDStatus(cdrom)) ) {
//    CDPlayTracks(cdrom, cdrom->numtracks-1, 0, 0, 0);
//  }
// Play first and second track and 10 seconds of third track:
//  if ( CD_INDRIVE(SDL_CDStatus(cdrom)) ) {
//    CDPlayTracks(cdrom, 0, 0, 2, 10); } 
//
// This function returns 0, or -1 if there was an error.
func CDPlayTracks(cdrom * C.SDL_CD, start_track, start_frame, 
                  ntracks, nframes int) (int) {           
  return int(C.SDL_CDPlayTracks(cdrom, C.int(start_track), C.int(start_frame),
                               C.int(ntracks), C.int(nframes));
}

// Play the given CD starting at 'start' frame for 'length' frames.
// It returns 0, or -1 if there was an error.
CDPlay(cdrom * C.SDL_CD, start, length) (int) {
  return int(C.SDL_CDPlay(cdrom, C.int(start), C.int(length)); 
} 

/* Pause play -- returns 0, or -1 on error */
CDPause(cdrom * C.SDL_CD) (int) {
  return int(C.SDL_CDPause(cdrom))
}  

/* Resume play -- returns 0, or -1 on error */
CDResume(cdrom * C.SDL_CD) (int) {
  return int(C.SDL_CDResume(cdrom))
}  


extern DECLSPEC int SDLCALL SDL_CDResume(SDL_CD *cdrom);

/* Stop play -- returns 0, or -1 on error */
CDStop(cdrom * C.SDL_CD) (int) {
  return int(C.SDL_CDStop(cdrom))
}  

/* Eject CD-ROM -- returns 0, or -1 on error */
CDEject(cdrom * C.SDL_CD) (int) {
  return int(C.SDL_CDEject(cdrom))
}  

/* Closes the handle for the CD-ROM drive */
CDClose(cdrom * C.SDL_CD) {
  C.SDL_CDClose(cdrom)
}  


