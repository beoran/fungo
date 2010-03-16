# Copyright 2010 Beoran. 
# 
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

all: libs test-fungo test-gui

libs:
	make -C sdl install
	make -C draw install
	make -C gui install
	make -C gl install
	make -C tile install 

test-fungo: test-fungo.go libs
	$(GC) test-fungo.go
	$(LD) -o $@ test-fungo.$(O)
	
test-gui: test-gui.go libs
	$(GC) test-gui.go
	$(LD) -o $@ test-gui.$(O)	

clean:
	make -C sdl clean
	rm -f -r *.8 *.6 *.o */*.8 */*.6 */*.o */_obj test-fungo test-gui
	
