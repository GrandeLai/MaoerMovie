package esClient

import (
	"MaoerMovie/common/kqueue"
	"MaoerMovie/common/utils"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"net/http"
	"strings"
	"time"
)

const FilmIndexName = "film"

type FilmEsModel struct {
	FilmId        int64
	FilmName      string
	FilmTime      time.Time
	FilmCategory  string
	FilmScore     float64
	FilmArea      string
	FilmCoverUrl  string
	ActorNameList string
}

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
		Index: FilmIndexName,
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

// insertFilmToES 将数据插入到Elasticsearch中
func (esModel *EsModel) InsertFilmToES(message kqueue.FilmInsertMessage) error {

	film := FilmEsModel{
		FilmId:        message.FilmId,
		FilmArea:      message.FilmArea,
		FilmCategory:  message.FilmCategory,
		FilmCoverUrl:  message.FilmCoverUrl,
		FilmName:      message.FilmName,
		FilmTime:      utils.StringToTime(message.FilmTime),
		ActorNameList: message.ActorNameList,
	}
	filmJSON, err := json.Marshal(film)
	if err != nil {
		log.Fatalf("Error marshaling film: %v", err)
		return err
	}

	req := esapi.IndexRequest{
		Index:      FilmIndexName,
		Body:       strings.NewReader(string(filmJSON)),
		DocumentID: fmt.Sprintf("%d", message.FilmId),
	}

	res, err := req.Do(context.Background(), esModel.es)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error indexing document ID=%d: %s", message.FilmId, res.String())
		return err
	}
	return nil
}

// 更新文档
func (esModel *EsModel) UpdateFilmToES(message kqueue.FilmUpdateMessage) error {
	film := FilmEsModel{
		FilmId:        message.FilmId,
		ActorNameList: message.ActorNameList,
		FilmArea:      message.FilmArea,
		FilmCategory:  message.FilmCategory,
		FilmCoverUrl:  message.FilmCoverUrl,
		FilmName:      message.FilmName,
		FilmTime:      utils.StringToTime(message.FilmTime),
		FilmScore:     utils.StringToFloat64(message.FilmScore),
	}
	update := map[string]interface{}{
		"doc": film,
	}
	updateJSON, err := json.Marshal(update)
	if err != nil {
		log.Fatalf("Error marshaling update: %v", err)
		return err
	}

	req := esapi.UpdateRequest{
		Index:      "user",
		DocumentID: fmt.Sprintf("%d", message.FilmId),
		Body:       strings.NewReader(string(updateJSON)),
	}

	res, err := req.Do(context.Background(), esModel.es)
	if err != nil {
		log.Printf("Failed to update document ID=%d: %s", message.FilmId, err)
		return err
	}
	defer res.Body.Close()
	return nil
}
