# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=fungo/draw

GOFILES:=line.go
# audio.go
# CGOFILES:=

# CGO_CFLAGS:=`pkg-config --cflags sdl`
# CGO_CFLAGS:=`sdl-config --cflags`
# CGO_LDFLAGS:=`pkg-config --libs sdl` `sdl-config --libs`-lSDL_image 
# CGO_LDFLAGS:=`sdl-config --libs` -lSDL_image -lSDL_ttf
CLEANFILES+=draw

include $(GOROOT)/src/Make.pkg

