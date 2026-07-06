package main

import (
	"fmt"
)

type the struct{
	best string
}

type are struct{
	the the
}

type we struct{
	are are
}

type hello struct{
	world string
}


func main()  {
	we := we{
		are: are{
			the: the{
				best: "Koda",
			},
		},
	}

	hello := hello{
		world: "Hello World",
	}
	fmt.Println(we.are.the.best)
	fmt.Println(hello.world)



}