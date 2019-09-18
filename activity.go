
package grovepiDigitalRead

import (


	"sync"
	"time"


	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/mrmorphic/hwio"
)

// log is the default package logger
var log = logger.GetLogger("activity-tibco-GrovePi")



const (

	ivPin     = "pin" 
	ovResult = "result"
	
	
	//Cmd format
	DIGITAL_READ = 1
	PIN_MODE      = 5
)



type GrovePi struct {
	
	i2cmodule hwio.I2CModule
	i2cDevice hwio.I2CDevice
}



// Activity is a Activity implementation
type grovePiDRActivity struct {
	sync.Mutex
	metadata *activity.Metadata
}



// NewActivity creates a new Activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &grovePiDRActivity{metadata: metadata}
}



// Metadata implements activity.Activity.Metadata
func (a *grovePiDRActivity) Metadata() *activity.Metadata {
	return a.metadata
}



// Eval → the final output
// Eval implements activity.Activity.Eval
func (a *grovePiDRActivity) Eval(context activity.Context) (done bool, err error) {

	var pin byte
	// var value bool

	log.Debug("Starting Pin Read")
	if context.GetInput(ivPin) != nil {
		pin = byte(context.GetInput(ivPin).(int))
	}

	var g *GrovePi
	g = InitGrovePi(0x04)
	result, err = g.DigitalRead(pin)
//g.PinMode(pin, "input")
	if err != nil {
		log.Error("GrovePi :: Set PinMode Error", err)
	}

	//read to GrovePi
	// if value {
	//	g.DigitalRead(pin, 1)
//	} else {
//		g.DigitalRead(pin, 0) 
//	}




	if result {
	context.SetOutput(ovResult, true)
	} else {
	context.SetOutput(ovResult, false)
	}


// return true → return it as the job is “done” 

	return true, nil
}



func InitGrovePi(address int) *GrovePi {
	grovePi := new(GrovePi)
	m, err := hwio.GetModule("i2c")
	if err != nil {
		log.Error("GrovePi :: could not get i2c module Error", err)
		//fmt.Printf("could not get i2c module: %s\n", err)
		return nil
	}
	grovePi.i2cmodule = m.(hwio.I2CModule)
	grovePi.i2cmodule.Enable()
	grovePi.i2cDevice = grovePi.i2cmodule.GetDevice(address)


	return grovePi
}



func (grovePi GrovePi) CloseDevice() {
	grovePi.i2cmodule.Disable()
}



func (grovePi GrovePi) DigitalRead(pin byte) (bool, error) {
	b := []byte{DIGITAL_READ, pin,1, 0}
	result, err := grovePi.i2cDevice.Read(1, b)
	if err != nil {
		log.Error("GrovePi :: DigitalRead Error", err)

		return 0, err
	}
	
	
	time.Sleep(100 * time.Millisecond)

	return nil
}



func (grovePi GrovePi) PinMode(pin byte, mode string) error {


	var b []byte


	if mode == "input" {
		b = []byte{PIN_MODE, pin, 1, 0}
	} else {
		b = []byte{PIN_MODE, pin, 0, 0}
	}


	result, err := grovePi.i2cDevice.Read(1, b)


	if err != nil {
		log.Error("GrovePi :: i2cDevice.Read Error", err)

		Return 0, err
	}


	time.Sleep(100 * time.Millisecond)


	return result, nil


}
