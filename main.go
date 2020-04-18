package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MEMORY_LENGTH = 0x100
const IP = MEMORY_LENGTH - 1
const EQ = MEMORY_LENGTH - 2
const LT = MEMORY_LENGTH - 3
const GT = MEMORY_LENGTH - 4

func boolToInt(b bool) uint8 {
	if b {
		return 1
	} else {
		return 0
	}
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	memory := [MEMORY_LENGTH]uint8{}

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hexes := strings.Split(scanner.Text(), " ")
		for j := range hexes {
			fmt.Sscanf(hexes[j], "%x", &memory[i])
			i++
		}
	}

	for true {
		ip := memory[IP]
		switch memory[ip] {
		case 0:
			return
		case 1:
			memory[IP] += 3
			memory[memory[ip + 1]] = memory[ip + 2]
		case 2:
			memory[IP] += 3
			memory[memory[ip + 1]] = memory[memory[ip + 2]]
		case 3:
			memory[IP] += 3
			a, b := memory[memory[ip + 1]], memory[memory[ip + 2]]
			memory[LT] = boolToInt(a < b)
			memory[GT] = boolToInt(a > b)
			memory[EQ] = boolToInt(a == b)
		case 4:
			memory[IP] += 2
			var b []byte = make([]byte, 1)
			b[0] = memory[memory[ip + 1]]
			os.Stdout.Write(b)
		case 5:
			memory[IP] += 2
			var b []byte = make([]byte, 1)
			os.Stdin.Read(b)
			memory[memory[ip + 1]] = b[0]
		case 6:
			memory[IP] += 3
			memory[memory[ip + 1]] = memory[memory[ip + 1]] + memory[memory[ip + 2]]
		case 7:
			memory[IP] += 3
			memory[memory[ip + 1]] = memory[memory[ip + 1]] * memory[memory[ip + 2]]
		default:
			memory[IP] += 1
		}
	}
}
