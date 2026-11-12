# Lexbor HTML parser

`lexbor.h` is the HTML module's official single-header distribution, generated
from the commit recorded in `PROVENANCE.json`. It includes the upstream license
header; `LICENSE` and `NOTICE` are also retained.

To regenerate, in a separate checkout of https://github.com/lexbor/lexbor:

```sh
git checkout e651d6bce83a6601dff37261374f0545e910b1c1
perl single.pl html
```

Then run, from this repository:

```sh
bun scripts/prepare-lexbor.ts /path/to/lexbor/html-single.h
```

There is one deliberate adaptation: the element-construction definition is
renamed so `src/html/parser.c` can wrap it and attach source locations (including
reconstructed formatting nodes). The wrapper delegates to the original function;
HTML parsing and tree repair stay in Lexbor. Tests reverse the rename and verify
the original distribution's hash. Update the pinned commit, both hashes, and
review the wrapper when upgrading.

`compat/float.h` allows compilation with Bun's bundled C compiler without GCC's
headers. It is intentionally empty: the selected module includes that header but
uses none of its constants. The preparation script and tests verify this.

The parser is compiled with `bun:ffi` `cc()` on import. There is no external
compiler, downloaded binary, runtime network request, or package install step.
Sanitizer policy is maintained separately in `src/html/sanitize.c`.
