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


func main()  {
	we := we{
		are: are{
			the: the{
				best: "Koda",
			},
		},
	}
	fmt.Println(we.are.the.best)

	
}