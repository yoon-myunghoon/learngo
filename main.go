package main //작성할 패키지의 이름을 써줌

import (
	"fmt"

	"github.com/yoon-myunghoon/learngo/mydict"
) //fommating 관련 패키지, go가 가지고 있음

// func multiply(a int, b int) int {
// 	return a * b
// }

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

// func lenAndUpperNaked(name string) (length int, uppercase string) {
// 	defer fmt.Println("I'm done!")
// 	length = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }

// func superAdd(numbers ...int) int {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }

// func canIDrink(age int) bool {
// 	// if koreanAge := age + 2; koreanAge < 18 {
// 	// 	return false
// 	// } else {
// 	// 	return true
// 	// }

// 	switch age {
// 	case 10:
// 		return false
// 	case 18:
// 		return true
// 	}
// 	return false
// }

// type person struct {
// 	name    string
// 	age     int
// 	favFood []string
// }

// [chapter 1]
// func main() { //main function 이 있어야함, 여기가 시작점
// 	// fmt.Println("Hello World!") //function을 export하고 싶으면 첫번째문자를 대문자로 써주면됨
// 	// something.SayHello()        //sayBye는 외부에서 사용할 수 없음

// 	//const const_name string = "nico" // 상수 정의
// 	// name := "nico" // var name string = "nico" // 변수 정의
// 	// name = "lynn"
// 	// fmt.Println(name)

// 	// fmt.Println(multiply(2, 2))

// 	// totalLenght, upperName := lenAndUpper("nico")
// 	// fmt.Println(totalLenght, upperName)

// 	// totalLenght1, _ := lenAndUpper("nico")
// 	// fmt.Println(totalLenght1)

// 	// repeatMe("nico", "lynn", "dal", "marl", "flynn")

// 	// totalLenghtNaked, upperNameNaked := lenAndUpperNaked("nico")
// 	// fmt.Println(totalLenghtNaked, upperNameNaked)

// 	// total := superAdd(1, 2, 3, 4, 5, 6, 7)
// 	// fmt.Println(total)

// 	// fmt.Println(canIDrink(18))

// 	// a := 2
// 	// b := &a
// 	// a = 5
// 	// fmt.Println(*b)

// 	// *b = 10
// 	// fmt.Println(a)

// 	// names := [5]string{"ddd", "dasda", "sfsf"} // array
// 	// fmt.Println(names)

// 	// names_slice := []string{"ddd", "dasda", "sfsf"} // slice
// 	// fmt.Println(names_slice)

// 	// new_names := append(names_slice, "asdfasdfasdf") // append를 사용하면 값이 추가되어 변경되는 것이 아니라 추가된 새로운 slice가 만들어짐
// 	// fmt.Println(new_names)

// 	// nico := map[string]string{"name": "nico", "age": "12"}
// 	// fmt.Println(nico)
// 	// for key, value := range nico {
// 	// 	fmt.Println(key, value)
// 	// }

// 	// favFood := []string{"kimchi", "ramen"}
// 	// nico := person{name: "nico", age: 18, favFood: favFood} // nico := person{"nico", 18, favFood}
// 	// fmt.Println(nico)

// }

// [chapter 2]
// func main() {
// 	account := accounts.NewAccount("nico")
// 	account.Deposit(10)
// 	// err := account.Withdraw(20)
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// }
// 	// fmt.Println(account.Balance(), account.Owner())
// 	fmt.Println(account)

// }

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}

}
