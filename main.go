package main

import (
	"os"

	commandLine "github.com/jpdejavite/messari-THA-solution/internal/application/command_line"
)

func main() {
	commandLine.ReadFromCommanLine(os.Stdin)
}
