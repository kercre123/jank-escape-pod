#!/bin/bash

function sstCmd() {
STT_TFLITE_DELEGATE=gpu ./stt/stt --model ./stt/model.tflite --scorer ./stt/model.scorer --audio ./stt/test.wav
}

function awesome() {
sleep 1
cd ../
rm -r ./stt/test.wav
ffmpeg -y -i /tmp/test.ogg ./stt/test.wav
sstCmd > testutterance1
sleep 0.5
ffmpeg -y -i /tmp/test.ogg ./stt/test.wav
sstCmd > testutterance2
#sleep 0.3
ffmpeg -y -i /tmp/test.ogg ./stt/test.wav
sstCmd > testutterance3
#sleep 0.3
ffmpeg -y -i /tmp/test.ogg ./stt/test.wav
sstCmd > testutterance4
#sleep 0.3
ffmpeg -y -i /tmp/test.ogg ./stt/test.wav
sstCmd > testutterance5
sleep 1
rm -rf ./testutterance*
}

awesome &
