package main
import ("fmt")


func fakt(x int) ( p [][]int) {



for t:=2; x!=1; t+=0{

if x%t == 0{
l := 0

for i:=0; x%t==0; l++{
i++
x /= t}

pp := []int{t,l}

p = append(p,pp)
}else{
t++}

}
return 
}

func main() {

var z int = 18747990

fmt.Println(z)

p1 := fakt(z)
fmt.Println(p1)

str := ""
 for i:=0; i<len(p1); i++{
 if i != 0 && i != len(p1) {
 str += " + "}
 str += fmt.Sprint(p1[i][0]) + " ^ " + fmt.Sprint(p1[i][1])}
 
 fmt.Println(str)
}