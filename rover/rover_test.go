package rover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertThatCommandYieldsPositionOnGrid(t *testing.T, command string, expectedPosition string, grid Grid) {
	// Arrange
	rover := NewRover(grid)

	// Act
	actualPosition := rover.Execute(command)

	// Assert
	assert.Equal(t, expectedPosition, actualPosition)
}

var tenByTenGrid = NewGrid(10, 10)

func TestWhenTheRoverIsNotMovedThenItIsAtTheStartingPosition(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "", "0:0:N", tenByTenGrid)
}

func TestWhenTheRoverIsRotatedLeftOneTimeThenItFacesWest(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "L", "0:0:W", tenByTenGrid)
}

func TestWhenTheRoverIsRotatedLeftTwoTimesThenItFacesSouth(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "LL", "0:0:S", tenByTenGrid)
}
func TestWhenTheRoverIsRotatedLeftThreeTimesThenItFacesEast(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "LLL", "0:0:E", tenByTenGrid)
}

func TestWhenTheRoverIsRotatedLeftFourTimesThenItFacesNorth(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "LLLL", "0:0:N", tenByTenGrid)
}

func TestWhenTheRoverIsRotatedRightOneTimeThenItFacesEast(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "R", "0:0:E", tenByTenGrid)
}
func TestWhenTheRoverIsRotatedRightTwoTimesThenItFacesSouth(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "RR", "0:0:S", tenByTenGrid)
}

func TestWhenTheRoverIsRotatedRightThreeTimesThenItFacesWest(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "RRR", "0:0:W", tenByTenGrid)
}

func TestWhenTheRoverIsRotatedRightFourTimesThenItFacesNorth(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "RRRR", "0:0:N", tenByTenGrid)
}

func TestWhenTheRoverMovesNorthThenItsPositionIncrements(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "M", "0:1:N", tenByTenGrid)
}
func TestWhenTheRoverMovesEastThenItsPositionIncrements(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "RM", "1:0:E", tenByTenGrid)
}
func TestWhenTheRoverMovesSouthThenItsPositionDecrements(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "MRRM", "0:0:S", tenByTenGrid)
}

func TestWhenTheRoverMovesWestThenItsPositionIncrements(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "RMRRM", "0:0:W", tenByTenGrid)
}

func TestWhenTheRoverReachesTheNorthSideOfTheGridThenItWrapsAround(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "MMMMMMMMMM", "0:0:N", tenByTenGrid)
}
func TestWhenTheRoverReachesTheSouthSideOfTheGridThenItWrapsAround(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "LLM", "0:9:S", tenByTenGrid)
}

func TestWhenTheRoverReachesTheWestSideOfTheGridThenItWrapsAround(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "LM", "9:0:W", tenByTenGrid)
}

func TestWhenTheRoverReachesTheEastSideOfTheGridThenItWrapsAround(t *testing.T) {
	assertThatCommandYieldsPositionOnGrid(t, "RMMMMMMMMMM", "0:0:E", tenByTenGrid)
}
