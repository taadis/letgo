#include <stdio.h>
#include "struct_size.h"

int main(int argc, const char * argv[]) {
    struct DVDInfo myInfo;

    printf("rating field:    %ld byte\n", sizeof(myInfo.rating));

    printf("title field:     %ld bytes\n", sizeof(myInfo.title));

    printf("comment field:   %ld bytes\n", sizeof(myInfo.comment));

    printf("                 ---------\n");

    printf("myInfo struct:   %ld bytes", sizeof(myInfo));

    return 0;
}

