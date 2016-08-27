package display

func Example_recursive() {
	type LinkedList struct {
		val  int
		next *LinkedList
	}

	var list LinkedList
	list.next = &list

	Display("list", list, 3)
	// Output:
	// Display list (display.LinkedList):
	// list.val = 0
	// (*list.next).val = 0
	// (*(*list.next).next).val = 0
	// ...
}
