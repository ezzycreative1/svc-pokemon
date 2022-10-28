package external

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ezzycreative1/svc-pokemon/internal/core/domain"
)

type externalPokeRepo struct{}

func NewExternalPokeRepo() externalPokeRepo {
	return externalPokeRepo{}
}

func (ep *externalPokeRepo) GetPokemonByName(ctx context.Context, name string) (*domain.PokemonExternal, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + name)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var dataPoke domain.PokemonExternal
	json.Unmarshal(bodyBytes, &dataPoke)

	return &dataPoke, nil
}
