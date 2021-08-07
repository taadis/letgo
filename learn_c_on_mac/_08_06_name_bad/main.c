#include <stdio.h>
// This is to bring in the declaration of strlen()
#include <string.h>

int main(int argc, const char * argv[]){
    char name[50];
    printf("Type your first name, please: ");
    
    // scanf() is bad.
    // Do NOT use it.
    // if you type in more than 50 chars,
    // you'll overflow the buffer and drop into gdb.
    scanf("%s", name);
    printf("Welcome, %s.\n", name);
    printf("Your name is %d characters long.", (int)strlen(name));

    return 0;
}
