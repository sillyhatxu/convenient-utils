package esclient

import (
	"context"
	"errors"
	"github.com/olivere/elastic"
	"log"
	"os"
	"time"
)

type Client struct {
	URL string

	ElasticIndex string

	ElasticType string

	CTX context.Context
}

type BulkEntity struct {
	Id string

	Data interface{}

	IsDelete bool
}

func NewClient(url string, elasticIndex string, elasticType string) *Client {
	return &Client{
		URL:          url,
		ElasticIndex: elasticIndex,
		ElasticType:  elasticType,
		CTX:          context.Background(),
	}
}

func (client *Client) NewElasticClient() (*elastic.Client, error) {
	elasticClient, err := elastic.NewClient(
		elastic.SetURL(client.URL),
		elastic.SetSniff(false),
		elastic.SetHealthcheckTimeout(60*time.Second),
		elastic.SetHealthcheckInterval(60*time.Second),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		return nil, err
	}
	return elasticClient, nil
}

func (client *Client) IndexExists() (bool, error) {
	elasticClient, err := client.NewElasticClient()
	if err != nil {
		return false, err
	}
	exists, err := elasticClient.IndexExists(client.ElasticIndex).Do(client.CTX)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (client *Client) CreateIndex() (bool, error) {
	elasticClient, err := client.NewElasticClient()
	if err != nil {
		return false, err
	}
	index, err := elasticClient.CreateIndex(client.ElasticIndex).Do(client.CTX)
	if err != nil {
		return false, err
	}
	if !index.Acknowledged {
		return false, errors.New("Not acknowledged.")
	}
	return true, nil
}

func (client *Client) DeleteIndex() (bool, error) {
	elasticClient, err := client.NewElasticClient()
	if err != nil {
		return false, err
	}
	deleteIndex, err := elasticClient.DeleteIndex(client.ElasticIndex).Do(client.CTX)
	if err != nil {
		return false, err
	}
	if !deleteIndex.Acknowledged {
		return false, errors.New("Not acknowledged")
	}
	return true, nil
}

func (client *Client) BulkAll(bulkEntityArray []BulkEntity) (*elastic.BulkResponse, error) {
	elasticClient, err := client.NewElasticClient()
	if err != nil {
		return nil, err
	}
	bulk := elasticClient.Bulk()
	for _, bulkEntity := range bulkEntityArray {
		request := elastic.NewBulkIndexRequest().Index(client.ElasticIndex).Type(client.ElasticType).Id(bulkEntity.Id).Doc(bulkEntity.Data)
		bulk.Add(request)
	}
	bulkResponse, err := bulk.Do(client.CTX)
	if err != nil {
		return nil, err
	}
	return bulkResponse, nil
}

func (client *Client) Bulk(bulkEntityArray []BulkEntity) (*elastic.BulkResponse, error) {
	elasticClient, err := client.NewElasticClient()
	if err != nil {
		return nil, err
	}
	bulk := elasticClient.Bulk()
	mgetService := elasticClient.MultiGet()
	for _, bulkEntity := range bulkEntityArray {
		mgetService = mgetService.Add(elastic.NewMultiGetItem().Index(client.ElasticIndex).Type(client.ElasticType).Id(bulkEntity.Id))
	}
	response, err := mgetService.Do(client.CTX)
	if err != nil {
		return nil, err
	}
	for _, bulkEntity := range bulkEntityArray {
		if checkExists(response.Docs, bulkEntity.Id) {
			if bulkEntity.IsDelete {
				request := elastic.NewBulkDeleteRequest().Index(client.ElasticIndex).Type(client.ElasticType).Id(bulkEntity.Id)
				bulk.Add(request)
			} else {
				request := elastic.NewBulkUpdateRequest().Index(client.ElasticIndex).Type(client.ElasticType).Id(bulkEntity.Id).Doc(bulkEntity.Data)
				bulk.Add(request)
			}
		} else {
			if !bulkEntity.IsDelete {
				request := elastic.NewBulkIndexRequest().Index(client.ElasticIndex).Type(client.ElasticType).Id(bulkEntity.Id).Doc(bulkEntity.Data)
				bulk.Add(request)
			}
		}
	}
	bulkResponse, err := bulk.Do(client.CTX)
	if err != nil {
		return nil, err
	}
	return bulkResponse, nil
}

func (client *Client) Get(id string) (*elastic.GetResult, error) {
	elasticClient, err := client.NewElasticClient()
	if err != nil {
		return nil, err
	}
	getResult, err := elasticClient.Get().Index(client.ElasticIndex).Type(client.ElasticType).Id(id).Do(client.CTX)
	if err != nil {
		return nil, err
	}
	return getResult, nil
}

type SearchParams struct {
	Index     []string
	TermQuery elastic.Query
	Sort      []map[string]string
	From      int
	Size      int
	Pretty    bool
}

func (client *Client) Search(params SearchParams) (*elastic.SearchResult, error) {
	elasticClient, err := client.NewElasticClient()
	if err != nil {
		return nil, err
	}
	//searchResult, err := elasticClient.Search(params.Index...).Query(params.TermQuery).Sort("user", true).From(0).Size(500).Pretty(true).Do(client.CTX)
	searchResult, err := elasticClient.Search(params.Index...).Query(params.TermQuery).From(params.From).Size(params.Size).Pretty(params.Pretty).Do(client.CTX)
	if err != nil {
		return nil, err
	}
	return searchResult, nil
}

func (client *Client) Query() (*elastic.MgetResponse, error) {
	elasticClient, err := client.NewElasticClient()
	if err != nil {
		return nil, err
	}
	//query := elastic.NewBoolQuery()
	//query = query.Must(elastic.NewTermQuery("user", "olivere"))
	//query = query.Filter(elastic.NewTermQuery("account", 1))
	//src, err := query.Source()
	//if err != nil {
	//	return nil, err
	//}
	//data, err := json.MarshalIndent(src, "", "  ")
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println(string(data))
	response, err := elasticClient.MultiGet().Do(client.CTX)
	//mgetResponse, err := elasticClient.MultiGet().Do(client.CTX)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (client *Client) Update(id string, doc interface{}) (*elastic.UpdateResponse, error) {
	elasticClient, err := client.NewElasticClient()
	if err != nil {
		return nil, err
	}
	update, err := elasticClient.Update().Index(client.ElasticIndex).Type(client.ElasticType).Id(id).Upsert(doc).Do(client.CTX)
	if err != nil {
		return nil, err
	}
	return update, nil
}

func (client *Client) Delete(id string) (int64, error) {
	termQuery := elastic.NewTermQuery("id", id)
	elasticClient, err := client.NewElasticClient()
	if err != nil {
		return 0, err
	}
	deleteResponse, err := elasticClient.DeleteByQuery(client.ElasticIndex).Query(termQuery).Do(client.CTX)
	if err != nil {
		return 0, err
	}
	return deleteResponse.Deleted, nil
}

func checkExists(resultArray []*elastic.GetResult, id string) bool {
	for _, result := range resultArray {
		if result.Id == id {
			return result.Found
		}
	}
	return false
}
