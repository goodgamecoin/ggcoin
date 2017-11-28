
#ifndef __CRYPTONIGHT_H_
#define __CRYPTONIGHT_H_

#include <unistd.h>

void cn_slow_hash(const void * pptr, size_t dlen, char * h);

void cn_fast_hash(const void * pptr, size_t dlen, char * h);

#endif