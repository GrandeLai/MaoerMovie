package esClient

import (
	"MaoerMovie/common/utils"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"net/http"
	"strings"
	"time"
)

const FilmIndexName = "film"

type EsModel struct {
	es        *elasticsearch.Client
	indexName string
}

type FilmPreview struct {
	FilmName      string
	FilmTime      time.Time
	FilmCategory  string
	FilmScore     float64
	FilmArea      string
	FilmCoverUrl  string
	ActorNameList string
}

type SearchFilterFactor struct {
	Keyword    string
	Category   string
	Page       int64
	Size       int64
	Area       string
	SortedType int64
	TimeStart  int64
	TimeEnd    int64
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
		indexName: FilmIndexName,
	}
}

// createIndex创建Elasticsearch索引
func createIndex(es *elasticsearch.Client) {
	indexMapping := `{
	  "mappings": {
		"properties": {
		  "FilmName": {
			"type": "text",
			"analyzer": "ik_max_word",
			"search_analyzer": "ik_smart"
		  },
		  "FilmTime": {
			"type": "date"
		  },
		  "FilmCategory": {
			"type": "keyword"
		  },
		  "FilmScore": {
			"type": "float"
		  },
		  "FilmArea": {
			"type": "keyword"
		  },
		  "FilmCoverUrl": {
			"type": "keyword"
		  },
		  "ActorNameList": {
			"type": "text",
			"analyzer": "ik_max_word",
			"search_analyzer": "ik_smart"
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

func (esModel *EsModel) SearchFilm(ctx context.Context, f *SearchFilterFactor) ([]*FilmPreview, int64, error) {
	var buf bytes.Buffer
	var sortType string
	if f.SortedType == 0 {
		sortType = "FilmTime"
	} else if f.SortedType == 1 {
		sortType = "FilmScore"
	}

	// 搜索请求
	searchRequest := map[string]interface{}{
		"from": f.Page - 1,
		"size": f.Size,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []any{
					map[string]any{
						"match": map[string]any{
							"FilmName": map[string]any{
								"query":    f.Keyword,
								"analyzer": "ik_smart",
							},
						},
					},
				},
				"filter": []interface{}{
					map[string]any{
						"term": map[string]any{
							"FilmCategory": f.Category,
						},
					},
					map[string]interface{}{
						"range": map[string]interface{}{
							"FilmTime": map[string]string{
								"gte": utils.Int64ToString(f.TimeStart),
								"lte": utils.Int64ToString(f.TimeEnd),
							},
						},
					},
				},
			},
		},
		"sort": map[string]any{
			sortType: map[string]any{
				"order": "desc",
			},
		},
	}
	err := json.NewEncoder(&buf).Encode(searchRequest)
	if err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	res, err := esModel.es.Search(
		esModel.es.Search.WithContext(context.Background()),
		esModel.es.Search.WithIndex("film"),
		esModel.es.Search.WithBody(&buf),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil, 0, err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error searching documents: %s", res.String())
		return nil, 0, errors.New(res.String())
	}

	var results map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		log.Fatalf("Error parsingthe response body: %s", err)
		return nil, 0, err
	}

	fmt.Printf("Found %d results\n", int(results["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)))

	var films []*FilmPreview
	for _, hit := range results["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		time := source["FilmTime"].(string)
		film := &FilmPreview{
			FilmName:      source["FilmName"].(string),
			FilmScore:     source["FilmScore"].(float64),
			FilmTime:      utils.StringToTimeAtLocal(time),
			FilmCategory:  source["FilmCategory"].(string),
			FilmArea:      source["FilmArea"].(string),
			FilmCoverUrl:  source["FilmCoverUrl"].(string),
			ActorNameList: source["ActorNameList"].(string),
		}
		films = append(films, film)
		//fmt.Printf("User ID: %d, Name: %s, Email: %s\n", int(source["ID"].(float64)), source["Name"].(string), source["Email"].(string))
	}
	return films, int64(len(films)), err
}
