package main

import (
	"fmt"
	"math"
)

func paddingPKCS7(plainText []byte , blockSize int) []byte{
	padding :=blockSize - int(math.Mod(float64(len(plainText)),float64(blockSize)))
	for i:= 0;i<padding;i++{
		plainText= append(plainText,byte(padding))
	}
	return plainText
}

func main(){
	blockSize := 20
	plainText := "YELLOW SUBMARINE"	
	fmt.Println(paddingPKCS7([]byte(plainText),blockSize))
}
