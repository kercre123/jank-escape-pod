#!/bin/bash

botNum=$1

function sstCmd() {
STT_TFLITE_DELEGATE=gpu ./stt/stt --model ./stt/model.tflite --scorer ./stt/large_vocabulary.scorer --audio /tmp/${botNum}voice.wav
}

function ffmpegCmd() {
ffmpeg -y -i /tmp/${botNum}voice.ogg /tmp/${botNum}voice.wav
}

function doSttARM() {
sleep 0.8
cd ../
rm -r /tmp/*voice.wav
ffmpegCmd
sstCmd > /tmp/${botNum}utterance1
touch /tmp/sttDone
sleep 0.5
rm -f /tmp/*utterance*
rm -f /tmp/*voice.wav
rm -f /tmp/*sttDone
}

function doSttAMD64() {
sleep 0.8
cd ../
rm -r /tmp/*voice.wav
ffmpegCmd
sstCmd > /tmp/${botNum}utterance1
sleep 0.5
ffmpegCmd
sstCmd > /tmp/${botNum}utterance2
ffmpegCmd
sstCmd > /tmp/${botNum}utterance3
ffmpegCmd
sstCmd > /tmp/${botNum}utterance4
sleep 0.5
rm -rf /tmp/*utterance*
rm -rf /tmp/*voice.wav
}

UNAME=$(uname -a)
if [[ "${UNAME}" == *"x86_64"*  ]]; then
  doSttAMD64 &
else
  doSttARM &
fi
