package main

import (
	"jvmgo/ch05/classfile"
	"jvmgo/ch05/rtda"
	"fmt"
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/instructions"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(uint(maxLocals), uint(maxStack))
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
}
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	var n = 1
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		fmt.Printf("loop: %d", n)
		reader.Reset(bytecode, pc)
		opcode := reader.ReadInt8()
		inst := instructions.NewInstruction(byte(opcode))
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
		n++
	}
}
