package grovepiDigitalRead

import (

  "github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/mrmorphic/hwio"
  )
  
  
  // log is the default package logger
var log = logger.GetLogger("activity-tibco-GrovePi")

const (

  ivPin     = "pin"
	ivValue   = "value"
	ovSuccess = "success"

	//Cmd format
	DIGITAL_READ = 2
	PIN_MODE      = 5

)


type GrovePi struct {
	i2cmodule hwio.I2CModule
	i2cDevice hwio.I2CDevice
	
}


//activity is Activity implementation
type grovePiDWActivity struct {
	sync.Mutex	
	metadata *activity.Metadata
}

// NewActivity creates a new Activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &grovePiDWActvitiy{metadata: metadata}	
}

// Metadata implements.Activity.Metadata
func (a *grovePiDWActivity) Eval(context activity.Context) (done bool, err error) {
	
	
	var pin byte
	var value bool
	
	log.Debug("Starting Pin Write")
	
	// Reading if the input is empty or not empty
	if context.GetInput(ivPin) != nil {
		pin = byte(context.GetInput(ivPin).(int))
	}
	if context.GetInput(ivValue) != nil {
		value = context.GetInput(ivValue).(bool)
	}
	
	var g *GrovePi
	g = InitGrovePi(0x04)
	err = g.PinMode(pin, "input")
	if err != nil {
		log.Error("GrovePi :: Set PinMode Error", err)
	}
	
	//read the GrovePi
	
	
	
