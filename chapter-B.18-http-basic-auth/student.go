package main

import "log"

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			log.Println("EACH ==> ", each)
			return each
		}
	}
	log.Println("NIL ==> ", nil)
	return nil
}

func init() {
	students = append(students, &Student{Id: "s001", Name: "bourne", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "ethan", Grade: 2})
	students = append(students, &Student{Id: "s003", Name: "wick", Grade: 3})
	// for k, v := range students {
	// 	// fmt.Println("k:", k, "v:", v)
	// 	fmt.Println("key[%s] value[%s]\n", k, v)
	// }

	// for k := range students {
	// 	fmt.Printf("key[%s] value[%s]\n", k, students[k])
	// }
}
