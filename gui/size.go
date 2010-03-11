package  gui

// The size struct 
type Size struct {
  X         int
  Y         int
  W         int
  H         int
  Margin    int
  Padding   int
} 


// Outside width, that is, width plus the margins
func (s Size) OutW() (int) {
  return s.W + (s.Margin * 2)
}

// Inside width, that is, width minus the paddings
func (s Size) InW() (int) {
  return s.W - (s.Padding * 2)
}

// Outside height, that is, height plus the margins
func (s Size) OutH() (int) {
  return s.H + (s.Margin * 2)
}

// Inside height, that is, height minus the paddings
func (s Size) InH() (int) {
  return s.H - (s.Padding * 2)
}

// Left hand side of the Sized object 
func (s Size) Left() (int) { 
  return s.X
}
    
// Right hand side of the widget
func (s Size) Right() (int) {
  return s.X + s.W
}
    
// Top side of the widget 
func (s Size) Top() (int) {
  return s.Y
}    
 
// Bottom side of the widget (position plus height)
func (s Size) Bottom() (int) {
  return s.Y + s.H
}

// Middle of the widget on the x axis
func (s Size) Middle() (int) {
  return s.X + (s.W / 2)
}

// Center of the widget on the y axis
func (s Size) Center() (int) {
  return s.Y + (s.H / 2)
}

// Outside left hand side
func (s Size) OutLeft() (int) {
  return s.X - s.Padding
}

// Inside left hand side
func (s Size) InLeft() (int) {
  return s.X + s.Margin
}

// Outside right hand side
func (s Size) OutRight() (int) {
  return s.X + s.W - s.Padding
}

// Inside right hand side
func (s Size) InRight() (int) {
  return s.X + s.W + s.Margin
}

// Outside top side
func (s Size) OutTop() (int) {
  return s.Y - s.Padding
}

// Inside top side
func (s Size) InTop() (int) {
  return s.Y + s.Margin
}

// Outside bottom side
func (s Size) OutBottom() (int) {
  return s.Y + s.H - s.Padding
}

// Inside bottom side
func (s Size) InBottom() (int) {
  return s.Y + s.H + s.Margin
}

