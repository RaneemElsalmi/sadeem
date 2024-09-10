package controllers

import (
	"net/http"
	
	"os"
	"strings"
	"test/controllers/utils"
	"test/controllers/utils/models"
)

var (
	Domainv = os.Getenv("DOMAIN")
	vendor_columns = []string{
		"id",
		"name",
		"img",
		"description",
		"created_at",
		"updated_at",
	}
)

func IndexvendorHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, vendors!")) 
}
func ShowvendorHandler(w http.ResponseWriter, r *http.Request) {
	var vendor []models.Vendor
	id := r.PathValue("id")
	query, args, err := QB.Select(strings.Join(vendor_columns, ", ")).
	From("vendor").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := db.Get(&vendor, query, args...); err != nil {
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendJSONResponse(w, http.StatusOK, vendor)
}