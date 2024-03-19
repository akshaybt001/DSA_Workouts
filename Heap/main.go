package main

type Heap struct {
	array []int
}

func (h *Heap) Insert(value int) {
	h.array = append(h.array, value)
	h.heapifyUp(len(h.array) - 1)

}
func (h *Heap) heapifyUp(i int) {
	for h.array[parent(i)] < h.array[i] {
		swap(h.array, parent(i), i)
		i = parent(i)
	}
}
func (h *Heap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.heapifyDown(0)
	return extracted
}

func (h *Heap) heapifyDown(i int) {
	lastIndex := len(h.array) - 1
	l, r := left(i), right(i)
	childToCompare:=0
	for l<=lastIndex{
		if l==lastIndex{
			childToCompare=l
		}else if h.array[l]>h.array[r]{
			childToCompare=l
		}else{
			childToCompare=r
		}
		if h.array[i]<h.array[childToCompare]{
			swap(h.array,i,childToCompare)
			i=childToCompare
			l,r=left(i),right(i)
		}else{
			break
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return (2 * i) + 1
}
func right(i int) int {
	return (2 * i) + 2
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
