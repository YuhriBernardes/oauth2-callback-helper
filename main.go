package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/enescakir/emoji"
	"github.com/mgutz/ansi"
	"github.com/yuhribernardes/oauth2-callback-helper/internal/server"
)

var (
	stopChan = make(chan os.Signal, 1)
)

func main() {
	logger := log.Default()

	signal.Notify(stopChan, syscall.SIGABRT, syscall.SIGINT, syscall.SIGKILL)

	srvr := server.Create(server.Options{
		ShowQuery: true,
	})

	if err := srvr.Start(); err != nil {
		logger.Println(ansi.Red, emoji.Warning, emoji.Fire, "Error when starting server", err, emoji.Fire, emoji.Warning, ansi.Reset)
	}

	logger.Println(emoji.CheckMarkButton, ansi.Green+"Server runnning at", ansi.ColorCode("green:black+b")+srvr.Addr()+ansi.Reset)
	logger.Println(emoji.CheckMarkButton, "You can use", ansi.ColorCode("white:black+b")+"http://"+srvr.Addr()+ansi.Reset, "as your callback")

	<-stopChan

	logger.Println(ansi.Yellow, emoji.Warning, "Shutting down server", ansi.Reset)

	if err := srvr.Stop(); err != nil {
		logger.Println(ansi.Red, emoji.Warning, emoji.Fire, "Error when shutting down server", err, emoji.Fire, emoji.Warning)
	}
}
