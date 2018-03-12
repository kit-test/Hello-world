package File

import (

	"fmt"
	"io/ioutil"
)

func check(e error)  {
	if e != nil{
		panic(e)
	}
}
func main()  {

	date, err := ioutil.ReadFile("../abc.txt")
	fmt.Printf("=========")
	check(err)
	ioutil.WriteFile("../abc_copy.txt",date,666)

}