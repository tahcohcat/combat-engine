package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tahcohcat/combat-engine/entities"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var testGame entities.Game

func setup() {
	fmt.Println("In setup")

	playerOne := entities.Player{
		Name: "player-one",
		Number: 1,
		Characters: make([]entities.Character, 0),
		IsAlive : true}

	playerTwo := entities.Player{
		Name: "player-two",
		Number: 2,
		Characters: make([]entities.Character, 0),
		IsAlive : true}

	players := make([]entities.Player, 0)
	players = append(players, playerOne)
	players = append(players, playerTwo)

	testGame = entities.Game{
		IsRunning : false,
		Players : players,
		RoundNumber: 0,
		MaxRounds: 10 }
}


func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func checkResponse(
	t *testing.T,
	rr *httptest.ResponseRecorder,
	expectedStatusCode int,
	expectedBody string) {

	if status := rr.Code; status != expectedStatusCode {
	t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
	}

	returned := rr.Body.String()
	returned = returned[0: len(returned)-1] //todo: why is there an extra newline appended?

	if returned != expectedBody {
		t.Errorf("handler returned unexpected body: got:\n[%v] want:\n[%v]",
			returned, expectedBody)
	}
}

func sendRequest(method string, url string, body io.Reader)  (*httptest.ResponseRecorder, *http.Request) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	return httptest.NewRecorder(), req
}

func TestHandler(t *testing.T) {

	jsonGame, err := json.Marshal(testGame)
	if err != nil { t.Fatal(err) }

	rr, req := sendRequest("POST", "/games", bytes.NewBuffer([]byte(jsonGame)))

	handler := http.HandlerFunc(Game_Handler_POST)

	handler.ServeHTTP(rr, req)
	checkResponse(t, rr, 200, `{"id":"\u0000","message":"Game created successfully"}`)
}

