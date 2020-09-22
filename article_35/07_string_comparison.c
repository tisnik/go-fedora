/*
 Seriál "Programovací jazyk Go"

 Třicátá pátá část
    Programovací jazyk Go pro skalní céčkaře (2.část)
    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
*/

#include <stdio.h>
#include <string.h>

int main(void) {
	printf("%d\n", strcmp("aa", "ab"));
	printf("%d\n", strcmp("aa", "aa"));

	printf("%d\n", strcmp("e", "é"));
	printf("%d\n", strcmp("e", "ě"));
	printf("%d\n", strcmp("ě", "é"));

	printf("%d\n", strcmp("z", "é"));
	printf("%d\n", strcmp("z", "ě"));

	printf("%d\n", strcmp("h", "ch"));
	printf("%d\n", strcmp("ch", "i"));

	printf("%d\n", strcmp("Hrdina", "Chocholoušek"));
	printf("%d\n", strcmp("Crha", "Chocholoušek"));

        return 0;
}
