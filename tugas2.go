package main

import "fmt"

func main()  {
	for i:=1;i<10; i++ {
		fmt.Print("masukan angka : ")
		fmt.Scan(&i)
		if i <=10{
			if  i%2 == 0 {
				fmt.Println("adalah bilangan genap")
			}else{
				fmt.Println("adalah bilangan ganjil")
			}
		}
	}
}