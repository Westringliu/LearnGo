#include <stdlib.h>
#include <stdio.h>

float myFunc(float benjin,float lixi,int time)
{
    if(time == 0)
    {
        return benjin;
    }
    time--;
    benjin *= lixi + 1;
    return myFunc(benjin,lixi,time);
}

int main()
{
    float ss = myFunc(10000.0,0.04,20);
    printf("%f",ss);
    return 0;
}