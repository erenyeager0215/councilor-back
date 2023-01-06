package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	//引数にとった各writerへ書き込まれるwriterを戻り値として返す
	multiLogFile := io.MultiWriter(os.Stdout, logfile)

	//logのフォーマットを指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//引数にio.writerをしてい
	// logの出力先を指定
	log.SetOutput(multiLogFile)
}
