package esClient

import (
	"MaoerMovie/service/user-rpc/model"
	"bytes"
	"context"
	"crypto/tls"
	"database/sql"
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
				"id": {
					"type": "integer"
				},
				"name": {
					"type": "text"
				},
				"email": {
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

func (esModel *EsModel) InsertData(user *model.User) error {
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

func (esModel *EsModel) SearchUser(ctx context.Context, name string, skip, count int64) ([]*model.User, int64, error) {
	var buf bytes.Buffer
	queryJSON := map[string]interface{}{
		"from": skip - 1,
		"size": count,
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"Name": name,
			},
		},
	}
	err := json.NewEncoder(&buf).Encode(queryJSON)
	if err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := esModel.es.Search(
		esModel.es.Search.WithContext(context.Background()),
		esModel.es.Search.WithIndex("user"),
		esModel.es.Search.WithBody(&buf),
		esModel.es.Search.WithTrackTotalHits(true),
		esModel.es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil, 0, err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error searching documents: %s", res.String())
		return nil, 0, err
	}

	var results map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		log.Fatalf("Error parsingthe response body: %s", err)
		return nil, 0, err
	}

	fmt.Printf("Found %d results\n", int(results["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)))

	var users []*model.User
	for _, hit := range results["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		user := &model.User{
			Id:   int64(source["Id"].(float64)),
			Name: source["Name"].(string),
			AvatarUrl: sql.NullString{
				String: source["AvatarUrl"].(string),
				Valid:  true, // 表示字符串有效
			},
		}
		users = append(users, user)
		//fmt.Printf("User ID: %d, Name: %s, Email: %s\n", int(source["ID"].(float64)), source["Name"].(string), source["Email"].(string))
	}
	return users, int64(len(users)), err
}
