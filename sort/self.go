package main
import "fmt"

func main(){
	//a := []int{3,1,2,5,9,6,8}
	//a = bubble(a)
	a , b := 3 , 5

	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a,b)
	//p(a)
}

func bubble(a []int) []int{
	for i := range a{
		t := false
		for j := i ; j < len(a) - 1; j++{
			if a[j] > a[j + 1]{
				a[j] , a[j + 1] = a[j + 1] , a[j]
				t = true
			}
		}
		if t{
			break
		}
	}
	return a
}

func select(a []int) []int{
	
}

func p(a []int) {
	for _ , i := range a{
		fmt.Println(i)
	}
}
