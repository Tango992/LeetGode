package main

func main() {

}

func twoSum(numbers []int, target int) []int {
    for i, j := 0, len(numbers)-1; i < j; {
        currentTotal := numbers[i] + numbers[j]
        if currentTotal == target {
            return []int{i+1, j+1}
        } else if currentTotal < target {
            i++
            continue
        }
        j--
    }
    return []int{}
}