package grovepiDigitalRead

import (
	"testing"
	"fmt"
	
	"github.com/stretchr/testify/assert"

	"io/ioutil"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"

	
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestGrovePiDR(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs GrovePi Tester, using Pin 3
	tc.SetInput(ivPin, 3)
	

	act.Eval(tc)

	result := tc.GetOutput(ovResult).(bool)
	
	assert.NotNil(t, result)
	
	fmt.Printf("Result: %t", result)
}
