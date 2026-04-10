package main 
import ("os"
		"fmt"
		"encoding/csv")




func main()  {
	filePath:= "CsvRead.csv"
	file, err := os.Open(filePath)

	if err != nil{
		fmt.Println(err)
	}

	defer file.Close()

	Reader:= csv.NewReader(file)
	records, err := Reader.ReadAll()
	
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(records)
}