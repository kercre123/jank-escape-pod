package noop

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	pb "github.com/digital-dream-labs/api/go/chipperpb"
	"github.com/digital-dream-labs/chipper/pkg/vtt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var intent pb.IntentResponse
var matched int = 0
var intentNum int = 0
var botNum int = 0

func IntentPass(req *vtt.IntentRequest, intentThing string) (*vtt.IntentResponse, error) {
	intent = pb.IntentResponse{
		IsFinal: true,
		IntentResult: &pb.IntentResult{
			Action: intentThing,
		},
	}
	if err := req.Stream.Send(&intent); err != nil {
		return nil, err
	}
	r := &vtt.IntentResponse{
		Intent: &intent,
	}
	log.Println("Intent Sent: " + intentThing)
	return r, nil
}

func processTextAll(req *vtt.IntentRequest, voiceText string, listOfLists [][]string, intentList []string) {
	intentNum = 0
	for _, b := range listOfLists {
		for _, c := range b {
			if strings.Contains(voiceText, c) {
				matched = 1
				IntentPass(req, intentList[intentNum])
				break
			}
		}
		intentNum = intentNum + 1
	}
	return
}

func (s *Server) ProcessIntent(req *vtt.IntentRequest) (*vtt.IntentResponse, error) {
	var finished1 string
	var finished2 string
	var finished3 string
	var finished4 string
	var transcribedText string
	matched = 0
	botNum = botNum + 1
	log.Println(botNum)
	log.Println(strconv.Itoa(botNum))
	log.Println(strconv.Itoa(botNum))
	f, err := os.Create("/tmp/" + strconv.Itoa(botNum) + "voice.ogg")
	check(err)
	cmd1 := exec.Command("/bin/bash", "../stt.sh", strconv.Itoa(botNum))
	data := []byte{}
	data = append(data, req.FirstReq.InputAudio...)
	cmd1.Run()
	f.Write(data)
	for {
		chunk, err := req.Stream.Recv()
		if err != nil {
			if err == io.EOF {
				IntentPass(req, "intent_system_noaudio")
				break
			}

		}
		data = append(data, chunk.InputAudio...)
		f.Write(chunk.InputAudio)
		fileBytes1, _ := ioutil.ReadFile("/tmp/" + strconv.Itoa(botNum) + "utterance1")
		transcribedText1 := strings.TrimSpace(string(fileBytes1))
		if _, err := os.Stat("/tmp/" + strconv.Itoa(botNum) + "utterance1"); err == nil {
			finished1 = transcribedText1
		}
		if _, err := os.Stat("./armarch"); err == nil {
			if _, err := os.Stat("/tmp/" + strconv.Itoa(botNum) + "sttDone"); err == nil {
				transcribedText = finished1
				log.Println("aarch: Speech has stopped, transcribed text is: " + finished1)
				break
			}
		} else {
			fileBytes2, _ := ioutil.ReadFile("/tmp/" + strconv.Itoa(botNum) + "utterance2")
			transcribedText2 := strings.TrimSpace(string(fileBytes2))
			if _, err := os.Stat("/tmp/" + strconv.Itoa(botNum) + "utterance2"); err == nil {
				finished2 = transcribedText2
				if finished1 == finished2 {
					transcribedText = finished2
					log.Println("2: Speech has stopped, transcribed text is: " + finished2)
					break
				}
			}
			fileBytes3, _ := ioutil.ReadFile("/tmp/" + strconv.Itoa(botNum) + "utterance3")
			transcribedText3 := strings.TrimSpace(string(fileBytes3))
			if _, err := os.Stat("/tmp/" + strconv.Itoa(botNum) + "utterance3"); err == nil {
				finished3 = transcribedText3
				if finished2 == finished3 {
					transcribedText = finished3
					log.Println("3: Speech has stopped, transcribed text is: " + finished3)
					break
				}
			}
			fileBytes4, _ := ioutil.ReadFile("/tmp/" + strconv.Itoa(botNum) + "utterance4")
			transcribedText4 := strings.TrimSpace(string(fileBytes4))
			if _, err := os.Stat("/tmp/" + strconv.Itoa(botNum) + "utterance4"); err == nil {
				finished4 = transcribedText4
				if finished3 == finished4 {
					transcribedText = finished4
					log.Println("4: Speech has stopped, transcribed text is: " + finished4)
					break
				}
			}
		}
	}

	processTextAll(req, transcribedText, matchListList, intentsList)
	botNum = botNum - 1

	if matched == 0 {
		log.Println("No intent was matched.")
		IntentPass(req, "intent_system_noaudio")
	}
	return nil, nil
}
