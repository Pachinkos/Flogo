package sendmailevent

import (
	"github.com/project-flogo/core/data/coerce"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

type Settings struct {
	EmailConfig map[string]interface{} `md:"emailConfig"`
}

type Input struct {
	ToEmail string `md:"toEmail"`
	Event map[string]string `md:"event"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"toEmail": i.ToEmail,
		"event":  i.Event,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
	i.ToEmail, err = coerce.ToString(values["toEmail"])
	if err != nil {
		return err
	}
	i.Event, err = coerce.ToParams(values["event"])

	return err
}

type Output struct {
	Success bool `md:"success"` // The result of the counter operation
}

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

func New(ctx activity.InitContext) (activity.Activity, error) {
	settings := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), settings, true)
	if err != nil {
		return nil, err
	}
	act := &Activity{settings: settings}
	//logger := ctx.Logger()

	return act, nil
}

// Activity is an activity that is used to invoke a REST Operation
// settings : {method, uri, headers, proxy, skipSSL}
// input    : {pathParams, queryParams, headers, content}
// outputs  : {status, result}
type Activity struct {
	settings      *Settings
}

func (a *Activity) SendMail() (bool, error){
	return true, nil
}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(context activity.Context) (done bool, err error) {
	input := &Input{}
	err = context.GetInputObject(input)
	if err != nil {
		return false, err
	}

	var success bool
	success, err = a.SendMail()
	err = context.SetOutput("success", success)
	if err != nil {
		return false, err
	}

	return true, nil
}
