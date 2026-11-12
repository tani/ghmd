#if defined(__wasip1__) || defined(__wasi__)
/* Some WASI libc builds omit tolower(), which cmark calls from its ASCII
 * case-folding paths. Keep this translation unit exclusively for WASI. */
int tolower(int c) {
  return c >= 'A' && c <= 'Z' ? c + ('a' - 'A') : c;
}
#endif
