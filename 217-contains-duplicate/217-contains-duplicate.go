func containsDuplicate(nums []int) bool {
    set := make(map[int]struct{})
    
    for _, num := range nums {
        _, present := set[num]
        if present{
            return true
        } else {
            set[num]=struct{}{}
        }
    }
    
    return false
}