package main
 
import (
   "fmt"
)
 
// quick sort is efficiency for large scale array
func quickSort(a []int, length int) {
   if length < 2 {
      return
   }
   max := 0
   for i := 0; i < length; i++ {
      if a[max] <= a[i] {
         max = i
      }
   }
   a[length-1], a[max] = a[max], a[length-1]
   quickSort(a, length-1)
}
 
func swap(a []int, i, j int) {
   a[i], a[j] = a[j], a[i]
}
 
func quickSort3Way(a []int, first, end int) {
   if first >= end {
      return
   }
   lt := first
   i := first + 1
   gt := end
   bound := a[first]
   for i <= gt {
      if a[i] < bound {
         swap(a, lt, i)
         lt++
         i++
      } else if a[i] > bound {
         swap(a, i, gt)
         gt--
      } else {
         i++
      }
   }
   quickSort3Way(a, first, lt-1)
   quickSort3Way(a, gt+1, end)
}
 
func main() {
   a := []int{2, 9, 9, 1, 10, 10}
   b := []int{1, 4, 6, 3, 2, 7, 7, 7, 7, 4}
   quickSort(a, len(a))
   quickSort3Way(b, 0, len(b)-1)
   fmt.Println(a)
   fmt.Println(b)
}
