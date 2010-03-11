# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=fungo/sdl

GOFILES:=structs.$(O).go
# audio.go
CGOFILES:=audio.go events.go cdrom.go joystick.go keyboard.go\
ttf.go mixer.go video.go wrap.go sdl.go

# CGO_CFLAGS:=`pkg-config --cflags sdl`
CGO_CFLAGS:=`sdl-config --cflags`
# CGO_LDFLAGS:=`pkg-config --libs sdl` `sdl-config --libs`-lSDL_image 
CGO_LDFLAGS:=`sdl-config --libs` -lSDL_image -lSDL_ttf -lSDL_mixer

CLEANFILES+=sdl

include $(GOROOT)/src/Make.pkg

structs.$(O).go: structs.c
	godefs -g sdl structs.c > structs.$(O).go
	gofmt -w structs.$(O).go
