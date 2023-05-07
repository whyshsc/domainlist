package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    ch := make(chan string)
    go func() {
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            ch <- scanner.Text()
        }
        close(ch)
    }()
    for line := range ch {
       domain:= strings.Split(line, ".")
	   for _,list := range domain[:len(domain)-2]{
		fmt.Println(list)

	   }
    }
}