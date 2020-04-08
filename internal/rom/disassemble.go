package rom

import (
	"fmt"

	"github.com/danmrichards/disassemble8080/internal/cpu"
)

// Disassemble iterates over the given ROM bytes and prints out the resulting assembly code.
func Disassemble(rb []byte) {
	var pc int64
	for pc < int64(len(rb)) {
		asm, opb := cpu.Op(rb, pc)

		fmt.Println(asm)

		pc += opb
	}
}
