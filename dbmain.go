package main

import "fmt"
import "example/kdill/hello/dbscratch"

func main() {
  //don't know what yet

  somethingNew := "this is some dumb bullshit\n";
  fileName := "/home/kdill/code/Tutorials/DbFromScratch/testFile.txt";
  
  err := dbscratch.SaveData1(fileName, []byte(somethingNew));

  if err != nil {
    fmt.Println("Failed to write")
    return
  }


}
