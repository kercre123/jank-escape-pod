#!/bin/bash

function sstCmd() {
STT_TFLITE_DELEGATE=gpu ./stt/stt --model ./stt/model.tflite --scorer ./stt/large_vocabulary.scorer --audio /tmp/voice.wav
}

function ffmpegCmd() {
ffmpeg -y -i /tmp/voice.ogg /tmp/voice.wav
}

function doSttAARCH64() {
sleep 0.8
cd ../
rm -r /tmp/voice.wav
ffmpegCmd
sstCmd > /tmp/utterance1
sleep 0.5
rm -rf /tmp/utterance*
rm -rf /tmp/voice.wav
}

function doSttAMD64() {
sleep 0.8
cd ../
rm -r /tmp/voice.wav
ffmpegCmd
sstCmd > /tmp/utterance1
sleep 0.5
ffmpegCmd
sstCmd > /tmp/utterance2
ffmpegCmd
sstCmd > /tmp/utterance3
ffmpegCmd
sstCmd > /tmp/utterance4
sleep 0.5
rm -rf /tmp/utterance*
rm -rf /tmp/voice.wav
}

if [[ $(arch) == "aarch64" ]]; then
  doSttAARCH64 &
else
  doSttAMD64 &
fi
