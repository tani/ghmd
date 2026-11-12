#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#ifndef OFFSET
#define OFFSET 0
#endif
#ifndef MULTIPLIER
#define MULTIPLIER 1
#endif

int ffi_add(int a, int b) {
  return (a + b + OFFSET) * MULTIPLIER;
}
double ffi_half(double value) {
  return value / 2;
}
bool ffi_not(bool value) {
  return !value;
}
int64_t ffi_signed(int64_t value) {
  return value;
}
uint64_t ffi_unsigned(uint64_t value) {
  return value;
}
const char *ffi_greeting(void) {
  return "Hello 日本語 😀";
}
uint64_t ffi_sum(const uint8_t *bytes, uint64_t length) {
  uint64_t result = 0;
  for (uint64_t i = 0; i < length; i++)
    result += bytes[i];
  return result;
}
void ffi_fill(uint8_t *bytes, uint32_t length, uint8_t value) {
  for (uint32_t i = 0; i < length; i++)
    bytes[i] = value;
}
void *ffi_pointer(void *pointer) {
  return pointer;
}
uint8_t ffi_byte(const uint8_t *pointer) {
  return *pointer;
}
