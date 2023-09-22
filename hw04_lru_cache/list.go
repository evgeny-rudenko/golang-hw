package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

func NewList() List {
	return new(list)
}

type list struct {
	front *ListItem
	back  *ListItem
	cnt   int
}

func (l *list) Len() int {
	return l.cnt
	/*
		// вариант не подошел - так как перебор всех элементов
		if l.front == nil {
				return 0
			}
			count := 1
			current := l.front
			for current.Next != nil {
				count++
				current = current.Next
			}
			return count
	*/
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Next: l.front}
	if l.front != nil {
		l.front.Prev = newItem
	} else {
		l.back = newItem
	}
	l.front = newItem
	l.cnt++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Prev: l.back}
	if l.back != nil {
		l.back.Next = newItem
	} else {
		l.front = newItem
	}
	l.back = newItem
	l.cnt++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}
	l.cnt--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}
	l.Remove(i)
	l.PushFront(i.Value)
}
