package main
import "fmt"

type contactInfo struct{
	email string
	zipCode string
}


type person struct{
	fname string
	lname string
	contact contactInfo
}

func main(){
	jim:= person{
		fname: "Ibrahim",
		lname: "kamal",
		contact: contactInfo{
			email: "i@gmail.com",
			zipCode: "1234",
		},
	}
	jimPointer := &jim
	jimPointer.updateName("Ibrahim")
	jim.print()
}

func(pointerToPerson *person) updateName(newFirstNAme string){
	(*pointerToPerson).fname=newFirstNAme
}
func(p person)print(){
	fmt.Print("%v",p)
}