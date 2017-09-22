#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

int base64_encode(const uint8_t* input, int input_len, char** output, int* output_len);
int base64_decode(const char* input, int input_len, uint8_t* output, int output_len);

#ifdef __cplusplus
}
#endif
