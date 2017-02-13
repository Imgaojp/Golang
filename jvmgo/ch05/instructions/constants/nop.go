package constants

import (
	"jvmgo/ch05/rtda"
	"jvmgo/ch05/instructions/base"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP)Execute(frame *rtda.Frame)  {
	//Do Nothing
}
