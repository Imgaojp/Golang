package main

import (
	"jvmgo/ch03/classfile"
	"fmt"
	"jvmgo/ch03/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
	//fmt.Printf("Class Name:%v\n",className)
	//classData,_,err:=cp.ReadClass(className)
	//if err!= nil {
	//	fmt.Printf("Could not find or load main class %s\n",cmd.class)
	//	return
	//}
	//fmt.Printf("Class data:%x\n",classData)
	//fmt.Printf("classpath: %s class: %s args :%v\n", cmd.cpOption, cmd.class, cmd.args)
}
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}
func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("Version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("Constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0X%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("	%s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("	%s\n", m.Name())
	}

}
