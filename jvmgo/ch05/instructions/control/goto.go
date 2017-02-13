package control

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
type GOTO_W struct {
	offset int
}

func (self *GOTO_W)FetchOperands(reader *base.BytecodeReader)  {
	self.offset=int(reader.ReadInt32())
}
func (self *GOTO_W)Execute(frame *rtda.Frame)  {
	base.Branch(frame,self.offset)
}
