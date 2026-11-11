

## AUDITBOUNDARY-cjk-emphasis

これは**重要。**です

これは**重要**です。

日本語 $`x`$。

## AUDITBOUNDARY-table-cell-count

| A | B |
| --- | --- |
| 1 | 2 | 3 |
| short |

## AUDITBOUNDARY-table-pipes

| A |
| --- |
| `a\|b` |
| C \\| D |
| E \\\\| F |

## AUDITBOUNDARY-table-unicode-empty

| 日本語 | 😀 |
| --- | --- |

## AUDITBOUNDARY-autolink-brackets

[ www.example.com

[ https://example.com

a www. b

WWW.COMMONMARK.ORG

## AUDITBOUNDARY-links-relative-unicode

日本語https://example.com/a。

[relative](../docs/a.md)

https://example.com/\(foo\)

## AUDITBOUNDARY-tasks-markers-nesting

- [ ] `[x]`
- [x] done
  - [ ] nested
- [ ] literal [x] appears later

## AUDITBOUNDARY-strike-delimiters

~single~ ~~double~~

a~~***nested***~~a

~~across
lines~~