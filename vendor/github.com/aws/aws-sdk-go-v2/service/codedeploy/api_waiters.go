// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilDeploymentSuccessful uses the CodeDeploy API operation
// GetDeployment to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDeploymentSuccessful(ctx context.Context, input *GetDeploymentInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDeploymentSuccessful",
		MaxAttempts: 120,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "deploymentInfo.status",
				Expected: "Succeeded",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "deploymentInfo.status",
				Expected: "Failed",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "deploymentInfo.status",
				Expected: "Stopped",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *GetDeploymentInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.GetDeploymentRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}