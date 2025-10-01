package main 
import "fmt"
func main(){
fueltypes:=[] string{"Electrolic","gasoline","hybrid"}
fmt.Println(fueltypes)
fueltypes = append(fueltypes, "diesel")
fmt.Println(fueltypes)

//creating a slice or array with zero values and size..

clubs:=make([]string,3)
fmt.Printf("before adding values:%s\n",clubs)
clubs[0]="barcelona"
clubs[1]="madrid"
clubs[2]="mancity"
fmt.Println(clubs)
players:=[]string{"messi","ronaldo","maradona","neymar","pedri","lamine"}
GOAT:=players[:3]
fmt.Println(GOAT)
newgeneration:=players[len(GOAT):]
fmt.Println(newgeneration)
//2D slice 
rows,cols:=3,2
grid:=make([][] int,rows)
for i:=0;i<rows;i++{
	grid[i]=make([]int, cols)
	cols++
}
fmt.Println(grid)
count:=1
for i:=0;i<rows;i++{
	for j:=0;j<len(grid[i]);j++{
		grid[i][j]=count
		count++
	}

}
for i:=0;i<rows;i++{
	for j:=0;j<len(grid[i]);j++{
		fmt.Printf("%v ",grid[i][j])
		
	}
	fmt.Println("\n")

}

}