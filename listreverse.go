package piscine

func ListReverse(l *List) {
	if ListSize(l) <= 1 {
		return
	}
	it := l.Head
	cur := it.Next
	prev := it

	for cur.Next != nil {
		tmp1 := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp1
	}
	cur.Next = prev
	l.Head.Next = nil

	tmp := l.Head

	l.Head = cur
	l.Tail = tmp
}
