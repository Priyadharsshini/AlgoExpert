package main

func OneEdit(stringOne string, stringTwo string) bool {
	madeEdit := false
    indexOne := 0
    indexTwo := 0
    if ((len(stringOne)-len(stringTwo) > 1) || (len(stringTwo) - len(stringOne) > 1)) {
        return false
    }
    for indexOne < len(stringOne) && indexTwo < len(stringTwo){
        if stringOne[indexOne] != stringTwo[indexTwo]{
            if madeEdit{
                return false
            }
            madeEdit = true
            if  len(stringOne) > len(stringTwo){
                indexOne++
            } else if len(stringTwo) > len(stringOne) {
                indexTwo++
            } else {
                indexOne++
                indexTwo++
            }
        } else {
            indexOne++
            indexTwo++
        }
    }
    
	return true
}
