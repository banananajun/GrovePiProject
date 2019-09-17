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
