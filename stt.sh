#!/bin/bash

function sstCmd() {
STT_TFLITE_DELEGATE=gpu ./stt/stt --model ./stt/model.tflite --scorer ./stt/large_vocabulary.scorer --audio /tmp/voice.wav
}

function ffmpegCmd() {
ffmpeg -y -i /tmp/voice.ogg /tmp/voice.wav
}

function awesome() {
sleep 0.5
cd ../
rm -r /tmp/voice.wav
ffmpegCmd
sstCmd > /tmp/utterance1
sleep 0.6
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

awesome &
