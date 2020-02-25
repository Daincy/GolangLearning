package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

//newStudent是一个student类型的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}
//学员管理的类型
type studentMgr struct{
	allStudent []*student
}
//构造函数newStudentMgr
func newStudentMgr()*studentMgr {
	return &studentMgr{
		allStudent: make([]*student, 0, 100),
	}
}
//添加学生
func (s *studentMgr)addStudent(newStu *student){
	s.allStudent = append(s.allStudent, newStu)
}
//编辑学生信息
func (s *studentMgr)modifyStudent(newStu *student){
	for i, v := range s.allStudent{
		if newStu.id == v.id{ //当学号相同时表示找到了对应的学生
			s.allStudent[i] = newStu //根据切片索引直接把新学生赋值
			return
		}
	}
	//如果没有找到学生
	fmt.Printf("系统中没有找到学号为 %d 对应的学生\n", newStu.id)
}
//展示学生
func (s *studentMgr)showStudent(){
	for _, v := range s.allStudent{
		fmt.Printf("学号：%d 姓名：%s 班级 %s \n\n", v.id, v.name, v.class )
	}
} 
