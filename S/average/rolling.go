package average

import (
	"fmt"

	"github.com/peter9207/F/S/shapes"
)

func Simple(input []int64) (mean float64) {

	sum := Sum(input)
	mean = float64(sum) / float64(len(input))
	return
}

func Sum(input []int64) (sum int64) {

	for _, v := range input {
		sum += v
	}
	return

}

func Rolling(input []int64, size int64) (averages []float64) {

	linkedList := shapes.NewLinkedList()
	bucketSum := int64(0)

	count := int64(0)

	for _, v := range input {
		bucketSum = bucketSum + v

		linkedList.AddFirst(v)

		count = count + 1
		if count > size {
			old := linkedList.RemoveFirst()
			bucketSum = bucketSum - old
			count = count - 1
		}

		average := float64(bucketSum) / float64(count)
		fmt.Printf("average: %f", average)
		averages = append(averages, average)

	}

	return
}
