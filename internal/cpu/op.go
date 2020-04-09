package cpu

import (
	"fmt"
	"strings"
)

// Op returns a string representing the assembler for the current opcode and the
// number of bytes the op takes.
//
// The opcode is determined by reading the bytes indicated by the given program
// counter from the given rom bytes.
func Op(rb []byte, pc int64) (string, int64) {
	// Get the opcode.
	opc := rb[pc]

	// Build the output.
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%04x ", pc))

	var (
		asm string
		opb = int64(1)
	)
	d, ok := disassemblers[opc]
	if ok {
		asm, opb = d(rb, pc)
		sb.WriteString(asm)
	} else {
		sb.WriteString(fmt.Sprintf("OPCODE: %x UNKNOWN", opc))
	}

	return sb.String(), opb
}
