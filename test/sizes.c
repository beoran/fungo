#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>
#include <SDL/SDL.h>




int main(void) {
  printf("SDL_Event Size: %d\n", sizeof(SDL_Event));
  printf("SDL_keysym Size: %d\n", sizeof(SDL_keysym));
  printf("SDLKey Size: %d\n", sizeof(SDLKey));
  printf("SDLMod Size: %d\n", sizeof(SDLMod));
  printf("Scancode: %d\n", offsetof(SDL_keysym, scancode));
  printf("Sym: %d\n", offsetof(SDL_keysym, sym));
  printf("Mod: %d\n", offsetof(SDL_keysym, mod));
  printf("Unicode: %d\n", offsetof(SDL_keysym, unicode));
}
