#include <SDL.h>


static int RWSeek(SDL_RWops * ctx, int offset, int whence) { 
  return SDL_RWseek(ctx, offset, whence); 
} 

