#!/bin/bash

function sstCmd() {
STT_TFLITE_DELEGATE=gpu ./stt/stt --model ./stt/model.tflite --scorer ./stt/model.scorer --audio ./stt/voice.wav
}

function ffmpegCmd() {
ffmpeg -y -i /tmp/voice.ogg ./stt/voice.wav
}

function awesome() {
cd ../
rm -r ./stt/voice.wav
sleep 1
ffmpegCmd
sstCmd > utterance1
sleep 0.5
ffmpegCmd
sstCmd > utterance2
ffmpegCmd
sstCmd > utterance3
ffmpegCmd
sstCmd > utterance4
ffmpegCmd
sstCmd > utterance5
sleep 1
rm -rf ./utterance*
rm -rf ./stt/voice.wav
}

awesome &
