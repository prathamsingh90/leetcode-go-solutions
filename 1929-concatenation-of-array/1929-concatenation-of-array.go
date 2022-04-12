func getConcatenation(nums []int) []int {
    ans := make([]int, 2*len(nums))
    size := len(nums)
    
    for i, _ := range nums {
        ans[i], ans[i+size] = nums[i], nums[i]
    }
    
    return ans
}