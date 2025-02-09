package main

import ("fmt")

func main() {

	a := 1237
    
    b := 11

  var x = a*b
  
  p := []int{}
  
  for t:= 2; x != 1 ; t+=0 {
  
  
  	if x % t == 0{
  
  		p = append(p, t)
        
        x /= t
        
    }else{
    
    	t++
        
        }
        
  }
  
  fmt.Println(p)

  
}