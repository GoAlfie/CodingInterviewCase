package queue

type Queue struct {
	items []any
}

func (q *Queue) Push(e any) {
	q.items = append(q.items, e)
}

func (q *Queue) Pop() any {
	if q.Empty() {
		return nil
	}
	res := q.items[0]
	q.items = q.items[1:]

	return res
}

func (q *Queue) Empty() bool {
	return len(q.items) == 0
}
