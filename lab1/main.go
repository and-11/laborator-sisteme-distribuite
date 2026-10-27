// Primul lucru care se precizeaza este package -ul din care fisierul face parte


package main

// A doua etapa este importarea pachetelor ( nefolosirea unui pachet genereaza o eroare )



import "fmt"
import "bufio"
import "strconv"
import "os"
// A treia etapa este scrierea functiilor si a tipurilor . Aici avem doar o functie main:


func ex1(){

	scanner := bufio.NewScanner(os.Stdin) 
	scanner.Scan()

	number, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	for i := 2; i < number - 1; i++ {
		if number % i == 0 {
			fmt.Println("nu e prim")
			return
		}
	} 

	fmt.Println("e prim")
}

func ex2(){
	scanner := bufio.NewScanner(os.Stdin) 
	scanner.Scan()

	var arr [100] int

	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	// fmt.Println(scanner.Text())
	
	for i := 0; i < n; i++ {
		scanner.Scan()
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}	

		arr[i] = number

	}
	
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++{
			if arr[i] < arr[j] {
				no := arr[j]
				arr[j] = arr[i]
				arr[i] = no
			}
		}
	}
	
	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
	}

}

func main () {

	// ex1()
	ex2()

}