package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"net/http"
	"time"
)

// Game représente l'"objet principal".
// Il contient les restos et la liaison avec le serveur
type Game struct {
	Url    string
	Restos []*Resto
}

// Effectue une map au serveur, retourne la map de la réponse
func (c Game) Req(ob map[string]interface{}) (map[string]interface{}, error) {
	msg, err := json.Marshal(ob)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.Url, bytes.NewBuffer(msg))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var repMap map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&repMap); err != nil {
		return nil, err
	}
	return repMap, nil
}

// Initialisation des restos, connection au serveur
func NewGame(width, height int, url string) *Game {
	game := Game{
		Url: url,
	}
	bonjour := make(map[string]interface{})
	bonjour["type"] = "bonjour"
	initMap, err := game.Req(bonjour)
	if err != nil {
		panic(err)
	}
	tmp := initMap["restos"].([]interface{})
	m := make([]map[string]interface{}, len(tmp))
	for i, v := range tmp {
		m[i] = v.(map[string]interface{})
	}
	for i, r := range m {
		win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
			Title:  fmt.Sprintf("La salle du resto %v oui", i),
			Bounds: pixel.R(0, 0, float64(width), float64(height)),
			VSync:  true,
		})
		if err != nil {
			panic(err)
		}

		hor := r["horaires"].([]interface{})
		h := make([][2]float64, len(hor))
		for i, v := range hor {
			t := v.([]interface{})
			for j, val := range t {
				h[i][j] = val.(float64)
			}
		}
		en := r["entrees"].([]interface{})
		pl := r["plats"].([]interface{})
		de := r["desserts"].([]interface{})
		e := make([]string, len(en))
		p := make([]string, len(pl))
		d := make([]string, len(de))
		intToStr(en, e)
		intToStr(pl, p)
		intToStr(de, d)
		game.Restos = append(game.Restos, NewResto(
			win, int(r["temps"].(float64)), int(r["acceleration"].(float64)), false, h,
			e, p, d,
		))
	}
	return &game
}

// Converts an array of interfaces to an array of strings
func intToStr(intefaceArray []interface{}, strArray []string) {
	for i, v := range intefaceArray {
		strArray[i] = v.(string)
	}
}