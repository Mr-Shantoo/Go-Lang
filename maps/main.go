package main 
import "fmt"
func main(){
	carinventory:=make(map[string]int) //string is key type, int value type
	carinventory["sedan"]=25
	carinventory["suv"]=15
	carinventory["gtr"]=20
	fmt.Println("car inventory:",carinventory)
	//directly intialization
	players:=map[string]int{
		"messi":1,
		"ronaldo":2,
		"maradona":3,
	}
	fmt.Println(players)
	delete(players,"ronaldo")
	fmt.Println("after deleting:", players)
	// check if a key is present in a map.. 
	players["neymar"]=4
	//go maps not only returns key value bt also returns a boolean.. which indicates if a key exists in the map or not..
	if goat,ok:=players["messi"];ok{
		fmt.Println("messi is number:",goat)
	}
	//these variables can't be used outside of the if condition..
	grades:=make(map[string]string)
	grades["shanto"]="A"
	grades["rahim"]="B"
	grades["karim"]="C"
	if exists,ok:=grades["shanto"];ok{
		fmt.Println("grade of shanto:",exists)
	}
}