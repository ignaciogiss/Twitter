package service

import (
	"github.com/ignaciogiss/twitter/src/domain"
	"os"
)

type TweetWriter interface {
	WriteTweet(tweet domain.Tweet )
}

type MemoryTweetWriter struct {
	tweets []domain.Tweet
}

func NewMemoryTweetWriter() *MemoryTweetWriter {

	return &MemoryTweetWriter{  }
}

func (memoryTweetWriter *MemoryTweetWriter) WriteTweet(tweet domain.Tweet ) {
	memoryTweetWriter.tweets = append( memoryTweetWriter.tweets, tweet )
}

func (memoryTweetWriter MemoryTweetWriter) GetLastSavedTweet() domain.Tweet {
	if (len(memoryTweetWriter.tweets ) == 0 ) {
		return nil
	}
	return memoryTweetWriter.tweets[ len(memoryTweetWriter.tweets ) - 1 ]
}

// FileTweetWriter
type FileTweetWriter struct {
	file *os.File
}

func NewFileTweetWriter() *FileTweetWriter {
	file, _ := os.OpenFile(
		"tweets.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)

	writer := new(FileTweetWriter)
	writer.file = file

	return writer
}

func (writer *FileTweetWriter) WriteTweet(tweet domain.Tweet) {

	go func() {
		if writer.file != nil {
			byteSlice := []byte(tweet.PrintableTweet() + "\n")
			writer.file.Write(byteSlice)
		}
	}()
}


func (fileTweetWriter FileTweetWriter) GetLastSavedTweet() domain.Tweet {
	return nil // TODO
}
