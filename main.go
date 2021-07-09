package main

import "fmt"

type Action struct {
	left, right int // кол-во шагов влево и вправо соответственноы
}

// человек - может шагать влево вправо определенное кол-во шагов
type Human struct {
	action Action
	name   string
	age    int
}

func main() {

	man := Human{
		action: Action{
			left:  10,
			right: 10,
		},
		name: "Alex007",
		age:  29,
	}

	fmt.Println(man.name, man.age, "Влево: ", man.action.left, " Вправо: ", man.action.right)
}
