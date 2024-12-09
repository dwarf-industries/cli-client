package components

import (
	"fmt"
	"time"

	"github.com/charmbracelet/glamour"
)

type Logo struct {
	elapsed        int
	render         string
	previousRender string
}

func (l Logo) Init() {
	for {
		l.tick()

		time.Sleep(1000 * time.Millisecond)
		if l.elapsed >= 15 {
			return
		}
		l.elapsed++
	}
}

func (l *Logo) tick() {
	if l.elapsed < 3 {
		l.render = animationPartOne()
	} else if l.elapsed < 6 {
		l.render = animationPartTwo()
	} else if l.elapsed < 9 {
		l.render = animationPartThree()
	} else if l.elapsed < 12 {
		l.render = animationPartFour()
	} else if l.elapsed > 12 {
		l.render = final()
	}

	if l.render != l.previousRender {
		l.View()
		l.previousRender = l.render
	}
}

func animationPartOne() string {
	return `
	 ▄█  ████████▄     ▄████████    ▄████████    ▄████████      ▀█████████▄  ▄██   ▄          ▄▄▄▄███▄▄▄▄      ▄████████ ███▄▄▄▄
     ███  ███   ▀███   ███    ███   ███    ███   ███    ███        ███    ███ ███   ██▄      ▄██▀▀▀███▀▀▀██▄   ███    ███ ███▀▀▀██▄
     ███▌ ███    ███   ███    █▀    ███    ███   ███    █▀         ███    ███ ███▄▄▄███      ███   ███   ███   ███    ███ ███   ███
     ███▌ ███    ███  ▄███▄▄▄       ███    ███   ███              ▄███▄▄▄██▀  ▀▀▀▀▀▀███      ███   ███   ███   ███    ███ ███   ███
     ███▌ ███    ███ ▀▀███▀▀▀     ▀███████████ ▀███████████      ▀▀███▀▀▀██▄  ▄██   ███      ███   ███   ███ ▀███████████ ███   ███
     ███  ███    ███   ███    █▄    ███    ███          ███        ███    ██▄ ███   ███      ███   ███   ███   ███    ███ ███   ███
     ███  ███   ▄███   ███    ███   ███    ███    ▄█    ███        ███    ███ ███   ███      ███   ███   ███   ███    ███ ███   ███
     █▀   ████████▀    ██████████   ███    █▀   ▄████████▀       ▄█████████▀   ▀█████▀        ▀█   ███   █▀    ███    █▀   ▀█   █▀

        ▄████████ ███▄▄▄▄   ████████▄        ▄██████▄  ███▄▄▄▄    ▄█       ▄██   ▄            ███        ▄█    █▄       ▄████████     ███
       ███    ███ ███▀▀▀██▄ ███   ▀███      ███    ███ ███▀▀▀██▄ ███       ███   ██▄      ▀█████████▄   ███    ███     ███    ███ ▀█████████▄
       ███    ███ ███   ███ ███    ███      ███    ███ ███   ███ ███       ███▄▄▄███         ▀███▀▀██   ███    ███     ███    ███    ▀███▀▀██
       ███    ███ ███   ███ ███    ███      ███    ███ ███   ███ ███       ▀▀▀▀▀▀███          ███   ▀  ▄███▄▄▄▄███▄▄   ███    ███     ███   ▀
     ▀███████████ ███   ███ ███    ███      ███    ███ ███   ███ ███       ▄██   ███          ███     ▀▀███▀▀▀▀███▀  ▀███████████     ███
       ███    ███ ███   ███ ███    ███      ███    ███ ███   ███ ███       ███   ███          ███       ███    ███     ███    ███     ███
       ███    ███ ███   ███ ███   ▄███      ███    ███ ███   ███ ███▌    ▄ ███   ███          ███       ███    ███     ███    ███     ███
       ███    █▀   ▀█   █▀  ████████▀        ▀██████▀   ▀█   █▀  █████▄▄██  ▀█████▀          ▄████▀     ███    █▀      ███    █▀     ▄████▀
                                                                 ▀
      ▄█     █▄   ▄█   ▄█        ▄█             ▄█          ▄████████    ▄████████     ███
     ███     ███ ███  ███       ███            ███         ███    ███   ███    ███ ▀█████████▄
     ███     ███ ███▌ ███       ███            ███         ███    ███   ███    █▀     ▀███▀▀██
     ███     ███ ███▌ ███       ███            ███         ███    ███   ███            ███   ▀
     ███     ███ ███▌ ███       ███            ███       ▀███████████ ▀███████████     ███
     ███     ███ ███  ███       ███            ███         ███    ███          ███     ███
     ███ ▄█▄ ███ ███  ███▌    ▄ ███▌    ▄      ███▌    ▄   ███    ███    ▄█    ███     ███
      ▀███▀███▀  █▀   █████▄▄██ █████▄▄██      █████▄▄██   ███    █▀   ▄████████▀     ▄████▀
                      ▀         ▀              ▀
	`
}

func animationPartTwo() string {
	return `
        ▄████████ ███▄▄▄▄   ████████▄        ▄██████▄   ▄█    █▄     ▄████████    ▄████████          ███      ▄█    ▄▄▄▄███▄▄▄▄      ▄████████
       ███    ███ ███▀▀▀██▄ ███   ▀███      ███    ███ ███    ███   ███    ███   ███    ███      ▀█████████▄ ███  ▄██▀▀▀███▀▀▀██▄   ███    ███
       ███    ███ ███   ███ ███    ███      ███    ███ ███    ███   ███    █▀    ███    ███         ▀███▀▀██ ███▌ ███   ███   ███   ███    █▀
       ███    ███ ███   ███ ███    ███      ███    ███ ███    ███  ▄███▄▄▄      ▄███▄▄▄▄██▀          ███   ▀ ███▌ ███   ███   ███  ▄███▄▄▄
     ▀███████████ ███   ███ ███    ███      ███    ███ ███    ███ ▀▀███▀▀▀     ▀▀███▀▀▀▀▀            ███     ███▌ ███   ███   ███ ▀▀███▀▀▀
       ███    ███ ███   ███ ███    ███      ███    ███ ███    ███   ███    █▄  ▀███████████          ███     ███  ███   ███   ███   ███    █▄
       ███    ███ ███   ███ ███   ▄███      ███    ███ ███    ███   ███    ███   ███    ███          ███     ███  ███   ███   ███   ███    ███
       ███    █▀   ▀█   █▀  ████████▀        ▀██████▀   ▀██████▀    ██████████   ███    ███         ▄████▀   █▀    ▀█   ███   █▀    ██████████
                                                                                 ███    ███
      ▄█     █▄     ▄████████  ▄█    █▄     ▄████████       ▄█          ▄████████    ▄████████    ▄████████ ███▄▄▄▄      ▄████████ ████████▄
     ███     ███   ███    ███ ███    ███   ███    ███      ███         ███    ███   ███    ███   ███    ███ ███▀▀▀██▄   ███    ███ ███   ▀███
     ███     ███   ███    █▀  ███    ███   ███    █▀       ███         ███    █▀    ███    ███   ███    ███ ███   ███   ███    █▀  ███    ███
     ███     ███  ▄███▄▄▄     ███    ███  ▄███▄▄▄          ███        ▄███▄▄▄       ███    ███  ▄███▄▄▄▄██▀ ███   ███  ▄███▄▄▄     ███    ███
     ███     ███ ▀▀███▀▀▀     ███    ███ ▀▀███▀▀▀          ███       ▀▀███▀▀▀     ▀███████████ ▀▀███▀▀▀▀▀   ███   ███ ▀▀███▀▀▀     ███    ███
     ███     ███   ███    █▄  ███    ███   ███    █▄       ███         ███    █▄    ███    ███ ▀███████████ ███   ███   ███    █▄  ███    ███
     ███ ▄█▄ ███   ███    ███ ███    ███   ███    ███      ███▌    ▄   ███    ███   ███    ███   ███    ███ ███   ███   ███    ███ ███   ▄███
      ▀███▀███▀    ██████████  ▀██████▀    ██████████      █████▄▄██   ██████████   ███    █▀    ███    ███  ▀█   █▀    ██████████ ████████▀
                                                           ▀                                     ███    ███
        ▄████████    ▄████████  ▄██████▄    ▄▄▄▄███▄▄▄▄            ███        ▄█    █▄       ▄████████         ▄███████▄    ▄████████    ▄████████     ███
       ███    ███   ███    ███ ███    ███ ▄██▀▀▀███▀▀▀██▄      ▀█████████▄   ███    ███     ███    ███        ███    ███   ███    ███   ███    ███ ▀█████████▄
       ███    █▀    ███    ███ ███    ███ ███   ███   ███         ▀███▀▀██   ███    ███     ███    █▀         ███    ███   ███    ███   ███    █▀     ▀███▀▀██
      ▄███▄▄▄      ▄███▄▄▄▄██▀ ███    ███ ███   ███   ███          ███   ▀  ▄███▄▄▄▄███▄▄  ▄███▄▄▄            ███    ███   ███    ███   ███            ███   ▀
     ▀▀███▀▀▀     ▀▀███▀▀▀▀▀   ███    ███ ███   ███   ███          ███     ▀▀███▀▀▀▀███▀  ▀▀███▀▀▀          ▀█████████▀  ▀███████████ ▀███████████     ███
       ███        ▀███████████ ███    ███ ███   ███   ███          ███       ███    ███     ███    █▄         ███          ███    ███          ███     ███
       ███          ███    ███ ███    ███ ███   ███   ███          ███       ███    ███     ███    ███        ███          ███    ███    ▄█    ███     ███
       ███          ███    ███  ▀██████▀   ▀█   ███   █▀          ▄████▀     ███    █▀      ██████████       ▄████▀        ███    █▀   ▄████████▀     ▄████▀
                    ███    ███
	`
}

func animationPartThree() string {
	return `
         ███        ▄█    █▄       ▄████████     ███          ███▄▄▄▄    ▄██████▄         ▄▄▄▄███▄▄▄▄      ▄████████ ███▄▄▄▄      ▄████████
     ▀█████████▄   ███    ███     ███    ███ ▀█████████▄      ███▀▀▀██▄ ███    ███      ▄██▀▀▀███▀▀▀██▄   ███    ███ ███▀▀▀██▄   ███    ███
        ▀███▀▀██   ███    ███     ███    ███    ▀███▀▀██      ███   ███ ███    ███      ███   ███   ███   ███    ███ ███   ███   ███    █▀
         ███   ▀  ▄███▄▄▄▄███▄▄   ███    ███     ███   ▀      ███   ███ ███    ███      ███   ███   ███   ███    ███ ███   ███   ███
         ███     ▀▀███▀▀▀▀███▀  ▀███████████     ███          ███   ███ ███    ███      ███   ███   ███ ▀███████████ ███   ███ ▀███████████
         ███       ███    ███     ███    ███     ███          ███   ███ ███    ███      ███   ███   ███   ███    ███ ███   ███          ███
         ███       ███    ███     ███    ███     ███          ███   ███ ███    ███      ███   ███   ███   ███    ███ ███   ███    ▄█    ███
        ▄████▀     ███    █▀      ███    █▀     ▄████▀         ▀█   █▀   ▀██████▀        ▀█   ███   █▀    ███    █▀   ▀█   █▀   ▄████████▀

        ▄████████  ▄█      ███              ███      ▄██████▄          ▄████████ ███    █▄   ▄█          ▄████████
       ███    ███ ███  ▀█████████▄      ▀█████████▄ ███    ███        ███    ███ ███    ███ ███         ███    ███
       ███    █▀  ███▌    ▀███▀▀██         ▀███▀▀██ ███    ███        ███    ███ ███    ███ ███         ███    █▀
      ▄███▄▄▄     ███▌     ███   ▀          ███   ▀ ███    ███       ▄███▄▄▄▄██▀ ███    ███ ███        ▄███▄▄▄
     ▀▀███▀▀▀     ███▌     ███              ███     ███    ███      ▀▀███▀▀▀▀▀   ███    ███ ███       ▀▀███▀▀▀
       ███        ███      ███              ███     ███    ███      ▀███████████ ███    ███ ███         ███    █▄
       ███        ███      ███              ███     ███    ███        ███    ███ ███    ███ ███▌    ▄   ███    ███
       ███        █▀      ▄████▀           ▄████▀    ▀██████▀         ███    ███ ████████▀  █████▄▄██   ██████████
                                                                      ███    ███            ▀
         ███        ▄█    █▄       ▄████████       ▄█     █▄   ▄██████▄     ▄████████  ▄█       ████████▄          ▄████████  ▄█        ▄██████▄  ███▄▄▄▄      ▄████████
     ▀█████████▄   ███    ███     ███    ███      ███     ███ ███    ███   ███    ███ ███       ███   ▀███        ███    ███ ███       ███    ███ ███▀▀▀██▄   ███    ███
        ▀███▀▀██   ███    ███     ███    █▀       ███     ███ ███    ███   ███    ███ ███       ███    ███        ███    ███ ███       ███    ███ ███   ███   ███    █▀
         ███   ▀  ▄███▄▄▄▄███▄▄  ▄███▄▄▄          ███     ███ ███    ███  ▄███▄▄▄▄██▀ ███       ███    ███        ███    ███ ███       ███    ███ ███   ███  ▄███▄▄▄
         ███     ▀▀███▀▀▀▀███▀  ▀▀███▀▀▀          ███     ███ ███    ███ ▀▀███▀▀▀▀▀   ███       ███    ███      ▀███████████ ███       ███    ███ ███   ███ ▀▀███▀▀▀
         ███       ███    ███     ███    █▄       ███     ███ ███    ███ ▀███████████ ███       ███    ███        ███    ███ ███       ███    ███ ███   ███   ███    █▄
         ███       ███    ███     ███    ███      ███ ▄█▄ ███ ███    ███   ███    ███ ███▌    ▄ ███   ▄███        ███    ███ ███▌    ▄ ███    ███ ███   ███   ███    ███
        ▄████▀     ███    █▀      ██████████       ▀███▀███▀   ▀██████▀    ███    ███ █████▄▄██ ████████▀         ███    █▀  █████▄▄██  ▀██████▀   ▀█   █▀    ██████████
                                                                           ███    ███ ▀                                      ▀
	`
}

func animationPartFour() string {
	return `
        ▄████████        ▄▄▄▄███▄▄▄▄      ▄████████ ███▄▄▄▄         ▄█     █▄   ▄█   ▄█        ▄█            ████████▄   ▄█     ▄████████
       ███    ███      ▄██▀▀▀███▀▀▀██▄   ███    ███ ███▀▀▀██▄      ███     ███ ███  ███       ███            ███   ▀███ ███    ███    ███
       ███    ███      ███   ███   ███   ███    ███ ███   ███      ███     ███ ███▌ ███       ███            ███    ███ ███▌   ███    █▀
       ███    ███      ███   ███   ███   ███    ███ ███   ███      ███     ███ ███▌ ███       ███            ███    ███ ███▌  ▄███▄▄▄
     ▀███████████      ███   ███   ███ ▀███████████ ███   ███      ███     ███ ███▌ ███       ███            ███    ███ ███▌ ▀▀███▀▀▀
       ███    ███      ███   ███   ███   ███    ███ ███   ███      ███     ███ ███  ███       ███            ███    ███ ███    ███    █▄
       ███    ███      ███   ███   ███   ███    ███ ███   ███      ███ ▄█▄ ███ ███  ███▌    ▄ ███▌    ▄      ███   ▄███ ███    ███    ███
       ███    █▀        ▀█   ███   █▀    ███    █▀   ▀█   █▀        ▀███▀███▀  █▀   █████▄▄██ █████▄▄██      ████████▀  █▀     ██████████
                                                                                    ▀         ▀
     ▀█████████▄  ███    █▄      ███          ███▄▄▄▄    ▄██████▄      ███             ▄█    █▄     ▄█     ▄████████
       ███    ███ ███    ███ ▀█████████▄      ███▀▀▀██▄ ███    ███ ▀█████████▄        ███    ███   ███    ███    ███
       ███    ███ ███    ███    ▀███▀▀██      ███   ███ ███    ███    ▀███▀▀██        ███    ███   ███▌   ███    █▀
      ▄███▄▄▄██▀  ███    ███     ███   ▀      ███   ███ ███    ███     ███   ▀       ▄███▄▄▄▄███▄▄ ███▌   ███
     ▀▀███▀▀▀██▄  ███    ███     ███          ███   ███ ███    ███     ███          ▀▀███▀▀▀▀███▀  ███▌ ▀███████████
       ███    ██▄ ███    ███     ███          ███   ███ ███    ███     ███            ███    ███   ███           ███
       ███    ███ ███    ███     ███          ███   ███ ███    ███     ███            ███    ███   ███     ▄█    ███
     ▄█████████▀  ████████▀     ▄████▀         ▀█   █▀   ▀██████▀     ▄████▀          ███    █▀    █▀    ▄████████▀

      ▄█  ████████▄     ▄████████    ▄████████    ▄████████
     ███  ███   ▀███   ███    ███   ███    ███   ███    ███
     ███▌ ███    ███   ███    █▀    ███    ███   ███    █▀
     ███▌ ███    ███  ▄███▄▄▄       ███    ███   ███
     ███▌ ███    ███ ▀▀███▀▀▀     ▀███████████ ▀███████████
     ███  ███    ███   ███    █▄    ███    ███          ███
     ███  ███   ▄███   ███    ███   ███    ███    ▄█    ███
     █▀   ████████▀    ██████████   ███    █▀   ▄████████▀

	`
}

func final() string {
	return `
	████████▄     ▄████████    ▄████████    ▄████████ ███▄▄▄▄   ████████▄
	███   ▀███   ███    ███   ███    ███   ███    ███ ███▀▀▀██▄ ███   ▀███
	███    ███   ███    █▀    ███    █▀    ███    █▀  ███   ███ ███    ███
	███    ███  ▄███▄▄▄      ▄███▄▄▄      ▄███▄▄▄     ███   ███ ███    ███
	███    ███ ▀▀███▀▀▀     ▀▀███▀▀▀     ▀▀███▀▀▀     ███   ███ ███    ███
	███    ███   ███    █▄    ███          ███    █▄  ███   ███ ███    ███
	███   ▄███   ███    ███   ███          ███    ███ ███   ███ ███   ▄███
	████████▀    ██████████   ███          ██████████  ▀█   █▀  ████████▀

	████████▄     ▄████████ ███▄▄▄▄   ▄██   ▄
	███   ▀███   ███    ███ ███▀▀▀██▄ ███   ██▄
	███    ███   ███    █▀  ███   ███ ███▄▄▄███
	███    ███  ▄███▄▄▄     ███   ███ ▀▀▀▀▀▀███
	███    ███ ▀▀███▀▀▀     ███   ███ ▄██   ███
	███    ███   ███    █▄  ███   ███ ███   ███
	███   ▄███   ███    ███ ███   ███ ███   ███
	████████▀    ██████████  ▀█   █▀   ▀█████▀

	████████▄     ▄████████    ▄███████▄  ▄██████▄     ▄████████    ▄████████
	███   ▀███   ███    ███   ███    ███ ███    ███   ███    ███   ███    ███
	███    ███   ███    █▀    ███    ███ ███    ███   ███    █▀    ███    █▀
	███    ███  ▄███▄▄▄       ███    ███ ███    ███   ███         ▄███▄▄▄
	███    ███ ▀▀███▀▀▀     ▀█████████▀  ███    ███ ▀███████████ ▀▀███▀▀▀
	███    ███   ███    █▄    ███        ███    ███          ███   ███    █▄
	███   ▄███   ███    ███   ███        ███    ███    ▄█    ███   ███    ███
	████████▀    ██████████  ▄████▀       ▀██████▀   ▄████████▀    ██████████

   `
}

func (l Logo) View() {
	fmt.Print("\033[H\033[2J")
	in := l.render

	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(220))

	out, _ := r.Render(in)
	fmt.Print(out)
}