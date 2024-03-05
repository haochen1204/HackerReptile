package saveFile

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// SaveCSV
// ************************************
//
//	@Description: 将用户提取的数据保存在csv表格中
//	@param valueList
//
// ************************************
func SaveCSV(valueList map[interface{}][]interface{}) {
	fileName := time.Now().Format("20060102150405") + ".csv"
	dirPath := "./csvFile"

	// check if directory exists
	_, err := os.Stat(dirPath)

	if os.IsNotExist(err) {
		// create directory
		errDir := os.MkdirAll(dirPath, 0755)
		if errDir != nil {
			fmt.Println("Error creating directory:", errDir)
			return
		}
	}

	fullPath := filepath.Join(dirPath, fileName)
	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// 创建新文件
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Writing the headers
	keys := make([]interface{}, 0, len(valueList))
	// 创建key列表
	for k := range valueList {
		keys = append(keys, k)
	}

	keyStrings := make([]string, len(keys))
	for i, key := range keys {
		keyStrings[i] = fmt.Sprintf("%v", key)
	}
	writer.Write(keyStrings)

	// Writing the data
	firstKey := valueList[keys[0]]
	length := len(firstKey)

	for i := 0; i < length; i++ {
		row := make([]string, len(keys))
		for j, key := range keys {
			// Added check here for writing the dereferenced pointer value
			// 存入的是地址，所以需要读取出地址中的内容
			switch v := (valueList[key][i]).(type) {
			case *string:
				if v != nil {
					row[j] = *v
				}
			case *int:
				if v != nil {
					row[j] = fmt.Sprintf("%d", *v)
				}
			default:
				row[j] = fmt.Sprintf("%v", valueList[key][i])
			}
		}
		writer.Write(row)
	}
}
