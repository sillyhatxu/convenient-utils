package esclient

import (
	"fmt"
	"github.com/olivere/elastic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientQuery(t *testing.T) {
	client := NewClient("http://127.0.0.1:9200", "shopintar", "product")
	termQuery := elastic.NewTermQuery("brand_id.keyword", "B3358614859922622059")
	response, err := client.Search(SearchParams{
		Index:     []string{"shopintar"},
		TermQuery: termQuery,
		From:      0,
		Size:      1000,
		Pretty:    true,
	})
	assert.Nil(t, err)
	fmt.Println(response)
	for _, searchHit := range response.Hits.Hits {
		dataJSON, err := searchHit.Source.MarshalJSON()
		assert.Nil(t, err)
		fmt.Println(string(dataJSON))
	}
}
