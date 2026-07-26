func search(nums []int, target int) int {

    start := 0
    end := len(nums) - 1

    for (start <= end) { // we want to allow a case where we've landed exactly on the target
        mid := (start + end) / 2 
        fmt.Printf("%d %d %d\n", start, mid, end)

        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            start = mid + 1
        } else {
            end = mid - 1
        }
    }

    return -1
}
