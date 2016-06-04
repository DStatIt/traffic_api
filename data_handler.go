package traffic_api

import (
	"fmt"
	"net/http"
)

func PrintOut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("%+v", r))
}
