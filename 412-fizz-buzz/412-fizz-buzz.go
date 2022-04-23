func fizzBuzz(n int) []string {
    
    result := make([]string, n)
    for i:=1; i<=n; i++ {
    var ans strings.Builder
        three, five := false, false
    if i%3==0{
        ans.WriteString("Fizz")
        three = true
    }
    if i%5==0{
            ans.WriteString("Buzz")
            five = true
    } 
    if three==false && five == false {
        num := strconv.Itoa(i)
        ans.WriteString(num)
    }
        result[i-1] = ans.String()
    }
    
    return result
}