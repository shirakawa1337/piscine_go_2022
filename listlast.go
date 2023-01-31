package piscine

func ListLast(l *List) interface{} {
	it := l.Head
	last := it
	for it != nil {
		last = it
		it = it.Next
	}
	if last != nil {
		return last.Data
	} else {
		return nil
	}
}
