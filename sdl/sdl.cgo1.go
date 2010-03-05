// Created by cgo - DO NOT EDIT
//line sdl.go:1
//
// Go Language wrappers around SDL
//
package sdl

// #include <SDL.h>

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
func cstrNew(size int) *_C_char {
	return (*_C_char)(unsafe.Pointer(_C_malloc(_C_size_t(size))))
}

// free is a method on C char * strings to method to free the associated memory
func (self *_C_char) free()	{ _C_free(unsafe.Pointer(self)) }

// cstring converts a string to a C string. This allocates memory,
// so don't forget to add a "defer s.free()"
func cstr(self string) *_C_char	{ return _C_CString(self) }

type mystring string

// Helper to convert strings to C strings
func (self mystring) cstr() *_C_char	{ return _C_CString(string(self)) }


// The available application states
const (
	APPMOUSEFOCUS	= SDL_APPMOUSEFOCUS	//0x01		// The app has mouse coverage */
	APPINPUTFOCUS	= SDL_APPINPUTFOCUS	//0x02		// The app has input focus */
	APPACTIVE	= SDL_APPACTIVE		//0x04		// The application is active */
)

//
// GetAppState returns the current state of the application, which is a
// bitwise combination of APPMOUSEFOCUS, APPINPUTFOCUS, and
// APPACTIVE.  If APPACTIVE is set, then the user is able to
// see your application, otherwise it has been iconified or disabled.
///
func GetAppState() _C_Uint8	{ return _C_Uint8(_C_SDL_GetAppState()) }


// The calculated values in this structure are calculated by SDL_OpenAudio()
/*
type AudioSpec struct {
	int freq;		// DSP frequency -- samples per second
	Uint16 format;		// Audio data format
	Uint8  channels;	// Number of channels: 1 mono, 2 stereo
	Uint8  silence;		// Audio buffer silence value (calculated)
	Uint16 samples;		// Audio buffer size in samples (power of 2)
	Uint16 padding;		// Necessary for some compile environments
	C.Uint32 size;		// Audio buffer size in bytes (calculated)
	// This function is called when the audio device needs more data.
	   'stream' is a pointer to the audio data buffer
	   'len' is the length of that buffer in bytes.
	   Once the callback returns, the buffer will no longer be valid.
	   Stereo samples are stored in a LRLRLR ordering.
	void (SDLCALL *callback)(void *userdata, Uint8 *stream, int len);
	void  *userdata;
} SDL_AudioSpec;
*/
// Audio format flags (defaults to LSB byte order)
const (
	FAUDIO_U8	= AUDIO_U8
	FAUDIO_U16LSB	= AUDIO_U16LSB
	FAUDIO_U16MSB	= AUDIO_U16LSB
	// FAUDIO_U16SYS  = C.AUDIO_U16SYS
	FAUDIO_S16LSB	= AUDIO_S16LSB
	FAUDIO_S16MSB	= AUDIO_S16LSB
	// FAUDIO_S16SYS  = C.AUDIO_S16SYS
	// FAUDIO_S16  	  = C.AUDIO_S16
	// FAUDIO_U16  	  = C.AUDIO_U16
	BYTEORDER	= SDL_BYTEORDER
	LIL_ENDIAN	= SDL_LIL_ENDIAN
)


// This function fills the given character buffer with the name of the
// current audio driver, and returns a pointer to it if the audio driver has
// been initialized.  It returns "" if no driver has been initialized.
func AudioDriverName() string {
	maxlen := 255
	namebuf := cstrNew(maxlen)
	defer namebuf.free()
	res := _C_SDL_AudioDriverName(namebuf, _C_int(maxlen))
	if res == nil {
		return ""
	}
	return _C_GoString(namebuf)
}

func AudioInit(drivername string) int {
	driver_name := cstr(drivername)
	defer driver_name.free()
	res := _C_SDL_AudioInit(driver_name)
	return int(res)
}

func AudioQuit()	{ _C_SDL_AudioQuit() }

// This function opens the audio device with the desired parameters, and
// returns 0 if successful, placing the actual hardware parameters in the
// structure pointed to by 'obtained'.  If 'obtained' is NULL, the audio
// data passed to the callback function will be guaranteed to be in the
// requested format, and will be automatically converted to the hardware
// audio format if necessary.  This function returns -1 if it failed
// to open the audio device, or couldn't set up the audio thread.
//
// When filling in the desired audio spec structure,
//  'desired->freq' should be the desired audio frequency in samples-per-second.
//  'desired->format' should be the desired audio format.
//  'desired->samples' is the desired size of the audio buffer, in samples.
//     This number should be a power of two, and may be adjusted by the audio
//     driver to a value more suitable for the hardware.  Good values seem to
//     range between 512 and 8096 inclusive, depending on the application and
//     CPU speed.  Smaller values yield faster response time, but can lead
//     to underflow if the application is doing heavy processing and cannot
//     fill the audio buffer in time.  A stereo sample consists of both right
//     and left channels in LR ordering.
//     Note that the number of samples is directly related to time by the
//     following formula:  ms = (samples*1000)/freq
//  'desired->size' is the size in bytes of the audio buffer, and is
//     calculated by SDL_OpenAudio().
//  'desired->silence' is the value used to set the buffer to silence,
//     and is calculated by SDL_OpenAudio().
//  'desired->callback' should be set to a function that will be called
//     when the audio device is ready for more data.  It is passed a pointer
//     to the audio buffer, and the length in bytes of the audio buffer.
//     This function usually runs in a separate thread, and so you should
//     protect data structures that it accesses by calling SDL_LockAudio()
//     and SDL_UnlockAudio() in your code.
//  'desired->userdata' is passed as the first parameter to your callback
//     function.
//
// The audio device starts out playing silence when it's opened, and should
// be enabled for playing by calling SDL_PauseAudio(0) when you are ready
// for your audio callback function to be called.  Since the audio driver
// may modify the requested size of the audio buffer, you should allocate
// any local mixing buffers after you open the audio device.
///
func OpenAudio(desired, obtained *_C_SDL_AudioSpec) int {
	res := _C_SDL_OpenAudio(desired, obtained)
	return int(res)
}


//
// Get the current audio state:
//

type AudioStatus int


const (
	SDL_AUDIO_STOPPED	= AudioStatus(iota)
	SDL_AUDIO_PLAYING
	SDL_AUDIO_PAUSED
)

func GetAudioStatus() AudioStatus	{ return AudioStatus(_C_SDL_GetAudioStatus()) }

//
// This function pauses and unpauses the audio callback processing.
// It should be called with a parameter of 0 after opening the audio
// device to start playing sound.  This is so you can safely initialize
// data for your callback function after opening the audio device.
// Silence will be written to the audio device during the pause.
///
func PauseAudio(pause_on bool) {
	pause := 0
	if pause_on {
		pause = 1
	}
	_C_SDL_PauseAudio(_C_int(pause))
}

// This function loads a WAVE from the data source, automatically freeing
// that source if 'freesrc' is non-zero.  For example, to load a WAVE file,
// you could do:
//	SDL_LoadWAV_RW(SDL_RWFromFile("sample.wav", "rb"), 1, ...);
//
// If this function succeeds, it returns the given SDL_AudioSpec,
// filled with the audio data format of the wave data, and sets
// 'audio_buf' to a malloc()'d buffer containing the audio data,
// and sets 'audio_len' to the length of that audio buffer, in bytes.
// You need to free the audio buffer with SDL_FreeWAV() when you are
// done with it.
//
// This function returns NULL and sets the SDL error message if the
// wave file cannot be opened, uses an unknown data format, or is
// corrupt.  Currently raw and MS-ADPCM WAVE files are supported.
func LoadWAV_RW(src *_C_SDL_RWops, freesrc bool, spec *_C_SDL_AudioSpec, audio_buf **_C_Uint8, audio_len *_C_Uint32) *_C_SDL_AudioSpec {
	free_src := _C_int(0)
	if freesrc {
		free_src = _C_int(1)
	}
	return _C_SDL_LoadWAV_RW(src, free_src, spec, audio_buf, audio_len)
}

// Compatibility convenience function -- loads a WAV from a file */
func LoadWav() {
	// SDL_LoadWAV(file, spec, audio_buf, audio_len) \
	// SDL_LoadWAV_RW(SDL_RWFromFile(file, "rb"),1, spec,audio_buf,audio_len)
}

//
// This function frees data previously allocated with SDL_LoadWAV_RW()
func FreeWav(audio_buf *_C_Uint8)	{ _C_SDL_FreeWAV(audio_buf) }


// This function takes a source format and rate and a destination format
// and rate, and initializes the 'cvt' structure with information needed
// by SDL_ConvertAudio() to convert a buffer of audio data from one format
// to the other.
// This function returns 0, or -1 if there was an error.
func BuildAudioCVT(cvt *_C_SDL_AudioCVT, src_format _C_Uint16, src_channels _C_Uint8, src_rate _C_int, dst_format _C_Uint16, dst_channels _C_Uint8, dst_rate _C_int) _C_int {
	return _C_SDL_BuildAudioCVT(cvt, src_format, src_channels, src_rate,
		dst_format, dst_channels, dst_rate)

}

// Once you have initialized the 'cvt' structure using SDL_BuildAudioCVT(),
// created an audio buffer cvt->buf, and filled it with cvt->len bytes of
// audio data in the source format, this function will convert it in-place
// to the desired format.
// The data conversion may expand the size of the audio data, so the buffer
// cvt->buf should be allocated after the cvt structure is initialized by
// SDL_BuildAudioCVT(), and should be cvt->len*cvt->len_mult bytes long.
///
func ConvertAudio(cvt *_C_SDL_AudioCVT) _C_int	{ return _C_SDL_ConvertAudio(cvt) }

// This takes two audio buffers of the playing audio format and mixes
// them, performing addition, volume adjustment, and overflow clipping.
// The volume ranges from 0 - 128, and should be set to SDL_MIX_MAXVOLUME
// for full audio volume.  Note this does not change hardware volume.
// This is provided for convenience -- you can mix your own audio data.
const MIX_MAXVOLUME = SDL_MIX_MAXVOLUME

func MixAudio(dst, src *_C_Uint8, length _C_Uint32, volume _C_int) {
	_C_SDL_MixAudio(dst, src, length, volume)
}

//
// The lock manipulated by these functions protects the callback function.
// During a LockAudio/UnlockAudio pair, you can be guaranteed that the
// callback function is not running.  Do not call these from the callback
// function or you will cause deadlock.
//
func LockAudio()	{ _C_SDL_LockAudio() }

func UnlockAudio()	{ _C_SDL_UnlockAudio() }


// This function shuts down audio processing and closes the audio device.
func CloseAudio()	{ _C_SDL_CloseAudio() }
