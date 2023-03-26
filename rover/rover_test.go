package rover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertCommandYieldsPositionOnGrid(t *testing.T, command string, expectedPosition string, grid Grid) {
	// Arrange
	rover := NewRover(grid)

	// Act
	actualPosition := rover.Execute(command)

	// Assert
	assert.Equal(t, expectedPosition, actualPosition)
}

var tenByTenGrid = NewGrid(10, 10)

func TestWhenTheRoverIsNotMovedThenItIsAtTheStartingPosition(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "", "0:0:N", tenByTenGrid)
}

func TestWhenTheRoverIsRotatedLeftOneTimeThenItFacesWest(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "L", "0:0:W", tenByTenGrid)
}

func TestWhenTheRoverIsRotatedLeftTwoTimesThenItFacesSouth(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "LL", "0:0:S", tenByTenGrid)
}
func TestWhenTheRoverIsRotatedLeftThreeTimesThenItFacesEast(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "LLL", "0:0:E", tenByTenGrid)
}

func TestWhenTheRoverIsRotatedLeftFourTimesThenItFacesNorth(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "LLLL", "0:0:N", tenByTenGrid)
}

func TestWhenTheRoverIsRotatedRightOneTimeThenItFacesEast(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "R", "0:0:E", tenByTenGrid)
}
func TestWhenTheRoverIsRotatedRightTwoTimesThenItFacesSouth(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "RR", "0:0:S", tenByTenGrid)
}

func TestWhenTheRoverIsRotatedRightThreeTimesThenItFacesWest(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "RRR", "0:0:W", tenByTenGrid)
}

func TestWhenTheRoverIsRotatedRightFourTimesThenItFacesNorth(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "RRRR", "0:0:N", tenByTenGrid)
}

func TestWhenTheRoverMovesNorthThenItsPositionIncrements(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "M", "0:1:N", tenByTenGrid)
}
func TestWhenTheRoverMovesEastThenItsPositionIncrements(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "RM", "1:0:E", tenByTenGrid)
}
func TestWhenTheRoverMovesSouthThenItsPositionDecrements(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "MRRM", "0:0:S", tenByTenGrid)
}

func TestWhenTheRoverMovesWestThenItsPositionIncrements(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "RMRRM", "0:0:W", tenByTenGrid)
}

func TestWhenTheRoverReachesTheNorthSideOfTheGridThenItWrapsAround(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "MMMMMMMMMM", "0:0:N", tenByTenGrid)
}
func TestWhenTheRoverReachesTheSouthSideOfTheGridThenItWrapsAround(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "LLM", "0:9:S", tenByTenGrid)
}

func TestWhenTheRoverReachesTheWestSideOfTheGridThenItWrapsAround(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "LM", "9:0:W", tenByTenGrid)
}

func TestWhenTheRoverReachesTheEastSideOfTheGridThenItWrapsAround(t *testing.T) {
	assertCommandYieldsPositionOnGrid(t, "RMMMMMMMMMM", "0:0:E", tenByTenGrid)
}
