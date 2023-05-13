package esClient

import (
	"MaoerMovie/common/kqueue"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"net/http"
	"strings"
)

const UserIndexName = "user"

type EsModel struct {
	es        *elasticsearch.Client
	indexName string
}

type UserEsModel struct {
	Id        int64
	Name      string
	AvatarUrl string
}

func NewEsModel(Addresses []string, Username string, Password string) *EsModel {
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: Addresses,
		Username:  Username,
		Password:  Password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	})
	if err != nil {
		return nil
	}
	return &EsModel{
		es:        esClient,
		indexName: UserIndexName,
	}
}

// createIndex创建Elasticsearch索引
func createIndex(es *elasticsearch.Client) {
	indexMapping := `{
		"mappings": {
			"properties": {
				"Id": {
					"type": "long"
				},
				"Name": {
					"type": "text"
				},
				"AvatarUrl": {
					"type": "keyword"
				}
			}
		}
	}`

	req := esapi.IndicesCreateRequest{
		Index: "user",
		Body:  strings.NewReader(indexMapping),
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error creating index: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error creating index: %s", res.String())
	}
}

func (esModel *EsModel) InsertUser(message kqueue.UserInsertMessage) error {
	user := UserEsModel{
		Id:   message.Id,
		Name: message.Name,
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error marshaling user: %v", err)
	}

	req := esapi.IndexRequest{
		Index:      "user",
		DocumentID: fmt.Sprintf("%d", user.Id),
		Body:       strings.NewReader(string(userJSON)),
	}

	res, err := req.Do(context.Background(), esModel.es)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error indexing document ID=%d: %s", user.Id, res.String())
		return err
	}
	return nil
}

// 更新文档
func (esModel *EsModel) UpdateUser(message kqueue.UserUpdateMessage) error {
	user := UserEsModel{
		Id:        message.Id,
		Name:      message.Name,
		AvatarUrl: message.AvatarUrl,
	}
	update := map[string]interface{}{
		"doc": user,
	}
	updateJSON, err := json.Marshal(update)
	if err != nil {
		log.Fatalf("Error marshaling update: %v", err)
		return err
	}

	req := esapi.UpdateRequest{
		Index:      "user",
		DocumentID: fmt.Sprintf("%d", user.Id),
		Body:       strings.NewReader(string(updateJSON)),
	}

	res, err := req.Do(context.Background(), esModel.es)
	if err != nil {
		log.Printf("Failed to update document ID=%d: %s", user.Id, err)
		return err
	}
	defer res.Body.Close()
	return nil
}
