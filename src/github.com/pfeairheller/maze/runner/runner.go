package main
import (
	"fmt"
	"os"
	"github.com/pfeairheller/maze"
	"log"
)

var currentMaze *maze.Maze

func main() {

	printCurrentMaze()

	var selection = -1
	var err error = nil
	for ; selection != 5 || err != nil; _, err = fmt.Scanln(&selection) {
		switch selection {
		case 1:
			printCurrentMaze()
		case 2:
			readMazeFromCLI()
		case 3:
			readMazeFromFile()
		case 4:
			findPathToExit()
		}
		printMenu()
	}

}

func printMenu() {
	fmt.Println("Please select from one of the following options:")
	fmt.Println("1. Print Current Maze")
	fmt.Println("2. Read New Maze from Command Like")
	fmt.Println("3. Read New Maze from File")
	fmt.Println("4. Find Path to Exit")
	fmt.Println("5. Exit")
	fmt.Print("> ")
}

func printCurrentMaze() {
	if currentMaze != nil {
		fmt.Println(currentMaze.String())
	} else {
		fmt.Println("No maze is currently selected.")
	}

}

func readMazeFromCLI() {
	var err error
	currentMaze, err = maze.NewMaze(os.Stdin)
	if err != nil {
		log.Fatalln("Unable to read maze, exiting", err)
	}
}

func readMazeFromFile() {
	fmt.Print("Please enter file name: ")

	var fileName string
	fmt.Scanln(&fileName)

	reader, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	currentMaze, err = maze.NewMaze(reader)
	if err != nil {
		log.Fatalln("Unable to read maze, exiting", err)
	}

}

func findPathToExit() {
	if currentMaze == nil {
		fmt.Println("You must create a current maze before finding exits")
		return
	}

	path, err := currentMaze.GetPathToExit(0, 1, 4, 2)
	if err != nil {
		log.Fatal(fmt.Sprint("Error finding exists ", err))
	}

	currentMaze.PrintPath(path)

}