package mysql

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	domain "github.com/ezzycreative1/svc-pokemon/internal/core/domain"
	"github.com/ezzycreative1/svc-pokemon/internal/data/customtype"
	"gorm.io/gorm"
)

const (
	conTimeout = 60

	// KeyTransaction concrete type for key context value transaction
	KTrans customtype.KeyTrx = customtype.KeyTrx("pokemon-trx")
)

type mysqlPokemonRepo struct {
	DB *gorm.DB
}

func NewMysqlPokemonRepo(db *gorm.DB) mysqlPokemonRepo {
	return mysqlPokemonRepo{
		DB: db,
	}
}

func (pr *mysqlPokemonRepo) FetchPokemons(ctx context.Context) ([]domain.Pokemons, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	ctxWT, cancel := context.WithTimeout(ctx, conTimeout*time.Second)
	defer cancel()

	var res []domain.Pokemons
	query := pr.DB.WithContext(ctxWT).Find(&res)
	if query.Error != nil {
		return nil, query.Error
	}

	return res, nil
}

func (pr *mysqlPokemonRepo) StorePokemon(ctx context.Context, input *domain.StorePokemonRequest) error {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var dataPoke domain.PokemonExternal
	json.Unmarshal(bodyBytes, &dataPoke)
	fmt.Printf("API Response as struct %+v\n", dataPoke)

	data := domain.Pokemons{
		Name:      dataPoke.Name,
		UserID:    input.UserID,
		Stock:     input.Stock,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	trx, ok := ctx.Value(KeyTransaction).(*gorm.DB)
	if !ok {
		trx = pr.DB
	}

	ctxWT, cancel := context.WithTimeout(ctx, connectTimeout*time.Second)
	defer cancel()

	query := trx.WithContext(ctxWT).Create(&data)
	if query.Error != nil {
		return query.Error
	}

	return nil
}
