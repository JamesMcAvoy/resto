// Partie "salle" du projet C#/.NET (sans .NET ni C#)
// Affiche et gère le restaurant, communique avec un serveur qui gère la cuisine.

package main

import (
	"github.com/JamesMcAvoy/resto/src/controller"
	"github.com/andlabs/ui"
	"github.com/faiface/pixel/pixelgl"
	"os"
	"time"
)

const (
	width   = 1280
	height  = 704
	adresse = "http://127.0.0.1:9090/"
	// Jusqu'à ce qu'on lie les 2 projets:
	acceleration = 60 // Accélération initiale du temps
	port         = 9090
)

func run() {
	// Jusqu'à ce qu'on lie les 2 projets:
	go Serv(port, acceleration)
	time.Sleep(50 * time.Millisecond)

	game := controller.NewGame(width, height, adresse)
	fin := make(chan bool)
	for i, r := range game.Restos {
		go func(i int, r *controller.Resto) {
			<-r.Win.Fin
			if len(game.Restos) > 1 {
				// Supprime le restaurant quand sa fenêtre est fermée
				game.Restos[i] = game.Restos[len(game.Restos)-1]
				game.Restos[len(game.Restos)-1] = nil
				game.Restos = game.Restos[:len(game.Restos)-1]
			} else {
				fin <- true
			}
		}(i, r)
	}
	<-fin
}

func main() {
	go func() {
		err := ui.Main(func() {})
		if err != nil {
			panic(err)
		}
	}()
	err := os.Chdir(os.Getenv("GOPATH") + "/src/github.com/JamesMcAvoy/resto")
	if err != nil {
		panic(err)
	}
	pixelgl.Run(run)
}
