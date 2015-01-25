package radio

import (
  "fmt"
  "github.com/the-experimenters/story-book-radio/barcode"
  "github.com/the-experimenters/story-book-radio/radiobot"
  "github.com/the-experimenters/story-book-radio/player"
)

func Printing() {
  fmt.Println("Radio package!")
  barcode.Read()
  radiobot.Blink()
  player.Load()
}