# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=fungo/sdl

# GOFILES:=constants.$(O).go

CGOFILES:=sdl.go

CGO_CFLAGS:=`pkg-config --cflags sdl`
CGO_LDFLAGS:=`pkg-config --libs sdl` -lSDL_image

CLEANFILES+=sdl

include $(GOROOT)/src/Make.pkg

constants.$(O).go: constants.c
	godefs -g opencv constants.c > constants.$(O).go
	gofmt -w constants.$(O).go
