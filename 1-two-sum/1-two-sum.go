func twoSum(nums []int, target int) []int {
    val := map[int]int{}
    ans := []int{}
    
    for i, num := range nums {
        pos, ok := val[target-num]
        if ok{
            ans = append(ans, i, pos)
        } else {
            val[num]=i
        }
        
        if len(ans)!=0{
            break
        }
    }
    
    return ans
    
}