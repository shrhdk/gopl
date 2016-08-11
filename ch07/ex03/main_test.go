package ex03

import (
	"fmt"
)

func ExampleTree_String() {
	tr :=
		&tree{
			1,
			&tree{
				2,
				nil,
				nil,
			},
			&tree{
				3,
				&tree{
					4,
					&tree{
						5,
						&tree{
							6,
							nil,
							nil,
						},
						nil,
					},
					&tree{
						7,
						nil,
						nil,
					},
				},
				&tree{
					8,
					nil,
					nil,
				},
			},
		}

	fmt.Printf("%s", tr)

	// Output: 12345678
}
