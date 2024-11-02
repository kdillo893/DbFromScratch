#include <assert.h>
#include <stdio.h>

//declaration thingy
#ifndef SSAVE_DATA
#include "simpleSave.h"
#endif /* ifndef SSAVE_DATA */

/** open fp, write new data, close fp...
 */
void simpleSaveData(char* path, int pathLength, void* data, int dataLength) {

  FILE* file;

  file = fopen(path, "a");

  assert(file != NULL);
  printf("writing to file: %s\n", path);

  //writing void data to the path? need to append maybe...
  fwrite(data, dataLength, sizeof(void), file);

  //we're done!
  fclose(file);

  printf("File written, returning\n");

}
