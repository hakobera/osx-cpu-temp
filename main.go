package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

func Connect(brokerUri string, clientId string, username string, password string) (*MQTT.MqttClient, error) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(brokerUri)
	opts.SetClientId(clientId)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(f)

	client := MQTT.NewClient(opts)
	_, err := client.Start()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func Publish(client *MQTT.MqttClient, topic string, payload []byte) {
	message := MQTT.NewMessage(payload)
	message.SetQoS(0)
	receipt := client.PublishMessage(topic, message)

	fmt.Println("publish:")
	fmt.Printf("TOPIC: %s\n", topic)
	fmt.Printf("MSG: %s\n", message.Payload())

	<-receipt
}

var f MQTT.MessageHandler = func(client *MQTT.MqttClient, message MQTT.Message) {
	fmt.Println("subscribe:")
	fmt.Printf("TOPIC: %s\n", message.Topic())
	fmt.Printf("MSG: %s\n", message.Payload())
}

func Subscribe(client *MQTT.MqttClient, topic string) {
	filter, _ := MQTT.NewTopicFilter(topic, byte(MQTT.QOS_ZERO))
	if receipt, err := client.StartSubscription(f, filter); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		<-receipt
	}
}

func Temp() ([]byte, error) {
	cmd := exec.Command("./osx-cpu-temp")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	r := bufio.NewReader(stdout)
	line, _, err := r.ReadLine()

	return line, nil
}

func main() {
	username := os.Getenv("SANGO_USERNAME")
	password := os.Getenv("SANGO_PASSWORD")

	cli, err := Connect("tcp://free.mqtt.shiguredo.jp:1883", "cputemp", username, password)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	temp, err := Temp()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	Subscribe(cli, username+"/#")
	Publish(cli, username+"/cputemp", temp)

	time.Sleep(3 * time.Second)
	cli.Disconnect(25)
}
