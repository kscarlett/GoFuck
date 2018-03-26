package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file")
		fmt.Println(err)
	} else {
		var src = string(content)

		run(src)
	}
}

func run(input string) {
    tape := [65535]uint8{0}
    var tapePointer uint16

    inputLength := len(input)

    for instructionPointer := 0; instructionPointer < inputLength; instructionPointer++ {
        switch input[instructionPointer] {
        case '>':
            tapePointer += 1
            if uint16(len(tape)) < tapePointer {
                fmt.Println("\nERROR - trying to access tape out of bounds")
            }

        case '<':
            if tapePointer > 0 {
                tapePointer -= 1
            }

        case '+':
            tape[tapePointer] += 1

        case '-':
            tape[tapePointer] -= 1

        case '.':
            fmt.Print(string(tape[tapePointer]))

        case ',':
            b := make([]byte, 1)
			os.Stdin.Read(b)
			tape[tapePointer] = b[0]

        case '[':
			if tape[tapePointer] == 0 {
				for depth := 1; depth > 0; {
					instructionPointer++
					char := input[instructionPointer]
					if char == '[' {
						depth++
					} else if char == ']' {
						depth--
					}
				}
			}

		case ']':
			for depth := 1; depth > 0; {
				instructionPointer--
				char := input[instructionPointer]
				if char == '[' {
					depth--
				} else if char == ']' {
					depth++
				}
			}
			instructionPointer--
		}
    }

    fmt.Print("\n\n")
}
