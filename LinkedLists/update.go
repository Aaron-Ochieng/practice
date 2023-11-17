package linkedlists

func (list *LinkedList) Update(oldData, newData interface{}) {
	current := list.Head
	for current != nil {
		if current.Data == oldData {
			current.Data = newData
			break
		}
		current = current.Next
	}
}
