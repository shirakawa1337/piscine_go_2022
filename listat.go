package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	it := l
	currentPosition := 0
	for it != nil {
		if currentPosition == pos {
			return it
		}

		currentPosition += 1
		it = it.Next
	}
	return nil
}
