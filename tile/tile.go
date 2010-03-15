// The tile package contains everything you need for that classic game 
// staple: a tile map.

package tile

import "fungo/sdl"

// One frame of animation of a tile or a sprite.
type Frame (*sdl.Surface)  

// A TileInfo flag contains info about the tile.
type TileInfo int32

// TileInfo constants
const NORMAL = TileInfo(0)
const (  
  SOLID  = TileInfo(1 << iota)
  WATER
)


// A tile is the elementary part of a tile map.
type Tile struct {
  // The frames of the Tile
  frames [] Frame
  // The current active frame
  active    Frame
  // The index of the active frame
  index     int
  // The capacity of this tile 
  capacity  int
  // How many are in use
  size int
  // The flags of the 
  info      TileInfo
}

func NewTile(info TileInfo, capacity int) (* Tile) {
  tile := &Tile{}
  tile.capacity = capacity
  tile.frames = make([]Frame, tile.capacity)
  tile.active = nil
  tile.index  = 0
  tile.size   = 0
  tile.info   = info
  return tile
}

// How many frames are in this tile
func (t * Tile) Size() (int) {
  return t.size
}

// How many frames can be added to this tile
func (t * Tile) Capacity() (int) {
  return t.capacity
}


func (t * Tile) Add(frame Frame) (bool) {
  last := t.Size()
  if last >= t.Capacity() {
    println("Could not add frame", t.size, t.capacity)
    return false;
  }  
  t.frames[last] = frame 
  t.index 	 = last
  t.active       = t.frames[last]
  t.size++  
  return true
}


func (t * Tile) Solid() (bool) {
  return (t.info & SOLID) != 0
}

func (t * Tile) Water() (bool) {
  return (t.info & WATER) != 0
}

func (t * Tile) Update() (int) {
  t.index++
  if t.index >= t.Size() {
    t.index = 0
  }  
  t.active = t.frames[t.index]
  return t.index
}

func (t * Tile) Draw(screen * sdl.Surface, x, y int) {
  if t == nil { return } // don't blit if the tile is nil.
  active := t.active
  if active == nil { return } // don't blit if no frame there yet.
  screen.Blit(active, x, y)
}


// A tileset is a set of tiles that is used within one or more maps.
// All tiles must be of the same size.
type TileSet struct {
  tiles map[int] *Tile
  // last index of added tile
  last int
  // How wide and high tiles in the tileset are. 
  // They must all be of the same size, although not neccesarily square.
  w, h int
}

func NewTileSet(wide, high int) (* TileSet){
  ts := &TileSet{}
  ts.tiles = make(map[int] * Tile)
  ts.last  = 0
  ts.w     = wide
  ts.h     = high
  return ts
}

func (ts * TileSet) Add(tile *Tile, i int) {
  if i < 0 { i = ts.last ; ts.last++ } 
  ts.tiles[i] = tile
}

type Layer struct {
  // width and height in number of tiles
  w, h int
  tiles [][]*Tile
  * TileSet
  // height and width of the tiles used
  tilewide, tilehigh int
  // Real height and withd in pixels
  realhigh, realwide int
}

func NewLayer(w, h, tw, th int) (* Layer) {
  layer        := &Layer{}
  layer.w 	= w
  layer.h 	= h
  layer.TileSet = NewTileSet(tw, th)
  layer.tiles 	= make([][]*Tile, layer.h)
  for y := 0 ; y < h ; y++ {
    layer.tiles[y] = make([]*Tile, layer.w)
  }
  // These fiels may seem unnnecessary, but they cache 
  // values that otherwise would need to be recalculated every time in draw()
  layer.tilewide= layer.TileSet.w
  layer.tilehigh= layer.TileSet.h
  layer.realwide= layer.tilewide * layer.w
  layer.realhigh= layer.tilehigh * layer.h
  return layer
}

// returns true if the tile with tile coordinates (tx, ty) is outside the tile map
func (l * Layer) Outside(tx, ty int) (bool) { 
  if tx < 0 || ty < 0 { return true } 
  if tx >= l.w || ty >= l.h { return true } 
  return false
} 

func (l * Layer) Set(x, y int, t * Tile) {
  if l.Outside(x, y) { return } 
  l.tiles[y][x] = t
}

// Gets the tile at the given tile indexes. Returns nil if there was no tile 
// there, or if the tile coordinaes were out of bounds.
// This makes drawing easier, in a  sense
func (l * Layer) Get(x, y int) (* Tile) {
  if l.Outside(x, y) { return nil } 
  return l.tiles[y][x]
}

func (l * Layer) Draw(screen * sdl.Surface, x, y int) {   
    txstart    := ( x / l.tilewide )
    tystart    := ( y / l.tilehigh )
    xtilestop  := (screen.W() / l.tilewide) + 1
    ytilestop  := (screen.H() / l.tilehigh) + 1
    txstop     := xtilestop + txstart
    tystop     := ytilestop + tystart
    drawx      := 0
    drawy      := 0
    // don't do anything of your draw indexes are 
    // outside the map on one side   
    if txstart >= l.w || tystart >= l.h { return }     
    // Clip the stop and start of tiles to the map tile coordinates
    if (txstart < 0)   { txstart = 0   }   
    if (tystart < 0)   { tystart = 0   }   
    if (txstop  > l.w) { txstop  = l.w }
    if (tystop  > l.h) { tystop  = l.h }
    // We start drawing here
    drawy       = -y + ( (tystart-1) * l.tilehigh )
    // start iterating
    tydex    :=  tystart
    for tydex < tystop {
      drawy  += l.tilehigh;
      drawx   = -x + ( (txstart-1) * l.tilewide )
      txdex  := txstart
      for txdex < txstop {
        drawx   += l.tilewide        
	// get the tile at this tile index
        aidtile := l.Get(txdex, tydex) 
	aidtile.Draw(screen, drawx, drawy) 
        txdex += 1
      }
      tydex  += 1
    }  
}













