#include "base64.h"

const uint8_t BAD = 1;

static const char base64val[] = {
    BAD,BAD,BAD,BAD, BAD,BAD,BAD,BAD, BAD,BAD,BAD,BAD, BAD,BAD,BAD,BAD,
    BAD,BAD,BAD,BAD, BAD,BAD,BAD,BAD, BAD,BAD,BAD,BAD, BAD,BAD,BAD,BAD,
    BAD,BAD,BAD,BAD, BAD,BAD,BAD,BAD, BAD,BAD,BAD, 62, BAD,BAD,BAD, 63,
    52, 53, 54, 55,  56, 57, 58, 59,  60, 61,BAD,BAD, BAD,BAD,BAD,BAD,
    BAD,  0,  1,  2,   3,  4,  5,  6,   7,  8,  9, 10,  11, 12, 13, 14,
    15, 16, 17, 18,  19, 20, 21, 22,  23, 24, 25,BAD, BAD,BAD,BAD,BAD,
    BAD, 26, 27, 28,  29, 30, 31, 32,  33, 34, 35, 36,  37, 38, 39, 40,
    41, 42, 43, 44,  45, 46, 47, 48,  49, 50, 51,BAD, BAD,BAD,BAD,BAD
};

static const char base64digits[] = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";

#define DECODE64(c)  (isascii(c) ? base64val[c] : BAD)

int base64_encode(const uint8_t* input, int input_len, char** output, int* output_len) {
    int len = input_len % 3 == 0 ? input_len / 3 * 4 : input_len / 3 * 4 + 4;
    *output_len = len;
    char* buf = (char*)malloc(len + 1);
    len = input_len;
    *output = buf;
    for (; len >= 3; len = len - 3) {
        *buf++ = base64digits[input[0] >> 2];
        *buf++ = base64digits[((input[0] << 4) & 0x30) | (input[1] >> 4)];
        *buf++ = base64digits[((input[1] << 2) & 0x3c) | (input[2] >> 6)];
        *buf++ = base64digits[input[2] & 0x3f];
        input = input + 3;
    }
    if (len > 0) {
        uint8_t fragment;
        *buf++ = base64digits[input[0] >> 2];
        fragment = (input[0] << 4) & 0x30;

        if (len > 0) {
            fragment |= input[1] >> 4;
        }

        *buf++ = base64digits[fragment];
        *buf++ = (len < 2) ? '=' : base64digits[(input[1] << 2) & 0x3c];
        *buf++ = '=';
    }
    *buf = 0;
    return 0;
}

int base64_decode(const char* input, int input_len, uint8_t** output, int* output_len) {
    return 0;
}
