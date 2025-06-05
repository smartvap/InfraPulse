package v1

import (
	"encoding/json"
	"github.com/emicklei/go-restful/v3"
	"k8s.io/klog/v2"
	"net/http"
)

type Healthz struct {
	Health string `json:"health"`
}

func Health(request *restful.Request, response *restful.Response) {
	klog.Info("health")
	health := &Healthz{
		Health: "up",
	}
	b, _ := json.Marshal(health)
	writeJSONResponse(response, b)
}

// Derived from go-restful writeJSON.
func writeJSONResponse(response *restful.Response, data []byte) {
	if data == nil {
		response.WriteHeader(http.StatusOK)
		// do not write a nil representation
		return
	}
	response.Header().Set(restful.HEADER_ContentType, restful.MIME_JSON)
	response.WriteHeader(http.StatusOK)
	if _, err := response.Write(data); err != nil {
		klog.ErrorS(err, "Error writing response")
	}

}
