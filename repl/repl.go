package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"Skaihen/Plep/vm"
)

type REPL struct {
	vm             vm.VM
	command_buffer []string
}

func RunREPL(repl *REPL) {
	fmt.Println("Welcome to Gongora v0.0.2!")
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Gongora> ")

		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		str = strings.Trim(str, "\n")

		repl.command_buffer = append(repl.command_buffer, str)

		switch str {
		case "exit":
			os.Exit(0)

		case "clear":
			os.Stdout.Write([]byte("\033[2J\033[1;1H"))

		case "history":
			for _, str := range repl.command_buffer {
				fmt.Println(str)
			}

		case "program":
			fmt.Println("Listing instructions currently in VM's program vector:")
			for _, instr := range repl.vm.Program {
				fmt.Println(instr)
			}
			fmt.Println("End of program listing.")

		case "registers":
			fmt.Println("Listing registers and all contents:")
			var i int32
			for _, reg := range repl.vm.Registers {
				fmt.Printf("Registers[%d]: %d\n", i, reg)
				i++
			}
			fmt.Println("End of register listing.")

		case "gracias":
			fmt.Println("De nada")

		default:
			results, err := Parse_hex(str)
			if err != nil {
				fmt.Println("Unable to decode hex string: ", err)
				continue
			}
			for _, bytes := range results {
				vm.Add_byte(&repl.vm, bytes)
			}
			vm.RunVM_once(&repl.vm)
		}
	}
}

func Parse_hex(i string) ([]uint8, error) {
	split := strings.Split(i, " ")
	var results []uint8
	for _, hex_string := range split {
		bytes, err := strconv.ParseUint(hex_string, 16, 8)
		results = append(results, byte(bytes))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
