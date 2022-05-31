rm -rf output.mov
perf stat -I 500 \
-e task-clock,cycles,instructions \
-e mem-loads,mem-stores \
-e block:block_rq_issue,block:block_rq_complete \
-e kmem:kfree,kmem:kmalloc \
ffmpeg -i ./input.mov -i ./watermark.png -filter_complex overlay=main_w/2-overlay_w/2:main_h/2-overlay_h/2 output.mov -hide_banner -loglevel error

rm -rf output.mov
perf mem record \
ffmpeg -i ./input.mov -i ./watermark.png -filter_complex overlay=main_w/2-overlay_w/2:main_h/2-overlay_h/2 output.mov -hide_banner -loglevel error