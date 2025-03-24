package go_stomp_websocket

import (
	"strings"
)

const (
	Receipt        = "receipt"
	Id             = "id"
	ReceiptId      = "receipt-id"
	Subscription_h = "subscription"
	Message        = "message"
)

const (
	// Connect commands.
	CONNECT = "CONNECT"

	// Client commands.
	SUBSCRIBE   = "SUBSCRIBE"
	UNSUBSCRIBE = "UNSUBSCRIBE"
	DISCONNECT  = "DISCONNECT"

	// Server commands.
	MESSAGE = "MESSAGE"
	RECEIPT = "RECEIPT"
	ERROR   = "ERROR"
)

type Frame struct {
	Command string
	Headers []string
	Body    string
}

func CreateFrame(command string, headers []string) *Frame {
	frame := &Frame{
		Command: command,
		Headers: headers,
	}
	return frame
}

func ReadFrame(data []byte) *Frame {
	frame := &Frame{}
	s := string(data)[3 : len(data)-2]
	s = strings.ReplaceAll(s, "\\"+"n", "\n")
	s = strings.ReplaceAll(s, "\\"+"\"", "\"")
	s = strings.ReplaceAll(s, "\\"+"u0000", "\u0000")
	sArray := strings.Split(s, "\n")
	frame.Command = sArray[0]
	isBody := false
	for i := 1; i < len(sArray); i++ {
		//read headers
		if sArray[i] != "" && !isBody {
			frame.Headers = append(frame.Headers, sArray[i])
		} else {
			//read body
			isBody = true
			frame.Body = strings.Trim(sArray[i], "\u0000")
		}
	}
	return frame
}

func (frame *Frame) Bytes() []byte {
	var resultSlice []string
	resultSlice = append(resultSlice, "[\"")
	resultSlice = append(resultSlice, frame.Command+"\\n")
	for _, header := range frame.Headers {
		resultSlice = append(resultSlice, header+"\\n")
	}
	resultSlice = append(resultSlice, "\\n")
	resultSlice = append(resultSlice, frame.Body)
	resultSlice = append(resultSlice, "\\u0000\"]")
	result := strings.Join(resultSlice, "")
	return []byte(result)
}

func (frame *Frame) Contains(header string) (string, bool) {
	for _, frameHeader := range frame.Headers {
		index := strings.Index(frameHeader, ":")
		key := frameHeader[:index]
		if key == header {
			value := frameHeader[index+1:]
			return value, true
		}
	}
	return "", false
}
