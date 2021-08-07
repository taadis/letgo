#include <stdio.h>
// This is to bring in the declaration of strlen()
#include <string.h>

int main(int argc, const char * argv[]){
    char name[50];
    int nameLength;
    printf("Type your first name, please: ");
    fgets(name, 50, stdin);
    nameLength = strlen(name);
    name[nameLength - 1] = 0;
    printf("Welcome, %s.\n", name);
    printf("Your name is %d characters long.", (int)strlen(name));

    return 0;
}
