package extended

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/instructions/loads"
	"jvmgo/ch06/instructions/math"
	"jvmgo/ch06/instructions/stores"
	"jvmgo/ch06/rtda"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0X15: //iload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0X16: //lload
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0X17: //fload
		inst := &loads.FLOAD{}

		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0X18: //dload
		inst := &loads.DLOAD{}

		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0X19: //aload
		inst := &loads.ALOAD{}

		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0X36: //istore
		inst := &stores.ISTORE{}

		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0X37: //lstore
		inst := &stores.LSTORE{}

		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0X38: //fstore
		inst := &stores.FSTORE{}

		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0X39: //dstore
		inst := &stores.DSTORE{}

		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0X3A: //astore
		inst := &stores.ASTORE{}

		inst.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = inst

	case 0X84: //iinc
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadInt16())
		inst.Const = int32(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0XA9: //ret
		panic("Unsupported opcode: 0XA9!")
	}
}
func (self *WIDE)Execute(frame *rtda.Frame)  {
	self.modifiedInstruction.Execute(frame)
}