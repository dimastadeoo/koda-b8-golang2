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

type tech struct{
	academy string
}
type man struct{
	tech tech 
}

type str struct{
	man []man
}

type obj struct{
	str [][][]str
}

type fruit struct{
	is string
}

type favourite struct{
	fruit fruit
}

type my struct{
	favourite []favourite
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

	// obj := obj{
	// 	str: make(str, 4),
	// }


	// obj.str[3] = make([][]struct{man []man}, 2)
	// obj.str[3][1] = make([]struct{man []man}, 3)
	// obj.str[3][1][2] = struct{
	// 	man []man
	// 	}{
	// 		man: []man{{
	// 			tech: tech{
	// 				academy: "Tech Academy",
	// 			},
	// 		}},
	// 	}
	obj := obj{
		str: [][][]str{3: {1:{2: {
			man: []man{{
				tech: tech{
					academy: "Tech Academy",
				},
			},},
		},},},},
	}

	my := []my{{
		favourite: []favourite{3: {
			fruit: fruit{
				is: "Apple",
			},
		},},
	},}

	fmt.Println(we.are.the.best)
	fmt.Println(hello.world)
	fmt.Println(obj.str[3][1][2].man[0].tech.academy)
	fmt.Println(my[0].favourite[3].fruit.is)


}