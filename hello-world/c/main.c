#include <stdio.h>

void f(){fprintf(stderr, "hello, world!\n");}

int main()
{
	char *s="#include <stdio.h>%c%cvoid f(){fprintf(stderr, %chello, world!%c%c%c);}%c%cint main()%c{%c%cchar *s=%c%s%c;%c%cf();%c%cprintf(s,10,10,34,92,110,34,10,10,10,10,9,34,s,34,10,9,10,9,10,9,10);%c%creturn 0;%c}";
	f();
	printf(s,10,10,34,92,110,34,10,10,10,10,9,34,s,34,10,9,10,9,10,9,10);
	return 0;
}
