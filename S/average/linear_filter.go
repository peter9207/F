package average

import (
	"github.com/peter9207/F/S/shapes"
)

func Box(data []float64, size int) (filtered []float64) {

	linkedList := shapes.NewLinkedList()
	bucketSum := float64(0)

	count := float64(0)

	//initialize with first couple elements
	for i := 0; i < int(size/2+1); i++ {
		v := data[i]
		linkedList.AddFirst(v)
		bucketSum = bucketSum + data[i]
		count = count + 1
	}

	for i, v := range data {

		if int(i+size) < len(data) {
			bucketSum = bucketSum + v
			linkedList.AddFirst(v)
			count = count + 1
		}

		if count > float64(size) {
			old := linkedList.RemoveFirst()
			bucketSum = bucketSum - old
			count = count - 1
		}

		average := float64(bucketSum) / float64(count)
		filtered = append(filtered, average)

	}
	return

}
