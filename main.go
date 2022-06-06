package main

import "fmt"

func main() {

	value := 0
	firstValueFound := false

	for {
		//Data Input
		exampleDatas := []interface{}{1, "X", 8, 17, "Y", "Z", 78, 113}
		////
		var diffLayers1, diffLayers2 []int
		for index, exampleData := range exampleDatas {
			if index != 0 {
				switch exampleData.(type) {
				case string:
					exampleData = value
					exampleDatas[index] = exampleData
				}
				diffLayers1 = append(diffLayers1, exampleData.(int)-exampleDatas[index-1].(int))
				if len(diffLayers1) >= 2 {
					diffLayers2 = append(diffLayers2, diffLayers1[len(diffLayers1)-1]-diffLayers1[len(diffLayers1)-2])
					if len(diffLayers2) >= 2 {
						if diffLayers2[len(diffLayers2)-1]-diffLayers2[len(diffLayers2)-2] != 1 {
							firstValueFound = false
							break
						} else {
							exampleDatas[index] = exampleData
							firstValueFound = true
							break
						}
					}
				}
			}
		}
		value++
		if firstValueFound {
			indexCurrent := len(diffLayers1)
			diffLayer1 := diffLayers1[len(diffLayers1)-1]
			diffLayer2 := diffLayers2[len(diffLayers2)-1]
			for i := indexCurrent; i < len(exampleDatas)-1; i++ {
				diffLayer2++
				diffLayer1 = diffLayer1 + diffLayer2
				if exampleDatas[i+1] != nil {
					exampleDatas[i+1] = diffLayer1 + exampleDatas[i].(int)
				}
			}
			fmt.Println("Data Output: ", exampleDatas)
			break
		}
	}
}
