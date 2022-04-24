func removeElement(nums []int, val int) int {
    low := 0
    high := len(nums)-1
    
    for high>=low{
        if nums[low]==val{
            nums[low], nums[high]=nums[high], nums[low]
            high--
        } else {
            low++
        }
    }
    return low
}