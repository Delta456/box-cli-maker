#!/usr/bin/env bash
# Regenerate every showcase screenshot in img/ (and the hero) from source.
set -euo pipefail
cd "$(dirname "$0")/../.."

subjects=(single single_double double double_single bold round hidden classic block
  top bottom top_center top_right bottom_center bottom_right
  inside_left inside_right left right)

for s in "${subjects[@]}"; do
  ./scripts/screenshots/shoot.sh "go run ./scripts/screenshots/showcase $s" "img/$s.png" --plain
  echo "img/$s.png"
done

./scripts/screenshots/shoot.sh "go run ./examples/readme/" img/hero.png
echo "img/hero.png"
