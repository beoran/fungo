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
  tile.size   = 0
  tile.info   = info
  tile.Rewind()
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

// Adds a frame of animation to this tile 
func (t * Tile) Add(frame Frame) (bool) {
  last := t.Size()
  if last >= t.Capacity() {
    println("Could not add frame", t.size, t.capacity)
    return false;
  }  
  t.frames[last] = frame 
  t.size++
  t.Rewind()  
  return true
}

// Retuns whether or not the tile is solid
func (t * Tile) Solid() (bool) {
  return (t.info & SOLID) != 0
}

// Returns whether or not the tile is a water tile 
func (t * Tile) Water() (bool) {
  return (t.info & WATER) != 0
}

// Rewind resets the tile to it's the first frame, which is frame 0 
func (t * Tile) Rewind() {
  t.index  = 0
  t.active = t.frames[t.index]
}

// Updates the tile, advancing it's animation, if any
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

// Adds a tile to the tileset
func (ts * TileSet) Add(tile *Tile, i int) {
  if i < 0 { i = ts.last ; ts.last++ } 
  ts.tiles[i] = tile
}


// A camera actively defines the current view of the visible 2D game world
type Camera struct {
  // Position of camera top left corner
  X, Y int
  // Field of view. Normally equal to screen size.
  W, H int
}



// A direction is a direction in which a Sprite is moving
type Direction int

const (  	
  NORTH          = Direction(1 << iota)
  EAST
  SOUTH
  WEST
  ALL_DIRECTIONS = NORTH | EAST | SOUTH | WEST
)

func (self Direction) Is(other Direction) (bool) {
  return (self & other) != 0
} 

// Offset describes where something, like a sprite should be drawn 
type Offset struct {
  X, Y int
}

// Moves the offset to x and y
func (o Offset) MoveTo(x, y int) {
  o.X = x
  o.Y = y
}

// Moves the offset by dx and dy
func (o Offset) MoveBy(dx, dy int) {
  o.X += dx
  o.Y += dy
}

// Adds this offset to the othe roffset and retuns a new offset 
// that combines the effects of the individual offsets
func (o1 Offset) AddOffset(o2 Offset) (Offset) {
  return Offset { o1.X + o2.X, o1.Y + o2.Y }  
}



// A Motion is a single set of frames which are displayed sequentially
// according to a next frame index list when the sprite performs 
// an action in a given direction.
// For example, it could contain a set of frames showing the character 
// walking north.
type Motion struct {
  frames 	        []Frame
  active   	      Frame
  size, capacity 	int
  index			      int
  direction		    Direction
  next			      map[int] int
  Offset
  // offset of the motion with regards to the action
}

// Creates a new, empty motion
func NewMotion(capacity int) (* Motion) {
  m         := new(Motion)
  m.frames   = make([]Frame, capacity)
  m.capacity = capacity
  m.size     = 0
  m.Rewind()
  return m
}

// Rewind resets the run to it's the first frame, which is frame 0 
func (f * Motion) Rewind() {
  f.index  = 0
  f.active = f.frames[f.index]
}

// Updates and animates the motion
func (f * Motion) Update() {  
  newindex, ok := f.next[f.index] // Get next index.
  // If not defined, rewind  
  if (!ok) { 
    f.Rewind()  
    return
  }
  f.index  = newindex
  f.active = f.frames[f.index]
}

// How many frames are in this tile
func (f * Motion) Size() (int) {
  return f.size
}

// How many frames can be added to this tile
func (f * Motion) Capacity() (int) {
  return f.capacity
}

// Adds a frame of animation to this tile 
func (f * Motion) Add(frame Frame) (bool) {
  last := f.Size()
  if last >= f.Capacity() {
    println("Could not add frame", f.size, f.capacity)
    return false;
  }  
  f.frames[last] = frame   
  f.size++  
  f.Rewind()
  
  return true
}

// Draws the motion's active bitmat at the given coordinates
func (m * Motion) Draw(screen * sdl.Surface, x, y int) {
  if m == nil { return } // don't blit if the tile is nil.
  active := m.active
  if active == nil { return } // don't blit if no frame there yet.
  screen.Blit(active, x, y)
}


// An action is a set of Runs which the sprite
// can perform in one or several directions.
type Action struct {
  motions map[Direction] * Motion  
  Offset
  // offset of the action with regards to the Fragment
}

// Creates a new action
func NewAction() (* Action) {
  a := new(Action)
  a.motions = make(map[Direction] * Motion) 
  return a 
}

// Adds an existing Motion to this action, in the given direction.
func (a * Action) Add(direction Direction, motion * Motion) {
  // set the motion in all possibly corresponding directories  
  if direction.Is(NORTH) { a.motions[NORTH] = motion }
  if direction.Is(EAST)  { a.motions[EAST]  = motion }
  if direction.Is(SOUTH) { a.motions[SOUTH] = motion }
  if direction.Is(WEST)  { a.motions[WEST]  = motion }
}

// Gets the motion that corresponds with the given direction.
func (a * Action) Get(direction Direction) (* Motion) {
  motion, ok := a.motions[direction]
  if !ok { return nil }
  return motion
}

// Activity is the type of action a sprite is engaging in
type Activity int

const (
  STAND = Activity(iota)
  WALK
  RUN
  ATTACK
  STRUCK
  KNEEL
  SWOON
)


// A Fragment is a part of a sprite. Sprites consist of one or more 
// fragments that move together but can move separately.
// Useful, for example, for a player character who holds a weapon, 
// etc 
type Fragment struct {
  offset  Offset        // Offset of the fragment relative to the sprite 
  hidden  bool
  actions map[Activity] Action
}


// A sprite is a visual, mobile representation of a game object.
type Sprite struct {
  id int                        // Sprite ID  
  fragments map[int] * Fragment // Fragments the sprite consists of 
  Offset
  // Offset of the Sprite with regards to the current Map
  layer, order int
  // Layer the sprite is on, and it's drawing order 
}


// A layer is a single layer of tiles in a Map
type Layer struct {
  // width and height in number of tiles
  w, h int
  tiles [][]*Tile
  * TileSet
  // height and width of the tiles used
  tilewide, tilehigh int
  // Real height and withd in pixels
  realhigh, realwide int
  // the sprites that are present on this layer of the map
  // This is a set of pointers, because a sprite can be
  // present in several layers 
  sprites map[int] Sprite  
}

func NewLayer(w, h, tw, th int) (* Layer) {
  layer        := &Layer{}
  layer.w   = w
  layer.h   = h
  layer.TileSet = NewTileSet(tw, th)
  layer.tiles   = make([][]*Tile, layer.h)
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

// Draws the layer to the creen using a given camera
func (l * Layer) DrawCamera(screen * sdl.Surface, cam * Camera) {   
    txstart    := ( cam.X / l.tilewide )
    tystart    := ( cam.Y / l.tilehigh )
    xtilestop  := ( cam.W / l.tilewide) + 1
    ytilestop  := ( cam.H / l.tilehigh) + 1
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
    drawy       = - cam.Y + ( (tystart-1) * l.tilehigh )
    // start iterating
    tydex    :=  tystart
    for tydex < tystop {
      drawy  += l.tilehigh;
      drawx   = -cam.X + ( (txstart-1) * l.tilewide )
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

// Draws the layer to the screen using x and y as the corner coordinates
func (l * Layer) Draw(screen * sdl.Surface, x, y int) {   
  cam := &Camera{x, y, screen.W(), screen.H()}
  l.DrawCamera(screen, cam)
}


// A Map consists of different Layers
type Map struct {
  // The layers
  layers [] *Layer;
  // Layers that are in the tile map and that can be in the tile map
  size, capacity int;  
  // amount of layers and capacity 
  // Sprites, by id 
}

// Makes a Map with the given amount of layers
func NewMap(size int) (* Map) {
  tm 	     := new(Map)
  tm.size   = size
  tm.layers = make([]*Layer, tm.size)
  return tm
}







