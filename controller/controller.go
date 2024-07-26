package controller

import (
	"encoding/json"
	"net/http"
	"test/entities"
	"test/models"
)

func TwoSumHandler(w http.ResponseWriter, r *http.Request) {
	var req entities.Input
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := models.TwoSum(req.Nums, req.Target)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func ThreeSumHandler(w http.ResponseWriter, r *http.Request) {
	var req entities.Input
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := models.ThreeSum(req.Nums)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func FindSubstringHandler(w http.ResponseWriter, r *http.Request) {
	var req entities.Input
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	res := models.FindSubstring(req.S, req.Words)


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}