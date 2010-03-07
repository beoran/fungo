# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=fungo/sdl

# GOFILES:=constants.$(O).go
 
CGOFILES:=audio.go error.go events.go rwops.go sdl.go cdrom.go cpuinfo.go\
image.go joystick.go keyboard.go

# CGO_CFLAGS:=`pkg-config --cflags sdl` 
CGO_CFLAGS:=`sdl-config --cflags`
# CGO_LDFLAGS:=`pkg-config --libs sdl` `sdl-config --libs`-lSDL_image 
CGO_LDFLAGS:=`sdl-config --libs` -lSDL_image

CLEANFILES+=sdl

include $(GOROOT)/src/Make.pkg

constants.$(O).go: constants.c
	godefs -g opencv constants.c > constants.$(O).go
	gofmt -w constants.$(O).go
