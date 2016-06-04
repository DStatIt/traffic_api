package traffic_api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ErrBadRequest = errors.New("Bad Request")
)

func UpdateHost(w http.ResponseWriter, r *http.Request) (int, error) {
	userID := r.FormValue("userID")
	if userID == "" {
		return 500, ErrBadRequest
	}
	url := r.FormValue("url")
	if url == "" {
		return 500, ErrBadRequest
	}

	db, err := sql.Open("mysql", "root:Danielkoz@tcp(107.191.48.191:3306)/userdata")
	defer db.Close()
	if err != nil {
		return 500, err
	}

	_, err = db.Exec("UPDATE userinfo SET host=? WHERE id=?", url, userID)
	if err != nil {
		return 500, err
	}
	fmt.Println("Host for User: ", userID)

	return 200, nil
}

func PrintOut(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, fmt.Sprintf("%+v", r))
}
