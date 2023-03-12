package handles

import (
	"fmt"
	"helloworld/model"
)

func SAM() {
	p := model.NewPerson("SAM")
	p.SetAge(20)
	p.SetSal(20000)
	fmt.Println(p)
	fmt.Println(p.Name, " age = ", p.GetAge(), " sal = ", p.GetSal())
}
