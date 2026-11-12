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

The distribution is copied byte-for-byte, and tests check its upstream SHA-256
directly. The preparation script verifies the input hash and copies the bytes.

`src/html/lexbor.c` surrounds the include with scoped preprocessor macros so
element creation passes through `src/html/parser.c` to attach source locations,
including reconstructed formatting nodes. The upstream implementation is given
a private name during compilation, while its internal calls retain the wrapper
name. No source file is rewritten and no callback field is added to Lexbor.

The macro adapter relies on the pinned source's first argument being
`lxb_html_tree_t *tree` in declarations/definitions and `tree` at call sites.
Unrecognized forms fail to compile. Update the pinned commit and upstream hash,
review the adapter, and run the location and rendering tests when upgrading.

`compat/float.h` allows compilation with Bun's bundled C compiler without GCC's
headers. It is intentionally empty: the selected module includes that header but
uses none of its constants. The preparation script and tests verify this.

The parser is compiled with `bun:ffi` `cc()` on import. There is no external
compiler, downloaded binary, runtime network request, or package install step.
Sanitizer policy is maintained separately in `src/html/sanitize.c`.
