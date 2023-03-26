// An implementation of the mars rover kata.
package rover

import "fmt"

// Grid represents the dimensions of the plane that the rover can traverse.
type Grid struct {
	x uint
	y uint
}

// NewGrid creates a new grid with the specified dimensions.
func NewGrid(x uint, y uint) Grid {
	return Grid{x, y}
}

// Rover represents the state of the mars rover.
type Rover struct {
	x           uint
	y           uint
	orientation rune
	grid        Grid
}

// NewRover creates a new rover instance with the specified grid.
func NewRover(grid Grid) Rover {
	return Rover{
		x:           0,
		y:           0,
		orientation: 'N',
		grid:        grid,
	}
}

// Execute the given command string. Command strings can consist of the
// characters: L (rotate the rover left), R (rotate the rover right), and
// M (move the rover one space forward).
func (r *Rover) Execute(command string) string {
	for _, c := range command {
		r.processCommand(c)
	}
	return fmt.Sprintf("%d:%d:%s", r.x, r.y, string(r.orientation))
}

func (r *Rover) processCommand(c rune) {
	r.handleRotation(c)
	r.handleMove(c)
}

func (r *Rover) handleRotation(c rune) {
	switch c {
	case 'L':
		r.rotateLeft()
	case 'R':
		r.rotateRight()
	}
}

func (r *Rover) handleMove(c rune) {
	if c == 'M' {
		r.move()
	}
}

func (r *Rover) rotateLeft() {
	switch r.orientation {
	case 'N':
		r.orientation = 'W'
	case 'W':
		r.orientation = 'S'
	case 'S':
		r.orientation = 'E'
	case 'E':
		r.orientation = 'N'
	}
}

func (r *Rover) rotateRight() {
	switch r.orientation {
	case 'N':
		r.orientation = 'E'
	case 'E':
		r.orientation = 'S'
	case 'S':
		r.orientation = 'W'
	case 'W':
		r.orientation = 'N'
	}
}

func (r *Rover) move() {
	switch r.orientation {
	case 'N':
		r.moveNorth()
	case 'S':
		r.moveSouth()
	case 'E':
		r.moveEast()
	case 'W':
		r.moveWest()
	}
}

func (r *Rover) moveNorth() {
	if r.atNorthSideOfGrid() {
		r.northWrapAround()
	} else {
		r.y += 1
	}
}

func (r *Rover) moveSouth() {
	if r.atSouthSideOfGrid() {
		r.southWrapAround()
	} else {
		r.y -= 1
	}
}

func (r *Rover) moveEast() {
	if r.atEastSideOfGrid() {
		r.eastWrapAround()
	} else {
		r.x += 1
	}
}

func (r *Rover) moveWest() {
	if r.atWestSideOfGrid() {
		r.westWrapAround()
	} else {
		r.x -= 1
	}
}

func (r *Rover) atNorthSideOfGrid() bool {
	return r.y == (r.grid.y - 1)
}

func (r *Rover) atSouthSideOfGrid() bool {
	return r.y == 0
}

func (r *Rover) atEastSideOfGrid() bool {
	return r.x == (r.grid.x - 1)
}

func (r *Rover) atWestSideOfGrid() bool {
	return r.x == 0
}

func (r *Rover) eastWrapAround() {
	r.x = 0
}

func (r *Rover) southWrapAround() {
	r.y = r.grid.y - 1
}

func (r *Rover) northWrapAround() {
	r.y = 0
}

func (r *Rover) westWrapAround() {
	r.x = r.grid.x - 1
}
