package rpc

import (
	"io"
	"encoding/json"
	"net/http"
	"../core"
)

var blockChain *core.Blockchain

func blockchainGetHandler(rw http.ResponseWriter, r *http.Request) {
	var (
		bytes []byte
		err error
	)
	if bytes, err = json.Marshal(blockChain); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(rw, string(bytes))
}

func blockchainUpdateHandler(rw http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockChain.AppendData(blockData)
	blockchainGetHandler(rw, r)
}

func Run()  {
	blockChain = core.InitBlockChain()
	http.HandleFunc("/blockchain/get", blockchainGetHandler)
	http.HandleFunc("/blockchain/update", blockchainUpdateHandler)

	http.ListenAndServe("0.0.0.0:8888", nil)
}
