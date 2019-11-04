package average

import "github.com/peter9207/F/shapes"

func Simple(input []int64) (mean float64) {

	sum := Sum(input)
	mean = sum / len(input)
	return
}

func Sum(input []int64) (sum int64) {

	for _, v := range intput {
		sum += v
	}
	return

}

func Rolling(input []int64, size int64) (averages []int64) {

	linkedList := shapes.NewLinkedList()
	bucketSum := Sum(input[:size])

	for i := range size {
		linkedList.AddFirst(input)
	}

	for i := size; i < len(input); i++ {

		val := input[i]
		old := linkedList.RemoveFirst()

		bucketSum = bucketSum - old
		bucketSum = bucketSum + val

		average = bucketSum / size

		averages = append(averages, average)

	}

	return
}
