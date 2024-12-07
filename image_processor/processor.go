package image_processor

import (
	"log"
)

func StartImageProcessor() {
	messages, err := Channel.Consume("image_queue", "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for msg := range messages {
			log.Println("Processing image:", string(msg.Body))
			// Implement download and compression logic here.
		}
	}()
}
