package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/digital-dream-labs/vector-cloud/internal/clad/cloud"
	"github.com/digital-dream-labs/vector-cloud/internal/cloudproc"
	"github.com/digital-dream-labs/vector-cloud/internal/config"
	"github.com/digital-dream-labs/vector-cloud/internal/ipc"
	"github.com/digital-dream-labs/vector-cloud/internal/jdocs"
	"github.com/digital-dream-labs/vector-cloud/internal/log"
	"github.com/digital-dream-labs/vector-cloud/internal/logcollector"
	"github.com/digital-dream-labs/vector-cloud/internal/robot"
	"github.com/digital-dream-labs/vector-cloud/internal/token"
	"github.com/digital-dream-labs/vector-cloud/internal/voice"

	"github.com/gwatts/rootcerts"
)

var checkDataFunc func() error // overwritten by platform_linux.go
var certErrorFunc func() bool  // overwritten by cert_error_dev.go, determines if error should cause exit
var platformOpts []cloudproc.Option


const awesomeCert = `-----BEGIN CERTIFICATE-----
MIIEijCCA3KgAwIBAgIUaI86chQQ0xvRBvpNjLIUDuaCmJwwDQYJKoZIhvcNAQEL
BQAwZjELMAkGA1UEBhMCVVMxDTALBgNVBAgMBElvd2ExEzARBgNVBAcMCkRlcyBN
b2luZXMxEDAOBgNVBAoMB0F3ZXNvbWUxITAfBgNVBAMMGGFua2l0ZXN0YW5raXRl
cy5tb29vLmNvbTAeFw0yMjA1MjMxNDA4MzZaFw0zMjA1MjAxNDA4MzZaMGYxCzAJ
BgNVBAYTAlVTMQ0wCwYDVQQIDARJb3dhMRMwEQYDVQQHDApEZXMgTW9pbmVzMRAw
DgYDVQQKDAdBd2Vzb21lMSEwHwYDVQQDDBhhbmtpdGVzdGFua2l0ZXMubW9vby5j
b20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDCTRPe3ETUkvj4UA/Q
yxXMk4HXRX2NUPGcfVbE3KaLCwCpTlHK8A38oBgUGvOlxuoujyhamyQq2iaL5Fzm
z+qbtbDepjt4c2KoGkdxAvbSBcBEm51NbUGTrNqhKOZWblc3XBjMe6XApT344Iu2
PWC/Ptg77rFVCA/a6ekv3idxrqEMteZKArYR80ffY+2rLXzDXmBP3FZ20Ga+/0tk
ORlIVvitoOg3giddjiRIQMrO8tl8EFqEqxyyiCpx465Gp4NSOgE2dZ0WTDkWtKxN
Jzb3IHuX1BL9pZXcUH9DLKGNnsxiRALP46zx2eJpWZy0RsaR65/6Mt+Lx7jmG02s
kLnnAgMBAAGjggEuMIIBKjAdBgNVHQ4EFgQUmV6fefbJoyOE07k8RmNh4E9fsikw
gaMGA1UdIwSBmzCBmIAUmV6fefbJoyOE07k8RmNh4E9fsimhaqRoMGYxCzAJBgNV
BAYTAlVTMQ0wCwYDVQQIDARJb3dhMRMwEQYDVQQHDApEZXMgTW9pbmVzMRAwDgYD
VQQKDAdBd2Vzb21lMSEwHwYDVQQDDBhhbmtpdGVzdGFua2l0ZXMubW9vby5jb22C
FGiPOnIUENMb0Qb6TYyyFA7mgpicMAwGA1UdEwQFMAMBAf8wCwYDVR0PBAQDAgL8
MCMGA1UdEQQcMBqCGGFua2l0ZXN0YW5raXRlcy5tb29vLmNvbTAjBgNVHRIEHDAa
ghhhbmtpdGVzdGFua2l0ZXMubW9vby5jb20wDQYJKoZIhvcNAQELBQADggEBAKV2
YenCiAx7P0TWfcQ955xLt3JLVn4ZfyETcyWDVYoi8oN/z4xIo4awEg/nfJpdcDJD
CftjUsk/S7SRvIPmYWs2E96linC4tSzHXX6S+tSo0PCsyub/E5iCt0mrEsMRr5CU
BXSJvhvOb/QNxuuwL0Xa9RvSqq9orTDbGuHStTKFAoMTbZ16gpvRuDgWg7fPIYrx
FIemBhUOuQOvj8U2SjoJsTeTZNDNWKgijSAQvkwFWSEWJsY4VkMc0FuR6Lqu3C+E
f19+BAqagUjCSKJ4Y6jfj2pMI32P2bhmM5gj5C9uf9QxeT1lIMmN6d3zgSvPJ8TC
U3tNlATAvpOg02N0vz4=
-----END CERTIFICATE-----`

               var pool = rootcerts.ServerCertPool()
        var _ = pool.AppendCertsFromPEM([]byte(awesomeCert))

func getSocketWithRetry(name string, client string) ipc.Conn {
	for {
		sock, err := ipc.NewUnixgramClient(name, client)
		if err != nil {
			log.Println("Couldn't create socket", name, "- retrying:", err)
			time.Sleep(5 * time.Second)
		} else {
			return sock
		}
	}
}

func getHTTPClient() *http.Client {
	// Create a HTTP client with given CA cert pool so we can use https on device
//		pool := rootcerts.ServerCertPool()
//        _ = pool.AppendCertsFromPEM([]byte(awesomeCert))
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: rootcerts.ServerCertPool(),
				InsecureSkipVerify: true,
			},
		},
	}
}

func testReader(serv ipc.Server, send voice.MsgSender) {
	for conn := range serv.NewConns() {
		go func(conn ipc.Conn) {
			for {
				msg := conn.ReadBlock()
				if msg == nil || len(msg) == 0 {
					conn.Close()
					return
				}
				var cmsg cloud.Message
				if err := cmsg.Unpack(bytes.NewBuffer(msg)); err != nil {
					log.Println("Test reader unpack error:", err)
					continue
				}
				send.Send(&cmsg)
			}
		}(conn)
	}
}

func main() {

	log.Println("Starting up")

	robot.InstallCrashReporter(log.Tag)

	// if we want to error, we should do it after we get socket connections, to make sure
	// vic-anim is running and able to handle it
	var tryErrorFunc bool
	if checkDataFunc != nil {
		if err := checkDataFunc(); err != nil {
			log.Println("CLOUD DATA VERIFICATION ERROR:", err)
			log.Println("(this should not happen on any DVT3 or later robot)")
			tryErrorFunc = true
		} else {
			log.Println("Cloud data verified")
		}
	}

	signalHandler()

	// don't yet have control over process startup on DVT2, set these as default
	test := false

	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "enable verbose logging")
	// var test bool
	// flag.BoolVar(&test, "test", false, "enable test channel")

	ms := flag.Bool("ms", false, "force microsoft handling on the server end")
	lex := flag.Bool("lex", false, "force amazon handling on the server end")

	awsRegion := flag.String("region", "us-west-2", "AWS Region")

	flag.Parse()

	micSock := getSocketWithRetry(ipc.GetSocketPath("mic_sock"), "cp_mic")
	defer micSock.Close()
	aiSock := getSocketWithRetry(ipc.GetSocketPath("ai_sock"), "cp_ai")
	defer aiSock.Close()

	// now that we have connection, we can error if necessary
	if tryErrorFunc && certErrorFunc != nil && certErrorFunc() {
		return
	}

	// set up test channel if flags say we should
	var testRecv *voice.Receiver
	if test {
		testSock, err := ipc.NewUnixgramServer(ipc.GetSocketPath("cp_test"))
		if err != nil {
			log.Println("Server create error:", err)
		}
		defer testSock.Close()

		var testSend voice.MsgIO
		testSend, testRecv = voice.NewMemPipe()
		go testReader(testSock, testSend)
		log.Println("Test channel created")
	}
	log.Println("Sockets successfully created")

	voice.SetVerbose(verbose)
	receiver := voice.NewIpcReceiver(micSock, nil)

	process := &voice.Process{}
	process.AddReceiver(receiver)
	if testRecv != nil {
		process.AddTestReceiver(testRecv)
	}
	process.AddIntentWriter(&voice.IPCMsgSender{Conn: aiSock})
	voiceOpts := []voice.Option{voice.WithChunkMs(120), voice.WithSaveAudio(true)}
	var options []cloudproc.Option
	options = append(options, platformOpts...)
	voiceOpts = append(voiceOpts, voice.WithCompression(true))
	if *ms {
		voiceOpts = append(voiceOpts, voice.WithHandler(voice.HandlerMicrosoft))
	} else if *lex {
		voiceOpts = append(voiceOpts, voice.WithHandler(voice.HandlerAmazon))
	}

	if err := config.SetGlobal(""); err != nil {
		log.Println("Could not load server config! This is not good!:", err)
		if certErrorFunc != nil && certErrorFunc() {
			return
		}
	}

	options = append(options, cloudproc.WithVoice(process))
	options = append(options, cloudproc.WithVoiceOptions(voiceOpts...))
	tokenOpts := []token.Option{token.WithServer()}
	options = append(options, cloudproc.WithTokenOptions(tokenOpts...))
	options = append(options, cloudproc.WithJdocs(jdocs.WithServer()))

	logcollectorOpts := []logcollector.Option{logcollector.WithServer()}
	logcollectorOpts = append(logcollectorOpts, logcollector.WithHTTPClient(getHTTPClient()))
	logcollectorOpts = append(logcollectorOpts, logcollector.WithS3UrlPrefix(config.Env.LogFiles))
	logcollectorOpts = append(logcollectorOpts, logcollector.WithAwsRegion(*awsRegion))
	options = append(options, cloudproc.WithLogCollectorOptions(logcollectorOpts...))

	cloudproc.Run(context.Background(), options...)

	robot.UninstallCrashReporter()

	log.Println("All processes exited, shutting down")
}

func signalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		fmt.Println("Received SIGTERM, shutting down immediately")
		robot.UninstallCrashReporter()
		os.Exit(0)
	}()
}
