package dbscratch

import "os"

func SaveData1(path string, data []byte) error {
  fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)

  if err != nil {
    return err
  }
  //defer the file close until after this function...
  defer fp.Close();

  _, err = fp.Write(data)
  if err != nil {
    return err
  }
  
  
  return fp.Sync()
}