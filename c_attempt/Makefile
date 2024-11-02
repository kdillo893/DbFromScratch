

dbmain : dbmain.o simpleSave.o
	gcc -o dbmain dbmain.o simpleSave.o

simpleSave.o : simpleSave.c
	gcc simpleSave.c -c

dbmain.o : dbmain.c
	gcc dbmain.c -c


clean :
	rm dbmain dbmain.o simpleSave.o

