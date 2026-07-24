package client

import (
	"fmt"
	"time"
	"io"
	"encoding/json"
)

type StreamClient struct {
	ApiRequest *ApiRequest
	ResponseChannel chan EndpointResponse
}

func NewStreamClient()