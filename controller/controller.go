package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"apirest.ofq/operation"
	"apirest.ofq/storages"
	"apirest.ofq/structs"
	"github.com/google/logger"
	"github.com/gorilla/mux"
)

func PostTopSecretHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("POST /topsecret/")
	var ofq structs.OFQ
	err := json.NewDecoder(r.Body).Decode(&ofq)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		logger.Error("Ha ocurrido un error decodificando el json")
		formatError(w, http.StatusInternalServerError, "Ha ocurrido un error decodificando el json")
		return
	}

	//Se guardan los satelites en el storage
	logger.Info("Se guardan los satelites en el Store")
	for _, v := range ofq.SateliteList {
		nameSat := strings.ToLower(v.Name)
		storages.SatelitesStore[nameSat] = v
	}

	//Se calcula la coordenadas de la señal
	logger.Info("Se calculan las coordenadas de la señal")
	distanceKenobi := storages.SatelitesStore["kenobi"].Distance
	distanceSkywalker := storages.SatelitesStore["skywalker"].Distance
	distanceSato := storages.SatelitesStore["sato"].Distance
	x, y, err := operation.GetLocation(distanceKenobi, distanceSkywalker, distanceSato)
	if err != nil {
		logger.Error("No se ha determinado la posicion de la señal: " + err.Error())
		formatError(w, http.StatusNotFound, "No se ha determinado la posicion de la señal")
		return
	}

	//Se descifra el mensaje
	logger.Info("Se descifra el mensaje enviado por la señal")
	messageKenobi := storages.SatelitesStore["kenobi"].Message
	messageSkywalker := storages.SatelitesStore["skywalker"].Message
	messageSato := storages.SatelitesStore["sato"].Message
	msg, err := operation.GetMessage(messageKenobi, messageSkywalker, messageSato)
	if err != nil {
		logger.Error("No se ha descifrar el mensaje recibido: " + err.Error())
		formatError(w, http.StatusNotFound, "No se ha descifrar el mensaje recibido")
		return
	}

	resp := new(structs.ResponseMessage)
	resp.Message = msg
	resp.Position.X = x
	resp.Position.Y = y

	j, err := json.Marshal(resp)
	if err != nil {
		logger.Error("Ha ocurrido un error decodificando el json")
		formatError(w, http.StatusInternalServerError, "Ha ocurrido un error codificando el json")
		return
	}
	//se retorna la respuesta al cliente
	logger.Info("Se envia la respuesta al cliente")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetTopSecretSplitHandler(w http.ResponseWriter, r *http.Request) {
	paramURL := mux.Vars(r)
	sateliteName := strings.ToLower(paramURL["satellite_name"])
	logger.Info("GET /topsecret_split/" + sateliteName)
	var satelite structs.Satelite

	w.Header().Set("Content-Type", "application/json")
	if _, ok := storages.SatelitesStore[sateliteName]; ok {
		satelite = storages.SatelitesStore[sateliteName]
	} else {
		logger.Error("No se ha encontrado el satelite")
		formatError(w, http.StatusNotFound, "No se ha encontrado el satelite")
		return
	}

	coord := operation.GetCoordSat(sateliteName)
	msg := strings.Join(satelite.Message, " ")
	if strings.Trim(msg, " ") == "" {
		logger.Error("No hay suficiente informacion")
		formatError(w, http.StatusNotFound, "No hay suficiente informacion")
		return
	}

	resp := new(structs.ResponseMessage)
	resp.Message = msg
	resp.Position.X = coord.X
	resp.Position.Y = coord.Y

	j, err := json.Marshal(resp)
	if err != nil {
		logger.Error("Ha ocurrido un error codificando el json")
		formatError(w, http.StatusInternalServerError, "Ha ocurrido un error codificando el json")
		return
	}
	//se retorna la respuesta al cliente
	logger.Info("Se envia la respuesta al cliente")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func PostTopSecretSplitHandler(w http.ResponseWriter, r *http.Request) {
	paramURL := mux.Vars(r)
	sateliteName := strings.ToLower(paramURL["satellite_name"])
	logger.Info("POST /topsecret_split/" + sateliteName)
	var sateliteInfo structs.SateliteInfo
	var satelite structs.Satelite

	err := json.NewDecoder(r.Body).Decode(&sateliteInfo)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		logger.Error("Ha ocurrido un error codificando el json")
		formatError(w, http.StatusInternalServerError, "Ha ocurrido un error codificando el json")
		return
	}

	if _, ok := storages.SatelitesStore[sateliteName]; ok {
		satelite = storages.SatelitesStore[sateliteName]
	} else {
		formatError(w, http.StatusNotFound, "No se ha encontrado el satelite")
		return
	}

	satelite.SateliteInfo = &sateliteInfo
	storages.SatelitesStore[sateliteName] = satelite

	//Se calcula la coordenadas de la señal
	logger.Info("Se calculan las coordenadas de la señal")
	distanceKenobi := storages.SatelitesStore["kenobi"].Distance
	distanceSkywalker := storages.SatelitesStore["skywalker"].Distance
	distanceSato := storages.SatelitesStore["sato"].Distance
	x, y, err := operation.GetLocation(distanceKenobi, distanceSkywalker, distanceSato)
	if err != nil {
		logger.Error("No se ha determinado la posicion de la señal: " + err.Error())
		formatError(w, http.StatusNotFound, "No se ha determinado la posicion de la señal")
		return
	}

	//Se descifra el mensaje
	logger.Info("Se descifra el mensaje enviado por la señal")
	messageKenobi := storages.SatelitesStore["kenobi"].Message
	messageSkywalker := storages.SatelitesStore["skywalker"].Message
	messageSato := storages.SatelitesStore["sato"].Message
	msg, err := operation.GetMessage(messageKenobi, messageSkywalker, messageSato)
	if err != nil {
		logger.Error("No se ha descifrar el mensaje recibido: " + err.Error())
		formatError(w, http.StatusNotFound, "No se ha descifrar el mensaje recibido")
		return
	}

	resp := new(structs.ResponseMessage)
	resp.Message = msg
	resp.Position.X = x
	resp.Position.Y = y

	fmt.Println(sateliteName)
	j, err := json.Marshal(resp)
	if err != nil {
		logger.Error("Ha ocurrido un error codificando el json")
		formatError(w, http.StatusInternalServerError, "Ha ocurrido un error codificando el json")
		return
	}
	//se retorna la respuesta al cliente
	logger.Info("Se envia la respuesta al cliente")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func formatError(w http.ResponseWriter, statusHttp int, messageError string) {
	w.WriteHeader(statusHttp)
	w.Write([]byte(messageError))
}
