package main
import ("os"
		"fmt"
		"encoding/json"
			"bufio")


type Config struct{
	
	DBHost        string `json:"database_host"`
    DBPort        int    `json:"database_port"`
    DBUsername    string `json:"database_username"`
    DBPassword    string `json:"database_password"`
    ServerPort    int    `json:"server_port"`
    ServerDebug   bool   `json:"server_debug"`
    ServerTimeout int    `json:"server_timeout"`

}


func main()  {
	filePath:= "config.json"
	var config Config
	file, err := os.Open(filePath)
	
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
    
	Reader:= NewDecoder(file)
	err = Reader.Decode(&config)

	if err != nil{
		fmt.Println(err)	
	}
	

}			