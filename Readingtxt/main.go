
//Reading txt file in go
//Date:09-04-2026
//Author: Oshim Odinaka Joshua

// We will be importing some new package that we never used before in handling file in go
//----------------------------------------------------------------------------
//1 os:This is a package that helps in file handling go. The package provides different kinds of functions.
//-----------------------------------------------------------------------------
//Look down below to see how this file is being used to read a txt file
//-----------------------------------------------------------------------------
//NOTE: To read a file the file must be available in the directory if npt the terminal will compile => "no such file or dirextory"

//in the package below the Readfile is a function that comes with the "os" package wich helps read the content of the file path.

package main
import ("os"
		"fmt")


func main()  {
filepath:="data.txt"
file, err := os.ReadFile(filepath)

if err != nil {
	fmt.Println("Encountering Error", err)
} else{
	fmt.Println(string(file))
}

}