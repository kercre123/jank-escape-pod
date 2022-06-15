#!/bin/bash

if [[ -f armarch ]]; then
  touch slowsys
fi

botNum=$1

function sstCmd() {
STT_TFLITE_DELEGATE=gpu ./stt/stt --model ./stt/model.tflite --scorer ./stt/large_vocabulary.scorer --audio /tmp/${botNum}voice.wav
}

function ffmpegCmd() {
ffmpeg -y -i /tmp/${botNum}voice.ogg /tmp/${botNum}voice.wav
}

function doSttSlow() {
sleep 2
cd ../
rm -rf /tmp/${botNum}voice.wav
ffmpegCmd
sstCmd > /tmp/${botNum}utterance1
touch /tmp/${botNum}sttDone
sleep 0.5
rm -f /tmp/${botNum}utterance*
rm -f /tmp/${botNum}voice.wav
rm -f /tmp/${botNum}sttDone
}

function doSttFast() {
sleep 0.8
cd ../
rm -rf /tmp/${botNum}voice.wav
ffmpegCmd
sstCmd > /tmp/${botNum}utterance1
sleep 0.5
ffmpegCmd
sstCmd > /tmp/${botNum}utterance2
ffmpegCmd
sstCmd > /tmp/${botNum}utterance3
ffmpegCmd
sstCmd > /tmp/${botNum}utterance4
touch /tmp/${botNum}sttDone
sleep 0.5
rm -rf /tmp/${botNum}utterance*
rm -rf /tmp/${botNum}voice.wav
rm -rf /tmp/${botNum}sttDone
}

if [[ -f slowsys ]]; then
  doSttSlow &
else
  doSttFast &
fi
