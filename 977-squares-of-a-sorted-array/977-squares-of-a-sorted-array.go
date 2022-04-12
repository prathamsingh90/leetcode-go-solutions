func square(val int) int {
    return abs(val) * abs(val)
}

func abs(val int) int {
    if val<0{
        return -1*val
    }
    return val
}

func sortedSquares(nums []int) []int {
   
    back := 0
    front := len(nums)-1
    pos := len(nums)-1
    
    result := make([]int, len(nums))
    
    for front>=back&&pos>=0{
        if abs(nums[front])>abs(nums[back]){
            result[pos]=square(nums[front])
            front--
            pos--
        } else {
            result[pos]=square(nums[back])
            back++
            pos--
        }
    }
    
    return result
}