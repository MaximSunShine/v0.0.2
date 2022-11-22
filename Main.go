package main

import (
	"fmt"
	"log"
	"regexp"
)

func convert(s string) string {

	reg, err := regexp.Compile("[Ee][Uu][Rr]|[$]|CDN[$]|AUD|CAD|GBP|[€]")
	if err != nil {
		log.Fatal(err)
	}
	nameMoney := reg.FindString(s)

	reg, err = regexp.Compile("[0-9[.,]+")
	if err != nil {
		log.Fatal(err)
	}
	digits := reg.FindString(s)
	//fmt.Println(digits)

	reg, err = regexp.Compile("[0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	digArr := reg.FindAllString(digits, -1)
	//fmt.Println(digArr)

	rez := ""
	if len(digArr) > 1 && len(digArr[len(digArr)-1]) == 2 {
		rez = "." + digArr[len(digArr)-1]
	} else {
		rez = digArr[len(digArr)-1]
	}
	for i := len(digArr) - 2; i >= 0; i-- {
		rez = digArr[i] + rez
	}
	if nameMoney == "" || rez == "" {
		return ""
	}

	return nameMoney + ":" + rez
}

func main() {
	fmt.Println(convert("$30.65 ($13.74 / Pound)"))

}
