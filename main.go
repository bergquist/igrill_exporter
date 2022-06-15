package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
)

var (
	tempGuage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "temperature",
		Help:      "",
	}, []string{"probe"})

	thresholdGuage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "threshold",
		Help:      "",
	}, []string{"probe"})

	systemInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "system_info",
		Help:      "",
	}, []string{"model", "firmwarerev", "hardwarerev", "manufacturer"})

	batteryLevel = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "battery",
		Help:      "",
	})

	up = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "up",
		Help:      "",
	})
)

// 70:91:8F:15:C2:E4
var bleAddress = flag.String("bleAddress", "", "bluetooth address of the igrill")

func main() {
	flag.Parse()

	bleAdaptor := ble.NewClientAdaptor(*bleAddress)
	battery := ble.NewBatteryDriver(bleAdaptor)
	info := ble.NewDeviceInformationDriver(bleAdaptor)
	igrill := NewIGrillDriver(bleAdaptor)

	work := func() {
		fmt.Println("Model number:", info.GetModelNumber())
		fmt.Println("Firmware rev:", info.GetFirmwareRevision())
		fmt.Println("Hardware rev:", info.GetHardwareRevision())
		fmt.Println("Manufacturer name:", info.GetManufacturerName())
		systemInfo.WithLabelValues(info.GetModelNumber(), info.GetFirmwareRevision(), info.GetHardwareRevision(), info.GetManufacturerName())

		gobot.Every(5*time.Second, func() {
			fmt.Println("collecting values")

			batteryLevel.Set(float64(battery.GetBatteryLevel()))

			tempGuage.WithLabelValues("probe1").Set(float64(igrill.GetProbe1Temp()))
			tempGuage.WithLabelValues("probe2").Set(float64(igrill.GetProbe2Temp()))
			tempGuage.WithLabelValues("probe3").Set(float64(igrill.GetProbe3Temp()))
			tempGuage.WithLabelValues("probe4").Set(float64(igrill.GetProbe4Temp()))

			thresholdGuage.WithLabelValues("probe1").Set(float64(igrill.GetThreshold1()))
			thresholdGuage.WithLabelValues("probe2").Set(float64(igrill.GetThreshold2()))
			thresholdGuage.WithLabelValues("probe3").Set(float64(igrill.GetThreshold3()))
			thresholdGuage.WithLabelValues("probe4").Set(float64(igrill.GetThreshold4()))
		})
	}

	robot := gobot.NewRobot("bleBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{battery, info, igrill},
		work,
	)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":9999", nil))
	}()

	robot.Start()
}
