package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/tejashwinn/splitwise/constants"
	"github.com/tejashwinn/splitwise/mappers"
	"github.com/tejashwinn/splitwise/repos"
	"github.com/tejashwinn/splitwise/types"
	"github.com/tejashwinn/splitwise/utils"
)

type CurrencyHandler struct {
	Repo    repos.CurrencyRepo
	JwtUtil utils.JwtUtil
}

func NewCurrencyHandler(
	repo repos.CurrencyRepo,
	jwtUtil *utils.JwtUtil,
) *CurrencyHandler {
	return &CurrencyHandler{Repo: repo, JwtUtil: *jwtUtil}
}

func (h *CurrencyHandler) ListCurrencies(w http.ResponseWriter, r *http.Request) {
	currencies, err := h.Repo.FindAllOrderByName(context.Background())
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to fetch currencies", http.StatusInternalServerError)
		return
	}
	currenciesRes := []types.CurrencyRes{}
	for _, currency := range currencies {
		currencyRes, err := mappers.CurrencyModelToCurrencyRes(
			&currency,
		)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}
		currenciesRes = append(currenciesRes, *currencyRes)
	}
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	json.NewEncoder(w).Encode(currenciesRes)
}
