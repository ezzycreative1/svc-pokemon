package external

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ezzycreative1/svc-pokemon/internal/core/domain"
)

type externalPokeTypeRepo struct{}

func NewExternalPokeTypeRepo() externalPokeTypeRepo {
	return externalPokeTypeRepo{}
}

func (ep *externalPokeTypeRepo) GetPokemonType(ctx context.Context) (*domain.Type, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/type/")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var dataPoke domain.Type
	json.Unmarshal(bodyBytes, &dataPoke)

	return &dataPoke, nil
}
