package grovepiDigitalRead

import (
	"sync"
	"time"

	"unsafe"
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
	// added ":" to define result
	//result, 
	result,err := g.DigitalRead(pin,"input") 
//g.DigitalRead(pin)
	if err != nil {
		log.Error("GrovePi :: DigitalRead issue ", err)
	}

	
	context.SetOutput(ovResult, result)

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

//func (grovePi GrovePi) DigitalRead(pin byte) (byte, error) {
//	b := []byte{DIGITAL_READ, pin, 0, 0}
//	result, err := grovePi.PinMode(b, "input")
//	if err != nil {
//		log.Error("GrovePi :: DigitalRead Error", err)

//		return 0, err
//	}
	
	
//	time.Sleep(100 * time.Millisecond)
	
//	return result, nil
//}
// val --> value
func (grovePi *GrovePi) DigitalRead(pin byte, mode string) (bool,error) {
	// b := []byte{DIGITAL_READ, pin, 1, 0}
	rawdata, err := grovePi.PinMode(pin, mode)
	if err != nil {
		return false, err
	}
	data := rawdata[1:5]

	dataInt := int32(data[0]) | int32(data[1])<<8 | int32(data[2])<<16 | int32(data[3])<<24
	d := (*(*int32)(unsafe.Pointer(&dataInt)))
	boolResult := !(d == 0) // come back to this later
	time.Sleep(100 * time.Millisecond)
	return boolResult,nil
}


func (grovePi GrovePi) PinMode(pin byte, mode string) ([]byte, error) {
	var b []byte
	if mode == "input" {
		b = []byte{PIN_MODE, pin, 0, 0}
	} else {
		b = []byte{PIN_MODE, pin, 1, 0}
	}
	err := grovePi.i2cDevice.Write(1, b)
	if err != nil {
		log.Error("GrovePi :: i2cDevice.Read Error", err)
		return nil, err
	}
	time.Sleep(100 * time.Millisecond)
	grovePi.i2cDevice.ReadByte(1)
	time.Sleep(100 * time.Millisecond)
	result, err := grovePi.i2cDevice.Read(1, 9)
	if err != nil {
		return nil,err
	}
	
	// if result {
	//return nil
	//}

	return result, nil
}
