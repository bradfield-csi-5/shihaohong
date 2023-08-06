package vm

const (
	Load  = 0x01
	Store = 0x02
	Add   = 0x03
	Sub   = 0x04
	Halt  = 0xff
)

// Stretch goals
const (
	Addi = 0x05
	Subi = 0x06
	Jump = 0x07
	Beqz = 0x08
)

// Given a 256 byte array of "memory", run the stored program
// to completion, modifying the data in place to reflect the result
//
// The memory format is:
//
// 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f ... ff
// __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ ... __
// ^==DATA===============^ ^==INSTRUCTIONS==============^
func compute(memory []byte) {

	registers := [3]byte{8, 0, 0} // PC, R1 and R2

	// Keep looping, like a physical computer's clock
clock:
	for {
		op := memory[registers[0]]
		registers[0]++

		var arg1, arg2 byte
		// decode and execute
		// TODO(shihaohong): error handle for registers that arents 0x01 and 0x02
		switch op {
		case Load:
			arg1 = memory[registers[0]]
			registers[0]++
			arg2 = memory[registers[0]]
			registers[0]++
			registers[arg1] = memory[arg2]
		case Store:
			arg1 = memory[registers[0]]
			registers[0]++
			arg2 = memory[registers[0]]
			registers[0]++

			memory[arg2] = registers[arg1]
		case Add:
			registers[1] = registers[1] + registers[2]
			registers[0] += 2
		case Sub:
			registers[1] = registers[1] - registers[2]
			registers[0] += 2
		case Halt:
			break clock
		}
	}
}
