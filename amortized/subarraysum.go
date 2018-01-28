package amortized


func SubarraySum(a []int, value int) (int,int){

  p1 := 0
  p2 := 1
  subSum := a[p1]+a[p2]
  for p1 < len(a) && p2 < len(a) {
      if subSum == value{
        return p1, p2
      }else if subSum < value {
        p2++
        subSum += a[p2]
      }else{
         subSum -= a[p1]
         p1++
      }
  }
  return -1,-1
}

func Solve2Sum(a []int, value int) (int, int){
  p1 := 0
  p2 := len(a)-1
  for p1 < p2 {
    if a[p1] + a[p2] < value {
      p1++
    }else if a[p1] + a[p2] > value{
      p2--
    }else if a[p1] + a[p2] == value{
       return p1, p2
    }
  }
  return -1, -1
}
