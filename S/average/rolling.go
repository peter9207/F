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
	bucketSum := Sum(input[:size])

	for i := 0; int64(i) < size; i++ {
		linkedList.AddFirst(input[i])
	}

	for i := int(size); i < len(input); i++ {

		val := input[i]
		old := linkedList.RemoveFirst()

		fmt.Println(old)

		bucketSum = bucketSum - old
		bucketSum = bucketSum + val

		average := float64(bucketSum) / float64(size)
		averages = append(averages, average)

	}

	return
}
