package main
import "fmt"
func main (){
	var arr [3] string
	arr[0]="hello"
	arr[1]="hi"
	arr[2]="meow"
	fmt.Println(arr)
	//another way of declaring
	nums:=[3]int{1,2,3}
	fmt.Println(nums)
	//2D arrays
	var matrix [2][2] int
	fmt.Println("enter  values for the matrix ")
	for i:=0;i<2;i++{
		for j:=0;j<2;j++{
			fmt.Scanln(&matrix[i][j])
		}
	}
	fmt.Println("the result matrix:")
	for i:=0;i<2;i++{
		for j:=0;j<2;j++{
			fmt.Print(matrix[i][j])
		}
		fmt.Printf("\n")
	}

}