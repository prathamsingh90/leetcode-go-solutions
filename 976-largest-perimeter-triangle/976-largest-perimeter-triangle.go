func largestPerimeter(nums []int) int {
    
    // Triangle have 3 sides
    if len(nums)<3{
        return 0
    }
    
    sort.Ints(nums)
    
    for i:=len(nums)-1; i>1; i--{
        if nums[i]<(nums[i-1]+nums[i-2]){
            return nums[i]+nums[i-1]+nums[i-2]
        }
    }
    
    return 0
}