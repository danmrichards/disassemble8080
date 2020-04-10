package rom

import (
	"fmt"

	"github.com/danmrichards/disassemble8080/pkg/dasm"
)

// Disassemble iterates over the given ROM bytes and prints out the resulting assembly code.
func Disassemble(rb []byte) {
	var pc int64
	for pc < int64(len(rb)) {
		asm, opb := dasm.Disassemble(rb, pc)

		fmt.Println(asm)

		pc += opb
	}
}
