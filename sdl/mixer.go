//
// Bindings to SDL_Mixer

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
import "runtime"

// The default mixer has 8 simultaneous mixing channels 
const CHANNELS = 8

// Good default values for a PC soundcard 
const DEFAULT_FREQUENCY = 22050
 
// const MIX_DEFAULT_FORMAT   = AUDIO_S16LSB
//MIX_DEFAULT_FORMAT = sdl.AUDIO_S16MSB
const MIX_DEFAULT_CHANNELS = 2
const MIX_MAX_VOLUME       = 128 // Volume of a chunk 


// The different fading types supported 
type Mix_Fading int 


const (
  NO_FADING  = Mix_Fading(iota) 
  FADING_OUT
  FADING_IN
)

type Mix_MusicType int

const (
  MUS_NONE  = Mix_MusicType(iota)
  MUS_CMD
  MUS_WAV
  MUS_MOD
  MUS_MID
  MUS_OGG
  MUS_MP3
)

var MixerOpened bool = false

func OpenMixerDefault() (bool) {
  format 	:= uint16(FAUDIO_U16LSB)
  chunksize 	:= 1024*8
  if BYTEORDER == BIG_ENDIAN { format = FAUDIO_U16MSB }   
  ok := (OpenMixer(DEFAULT_FREQUENCY, format, MIX_DEFAULT_CHANNELS, chunksize))
  return ok 
}

// Open the mixer with a certain audio format 
func OpenMixer(frequency int, format uint16, channels int, chunksize int) (bool) {     
  res := (int(C.Mix_OpenAudio(C.int(frequency), 
             C.Uint16(format), C.int( channels), C.int(chunksize))))
  ok  := (res == 0)
  MixerOpened  = ok 	     
  return ok
}

// Dynamically change the number of channels managed by the mixer.
// If decreasing the number of channels, the upper channels are
// stopped. This function returns the new number of allocated channels.
func AllocateChannels(numchans int) (int) {
  return int(C.Mix_AllocateChannels(C.int(numchans)))
}
 
// Find out what the actual audio device parameters are.
// This function returns 1 if the audio has been opened, 0 otherwise.
// it also return the specified values
func QuerySpec() (int, int, uint16, int) {
  var ok, frequency, channels int 
  var format uint16
  pfreq := (*C.int)(unsafe.Pointer(&frequency))
  pform := (*C.Uint16)(unsafe.Pointer(&format))
  pchan := (*C.int)(unsafe.Pointer(&channels)) 
  ok    = int(C.Mix_QuerySpec(pfreq, pform, pchan))
  return ok, frequency, format, channels
}

// Load a wave file or a music (.mod .s3m .it .xm) file
func LoadWAV_RW(src * C.SDL_RWops, freesrc int) (* C.Mix_Chunk) { 
  return C.Mix_LoadWAV_RW(src, C.int(freesrc))
}
 
func LoadWAV(filename string)  (* C.Mix_Chunk) {  
  rwops   := RWFromFile(filename, "r")
  
  
  if rwops == nil { return nil }
  wav := C.Mix_LoadWAV_RW(rwops, C.int(1))
  // wav := LoadWAV_RW(rwops, 1)
   
  println(filename, rwops, wav, GetError())
  
  return wav
}

func LoadMUS(res string) (* C.Mix_Music) {
  cres := cstr(res) ; defer cres.free()   
  return C.Mix_LoadMUS(cres)
}
// Load a wave file of the mixer format from a memory buffer 
// extern DECLSPEC Mix_Chunk * SDLCALL Mix_QuickLoad_WAV(Uint8 *mem);

// Load raw audio data of the mixer format from a memory buffer 
// extern DECLSPEC Mix_Chunk * SDLCALL Mix_QuickLoad_RAW(Uint8 *mem, 
// Uint32 len);

// Free an audio chunk previously loaded
func FreeSound(wave * C.Mix_Chunk) {
  C.Mix_FreeChunk(wave)
} 
 
// Free audio music previously loaded
func FreeMusic(music * C.Mix_Music) {
  C.Mix_FreeMusic(music)
} 

// Find out the music format of a mixer music, or the currently playing
// music, if 'music' is NULL.
func MusicType(music * C.Mix_Music) (Mix_MusicType) {
  return Mix_MusicType(C.Mix_GetMusicType(music)) 
}

// Many callback functions of Mixer are not supported.
// perhaps through callback.h later 

// Set the panning of a channel. The left and right channels are specified
// as integers between 0 and 255, quietest to loudest, respectively.
//
// Technically, this is just individual volume control for a sample with
// two (stereo) channels, so it can be used for more than just panning.
// If you want real panning, call it like this:
//
// Mix_SetPanning(channel, left, 255 - left);
//
// ...which isn't so hard.
//
// returns zero if error
// nonzero if panning effect enabled. Note that an audio device in mono
// mode is a no-op, but this call will return successful in that case.
// Error messages can be retrieved from Mix_GetError().
func SetPanning(channel int left, right uint8) (int) {
  return int(C.Mix_SetPanning(C.int(channel), C.Uint8(left), C.Uint8(right)))
}

// Set the position of a channel. (angle) is an integer from 0 to 360, that
// specifies the location of the sound in relation to the listener. (angle)
// will be reduced as neccesary (540 becomes 180 degrees, -100 becomes 260).
// Angle 0 is due north, and rotates clockwise as the value increases.
// For efficiency, the precision of this effect may be limited (angles 1
// through 7 might all produce the same effect, 8 through 15 are equal, etc).
// (distance) is an integer between 0 and 255 that specifies the space
// between the sound and the listener. The larger the number, the further
// away the sound is. Using 255 does not guarantee that the channel will be
// culled from the mixing process or be completely silent. For efficiency,
// the precision of this effect may be limited (distance 0 through 5 might
// all produce the same effect, 6 through 10 are equal, etc). Setting (angle)
// and (distance) to 0 unregisters this effect, since the data would be
// unchanged.
//
// If the audio device is configured for mono output, then you won't get
// any effectiveness from the angle; however, distance attenuation on the
// channel will still occur. While this effect will function with stereo
// voices, it makes more sense to use voices with only one channel of sound,
// so when they are mixed through this effect, the positioning will sound
// correct. You can convert them to mono through SDL before giving them to
// the mixer in the first place if you like.
// This is a convenience wrapper over SetDistance() and SetPanning().
//
// returns zero if error,
// nonzero if position effect is enabled.
// Error messages can be retrieved from Mix_GetError().
func SetPosition(channel int, angle int16, distance uint8) (int) {  
 return int(C.Mix_SetPosition(C.int(channel), C.Sint16(angle), 
              C.Uint8(distance)))
}

// Set the "distance" of a channel. (distance) is an integer from 0 to 255
// that specifies the location of the sound in relation to the listener.
// Distance 0 is overlapping the listener, and 255 is as far away as possible
// A distance of 255 does not guarantee silence; in such a case, you might
// want to try changing the chunk's volume, or just cull the sample from the
// mixing process with Mix_HaltChannel().
// For efficiency, the precision of this effect may be limited (distances 1
// through 7 might all produce the same effect, 8 through 15 are equal, etc).
// (distance) is an integer between 0 and 255 that specifies the space
// between the sound and the listener. The larger the number, the further
// away the sound is.
// Setting (distance) to 0 unregisters this effect, since the data would be
// unchanged.
// returns zero if error, nonzero if position effect is enabled.
// Error messages can be retrieved from Mix_GetError().
func SetDistance(channel int, distance uint8) (int) { 
  return int(C.Mix_SetDistance(C.int(channel), C.Uint8(distance)))
}



// Causes a channel to reverse its stereo. This is handy if the user has his
// speakers hooked up backwards, or you would like to have a minor bit of
// psychedelia in your sound code.  :)  Calling this function with (flip)
// set to non-zero reverses the chunks's usual channels. If (flip) is zero,
// the effect is unregistered. 
// returns zero if error nonzero if reversing effect is enabled. 
// Note that an audio device in mono mode is a no-op, but this call 
// will return successful in that case.
// Error messages can be retrieved from Mix_GetError().
func SetReverseStereo(channel int, flip int) (int)  { 
  return int(C.Mix_SetReverseStereo(C.int(channel), C.int(flip)))
}


// Reserve the first channels (0 -> n-1) for the application, i.e. 
// don't allocate them dynamically to the next sample if requested 
// with a -1 value below.  Returns the number of reserved channels.
func ReserveChannels(num int) (int) { 
  return int(C.Mix_ReserveChannels(C.int(num)))
}

// Channel grouping functions 
// Attach a tag to a channel. A tag can be assigned to several mixer
// channels, to form groups of channels.
// If 'tag' is -1, the tag is removed (actually -1 is the tag used to
// represent the group of all the channels).
// Returns true if everything was OK.
func GroupChannel(which, tag int) (int) {
  return int(C.Mix_GroupChannel(C.int(which), C.int(tag)))  
}
 
// Assign several consecutive channels to a group
func GroupChannels(from, to, tag int) (int) {
  return int(C.Mix_GroupChannels(C.int(from), C.int(to), C.int(tag)))  
}
// Finds the first available channel in a group of channels,
// returning -1 if none are available.
func GroupAvailable(tag int) (int) {
  return int(C.Mix_GroupAvailable(C.int(tag)))  
}

// Returns the number of channels in a group. This is also a subtle
// way to get the total number of channels when 'tag' is -1
// Finds the first available channel in a group of channels,
// returning -1 if none are available.
func GroupCount(tag int) (int) {
  return int(C.Mix_GroupCount(C.int(tag)))  
}
  
// Finds the "oldest" sample playing in a group of channels 
func GroupOldest(tag int) (int) {
  return int(C.Mix_GroupOldest(C.int(tag)))  
}

// Finds the "most recent" (i.e. last) sample playing in a group of channels 
func GroupNewest(tag int) (int) {
  return int(C.Mix_GroupNewer(C.int(tag)))  
}

// Play an audio chunk on a specific channel.
// If the specified channel is -1, play on the first free channel.
// If 'loops' is greater than zero, loop the sound that many times.
// If 'loops' is -1, loop inifinitely (~65000 times).
// Returns which channel was used to play the sound.
// The the sound is played at most 'ticks' milliseconds.
func PlayChannelTimed(channel int , chunk * C.Mix_Chunk, 
     loops, ticks int) (int) {
  return int(C.Mix_PlayChannelTimed(C.int(channel), chunk, 
    C.int(loops), C.int(ticks))) 
}

// Play an audio chunk on a specific channel.
// If the specified channel is -1, play on the first free channel.
// If 'loops' is greater than zero, loop the sound that many times.
// If 'loops' is -1, loop inifinitely (~65000 times).
// Returns which channel was used to play the sound.
func PlayChannel(channel int, chunk * C.Mix_Chunk, loops int) (int) {
  return PlayChannelTimed(channel, chunk, loops, -1)
}

// Play music
// If 'loops' is greater than zero, loop the sound that many times.
// If 'loops' is -1, loop inifinitely (~65000 times).
func PlayMusic(music * C.Mix_Music, loops int) (int) {
  return int(C.Mix_PlayMusic(music, C.int(loops)))
}

// Fade in music or a channel over "ms" milliseconds, same semantics 
// as the "Play" functions 
func FadeInMusic(music * C.Mix_Music, loops int, ms int) (int) {
  return int(C.Mix_FadeInMusic(music, C.int(loops), C.int(ms)))
}

func FadeInMusicPos(music * C.Mix_Music, loops int, 
                    ms int, pos float) (int) {
  return int(C.Mix_FadeInMusicPos(music, C.int(loops), C.int(ms), C.double(pos)))
}

func FadeInChannelTimed(channel int, chunk * C.Mix_Chunk, loops, ms, 
                  ticks int) (int) {
  return int(C.Mix_FadeInChannelTimed(C.int(channel), chunk, 
    C.int(loops), C.int(ms), C.int(ticks))) 
}

func FadeInChannel(channel int, chunk * C.Mix_Chunk, loops, ms, 
                  ticks int) (int) {
  return FadeInChannelTimed(channel, chunk, loops, ms, -1) 
}


// Set the volume in the range of 0-128 of a specific channel or chunk.
// If the specified channel is -1, set volume for all channels.
// Returns the original volume.
// If the specified volume is -1, just return the current volume.
func VolumeChannel(channel, volume int) (int) {
  return int(C.Mix_Volume(C.int(channel), C.int(volume)));
}

func VolumeChunk(chunk * C.Mix_Chunk, volume int) (int) {
  return int(C.Mix_VolumeChunk(chunk, C.int(volume)));
}

func VolumeMusic(volume int) (int) {
  return int(C.Mix_VolumeMusic(C.int(volume)));
}

// Halt playing of a particular channel
func HaltChannel(channel int) (int) { 
  return int(C.Mix_HaltChannel(C.int(channel)))
}

// Halt playing of a particular Group
func HaltGroup(tag int) (int) { 
  return int(C.Mix_HaltGroup(C.int(tag)))
}

// Halt playing of the music
func HaltMusic() (int) { 
  return int(C.Mix_HaltMusic())
}

// Change the expiration delay for a particular channel.
// The sample will stop playing after the 'ticks' milliseconds have elapsed,
// or remove the expiration if 'ticks' is -1
func ExpireChannel(channel, ticks int) (int) {
  return int(C.Mix_ExpireChannel(C.int(channel), C.int(ticks)))
}

// Halt a channel, fading it out progressively till it's silent
// The ms parameter indicates the number of milliseconds the fading
// will take.
// Fade out a particular channel
func FadeOutChannel(channel, ms int) (int) { 
  return int(C.Mix_FadeOutChannel(C.int(channel), C.int(ms)))
}

// Halt playing of a particular group
func FadeOutGroup(tag, ms int) (int) { 
  return int(C.Mix_FadeOutGroup(C.int(tag), C.int(ms)))
}

// Halt playing of the music
func FadeOutMusic(ms int) (int) {
  return int(C.Mix_FadeOutMusic(C.int(ms)))
}

// Query the fading status of a the music
func FadingMusic() (Mix_Fading) { 
  return Mix_Fading(C.Mix_FadingMusic()) 
}

// Query the fading status of a channel 
func FadingChannel(channel int) (Mix_Fading) { 
  return Mix_Fading(C.Mix_FadingChannel(C.int(channel))) 
}

// Pause/Resume a particular channel
func PauseChannel(channel int) { 
  C.Mix_Pause(C.int(channel))
}
  
func ResumeChannel(channel int) { 
  C.Mix_Resume(C.int(channel))
}

func PausedChannel(channel int) (int) { 
  return int(C.Mix_Paused(C.int(channel)))
}

// Pause the music stream
func PauseMusic() { 
  C.Mix_PauseMusic()
}

// Resume the music stream  
func ResumeMusic() { 
  C.Mix_ResumeMusic()
}

// Rewinds the music stream  
func RewindMusic() { 
  C.Mix_RewindMusic()
}

// Stops the music stream  
func StopMusic() { 
  HaltMusic()
}

// Returns nonzero if the music is paused
func PausedMusic() (int) { 
  return int(C.Mix_PausedMusic())
}

// Set the current position in the music stream.
// This returns 0 if successful, or -1 if it failed or isn't implemented.
// This function is only implemented for MOD music formats (set pattern
// order number) and for OGG music (set position in seconds), at the
// moment.
func SetMusicPosition(position float64) (int) {
  return int(C.Mix_SetMusicPosition(C.double(position)))
} 

// Check the status of a specific channel.
// If the specified channel is -1, check all channels.
func PlayingChannel(channel int) (int) { 
  return int(C.Mix_Playing(C.int(channel)))
}

// Check the status of the playing music.
func PlayingMusic() (int) { 
  return int(C.Mix_PlayingMusic())
}

// Stop music and set external music playback command
func SetMusicCommand(command string) (int) { 
  ccommand := cstr(command) ; defer ccommand.free()
  return int(C.Mix_SetMusicCMD(ccommand))
}
  
// Synchro value is set by MikMod from modules while playing
func SetSynchroValue(value int) (int) {
  return int(C.Mix_SetSynchroValue(C.int(value)))
} 

// Synchro value is set by MikMod from modules while playing
func GetSynchroValue() (int) {
  return int(C.Mix_GetSynchroValue())
} 

// Get the Mix_Chunk currently associated with a mixer channel
// Returns NULL if it's an invalid channel, or there's no chunk associated.
func GetChunk(channel int) (* C.Mix_Chunk) {
  return C.Mix_GetChunk(C.int(channel))
}

// Close the mixer, halting all playing audio
func CloseMixer() { 
  C.Mix_CloseAudio()
  MixerOpened = false
}

// Wrappers

// Mixer type 
type Mixer struct {
  nchannels int
} 

// Channel type
type Channel int

// Sound type, that is, a short sample of sound. 
// Many can play at the same time.  
type Sound struct { 
  chunk * C.Mix_Chunk
  channel Channel
}

// Music type. Only one music can play at the time. 
type Music struct {
  music * C.Mix_Music
}

// Loads the music from an .ogg .midi, .mod file 
func LoadMusic(filename string) (* Music) {
  result      := new(Music)
  result.music = LoadMUS(filename)
  if result.music == nil { return nil }
  clean           := func(m * Music) { m.Free() }  
  runtime.SetFinalizer(result, clean)  
  return result  
} 

// Returns true if the music is unusable of has alredy been freed
// Also returns true if the audio mixer has not been opened yet, in which case
// the music cannot be played nor deallocated
func (music * Music) Destroyed() (bool) {
  if !MixerOpened { return true } 
  if music == nil { return true } 
  if music.music == nil { return true}
  return false
}

// Frees the memory associated with this music.
// Only works if the mixer has not been closed yet!
func (music * Music) Free() {
  if music.Destroyed() { return }
  FreeMusic(music.music)
}

// Plays the music indefinitely 
func (music * Music) Play() {  
  if music.Destroyed() { return }
  PlayMusic(music.music, -1)
}

// Pauses the music
func (music * Music) Pause() {
  if music.Destroyed() { return }
  PauseMusic()
}

// Stops the music
func (music * Music) Stop() {
  if music.Destroyed() { return }
  StopMusic()
}

// Resumes the music
func (music * Music) Resume() {
  if music.Destroyed() { return }
  ResumeMusic()
}

// Rewinds the music
func (music * Music) Rewind() {
  if music.Destroyed() { return }
  RewindMusic()
}

// Fades out the music in ms milliseconds
func (music * Music) FadeOut(ms int) {
  if music.Destroyed() { return }
  FadeOutMusic(ms)
} 


// loads a wave file from a .wav or .ogg file 
// It must be stereo if you opened with OpenMixerDefault
func LoadSound(filename string) (* Sound) {
  result          := new(Sound)
  result.chunk     = LoadWAV(filename)
  result.channel   = -1
  if result.chunk == nil { return nil }
  clean           := func(s * Sound) { s.Free() }  
  runtime.SetFinalizer(result, clean)
  
  return result
} 

// Returns true if the sound is unusable of has alredy been freed
// Also returns true if the audio mixer has not been opened yet, in which case
// the music cannot be played nor deallocated
func (sound * Sound) Destroyed() (bool) {
  if !MixerOpened { return true }
  if sound == nil { return true } 
  if sound.chunk == nil { return true}
  return false
}


// Frees the mmemory associated with this wave
// XXX this crashes somehow when called through SetFinalizer()
func (wave * Sound) Free() {
  if wave.Destroyed() { return }
  FreeSound(wave.chunk)
}

// returns the channel the wave will play on, or -1 
// if it will be played on the first available 
func (wave * Sound) Channel() (Channel) {
  return wave.channel
}

// Sets the channel the wave will play on, or -1 
// if it will be played on the first available.
// If the channel is not available, it will use -1 
// by default 
func (wave * Sound) SetChannel(channel Channel) (Channel) {
  max := Channel(GroupAvailable(-1)) 
  if channel >= max || channel < -1 {
    channel = -1
  }
  wave.channel = channel
  return wave.channel
}

// Plays the wave one time on it's channel 
func (wave * Sound) Play() {
  if wave.Destroyed() { return }  
  PlayChannel(int(wave.channel), wave.chunk, 0)   
}






