package maze

import (
	"io"
	"bufio"
	"fmt"
	"log"
	"errors"
)


type Maze struct {
	Width, Height int
	Spots [][]Spot
}

func NewMaze(reader io.Reader) (*Maze, error) {
	out := &Maze{}

	var width, height int

	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d,%d", &width, &height)

	out.Width = width
	out.Height = height

	for len(out.Spots) < height && scanner.Scan() {
		line := scanner.Text()
		if len(line) != width {
			log.Fatalln("Line too short")
		}

		newLine := make([]Spot, width)
		for idx, char := range line {
			newLine[idx] = NewSpot(byte(char))
			newLine[idx].SetLocation(idx, len(out.Spots))
		}
		out.Spots = append(out.Spots, newLine)
	}

	return out, nil
}

func (m *Maze) String() string {
	out := ""
	for _, line := range m.Spots {
		for _, spot := range line {
			out = out + spot.String()
		}
		out = out + "\n"
	}
	return out
}

func (m *Maze) GetNeighbors(s Spot) []Spot {
	out := []Spot{}
	x, y := s.GetLocation()

	if x != 0 {
		out = append(out, m.Spots[y][x-1])
		if y != 0 {
			out = append(out, m.Spots[y - 1][x - 1])
		}
		if y < m.Height-1 {
			out = append(out, m.Spots[y+1][x-1])
		}
	}

	if x < m.Width-1 {
		out = append(out, m.Spots[y][x+1])
		if y != 0 {
			out = append(out, m.Spots[y-1][x+1])
		}
		if y < m.Height-1 {
			out = append(out, m.Spots[y+1][x+1])
		}
	}

	if y != 0 {
		out = append(out, m.Spots[y-1][x])
	}

	if y < m.Height-1 {
		out = append(out, m.Spots[y + 1][x])
	}

	return out
}

func (m *Maze) Initialize() {
	for _, line := range m.Spots {
		for _, spot := range line {
			spot.Initialize()
		}
	}
}

func (m *Maze) FindExits(x, y int) error {
	if x >= m.Width || y >= m.Height {
		return errors.New("Invalid starting spot")
	}

	m.Initialize()

	workingList := []Spot{}

	start := m.Spots[y][x]
	start.SetDistance(0)

	var current Spot
	workingList = append(workingList, start)

	for len(workingList) > 0 {
		current, workingList = workingList[0], workingList[1:]
		fmt.Println("Visiting", current)
		neighbors := m.GetNeighbors(current)
		for _, neighbor := range neighbors {
			if neighbor.String() != "X" && neighbor.GetDistance() == -1 {
				neighbor.SetDistance(current.GetDistance() + 1)
				neighbor.SetPath(current)
				workingList = append(workingList, neighbor)
			}
		}
	}

	return nil
}

func (m *Maze) GetPathToExit(startX, startY, exitX, exitY int) ([]Spot, error) {
	m.FindExits(startX, startY)

	if startX >= m.Width || startY >= m.Height {
		return nil, errors.New("Invalid starting spot")
	}
	if exitX >= m.Width || exitY >= m.Height {
		return nil, errors.New("Invalid exit spot")
	}

	finalPath := []Spot{}
	current := m.Spots[exitY][exitX]

	for current.GetPath() != nil {
		finalPath = append(finalPath, current)
		current = current.GetPath()
	}

	return finalPath, nil
}

func (m *Maze) PrintPath(path []Spot) {
	mazeStr := []byte(m.String())

	for _, s := range path {
		x, y := s.GetLocation()
		mazeStr[x*y] = '*'
	}

	start := path[len(path)-1]
	x, y := start.GetLocation()
	mazeStr[x*y] = 'S'

	fmt.Println(string(mazeStr))

}

func NewSpot(typ byte) Spot {
	switch typ {
	case 'X':
		return &Wall{}
	case 'E':
		return &Exit{}
	case ' ':
		return &Space{}
	default:
		log.Fatalln("Unknow Spot type")
	}
	return nil
}


type Spot interface {
	String() string
	Initialize()
	SetDistance(int)
	GetDistance()int
	SetLocation(x, y int)
	GetLocation() (int, int)
	SetPath(Spot)
	GetPath()Spot
}

type spot struct {
	X, Y int
	Distance int
	Path Spot
}

func (s *spot) SetLocation(x, y int) {
	s.X, s.Y = x, y
}

func (s *spot) GetLocation() (int, int) {
	return s.X, s.Y
}

func (s *spot) Initialize() {
	s.Distance = -1
	s.Path = nil
}

func (s *spot) SetDistance(d int) {
	s.Distance = d
}

func (s *spot) GetDistance() int {
	return s.Distance
}

func (s *spot) SetPath(p Spot) {
	s.Path = p
}

func (s *spot) GetPath() Spot{
	return s.Path
}

type Wall struct {
	spot
}

func (w *Wall) String() string {
	return "X"
}

type Space struct {
	spot
}

func (w *Space) String() string {
	return " "
}

type Start struct {
	spot
}

func (e *Start) String() string {
	return "S"
}

type Exit struct {
	spot
}

func (e *Exit) String() string {
	return "E"
}

