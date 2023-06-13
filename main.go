package main

import (
	"fmt"

	"github.com/alim7007/blockchain/chain"
	// "rsc.io/quote"
)

func main() {
	chain := chain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Frist")
	chain.AddBlock("Third Block after Second")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
	}
}

////////////////////////////////////////////////
////////////////////////////////////////////////
////////////////////////////////////////////////

// fmt.Println(quote.Hello())
/////////////////////////////
/////////////////////////////

// // type person struct {
// // 	name string
// //    }
// //    func main() {
// // 	p := person{"richard"}
// // 	p = rename(p)
// // 	x := []int{1,2}
// // 	x = append(x, 3)
// // 	x = append(x, 4)
// // 	fmt.Println(p,x)
// //    }
// //    func rename(p person) person {
// // 	p.name = "test"
// // 	return p
// //    }

/////////////////////////////
/////////////////////////////

// func main() {
// 	str := [][]byte{
// 		[]byte("iPhone"),
// 		[]byte("iWatch"),
// 		[]byte("iPad")}

// 	fmt.Printf("%q\n", str)

// 	// Separated by " "
// 	fmt.Printf("%s\n", bytes.Join(str, []byte(" ")))

// 	// Separated by ", "
// 	fmt.Printf("%s\n", bytes.Join(str, []byte(", ")))

// 	// Separated by "---"
// 	fmt.Printf("%s\n", bytes.Join(str, []byte("---")))
// }
