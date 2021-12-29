#include <stdio.h>

int main() {
	char *s="#include <stdio.h>%c%cint main() {%c%cchar *s=%c%s%c;%c%cfprintf(stderr, %chello, world!%c%c%c);%c%cprintf(s,10,10,10,9,34,s,34,10,9,34,34,9,10,9,10,9,10);%c%creturn 0;%c}";
	fprintf(stderr, "hello, world!\n");
	printf(s,10,10,10,9,34,s,34,10,9,34,92,110,34,10,9,10,9,10);
	return 0;
}
