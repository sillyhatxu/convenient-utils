package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type Todo struct {
	Name        string
	Description string
}

const (
	delete_sql = `
		DELETE FROM stock_keeping_unit 
		WHERE 
		{{if .ProductId }}
			product_id = {{ .ProductId }} AND order_id IS NULL
		{{else}}
			variation_id = {{ .VariationId }} AND order_id IS NULL LIMIT {{ .DeleteAmount }}
        {{end}}
	`
	get_stock_by_variation_ids_sql = `
		SELECT variation_id, product_id, count(*) as stockAmount
		FROM stock_keeping_unit
		WHERE order_id IS NULL 
		AND variation_id IN ({{ . }})
		GROUP BY variation_id, product_id
	`
)

type Stock struct {
	Id           string `mapstructure:"id"`
	VariationId  string `mapstructure:"variation_id"`
	ProductId    string `mapstructure:"product_id"`
	DeleteAmount int    `mapstructure:"product_id"`
}

type StockProductId struct {
	ProductId string `mapstructure:"product_id"`
}

type StockVariationId struct {
	VariationId  string `mapstructure:"variation_id"`
	DeleteAmount int    `mapstructure:"product_id"`
}

func main() {
	fmt.Println("--------------")
	fmt.Println(GetSQL(delete_sql, Stock{Id: "1", VariationId: "VARIATION_ID", ProductId: "321", DeleteAmount: 1}))
	fmt.Println("--------------")
	fmt.Println(GetSQL(delete_sql, Stock{VariationId: "VARIATION_ID", DeleteAmount: 1}))
	fmt.Println("--------------")
	fmt.Println(GetSQL(delete_sql, Stock{ProductId: "P1001"}))
	fmt.Println("--------------")
	src := [3]string{"VARIATION_ID_101", "VARIATION_ID_102", "VARIATION_ID_103"}
	var buffer bytes.Buffer
	for i := 0; i < len(src); i++ {
		buffer.WriteString(src[i])
		buffer.WriteString(",")
	}
	var variationIdsString = buffer.String()
	if last := len(variationIdsString) - 1; last >= 0 {
		variationIdsString = variationIdsString[:last]
	}
	fmt.Println(variationIdsString)
	fmt.Println(GetSQL(get_stock_by_variation_ids_sql, variationIdsString))
}

func GetSQL(sql string, data interface{}) (string, error) {
	temp, err := template.New("sql").Parse(sql)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := temp.Execute(&tpl, data); err != nil {
		return "", err
	}
	return tpl.String(), nil
}
