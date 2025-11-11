package main

import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	
	if n % 2 == 0 {
	    if n < 0 {
	        fmt.Println("negative even")
	    } else {
	        fmt.Println("positive even")
	    } 
	    
	} else if n < 0 {
	    fmt.Println("negative odd")
	} else {
	    fmt.Println("positive odd")
	}
}
