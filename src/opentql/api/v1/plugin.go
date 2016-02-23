package v1

import (
    _"encoding/json"
    "net/http"
    "opentql/plugins/openstack"
)

//return the list of OS falvours fro perticular account
func GetFlavors(w http.ResponseWriter, r *http.Request) {
	//get the account id param from request
	
	image_list := openstack.GetOSImages() //tested
    w.Header().Set("Content-Type", "application/json")
    w.Write(image_list)
}
