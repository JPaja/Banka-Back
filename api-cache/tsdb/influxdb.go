package tsdb

import (
	"context"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/pkg/errors"

	"xmudrii.com/api-cache/model"
)

type StockInfluxDBClient struct {
	client influxdb2.Client

	org    string
	bucket string
}

type ForexInfluxDBClient struct {
	client influxdb2.Client

	org    string
	bucket string
}

func NewStockInfluxDBClient(url, token, org, bucket string) TSDBStockClient {
	log.Println(url)
	client := influxdb2.NewClient(url, token)
	client.Options().WriteOptions().SetUseGZip(true)

	return &StockInfluxDBClient{
		client: client,
		org:    org,
		bucket: bucket,
	}
}

func (c StockInfluxDBClient) PushIntradayStockData(ticker string, data []model.IntradayStocks) error {
	writeAPI := c.client.WriteAPIBlocking(c.org, c.bucket)

	points := []*write.Point{}
	for _, d := range data {
		t, err := time.Parse("2006-01-02 15:04:05", d.Time)
		if err != nil {
			return errors.Wrap(err, "converting time")
		}

		p := influxdb2.NewPoint("stock_price_intraday",
			map[string]string{"ticker": ticker},
			map[string]interface{}{
				"open":   d.Open,
				"high":   d.High,
				"low":    d.Low,
				"close":  d.Close,
				"volume": d.Volume,
			},
			t)

		points = append(points, p)
	}

	if err := writeAPI.WritePoint(context.Background(), points...); err != nil {
		return errors.Wrap(err, "failed to push to influxdb")
	}

	log.Println("Push successful...")

	return nil
}

func (c StockInfluxDBClient) PushPeriodicStockData(reqType model.AlphaVantageStockRequestType, ticker string, data []model.PeriodicStocks) error {
	writeAPI := c.client.WriteAPIBlocking(c.org, c.bucket)

	points := []*write.Point{}
	for _, d := range data {
		t, err := time.Parse("2006-01-02", d.Timestamp)
		if err != nil {
			return errors.Wrap(err, "converting time")
		}

		p := influxdb2.NewPoint("stock_price_"+string(reqType),
			map[string]string{"ticker": ticker},
			map[string]interface{}{
				"open":   d.Open,
				"high":   d.High,
				"low":    d.Low,
				"close":  d.Close,
				"volume": d.Volume,
			},
			t)

		points = append(points, p)
	}

	if err := writeAPI.WritePoint(context.Background(), points...); err != nil {
		return errors.Wrap(err, "failed to push to influxdb")
	}

	log.Println("Push successful...")

	return nil
}

func NewForexInfluxDBClient(url, token, org, bucket string) TSDBForexClient {
	log.Println(url)
	client := influxdb2.NewClient(url, token)
	client.Options().WriteOptions().SetUseGZip(true)

	return &ForexInfluxDBClient{
		client: client,
		org:    org,
		bucket: bucket,
	}
}

func (c ForexInfluxDBClient) PushIntradayForexData(from string, to string, data []model.IntradayForex) error {
	writeAPI := c.client.WriteAPIBlocking(c.org, c.bucket)

	points := []*write.Point{}
	for _, d := range data {
		t, err := time.Parse("2006-01-02 15:04:05", d.Timestamp)
		if err != nil {
			return errors.Wrap(err, "converting time")
		}

		p := influxdb2.NewPoint("forex_price_from_intraday",
			map[string]string{"title": from + "_" + to},
			map[string]interface{}{
				"open":  d.Open,
				"high":  d.High,
				"low":   d.Low,
				"close": d.Close,
			},
			t)

		points = append(points, p)
	}

	if err := writeAPI.WritePoint(context.Background(), points...); err != nil {
		return errors.Wrap(err, "failed to push to influxdb")
	}

	log.Println("Push successful...")

	return nil
}

func (c ForexInfluxDBClient) PushPeriodicForexData(reqType model.AlphaVantageForexRequestType, from, to string, data []model.PeriodicForex) error {
	writeAPI := c.client.WriteAPIBlocking(c.org, c.bucket)

	points := []*write.Point{}
	for _, d := range data {
		t, err := time.Parse("2006-01-02", d.Timestamp)
		if err != nil {
			return errors.Wrap(err, "converting time")
		}

		p := influxdb2.NewPoint("forex_price_"+string(reqType),
			map[string]string{"title": from + "_" + to},
			map[string]interface{}{
				"open":  d.Open,
				"high":  d.High,
				"low":   d.Low,
				"close": d.Close,
			},
			t)

		points = append(points, p)
	}

	if err := writeAPI.WritePoint(context.Background(), points...); err != nil {
		return errors.Wrap(err, "failed to push to influxdb")
	}

	log.Println("Push successful...")

	return nil
}
