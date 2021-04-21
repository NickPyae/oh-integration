package main

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"time"
	"strings"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// Default values
const (
	defaultSmtpPort = "587"
	defaultInfluxBucket = "hello-sally-frk"
	defaultQuerySeconds = 15
	defaultAlertThreshold = 1000.0
)

type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

// Sends email using SMTP
func sendEmail(temperature float64, statusTime time.Time, alertTime time.Time) {
	senderEmail := os.Getenv("HS_SENDER_EMAIL")
	senderAppKey := os.Getenv("HS_SENDER_APP_KEY")
	userEmails := strings.Split(os.Getenv("HS_USER_EMAIL"), ",")
	smtpHost := os.Getenv("HS_SMTP_HOST")
	smtpPort := defaultSmtpPort

	// Override default value
	if os.Getenv("HS_SMTP_PORT") != "" {
		smtpPort = os.Getenv("HS_SMTP_PORT")
	}

	from := senderEmail
	password := senderAppKey
	// Receiver email addresses.
	to := userEmails
	// smtp server configuration.
	smtpServer := smtpServer{host: smtpHost, port: smtpPort}
	// Message.
	message := []byte(fmt.Sprintf("From: No-Reply: Hello Sally\r\n"+
		"To: %v\r\n"+
		"Subject: [ALERT] Critical Temperature\r\n"+
		"\r\n"+
		"[ALERT]\r\n"+
		"Device Name:	Random Temperature\r\n"+
		"Temperature:	%.1fF\r\n"+
		"Status:		Critical\r\n"+
		"Event Time:	%v\r\n"+
		"Alert Time:	%v\r\n",
		strings.Join(userEmails, ", "), temperature, statusTime.Format(time.RFC1123), alertTime.Format(time.RFC1123)))
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Email sent")
}

// Query InfluxDB Cloud
func queryInflux(querySeconds int) {
	influxUrl := os.Getenv("HS_INFLUX_URL")
	influxToken := os.Getenv("HS_INFLUX_TOKEN")
	influxOrg := os.Getenv("HS_INFLUX_ORG")
	influxBucket := defaultInfluxBucket
	alertThreshold := defaultAlertThreshold

	// Override default values
	if os.Getenv("HS_INFLUX_BUCKET") != "" {
		influxBucket = os.Getenv("HS_INFLUX_BUCKET")
	}

	if os.Getenv("HS_ALERT_THRESHOLD") != "" {
		var err error
		alertThreshold, err = strconv.ParseFloat(os.Getenv("HS_ALERT_THRESHOLD"), 64)
		if err != nil {
			log.Println("HS_ALERT_THRESHOLD must be float64. Reverting to default value")
			alertThreshold = defaultAlertThreshold
		}
	}

	log.Println("Connecting to InfluxDB Cloud...")
	// Create InfluxDb V2 client
	client := influxdb2.NewClient(influxUrl, influxToken)
	queryAPI := client.QueryAPI(influxOrg)
	// Flux query to get all values above threshold in the last (2 * querySeconds) seconds
	query := fmt.Sprintf(
		`from(bucket: "%s")
			|> range(start: -%ds)
			|> filter(fn: (r) => r["_measurement"] == "hello-sally")
			|> filter(fn: (r) => r["_field"] == "temperature")
			|> filter(fn: (r) => r["_value"] >= %f)`,
		influxBucket, querySeconds*2, alertThreshold)

	// Query results
	result, err := queryAPI.Query(context.Background(), query)
	if err == nil {
		if result.Next() {
			temperature, ok := result.Record().Value().(float64)
			timestamp := result.Record().Time()
			if !ok {
				panic(fmt.Sprintf("expected value of type float64 but got %T", result.Record().Value()))
			}

			log.Printf("[ALERT] Temperature: %.1fF, Time: %v\n", temperature, timestamp.Format(time.RFC1123))
			sendEmail(temperature, timestamp, time.Now())
		} else {
			// No results
			log.Println("No critical temperatures")
		}
	} else {
		panic(err)
	}

	// Close InfluxDB client
	client.Close()
}

func runTicker() {
	querySeconds := defaultQuerySeconds
	if os.Getenv("HS_QUERY_SECONDS") != "" {
		var err error
		querySeconds, err = strconv.Atoi(os.Getenv("HS_QUERY_SECONDS"))
		if err != nil {
			log.Println("HS_QUERY_SECONDS must be integer. Reverting to default value")
			querySeconds = defaultQuerySeconds
		}
	}

	// Ticker triggers goroutine every "tick"
	ticker := time.NewTicker(time.Second * time.Duration(querySeconds))
	defer ticker.Stop()

	log.Println("Email Alert Service started")

	for {
		// Ticker runs until program quits
		<-ticker.C
		go queryInflux(querySeconds)
	}
}

func main() {
	runTicker()
}
