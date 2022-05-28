package noop

import (
	pb "github.com/digital-dream-labs/api/go/chipperpb"
	"github.com/digital-dream-labs/chipper/pkg/vtt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var intent pb.IntentResponse

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

func (s *Server) ProcessIntent(req *vtt.IntentRequest) (*vtt.IntentResponse, error) {
	var finished1 string
	var finished2 string
	var finished3 string
	var finished4 string
	var transcribedText string
//	var aarchWait int = 0
	f, err := os.Create("/tmp/voice.ogg")
	check(err)
	cmd1 := exec.Command("/bin/bash", "../stt.sh")
	data := []byte{}
	data = append(data, req.FirstReq.InputAudio...)
	cmd1.Run()
	var intent pb.IntentResponse
	f.Write(data)
	for {
		chunk, err := req.Stream.Recv()
		if err != nil {
			if err == io.EOF {
				intent = pb.IntentResponse{
					IsFinal: true,
					IntentResult: &pb.IntentResult{
						Action: "intent_system_noaudio",
					},
				}
				r := &vtt.IntentResponse{
					Intent: &intent,
				}
				return r, nil
				break
			}

		}
		data = append(data, chunk.InputAudio...)
		f.Write(chunk.InputAudio)
		fileBytes1, err := ioutil.ReadFile("/tmp/utterance1")
		if err != nil {
			//nothing
		}
		transcribedText1 := strings.TrimSpace(string(fileBytes1))
		if _, err := os.Stat("/tmp/utterance1"); err == nil {
			finished1 = transcribedText1
		}
		if _, err := os.Stat("./armarch"); err == nil {
			if _, err := os.Stat("/tmp/sttDone"); err == nil {
				transcribedText = finished1
                                log.Println("aarch: Speech has stopped, transcribed text is: " + finished1)
                                break
			}
		} else {
		fileBytes2, err := ioutil.ReadFile("/tmp/utterance2")
		if err != nil {
			//nothing
		}
		transcribedText2 := strings.TrimSpace(string(fileBytes2))
		if _, err := os.Stat("/tmp/utterance2"); err == nil {
			finished2 = transcribedText2
			if finished1 == finished2  {
				transcribedText = finished2
				log.Println("2: Speech has stopped, transcribed text is: " + finished2)
				break
			}
		}
		fileBytes3, err := ioutil.ReadFile("/tmp/utterance3")
		if err != nil {
			//nothing
		}
		transcribedText3 := strings.TrimSpace(string(fileBytes3))
		if _, err := os.Stat("/tmp/utterance3"); err == nil {
			finished3 = transcribedText3
			if finished2 == finished3  {
				transcribedText = finished3
				log.Println("3: Speech has stopped, transcribed text is: " + finished3)
				break
			}
		}
		fileBytes4, err := ioutil.ReadFile("/tmp/utterance4")
		if err != nil {
			//nothing
		}
		transcribedText4 := strings.TrimSpace(string(fileBytes4))
		if _, err := os.Stat("/tmp/utterance4"); err == nil {
			finished4 = transcribedText4
			if finished3 == finished4  {
				transcribedText = finished4
				log.Println("4: Speech has stopped, transcribed text is: " + finished4)
				break
			}
		}
		}
	}
	if (strings.Contains(transcribedText, "good r") || strings.Contains(transcribedText, "awesome") || strings.Contains(transcribedText, "also") || strings.Contains(transcribedText, "as some") || strings.Contains(transcribedText, "of them") || strings.Contains(transcribedText, "battle") || strings.Contains(transcribedText, "t rob") || strings.Contains(transcribedText, "the ro")) {
		IntentPass(req, "intent_imperative_praise")
	} else if (strings.Contains(transcribedText, "bad") || strings.Contains(transcribedText, "that ro") || strings.Contains(transcribedText, "ad ro") || strings.Contains(transcribedText, "a root")) {
		IntentPass(req, "intent_imperative_abuse")
	} else if (strings.Contains(transcribedText, "eye color") || strings.Contains(transcribedText, "colo") || strings.Contains(transcribedText, "i call her") || strings.Contains(transcribedText, "i foller") || strings.Contains(transcribedText, "icolor") || strings.Contains(transcribedText, "ecce") || strings.Contains(transcribedText, "erior") || strings.Contains(transcribedText, "ichor") || strings.Contains(transcribedText, "agricola") || strings.Contains(transcribedText, "change")) {
		IntentPass(req, "intent_imperative_eyecolor")
	} else if (strings.Contains(transcribedText, "older") || strings.Contains(transcribedText, "how old") || strings.Contains(transcribedText, "old are you") || strings.Contains(transcribedText, "old or yo")) {
		IntentPass(req, "intent_character_age")
	} else if (strings.Contains(transcribedText, "start") || strings.Contains(transcribedText, "plor") || strings.Contains(transcribedText, "owing") || strings.Contains(transcribedText, "tailoring") || strings.Contains(transcribedText, "oding") || strings.Contains(transcribedText, "oring") || strings.Contains(transcribedText, "pling")) {
		IntentPass(req, "intent_explore_start")
	} else if (strings.Contains(transcribedText, "charge") || strings.Contains(transcribedText, "home") || strings.Contains(transcribedText, "go to your") || strings.Contains(transcribedText, "church")) {
		IntentPass(req, "intent_system_charger")
	} else if (strings.Contains(transcribedText, "flee") || strings.Contains(transcribedText, "sleep") || strings.Contains(transcribedText, "sheep")) {
		IntentPass(req, "intent_system_sleep")
	} else if (strings.Contains(transcribedText, "morning") || strings.Contains(transcribedText, "mourning") || strings.Contains(transcribedText, "mooning") || strings.Contains(transcribedText, "it bore")) {
		IntentPass(req, "intent_greeting_goodmorning")
	} else if (strings.Contains(transcribedText, "night") || strings.Contains(transcribedText, "might")) {
		IntentPass(req, "intent_greeting_goodnight")
	} else if (strings.Contains(transcribedText, "time is it") || strings.Contains(transcribedText, "the time") || strings.Contains(transcribedText, "what time")) {
		IntentPass(req, "intent_clock_time")
	} else if (strings.Contains(transcribedText, "good bye") || strings.Contains(transcribedText, "good by") || strings.Contains(transcribedText, "good buy") || strings.Contains(transcribedText, "goodbye")) {
		IntentPass(req, "intent_greeting_goodbye")
	} else if (strings.Contains(transcribedText, "fireworks") || strings.Contains(transcribedText, "new year") || strings.Contains(transcribedText, "happy new") || strings.Contains(transcribedText, "happy to") || strings.Contains(transcribedText, "have been") || strings.Contains(transcribedText, "i now you") || strings.Contains(transcribedText, "no year") || strings.Contains(transcribedText, "enee") || strings.Contains(transcribedText, "i never") || strings.Contains(transcribedText, "knew her")) {
		IntentPass(req, "intent_seasonal_happynewyear")
	} else if (strings.Contains(transcribedText, "he holds") || strings.Contains(transcribedText, "christmas") || strings.Contains(transcribedText, "behold") || strings.Contains(transcribedText, "holiday")) {
		IntentPass(req, "intent_seasonal_happyholidays")
	} else if (strings.Contains(transcribedText, "hello") || strings.Contains(transcribedText, "are you") || strings.Contains(transcribedText, "high") || strings.Contains(transcribedText, "hi") || strings.Contains(transcribedText, "below") || strings.Contains(transcribedText, "little") || strings.Contains(transcribedText, "follow")) {
		IntentPass(req, "intent_greeting_hello")
	} else if (strings.Contains(transcribedText, "in intellect") || strings.Contains(transcribedText, "fine in electa") || strings.Contains(transcribedText, "in alexa") || strings.Contains(transcribedText, "in an elect") || strings.Contains(transcribedText, "to alexa") || strings.Contains(transcribedText, "in angelica")) {
		IntentPass(req, "intent_amazon_signin")
	} else if (strings.Contains(transcribedText, "in outlet") || strings.Contains(transcribedText, "i now of elea") || strings.Contains(transcribedText, "out alexa") || strings.Contains(transcribedText, "out of ale")) {
		IntentPass(req, "intent_amazon_signin")
	} else if (strings.Contains(transcribedText, "love") || strings.Contains(transcribedText, "dove")) {
		IntentPass(req, "intent_imperative_love")
	} else {
		log.Println("Did not match an intent.")
		log.Println("Intent Sent: intent_system_noaudio")
	}
	intent = pb.IntentResponse{
		IsFinal: true,
		IntentResult: &pb.IntentResult{
			Action: "intent_system_noaudio",
		},
	}
	if err := req.Stream.Send(&intent); err != nil {
		return nil, err
	}
	r := &vtt.IntentResponse{
		Intent: &intent,
	}
	return r, nil
}

