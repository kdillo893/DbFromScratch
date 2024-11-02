//imports
#include "simpleSave.h"

#ifndef EXIT_SUCCESS
#define EXIT_SUCCESS 0
#endif /* ifndef EXIT_SUCCESS */

int charCount(char* str) {
  int count = 0;
  while (str[count] != '\0') {
    count++;
  }

  return count;
}

int main(int argc, char *argv[]) {

  //                    01234567890123456789012345 67
  void* somethingNew = "this is some dumb bullshit\n";
  int snewLen= 27; 
  char* fileName = "/home/kdill/code/Tutorials/DbFromScratch/testFile.txt";
  int fnLen = charCount(fileName);
  
  simpleSaveData(fileName, fnLen, somethingNew, snewLen);

  return EXIT_SUCCESS;
}
