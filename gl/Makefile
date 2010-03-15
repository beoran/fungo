# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=fungo/gl

# GOFILES:=structs.$(O).go
# audio.go
CGOFILES:=gl.go

# CGO_CFLAGS:=`pkg-config --cflags sdl`
CGO_CFLAGS:=-I/usr/lib/GL
# CGO_LDFLAGS:=`pkg-config --libs sdl` `sdl-config --libs`-lSDL_image 
CGO_LDFLAGS:=-lGL

CLEANFILES+=gl

include $(GOROOT)/src/Make.pkg

#structs.$(O).go: structs.c
#	godefs -g sdl structs.c > structs.$(O).go
#	gofmt -w structs.$(O).go
