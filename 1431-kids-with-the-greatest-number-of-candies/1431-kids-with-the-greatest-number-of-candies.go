func kidsWithCandies(candies []int, extraCandies int) []bool {
    result := make([]bool, len(candies))
    max := math.MinInt
    
    // Find maximum value in candies array
    for i, _ := range candies {
        if candies[i]>max{
            max=candies[i]
        }
    }
    
    // populate result array
    for i, _ := range candies {
        if (candies[i]+extraCandies)>=max{
            result[i]=true
        } else {
            result[i]=false
        }
    }
    
    return result
}