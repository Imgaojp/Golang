package constants

import (
	"jvmgo/ch06/rtda"
	"jvmgo/ch06/instructions/base"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP)Execute(frame *rtda.Frame)  {
	//Do Nothing
}
