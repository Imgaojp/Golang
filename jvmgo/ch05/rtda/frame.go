package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPc       int
}

//func NewFrame(maxLocals, maxStack uint) *Frame {
//	return &Frame{
//		localVars:   newLocalVars(maxLocals),
//		operandStack:newOperandStack(maxStack),
//	}
//}
func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:      thread,
		localVars:   newLocalVars(maxLocals),
		operandStack:newOperandStack(maxStack),
	}
}
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) NextPC() int {
	return self.nextPc
}
func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPc = nextPC
}
