package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var body = "{\"payload\":\"{\\\"afterSaleNo\\\":\\\"AFTER17129786130000\\\",\\\"afterSaleType\\\":2,\\\"bizCode\\\":\\\"mall\\\",\\\"erpAfterSaleStage\\\":0,\\\"mqOrderAddress\\\":{\\\"addressDetail\\\":\\\"\\\",\\\"addressId\\\":\\\"563\\\",\\\"addressName\\\":\\\"\\\",\\\"cityId\\\":\\\"0\\\",\\\"cityName\\\":\\\"Tarragona\\\",\\\"consigneeMobile\\\":\\\"+34 654456652\\\",\\\"consigneeName\\\":\\\"carmen\\\",\\\"country\\\":\\\"ES\\\",\\\"deliverable\\\":false,\\\"districtId\\\":\\\"0\\\",\\\"districtName\\\":\\\"\\\",\\\"extInfo\\\":\\\"{\\\\\\\"addressRemark\\\\\\\":\\\\\\\"\\\\\\\",\\\\\\\"beforeDoor\\\\\\\":false}\\\",\\\"gender\\\":0,\\\"location\\\":\\\"1.2417141,41.1199637\\\",\\\"orderNo\\\":\\\"OD17129766110000\\\",\\\"orderSubNo\\\":\\\"OS17129766110001\\\",\\\"postCode\\\":\\\"43005\\\",\\\"provinceId\\\":\\\"0\\\",\\\"provinceName\\\":\\\"CT\\\",\\\"streetId\\\":\\\"0\\\",\\\"streetName\\\":\\\"Comissaria de la Policia Nacional\\\",\\\"userId\\\":\\\"200236\\\"},\\\"operatorId\\\":\\\"200236\\\",\\\"operatorName\\\":\\\"user\\\",\\\"orderMode\\\":1,\\\"orderNo\\\":\\\"OD17129766110000\\\",\\\"orderSkuList\\\":[{\\\"afterSaleNo\\\":\\\"AFTER17129786130000\\\",\\\"attrInfoList\\\":[{\\\"attrName\\\":\\\"colour\\\",\\\"attrValueName\\\":\\\"Light Brown 50CM\\\"}],\\\"categoryLevel1\\\":28,\\\"given\\\":false,\\\"orderNo\\\":\\\"OD17129766110000\\\",\\\"orderSubNo\\\":\\\"OS17129766110001\\\",\\\"refundAmount\\\":39.00,\\\"returnNum\\\":2,\\\"skuId\\\":\\\"SKU0317127999350101\\\",\\\"skuNum\\\":2}],\\\"orderSubNo\\\":\\\"OS17129766110001\\\",\\\"orderTime\\\":1712976611000,\\\"payNo\\\":\\\"T1PT00120240413025011000001\\\",\\\"refundCoupon\\\":true,\\\"refundMode\\\":1,\\\"refundNo\\\":\\\"RF17129786130001\\\",\\\"refundType\\\":2,\\\"returnFinish\\\":false,\\\"returnInfo\\\":{},\\\"sendMqTimestamp\\\":1712978613087}\",\"headers\":{\"tag\":\"create_after_sale\",\"id\":\"51f88707-efe4-2cae-cc99-36ee0d4113d8\",\"keys\":\"OS17129766110001\",\"timestamp\":1712978613087}}"

// Package main implements a simple producer to send message.
func main() {
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"http://47.92.147.114:9876"})),
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
	topic := "TOPIC_REFUND_CALLBACK_test"

	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte(body),
	}
	res, err := p.SendSync(context.Background(), msg)

	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}

	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
