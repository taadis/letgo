#include <stdio.h>
#include "param_address.h"

int main(int argc, const char * argv[]) {
    struct DVDInfo myDVD;

    printf("Address of myDVD.rating in main(): %p\n", &(myDVD.rating));

    PrintParamInfo(&myDVD, myDVD);

    return 0;
}

void PrintParamInfo(struct DVDInfo *myDVDPtr, struct DVDInfo myDVDCopy) {
    printf("Address of myDVDPtr->rating in PrintParamInfo(): %p\n", &(myDVDPtr->rating));

    printf("Address of myDVDCopy.rating in PrintParamInfo(): %p\n", &(myDVDCopy.rating));
}

