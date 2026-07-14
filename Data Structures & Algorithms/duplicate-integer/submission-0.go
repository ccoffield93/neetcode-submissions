func hasDuplicate(nums []int) bool {
    m := make(map[int]int)
    for i:= 0; i < len(nums); i++ {
        _, prs := m[nums[i]]
        if prs {
            return true
        }
        m[nums[i]] = 1
    }
    return false
}
