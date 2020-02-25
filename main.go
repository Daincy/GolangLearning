package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("欢迎使用学生信息管理系统")
	fmt.Println("----------------------------")
	fmt.Println("1、添加学员")
	fmt.Println("2、修改学员信息")
	fmt.Println("3、展示学员信息")
	fmt.Println("4、退出系统")
	fmt.Println("请输入序号继续：")
}

func getInput()*student{
	var (
		id int
		name string
		class string
	)
	fmt.Println("前要求输入学生信息")
	fmt.Print("请输入学生学号：")
	fmt.Scanf("%d\n",&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanf("%s\n",&name)
	fmt.Print("请输入学生班级：")
	fmt.Scanf("%s\n",&class)
	stu := newStudent(id, name, class)//调用构造函数
	return stu
}

func main() {

	sm := newStudentMgr()
	for {
		showMenu()
		var input int
		fmt.Scanf("%d\n",&input)
		switch input{
			case 1:
				//添加学员
				stu := getInput()
				sm.addStudent(stu)
			case 2:
				//修改学员信息
				stu := getInput()
				sm.modifyStudent(stu)
			case 3:
				//展示所有学员信息
				sm.showStudent()
			case 4:
				//退出
				os.Exit(0)
		}
	}
}	
 

