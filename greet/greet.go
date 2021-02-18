package greet

import "fmt"

type  Salutation string

type  Salu struct {
	Name string
	Company string
}

const (
	PI = 3.14
	Launguage = "Go"
	A = iota
	B = iota
	C = iota
)

const (
	D = iota
	E = iota
	F = iota
)
const (
	G = iota
	H
	I

)

type Printer func(string) ()

func TypeSwitchFunc (x interface{}) {

	switch t := x.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
			fmt.Println(t)
		case Salu:
			fmt.Println("Salu")
		default:
			fmt.Println("unknown")
	}

}
type Salutations []Salu

type  Renamable interface {
	Rename(newName string)
}


func(sal *Salu) Rename (newname string) {
	sal.Name = newname
}

func (sal *Salu) Write(b []byte)(n int, err error) {
	s := string(b)
	sal.Rename(s)
	n = len(s)
	err = nil
	return
}

func(sal Salutations) Greet5(do Printer, isFormal bool, times int) {

	for _, s := range  sal {

		mesg1, mesg2 := CreateMultiMessage3(s.Name, s.Company, "param3")
		fmt.Println(isFormal)

		if prefix := GetPrefix4(s.Name); isFormal {
			do(prefix + mesg1)
		} else {
			do(mesg2)
		}
	}

}

func (sal Salutations) ChannelGreeter(c chan Salu) {

	for _, s:= range sal {
		c <- s
	}
	close(c)
}

func Greet4(sal Salutations, do Printer, isFormal bool, times int) {

	for _, s := range  sal {

		mesg1, mesg2 := CreateMultiMessage3(s.Name, s.Company, "param3")
		fmt.Println(isFormal)

		if prefix := GetPrefix4(s.Name); isFormal {
			do(prefix + mesg1)
		} else {
			do(mesg2)
		}
	}

}

func Greet3(sal Salu, do Printer, isFormal bool, times int) {
	mesg1, mesg2 := CreateMultiMessage3(sal.Name, sal.Company, "param3")
	fmt.Println(isFormal)

	i := 0
	//for i:=0; i < times; i++ {
	for i < times {

		if prefix := GetPrefix2(sal.Name); isFormal {
			do(prefix + mesg1)
		} else {
			do(mesg2)
		}
		i++
	}

}

func GetPrefix(name string) (prefix string) {

	switch name {
		case "param1":
			prefix = "Mr "
			fallthrough
		case "param2", "param4": prefix = "Dr "
		case "param3": prefix = "Mrs"
		default: prefix = "Dude"
	}

	return
}

func GetPrefix4(name string) (prefix string) {

	prefixMap := map[string]string {
		"k1" : "Mr " ,
		"k2" : "Dr " ,
		"k3" : "Dr " ,
		"k4" : "Mrs " ,
	}

	//prefixMap["k3"] = "", false
	delete (prefixMap, "k3")

	if value, exists := prefixMap[name]; exists {
		return value
	}
	return "Dude "
}

func GetPrefix3(name string) (prefix string) {

    var prefixMap map[string]string

    prefixMap = make(map[string]string)

    prefixMap["k1"] = "Mr "
	prefixMap["k2"] = "Dr "
	prefixMap["k3"] = "Dr "
	prefixMap["k4"] = "Mrs "



	return prefixMap[name]
}

func GetPrefix2(name string) (prefix string) {

	switch  {
	case name == "param1":
		prefix = "Mr "
	case name == "param2", name  == "param4", len(name) == 10:
		prefix = "Dr "
	case name == "param3":
		prefix = "Mrs"
	default:
		prefix = "Dude"
	}

	return
}

func Greet2(sal Salu, do func(string)) {
	mesg1, mesg2 := CreateMultiMessage3(sal.Name, sal.Company, "param3")
	do(mesg1)
	do(mesg2)

}

func Greet(sal Salu)  {
	fmt.Println(sal.Name)
	fmt.Println(sal.Company)
	fmt.Println(CreateMessage(sal.Name, sal.Company))

	_, mesg2 := CreateMultiMessage1(sal.Name, sal.Company)
	//fmt.Println(mesg1)
	fmt.Println(mesg2)

	mesg1, mesg2 := CreateMultiMessage2(sal.Name, sal.Company)
	fmt.Println(mesg1)
	fmt.Println(mesg2)

	mesg1, mesg3 := CreateMultiMessage3(sal.Name, sal.Company, "param3")
	fmt.Println(mesg1)
	fmt.Println(mesg3)
}

func CreateMultiMessage1(Name, Company string) (string, string) {
	return Name + " CreateMultiMessage1 " + Company, Company + " CreateMultiMessage1 " + Name

}


func CreateMultiMessage2(Name, Company string) (mesg1 string, mesg2 string) {
	mesg1 = Name + " CreateMultiMessage2 " + Company
	mesg2 = Company + " CreateMultiMessage2 " + Name
	return

}

func CreateMultiMessage3(Name string, Company ...string) (mesg1 string, mesg2 string) {
	fmt.Println(len(Company))
	mesg1 = Name + " CreateMultiMessage3 " + Company[0]
	mesg2 = Company[1] + " CreateMultiMessage3 " + Name
	return

}


func CreateMessage(Name, Company string) string {
	return Name + " " + Company
}

func Print(s string)  {
	fmt.Print(s)
}

func PrintLine(s string)  {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {

	return  func(s string) {
		fmt.Println(s + custom)
	}

}
