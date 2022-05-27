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

var returnThing *vtt.IntentResponse

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (s *Server) ProcessIntent(req *vtt.IntentRequest) (*vtt.IntentResponse, error) {
	var finished1 string
	var finished2 string
	var finished3 string
	var finished4 string
	var finished5 string
	var transcribedText string
	f, err := os.Create("/tmp/test.ogg")
	check(err)
	cmd1 := exec.Command("/bin/bash", "../test.sh")
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
		fileBytes1, err := ioutil.ReadFile("../testutterance1")
		if err != nil {
			//nothing
		}
		transcribedText1 := strings.TrimSpace(string(fileBytes1))
		    if _, err := os.Stat("../testutterance1"); err == nil {
			finished1 = transcribedText1
		    }
                fileBytes2, err := ioutil.ReadFile("../testutterance2")
                if err != nil {
             		//nothing
		}
                transcribedText2 := strings.TrimSpace(string(fileBytes2))
                    if _, err := os.Stat("../testutterance2"); err == nil {
			finished2 = transcribedText2
			if finished1 == finished2  {
                                transcribedText = finished2
				log.Println("2: Speech has stopped, transcribed text is: " + finished2)
				break
			}
		}
                fileBytes3, err := ioutil.ReadFile("../testutterance3")
                if err != nil {
                        //nothing
                }
                transcribedText3 := strings.TrimSpace(string(fileBytes3))
                    if _, err := os.Stat("../testutterance3"); err == nil {
                        finished3 = transcribedText3
                        if finished2 == finished3  {
				transcribedText = finished3
                                log.Println("3: Speech has stopped, transcribed text is: " + finished3)
                                break
                        }
                }
                fileBytes4, err := ioutil.ReadFile("../testutterance4")
                if err != nil {
                        //nothing
                }
                transcribedText4 := strings.TrimSpace(string(fileBytes4))
                    if _, err := os.Stat("../testutterance4"); err == nil {
                        finished4 = transcribedText4
                        if finished3 == finished4  {
                                transcribedText = finished4
                                log.Println("4: Speech has stopped, transcribed text is: " + finished4)
                                break
                        }
                }
                fileBytes5, err := ioutil.ReadFile("../testutterance5")
                if err != nil {
                        //nothing
                }
                transcribedText5 := strings.TrimSpace(string(fileBytes5))
                    if _, err := os.Stat("../testutterance5"); err == nil {
                        finished5 = transcribedText5
                        if finished4 == finished5  {
                                transcribedText = finished5
                                log.Println("5: Speech has stopped, transcribed text is: " + finished5)
                                break
                        }
                }
}
		if (strings.Contains(transcribedText, "good") || strings.Contains(transcribedText, "awesome") || strings.Contains(transcribedText, "also") || strings.Contains(transcribedText, "as some") || strings.Contains(transcribedText, "of them") || strings.Contains(transcribedText, "battle") || strings.Contains(transcribedText, "t rob") || strings.Contains(transcribedText, "the root")) {
			intent = pb.IntentResponse{
				IsFinal: true,
				IntentResult: &pb.IntentResult{
					Action: "intent_imperative_praise",
				},
			}
			if err := req.Stream.Send(&intent); err != nil {
				return nil, err
			}
			r := &vtt.IntentResponse{
				Intent: &intent,
			}
			return r, nil
		} else if (strings.Contains(transcribedText, "bad") || strings.Contains(transcribedText, "that ro") || strings.Contains(transcribedText, "ad ro")) {
			intent = pb.IntentResponse{
				IsFinal: true,
				IntentResult: &pb.IntentResult{
					Action: "intent_imperative_abuse",
				},
			}
			if err := req.Stream.Send(&intent); err != nil {
				return nil, err
			}
			r := &vtt.IntentResponse{
				Intent: &intent,
			}
			return r, nil
		} else if (strings.Contains(transcribedText, "eye color") || strings.Contains(transcribedText, "colo") || strings.Contains(transcribedText, "i call her") || strings.Contains(transcribedText, "i foller") || strings.Contains(transcribedText, "icolor") || strings.Contains(transcribedText, "ecce")) {
			intent = pb.IntentResponse{
				IsFinal: true,
				IntentResult: &pb.IntentResult{
					Action: "intent_imperative_eyecolor",
				},
			}
			if err := req.Stream.Send(&intent); err != nil {
				return nil, err
			}
			r := &vtt.IntentResponse{
				Intent: &intent,
			}
			return r, nil
		} else if (strings.Contains(transcribedText, "older") || strings.Contains(transcribedText, "how old") || strings.Contains(transcribedText, "old are you") || strings.Contains(transcribedText, "old or yo")) {
			intent = pb.IntentResponse{
				IsFinal: true,
				IntentResult: &pb.IntentResult{
					Action: "intent_character_age",
				},
			}
			if err := req.Stream.Send(&intent); err != nil {
				return nil, err
			}
			r := &vtt.IntentResponse{
				Intent: &intent,
			}
			log.Println("old")
			return r, nil
 		} else if (strings.Contains(transcribedText, "start") || strings.Contains(transcribedText, "plor") || strings.Contains(transcribedText, "owing") || strings.Contains(transcribedText, "tailoring") || strings.Contains(transcribedText, "oding") || strings.Contains(transcribedText, "oring") || strings.Contains(transcribedText, "pling")) {
			intent = pb.IntentResponse{
				IsFinal: true,
				IntentResult: &pb.IntentResult{
					Action: "intent_explore_start",
				},
			}
			if err := req.Stream.Send(&intent); err != nil {
				return nil, err
			}
			r := &vtt.IntentResponse{
				Intent: &intent,
			}
			log.Println("old")
			return r, nil
		} else if (strings.Contains(transcribedText, "charge") || strings.Contains(transcribedText, "home") || strings.Contains(transcribedText, "go to your") || strings.Contains(transcribedText, "church")) {
			intent = pb.IntentResponse{
				IsFinal: true,
				IntentResult: &pb.IntentResult{
					Action: "intent_system_charger",
				},
			}
			if err := req.Stream.Send(&intent); err != nil {
				return nil, err
			}
			r := &vtt.IntentResponse{
				Intent: &intent,
			}
			log.Println("old")
			return r, nil
		} else if (strings.Contains(transcribedText, "flee") || strings.Contains(transcribedText, "sleep") || strings.Contains(transcribedText, "sheep")) {
			intent = pb.IntentResponse{
				IsFinal: true,
				IntentResult: &pb.IntentResult{
					Action: "intent_system_sleep",
				},
			}
			if err := req.Stream.Send(&intent); err != nil {
				return nil, err
			}
			r := &vtt.IntentResponse{
				Intent: &intent,
			}
			log.Println("old")
			return r, nil
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
                        log.Println("nothing")
                        return r, nil
	}

