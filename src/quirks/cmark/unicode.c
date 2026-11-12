/* Unicode default lowercase using pinned Unicode 17 data and cmark UTF-8. */
#include "unicode-lower-data.h"

static int gfm_unicode_in_ranges(uint32_t cp, const gfm_unicode_range *ranges, size_t count) {
  size_t lo = 0, hi = count;
  while (lo < hi) {
    size_t mid = lo + (hi - lo) / 2;
    if (cp < ranges[mid].first)
      hi = mid;
    else if (cp > ranges[mid].last)
      lo = mid + 1;
    else
      return 1;
  }
  return 0;
}
static int gfm_unicode_cased(uint32_t cp) {
  return gfm_unicode_in_ranges(cp, gfm_cased_ranges,
                               sizeof(gfm_cased_ranges) / sizeof(gfm_cased_ranges[0]));
}
static int gfm_unicode_ignorable(uint32_t cp) {
  return gfm_unicode_in_ranges(cp, gfm_ignorable_ranges,
                               sizeof(gfm_ignorable_ranges) / sizeof(gfm_ignorable_ranges[0]));
}
static int gfm_unicode_followed_by_cased(const unsigned char *s, size_t length, size_t pos) {
  while (pos < length) {
    int32_t cp;
    int width = cmark_utf8proc_iterate(s + pos, length - pos, &cp);
    if (width <= 0)
      return 0;
    if (!gfm_unicode_ignorable(cp))
      return gfm_unicode_cased(cp);
    pos += (size_t)width;
  }
  return 0;
}
static void gfm_unicode_lower_into(cmark_strbuf *out, const unsigned char *s, size_t length) {
  int preceding_cased = 0;
  for (size_t pos = 0; pos < length;) {
    int32_t cp;
    int width = cmark_utf8proc_iterate(s + pos, length - pos, &cp);
    if (width <= 0) {
      /* Inputs from JavaScript are UTF-8; replace malformed bytes defensively. */
      static const unsigned char replacement[] = {0xef, 0xbf, 0xbd};
      cmark_strbuf_put(out, replacement, sizeof(replacement));
      preceding_cased = 0;
      pos++;
      continue;
    }
    if (cp == 0x3a3 && preceding_cased && !gfm_unicode_followed_by_cased(s, length, pos + width)) {
      static const unsigned char final_sigma[] = {0xcf, 0x82};
      cmark_strbuf_put(out, final_sigma, sizeof(final_sigma));
    } else {
      size_t lo = 0, hi = sizeof(gfm_lower_mappings) / sizeof(gfm_lower_mappings[0]);
      while (lo < hi) {
        size_t mid = lo + (hi - lo) / 2;
        if ((uint32_t)cp > gfm_lower_mappings[mid].cp)
          lo = mid + 1;
        else
          hi = mid;
      }
      if (lo < sizeof(gfm_lower_mappings) / sizeof(gfm_lower_mappings[0]) &&
          gfm_lower_mappings[lo].cp == (uint32_t)cp)
        cmark_strbuf_put(out, gfm_lower_bytes + gfm_lower_mappings[lo].offset,
                         gfm_lower_mappings[lo].length);
      else
        cmark_strbuf_put(out, s + pos, width);
    }
    if (!gfm_unicode_ignorable(cp))
      preceding_cased = gfm_unicode_cased(cp);
    pos += (size_t)width;
  }
}
