package grovepiDigitalRead

import (
	"sync"
	"time"


	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/mrmorphic/hwio"
)

// log is the default package logger
var log = logge }



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
	// added ":" to define result
	
	result, err := g.DigitalRead(pin)

	if err != nil {
		log.Error("GrovePi :: DigitalRead issue ", err)
	}
	
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
	b := []bool{DIGITAL_READ, pin, 1, 0}
	result, err := grovePi.i2cDevice.Read(1, b)
	if err != nil {
		log.Error("GrovePi :: DigitalRead Error", err)

		return false, err
	}
	
	
	time.Sleep(100 * time.Millisecond)

	return result, nil
}



func (grovePi GrovePi) PinMode(pin byte, mode string) (bool, error) {


	var b []bool

	if mode == "input" {
		b = []byte{PIN_MODE, pin, 1, 0}
	} else {
		b = []byte{PIN_MODE, pin, 0, 0}
	}


	result, err := grovePi.i2cDevice.Read(1, b)


	if err != nil {
		log.Error("GrovePi :: i2cDevice.Read Error", err)

		return false, err
	}


	time.Sleep(100 * time.Millisecond)


	return result, nil


}
