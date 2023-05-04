package main

func SpiralTraverse(array [][]int) []int {
    var finalarray []int
    rows, columns := len(array), len(array[0])
    sr := 0
    sc := 0
    er := rows-1
    ec := columns-1
    for sc <= ec && sr <= er{
        for j := sc; j <= ec; j++ {
            finalarray = append(finalarray, array[sr][j])
        }
        for i := sr+1; i <= er; i++ {
            finalarray = append(finalarray, array[i][ec])
        }
        if sr < er{
            for j := ec-1; j >= sc; j-- {
                finalarray = append(finalarray, array[er][j])
            }
        }
        if sc < ec{
            for i := er-1; i > sr; i-- {
                finalarray = append(finalarray, array[i][sc])
            }
        }
        sr++
        sc++
        er--
        ec--
    }
	return finalarray
}
