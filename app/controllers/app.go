package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	endpoint := revel.Config.StringDefault("brawlstars.apiendpoint", "")
	api_key := "Bearer " + revel.Config.StringDefault("brawlstars.apikey", "")

	log := c.Log.New("endpoint", endpoint)
	log.Debug("Reading the output")
	log.Debug(endpoint)
	log.Debug(api_key)

	playerEndpoint := endpoint + "/players/%2320UGVV22Q"
	req, err := http.NewRequest("GET", playerEndpoint, nil)
	if err != nil {
		return c.Render("Something wrong")
	}
	req.Header.Set("Authorization", api_key)

	client := &http.Client{}
	resp, _ := client.Do(req)

	type Player struct {
		Name      string
		NameColor string
		Icon      struct {
			Id string
		}
		Trophies        string
		HighestTrophies string
	}
	// Lets use lumen next time
	newPlayer := new(Player)
	json.NewDecoder(resp.Body).Decode(&newPlayer)
	return c.RenderJSON(newPlayer)
}

func (c App) Hello(myName string) revel.Result {

	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	// 	return c.RenderTemplate("App/Something/New.html")

	return c.Render(myName)
}
