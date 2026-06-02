package raft

// THIS FILE IS FOR THE SENDING AND RECIEVING OF MESSAGES USING RPC

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"encoding/binary"
)


func sendMessage(conn net.Conn, v any) (int, error){
	// marshal v to JSON
	messageJSON, err := json.Marshal(v) // returns []byte

	if err != nil {
		fmt.Printf("failed to convert the message to JSON format: %v", err)
	}
	
	// capture the length into 4 bytes buffer or uint32
	// create the buffer
	buf := make([]byte, 4)
	
	// put the length into the buffer using bigendian encoding
	binary.BigEndian.PutUint32(buf, uint32(len(messageJSON)))

	// add length prefixing
	_, err = conn.Write(buf) // expects bytes not uint32
	
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	
	// send the message 
	_, err = conn.Write(messageJSON)

	// handle error if message fails to send
	if err != nil {
		fmt.Printf("error: %v", err)
	}


	return len(messageJSON), nil
}

func recieveMessage(conn net.Conn, v any) error {
	// create a 4 byte buffer
	buf := make([]byte, 4)

	// read straight from connection
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	// decode the length from what you read
	length := binary.BigEndian.Uint32(buf)

	// read exactly the number of bytes the length denotes
	payload := make([]byte, length)
	_, err = io.ReadFull(conn, payload)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	// unmarshall into JSON
	err = json.Unmarshal(payload, v)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	// return true for success
	return nil
}