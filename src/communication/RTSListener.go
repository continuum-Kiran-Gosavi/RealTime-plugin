package communication

import (
	"fmt"
	nh "net/http"
	"strings"
	"sync"

	//wmi "github.com/ContinuumLLC/platform-common-lib/src/plugin/wmi"
	webClient "github.com/ContinuumLLC/platform-common-lib/src/webClient"
)

//httpConnectionState This flag is set to true if the plugin is able to communicate to RTS Microservice or else set it to false
var httpConnectionState *connectionState

//RTSListener handles put request
type RTSListener struct {
	dep dependancies
	//de  wmi.Wrapper
}

type dependancies struct {
	webClient.ClientFactoryImpl
	webClient.HTTPClientFactoryImpl
}
type connectionState struct {
	online bool
	mu     sync.Mutex
}

func (r *connectionState) isOnline() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.online
}

func (r *connectionState) setOnlineState(pFlag bool) (success bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	success = pFlag != r.online
	r.online = pFlag
	return
}

//SendMessage handles put call to RTS MS
func (rl *RTSListener) SendMessage(msg string) error {
	fmt.Printf("Sending data: %s", msg)
	err := rl.onlinePost(msg)
	if err != nil {
		return err
	}

	return nil
}

func (rl *RTSListener) onlinePost(msg string) error {
	//fmt.Printf("In OnlinePost ")
	url := "http://internal-realtimealb-1281021573.us-east-1.elb.amazonaws.com/realtime/hostname/realtimeMS-B/endpointinfo/endpointid/1/realtimeid/1"
	request, err := nh.NewRequest(nh.MethodPut, url, strings.NewReader(msg))
	if err != nil {
		return err
	}

	client := rl.dep.GetClientServiceByType(webClient.TlsClient, webClient.ClientConfig{IdleConnTimeoutMinute: 1,
		MaxIdleConns:                1,
		MaxIdleConnsPerHost:         1,
		TimeoutMinute:               5,
		DialKeepAliveSecond:         5,
		DialTimeoutSecond:           5,
		TLSHandshakeTimeoutSecond:   5,
		ExpectContinueTimeoutSecond: 5})

	_, err = client.Do(request)
	if err != nil {
		return err
	}
	fmt.Println("Successful put methode call")
	return nil

}
