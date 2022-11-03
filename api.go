package lambda

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

func api(h func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) Response {
	response := Response{
		Payload: Payload{},
	}

	event := events.APIGatewayProxyRequest{}
	err := decode(os.Getenv("LAMBDA_EVENT"), &event)
	if err != nil {
		response.Payload.Error = err.Error()
		return response
	}

	ctxArgs := map[string]interface{}{}
	err = decode(os.Getenv("LAMBDA_CONTEXT"), &ctxArgs)
	if err != nil {
		response.Payload.Error = err.Error()
		return response
	}

	ctx := lambdacontext.NewContext(context.Background(), &lambdacontext.LambdaContext{
		AwsRequestID:       ctxArgs["awsRequestId"].(string),
		InvokedFunctionArn: ctxArgs["invokedFunctionArn"].(string),
	})

	res, err := h(ctx, event)
	if err != nil {
		response.Payload.Error = err.Error()
		return response
	}

	response.Payload.Success = res
	return response
}

func apiV2(h func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error)) Response {
	response := Response{
		Payload: Payload{},
	}

	event := events.APIGatewayV2HTTPRequest{}
	err := decode(os.Getenv("LAMBDA_EVENT"), &event)
	if err != nil {
		response.Payload.Error = err.Error()
		return response
	}

	ctxArgs := map[string]interface{}{}
	err = decode(os.Getenv("LAMBDA_CONTEXT"), &ctxArgs)
	if err != nil {
		response.Payload.Error = err.Error()
		return response
	}

	ctx := lambdacontext.NewContext(context.Background(), &lambdacontext.LambdaContext{
		AwsRequestID:       ctxArgs["awsRequestId"].(string),
		InvokedFunctionArn: ctxArgs["invokedFunctionArn"].(string),
	})

	res, err := h(ctx, event)
	if err != nil {
		response.Payload.Error = err.Error()
		return response
	}

	response.Payload.Success = res
	return response
}
