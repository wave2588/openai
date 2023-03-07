package main

import (
	"context"
	"fmt"
	"github.com/wave2588/openai"
	"time"
)

const ASAK = "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"

func main() {

	ctx := context.Background()

	text := "你好"
	chat := openai.NewChatGPT(ASAK, "123321", 180*time.Second)
	defer chat.Close()
	//resp, err := chat.ChatWithContext(text)
	//if err != nil {
	//	fmt.Println(111, err)
	//	return
	//}
	//fmt.Println("result--->: ", resp)

	resp, err := chat.ChatV2(ctx, text)
	if err != nil {
		fmt.Println(111, err)
		return
	}
	fmt.Println("result--->: ", resp)
}
