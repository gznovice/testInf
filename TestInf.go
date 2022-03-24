package main

import(
	"fmt"
	"math/rand"
	"time"
)

type animal interface{
	say()
}


type people struct{
	name string
}

func (person people) say(){
	fmt.Println(person.name, " says hello")
}

type cat struct{	
}

func (acat cat) say() {
	fmt.Println("cat says meow")
}

type dog struct{
}

func (adog dog) say() {
	fmt.Println("cat says wow")
}

func main(){
	rand.Seed(time.Now().Unix())
	createtures := make([]animal, 3)
	createtures[0] = people{"tom"}
	createtures[1] = cat{}
	createtures[2] = dog{}
	
	for cnt:=0; cnt < 10; cnt++ {
		i := rand.Intn(len(createtures))
		createtures[i].say()
	}
}