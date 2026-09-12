#!/usr/bin/env bash
# Render a command's ANSI output to a framed terminal-window PNG.
#
#   usage: shoot.sh "<command>" <output.png> [--plain]
#
# --plain skips the window title bar (used for the compact showcase
# specimens; the hero keeps the full window chrome).
#
# Requires: firefox, imagemagick, python3, and fontconfig with a mono font
# plus Noto Color Emoji / Noto Sans Mono CJK for emoji and wide characters.
set -euo pipefail

cmd=$1
out=$2
here=$(cd "$(dirname "$0")" && pwd)
dir=$(mktemp -d)
trap 'rm -rf "$dir"' EXIT

CLICOLOR_FORCE=1 bash -c "$cmd" > "$dir/o.ans"
python3 "$here/ans2html.py" < "$dir/o.ans" > "$dir/o.html"
mkdir "$dir/ffprof"
firefox --headless --profile "$dir/ffprof" --screenshot "$dir/full.png" \
  --window-size=2600,2200 "file://$dir/o.html" 2>/dev/null

pad=52
if [ "${3:-}" = "--plain" ]; then
  pad=40
fi
convert "$dir/full.png" -trim +repage -bordercolor '#101014' -border "$pad" "$dir/body.png"
W=$(identify -format %w "$dir/body.png")
if [ "${3:-}" = "--plain" ]; then
  cp "$dir/body.png" "$dir/win.png"
else
  convert -size "${W}x92" xc:'#17171d' \
    -fill '#FF5F57' -draw "circle 46,46 46,31" \
    -fill '#FEBC2E' -draw "circle 92,46 92,31" \
    -fill '#28C840' -draw "circle 138,46 138,31" "$dir/bar.png"
  convert "$dir/bar.png" "$dir/body.png" -append "$dir/win.png"
fi
H=$(identify -format %h "$dir/win.png")
convert -size "${W}x${H}" xc:none -fill white \
  -draw "roundrectangle 0,0,$((W - 1)),$((H - 1)),28,28" "$dir/mask.png"
convert "$dir/win.png" "$dir/mask.png" -alpha off -compose CopyOpacity -composite "$dir/rounded.png"
convert "$dir/rounded.png" -stroke '#32323e' -strokewidth 2 -fill none \
  -draw "roundrectangle 1,1,$((W - 2)),$((H - 2)),28,28" "$out"
