#!/usr/bin/env python3
"""Convert simple truecolor ANSI output to an HTML terminal-cell grid.

Every character is placed in an explicit 1- or 2-cell box, so browser font
fallback (CJK, emoji) can never drift the columns out of alignment.
"""
import html
import re
import sys
import unicodedata

SGR = re.compile(r"\x1b\[([0-9;]*)m")


def cell_width(ch: str) -> int:
    if unicodedata.east_asian_width(ch) in ("W", "F"):
        return 2
    if 0x1F300 <= ord(ch) <= 0x1FAFF:
        return 2
    return 1


def cells(text: str, color: str | None) -> str:
    out = []
    style = f' style="color:{color}"' if color else ""
    for ch in text:
        cls = "c2" if cell_width(ch) == 2 else "c1"
        out.append(f'<span class="{cls}"{style}>{html.escape(ch)}</span>')
    return "".join(out)


def convert(ans: str) -> str:
    body = []
    for line in ans.rstrip("\n").split("\n"):
        pos = 0
        color = None
        parts = []
        for m in SGR.finditer(line):
            if m.start() > pos:
                parts.append(cells(line[pos : m.start()], color))
            params = m.group(1)
            if params.startswith("38;2;"):
                r, g, b = params.split(";")[2:5]
                color = f"rgb({r},{g},{b})"
            elif params in ("", "0"):
                color = None
            pos = m.end()
        if pos < len(line):
            parts.append(cells(line[pos:], color))
        body.append('<div class="row">' + "".join(parts) + "</div>")
    return TEMPLATE.replace("@BODY@", "\n".join(body))


TEMPLATE = """<!doctype html>
<html><head><meta charset="utf-8"><style>
  body { margin: 0; background: #101014; }
  .term {
    display: inline-block;
    padding: 64px 72px;
    font-family: "DejaVu Sans Mono", "Noto Sans Mono CJK JP", "Noto Color Emoji", monospace;
    font-size: 42px;
    line-height: 1.0;
    color: #C8CCD4;
  }
  .row { white-space: pre; height: 1.164em; line-height: 1.164em; }
  .c1, .c2 { display: inline-block; overflow: visible; vertical-align: top; }
  .c1 { width: 1ch; }
  .c2 { width: 2ch; }
</style></head><body><div class="term">
@BODY@
</div></body></html>
"""

if __name__ == "__main__":
    sys.stdout.write(convert(sys.stdin.read()))
