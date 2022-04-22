func moveZeroes(nums []int)  {
    orig := 0
    
    for i:=0; i<len(nums);{
        if orig<len(nums){
        if nums[i]!=0 {
            nums[i], nums[orig] = nums[orig], nums[i]
            orig++
        }
        i++
        }
    }
}