#!bin/bash

INPUT_VIDEO=$1
MARK_IMAGE=$2
OUTPUT_VIDEO=$3

ffmpeg \
-i "$INPUT_VIDEO" \
-i "$MARK_IMAGE" \
-filter_complex "overlay=main_w/2-overlay_w/2:main_h/2-overlay_h/2" \
-hide_banner -loglevel error \
$OUTPUT_VIDEO

