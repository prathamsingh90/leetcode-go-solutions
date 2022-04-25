func removeDuplicates(nums []int) int {
    if len(nums)==1 || len(nums)==0{
        return len(nums)
    }
    
    orig := 0
    
    for i:=1; i<=len(nums)-1;{
        if nums[orig]==nums[i]{
            i++
        } else {
            nums[orig+1],nums[i] = nums[i], nums[orig+1]
            orig++
            i++
        }
    }
    return orig+1
}