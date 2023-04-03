package main

func MoveElementToEnd(array []int, toMove int) []int {
  for i, j := 0, len(array)-1; i < j; {
    for array[i] != toMove && i < j{
      i++
    }
    for array[j] == toMove && i < j{
    j--
  }
  array[i], array[j] = array[j], array[i]

  }

return array
}
