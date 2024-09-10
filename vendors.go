package controllers

import (
	"fmt"
	"net/http"

	"os"
	"test/controllers/utils"
	"test/controllers/utils/models"
	"strings"
)

var (
	Domainv        = os.Getenv("DOMAIN")
	vendor_columns = []string{
		"id",
		"name",
		"description",
		"created_at",
		"updated_at",
		fmt.Sprintf("CASE WHEN NULLIF(img, '') IS NOT NULL THEN FORMAT('%s/%%s', img) ELSE NULL END AS img", Domainv),
	}
)

func IndexvendorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("vendor")
	var vendors []models.Vendor

	query, args, err := QB.Select(strings.Join(vendor_columns, ", ")).
		From("vendors").
		ToSql()

	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := db.Select(&vendors, query, args...); err != nil {
		utils.HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendJSONResponse(w, http.StatusOK, vendors)
}

func ShowvenderHandler(w http.ResponseWriter, r *http.Request) {
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