package internal

import (
	"errors"
	"strconv"
	"strings"
)

type point rune

const (
	steppingStone point = '.'
	wall          point = '#'
	placedWall    point = 'O'

	upwardGuard   point = '^'
	downwardGuard point = 'v'
	leftGuard     point = '<'
	rightGuard    point = '>'
)

type guardDirection int

const (
	up guardDirection = iota
	down
	left
	right
)

type mapPosition struct {
	i int // column
	j int // row
}

type guardMap struct {
	gMapMatrix [][]point

	steppedOn   map[mapPosition]bool
	turnedPoint map[mapPosition]bool

	history struct {
		matrix      [][]point
		positions   mapPosition
		steppedOn   map[mapPosition]bool
		turnedPoint map[mapPosition]bool
		// direction   guardDirection
	}

	gPosition mapPosition
	// walls     []mapPosition
}

func newGuardMap(input string) *guardMap {
	gMap := strings.Split(input, "\n")
	matrix := make([][]point, 0)
	// walls := make([]mapPosition, 0)
	steppedOn := make(map[mapPosition]bool)
	gI, gJ := 0, 0
	for i, row := range gMap {
		matrix = append(matrix, make([]point, len(row)))
		for j, point := range row {
			switch point {
			case '.':
				matrix[i][j] = steppingStone
			case '#':
				matrix[i][j] = wall
				// walls = append(walls, mapPosition{i, j})
			case '^':
				matrix[i][j] = upwardGuard
				steppedOn[mapPosition{i, j}] = true
				gI, gJ = i, j
			case 'v':
				matrix[i][j] = downwardGuard
				steppedOn[mapPosition{i, j}] = true
				gI, gJ = i, j
			case '<':
				matrix[i][j] = leftGuard
				steppedOn[mapPosition{i, j}] = true
				gI, gJ = i, j
			case '>':
				matrix[i][j] = rightGuard
				steppedOn[mapPosition{i, j}] = true
				gI, gJ = i, j
			}
		}
	}
	return &guardMap{
		gMapMatrix:  matrix,
		steppedOn:   steppedOn,
		gPosition:   mapPosition{gI, gJ},
		turnedPoint: make(map[mapPosition]bool),
		// walls:       walls,
	}
}

func (g *guardMap) isWall(i, j int) bool {
	return g.gMapMatrix[i][j] == wall || g.gMapMatrix[i][j] == placedWall
}

func (g *guardMap) directtion() guardDirection {
	switch g.gMapMatrix[g.gPosition.i][g.gPosition.j] {
	case upwardGuard:
		return up
	case downwardGuard:
		return down
	case leftGuard:
		return left
	case rightGuard:
		return right
	}
	return up
}

func (g *guardMap) nextDeadendPosition() (mapPosition, int, bool) {
	hasDeadend := true
	iPos := -1
	jPos := -1
	steps := 0

	switch g.directtion() {
	case up:
		for i := g.gPosition.i - 1; i >= 0; i-- {
			if g.isWall(i, g.gPosition.j) {
				hasDeadend = false
				break
			}
			iPos = i
			jPos = g.gPosition.j
			steps++
		}
	case down:
		for i := g.gPosition.i + 1; i < len(g.gMapMatrix); i++ {
			if g.isWall(i, g.gPosition.j) {
				hasDeadend = false
				break
			}
			iPos = i
			jPos = g.gPosition.j
			steps++
		}
	case left:
		for j := g.gPosition.j - 1; j >= 0; j-- {
			if g.isWall(g.gPosition.i, j) {
				hasDeadend = false
				break
			}
			iPos = g.gPosition.i
			jPos = j
			steps++
		}
	case right:
		for j := g.gPosition.j + 1; j < len(g.gMapMatrix[0]); j++ {
			if g.isWall(g.gPosition.i, j) {
				hasDeadend = false
				break
			}
			iPos = g.gPosition.i
			jPos = j
			steps++
		}
	}

	return mapPosition{
		i: iPos,
		j: jPos,
	}, steps, hasDeadend
}

func (g *guardMap) nextWallPosition() (mapPosition, int, bool) {
	hasNextWall := false
	iPos := -1
	jPos := -1
	steps := 0

	switch g.directtion() {
	case up:
		for i := g.gPosition.i - 1; i >= 0; i-- {
			if g.isWall(i, g.gPosition.j) {
				iPos = i
				jPos = g.gPosition.j
				hasNextWall = true
				break
			}
			steps++
		}
	case down:
		for i := g.gPosition.i + 1; i < len(g.gMapMatrix); i++ {
			if g.isWall(i, g.gPosition.j) {
				iPos = i
				jPos = g.gPosition.j
				hasNextWall = true
				break
			}
			steps++
		}
	case left:
		for j := g.gPosition.j - 1; j >= 0; j-- {
			if g.isWall(g.gPosition.i, j) {
				iPos = g.gPosition.i
				jPos = j
				hasNextWall = true
				break
			}
			steps++
		}
	case right:
		for j := g.gPosition.j + 1; j < len(g.gMapMatrix[0]); j++ {
			if g.isWall(g.gPosition.i, j) {
				iPos = g.gPosition.i
				jPos = j
				hasNextWall = true
				break
			}
			steps++
		}
	}

	return mapPosition{
		i: iPos,
		j: jPos,
	}, steps, hasNextWall
}

func (g *guardMap) turnRight() {
	switch g.directtion() {
	case up:
		g.gMapMatrix[g.gPosition.i][g.gPosition.j] = rightGuard
	case right:
		g.gMapMatrix[g.gPosition.i][g.gPosition.j] = downwardGuard
	case down:
		g.gMapMatrix[g.gPosition.i][g.gPosition.j] = leftGuard
	case left:
		g.gMapMatrix[g.gPosition.i][g.gPosition.j] = upwardGuard
	}
	g.turnedPoint[mapPosition{g.gPosition.i, g.gPosition.j}] = true
}

func (g *guardMap) moveGuard(step int) error {
	if step < 1 {
		return errors.New("step must be greater than 0")
	}
	switch g.directtion() {
	case up:
		newI := g.gPosition.i - step
		for i := g.gPosition.i - 1; i >= newI; i-- {
			if g.isWall(i, g.gPosition.j) {
				return errors.New("cannot move through walls")
			}
			g.steppedOn[mapPosition{i, g.gPosition.j}] = true
		}
		g.gMapMatrix[g.gPosition.i][g.gPosition.j] = steppingStone
		g.gPosition.i = newI
		g.gMapMatrix[g.gPosition.i][g.gPosition.j] = upwardGuard
	case down:
		newI := g.gPosition.i + step
		for i := g.gPosition.i + 1; i <= newI; i++ {
			if g.isWall(i, g.gPosition.j) {
				return errors.New("cannot move through walls")
			}
			g.steppedOn[mapPosition{i, g.gPosition.j}] = true
		}
		g.gMapMatrix[g.gPosition.i][g.gPosition.j] = steppingStone
		g.gPosition.i = newI
		g.gMapMatrix[g.gPosition.i][g.gPosition.j] = downwardGuard
	case left:
		newJ := g.gPosition.j - step
		for j := g.gPosition.j - 1; j >= newJ; j-- {
			if g.isWall(g.gPosition.i, j) {
				return errors.New("cannot move through walls")
			}
			g.steppedOn[mapPosition{g.gPosition.i, j}] = true
		}
		g.gMapMatrix[g.gPosition.i][g.gPosition.j] = steppingStone
		g.gPosition.j = newJ
		g.gMapMatrix[g.gPosition.i][g.gPosition.j] = leftGuard
	case right:
		newJ := g.gPosition.j + step
		for j := g.gPosition.j + 1; j <= newJ; j++ {
			if g.isWall(g.gPosition.i, j) {
				return errors.New("cannot move through walls")
			}
			g.steppedOn[mapPosition{g.gPosition.i, j}] = true
		}
		g.gMapMatrix[g.gPosition.i][g.gPosition.j] = steppingStone
		g.gPosition.j = newJ
		g.gMapMatrix[g.gPosition.i][g.gPosition.j] = rightGuard
	}
	return nil
}

func (g *guardMap) printMap() string {
	result := ""
	for i, row := range g.gMapMatrix {
		for j, point := range row {
			if g.steppedOn[mapPosition{i, j}] {
				if g.gPosition.i == i && g.gPosition.j == j {
					result += string(point)
				} else {
					result += "X"
				}

				continue
			}
			result += string(point)
		}
		result += "\n"
	}
	return result
}

func SolveDay06p1(input string) string {
	debugMode := false
	gMap := newGuardMap(input)
	for {
		_, steps, hasNextWall := gMap.nextWallPosition()
		if !hasNextWall {
			_, steps, hasDeadend := gMap.nextDeadendPosition()
			if !hasDeadend {
				return ""
			}

			gMap.moveGuard(steps)
			if debugMode {
				println("   =======")
				print(gMap.printMap())
				// fmt.Scanln()
			}
			break
		}

		gMap.moveGuard(steps)
		gMap.turnRight()
		if debugMode {
			println("   =======")
			print(gMap.printMap())
			// fmt.Scanln()
		}
	}
	result := 0
	for _, steppedOn := range gMap.steppedOn {
		if steppedOn {
			result++
		}
	}

	return strconv.Itoa(result)
}

func (g *guardMap) saveState() {
	g.history.turnedPoint = make(map[mapPosition]bool)
	g.history.steppedOn = make(map[mapPosition]bool)
	g.history.matrix = make([][]point, len(g.gMapMatrix))

	for i, row := range g.gMapMatrix {
		g.history.matrix[i] = make([]point, len(row))
		copy(g.history.matrix[i], row)
	}

	for k, v := range g.turnedPoint {
		g.history.turnedPoint[k] = v
	}

	for k, v := range g.steppedOn {
		g.history.steppedOn[k] = v
	}

	g.history.positions = g.gPosition
}

func (g *guardMap) restoreState() {
	g.turnedPoint = g.history.turnedPoint
	g.gMapMatrix = g.history.matrix

	g.gPosition = g.history.positions
	g.steppedOn = g.history.steppedOn
}

func (g *guardMap) hasLoopAndCanPlaceWall() bool {
	hasLooping := false
	canPlaceWall := false

	switch g.directtion() {
	case up:
		canPlaceWall = g.gPosition.i-1 >= 0 && g.gMapMatrix[g.gPosition.i-1][g.gPosition.j] == steppingStone
	case right:
		canPlaceWall = g.gPosition.j+1 < len(g.gMapMatrix[0]) && g.gMapMatrix[g.gPosition.i][g.gPosition.j+1] == steppingStone
	case down:
		canPlaceWall = g.gPosition.i+1 < len(g.gMapMatrix) && g.gMapMatrix[g.gPosition.i+1][g.gPosition.j] == steppingStone
	case left:
		canPlaceWall = g.gPosition.j-1 >= 0 && g.gMapMatrix[g.gPosition.i][g.gPosition.j-1] == steppingStone
	}
	if !canPlaceWall {
		return canPlaceWall
	}
	g.saveState()
	// place wall
	switch g.directtion() {
	case up:
		g.gMapMatrix[g.gPosition.i-1][g.gPosition.j] = placedWall
	case right:
		g.gMapMatrix[g.gPosition.i][g.gPosition.j+1] = placedWall
	case down:
		g.gMapMatrix[g.gPosition.i+1][g.gPosition.j] = placedWall
	case left:
		g.gMapMatrix[g.gPosition.i][g.gPosition.j-1] = placedWall
	}

	boxCounter := 0
	for {
		g.turnRight()
		_, _, hasDeadend := g.nextDeadendPosition()
		if hasDeadend {
			break
		}
		_, steps, hasNextWall := g.nextWallPosition()
		if !hasNextWall {
			break
		}
		g.moveGuard(steps)
		// fmt.Printf("======= %v - %v\nstep: %d\nboxCounter: %d\n%v\n", g.gPosition, g.directtion(), steps, boxCounter, g.printMap())
		// fmt.Scanln()

		if steps == 0 {
			boxCounter++
		} else {
			boxCounter = 0
		}

		if boxCounter > 4 {
			hasLooping = true
			break
		}

		if g.turnedPoint[g.gPosition] && steps > 0 {
			hasLooping = true
			break
		}

	}
	g.restoreState()
	return hasLooping
}

func SolveDay06p2(input string) string {
	result := 0
	gMap := newGuardMap(input)
	for {
		_, steps, hasNextWall := gMap.nextWallPosition()

		for i := 0; i < steps; i++ {
			// fmt.Println("checking", gMap.gPosition)
			if gMap.hasLoopAndCanPlaceWall() {
				// fmt.Printf("   Placing wall at %d, %d\n%v\n", gMap.gPosition.i, gMap.gPosition.j, gMap.printMap())
				result++
			}

			gMap.moveGuard(1)
			// fmt.Printf("moved: %v\n", gMap.gPosition)
		}
		if !hasNextWall {
			break
		}
		gMap.turnRight()
	}
	return strconv.Itoa(result)
}
