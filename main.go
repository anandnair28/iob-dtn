package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"iob-dtn/env"
	"iob-dtn/env/manager"
	"iob-dtn/env/delaymanager"
	"iob-dtn/env/policy"
	"iob-dtn/env/sensor"
	"iob-dtn/env/util"
)

func main() {
	var simulationTime int
	flag.IntVar(&simulationTime, "env-time", 30, "simulation time in second")
	flag.Parse()

	policies := []policy.Policy{
		policy.New(policy.KONP_POLICY),
		policy.New(policy.NP_POLICY),
		policy.New(policy.GPP_POLICY),
		policy.New(policy.LC_POLICY),
        policy.New(policy.RPP_POLICY),
        //policy.New(policy.RPP1_POLICY),
        //policy.New(policy.RPP2_POLICY),
        //policy.New(policy.RPP3_POLICY),
        //policy.New(policy.RPP4_POLICY),
	}

	names := []string{"konp", "np", "gpp", "lc", "rpp"}
    //names := []string{"konp"}
    //names := []string{"np"}
    //names := []string{"gpp"}
    //names := []string{"lc"}
    //names := []string{"rpp"}
    //names := []string{"rpp1", "rpp2", "rpp3", "rpp4"}
    //names := []string{"rpp1"}
    //names := []string{"rpp2"}
    //names := []string{"rpp3"}
    //names := []string{"rpp4"}

	for i, p := range policies {
		RunSimulationWithPolicy(p, simulationTime, names[i])
	}

}

func RunSimulationWithPolicy(p policy.Policy, simulationTime int, name string) {
	manager.Erase()
	delaymanager.Erase()

	spos := []util.Position{
		util.Position{X: 0, Y: 0},
		util.Position{X: 25, Y: 25},
		util.Position{X: 0, Y: 25},
		util.Position{X: 25, Y: 0},
		util.Position{X: 5, Y: 15},
		util.Position{X: 12, Y: 25},
		util.Position{X: 25, Y: 12},
		util.Position{X: 13, Y: 6},
		util.Position{X: 10, Y: 14},
	}

	stop := make(chan bool)
	var sensors []sensor.Sensor

	for i := 1; i <= len(spos)*int(env.Num_cycles_per_station); i++ {
		manager.Register(i)
		sensors = append(sensors, sensor.New(i, p))
	}

	envr := env.New(spos, sensors, stop)
	envr.StartSimulation()

	time.Sleep(time.Duration(simulationTime) * time.Second)
	for i := 1; i <= 7*int(env.Num_cycles_per_station); i++ {
		stop <- true
	}

	filename1 := fmt.Sprintf("result_lr_%s.csv", name)
	file1, err1 := os.Create(filename1)
	if err1 != nil {
		log.Printf("can not create %s\n", filename1)
		return
	}

	defer file1.Close()
	fmt.Fprint(file1, manager.Manager)
	fmt.Printf("Loss Rate result for %s written\n", name)

	filename2 := fmt.Sprintf("result_delay_%s.csv", name)
	file2, err2 := os.Create(filename2)
	if err2 != nil {
		log.Printf("can not create %s\n", filename2)
		return
	}

	defer file2.Close()
	fmt.Fprintf(file2, "%v\n",delaymanager.DelayManager)
	fmt.Printf("Delay result for %s written\n", name)
}
