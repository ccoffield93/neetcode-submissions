func twoSum(nums []int, target int) []int {
    vals := make(map[int]int)
    ret := []int{}
    for i := 0; i < len(nums); i++ {
        vals[nums[i]] = i
    }
 
    for i := 0; i < len(nums); i++ {
        goal := target - nums[i]
        index, found := vals[goal]
        if found && index != i { // can't have a duplicate
            ret = append(ret, i)
            ret = append(ret, index)
            return ret
        }
    }

    return ret
}
