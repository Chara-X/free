package free

import "container/heap"

type Heap struct{ offsets offsets }

func (h *Heap) Len() int          { return h.offsets.Len() }
func (h *Heap) Push(offset int64) { heap.Push(&h.offsets, offset) }
func (h *Heap) Peek() int64       { return h.offsets[h.Len()-1] }
func (h *Heap) Pop() int64        { return heap.Pop(&h.offsets).(int64) }

type offsets []int64

func (o offsets) Len() int           { return len(o) }
func (o offsets) Less(i, j int) bool { return o[i] < o[j] }
func (o offsets) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o *offsets) Push(offset any)   { *o = append(*o, offset.(int64)) }
func (o *offsets) Pop() any {
	var offset = (*o)[len(*o)-1]
	*o = (*o)[0 : len(*o)-1]
	return offset
}
