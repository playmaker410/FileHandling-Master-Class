//Golang line by line text reading
//We will be reading this with the function that comes with the "bufio " package which are the following
//Bufio: This help in reading of files. this reads big files. it reads and writes one byte at a time which makes program fast and efficient especially when dealing with large files.

//--------------------------------Some key Fuction of bufio---------------------------------------

//1. NewScanner():used in collecting or reading files line by line in go. it wraps Around with io.Reader(like os.Stdin ora file)
//-----------------------------------------------------------------------------------------------
// Here this program will be reading the last three row of the specified file

package main
import ("fmt"
		"os"
		"bufio"
	)


func PrintLastln(lines []string, num int)[]string{
	var Printlastline []string
	for i := len(lines) - num; i < len(lines); i++{
		Printlastline = append(Printlastline, lines[i])
	}
	return Printlastline
}



func main()  {
filepath:= "log.txt"
file, err:= os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644) //OR SIMPLY os.Open(filePath)


if err != nil{
	fmt.Println(err)
}
defer file.Close()

var lines []string
Read:= bufio.NewScanner(file)

for Read.Scan(){
	lines = append(lines, Read.Text())
}

if err := Read.Err(); err != nil{
	fmt.Println("Encountering erro while reading", err)

}

//Printing last three

lastthree:= PrintLastln(lines, 3)

for _, line :=  range lastthree{
	fmt.Println(line)
}

}	