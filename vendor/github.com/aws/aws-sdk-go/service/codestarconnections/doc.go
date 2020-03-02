// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codestarconnections provides the client and types for making API
// requests to AWS CodeStar connections.
//
// This AWS CodeStar Connections API Reference provides descriptions and usage
// examples of the operations and data types for the AWS CodeStar Connections
// API. You can use the Connections API to work with connections and installations.
//
// Connections are configurations that you use to connect AWS resources to external
// code repositories. Each connection is a resource that can be given to services
// such as CodePipeline to connect to a third-party repository such as Bitbucket.
// For example, you can add the connection in CodePipeline so that it triggers
// your pipeline when a code change is made to your third-party code repository.
// Each connection is named and associated with a unique ARN that is used to
// reference the connection.
//
// When you create a connection, the console initiates a third-party connection
// handshake. Installations are the apps that are used to conduct this handshake.
// For example, the installation for the Bitbucket provider type is the Bitbucket
// Cloud app. When you create a connection, you can choose an existing installation
// or create one.
//
// You can work with connections by calling:
//
//    * CreateConnection, which creates a uniquely named connection that can
//    be referenced by services such as CodePipeline.
//
//    * DeleteConnection, which deletes the specified connection.
//
//    * GetConnection, which returns information about the connection, including
//    the connection status.
//
//    * ListConnections, which lists the connections associated with your account.
//
// For information about how to use AWS CodeStar Connections, see the AWS CodePipeline
// User Guide (https://docs.aws.amazon.com/codepipeline/latest/userguide/welcome.html).
//
// See https://docs.aws.amazon.com/goto/WebAPI/codestar-connections-2019-12-01 for more information on this service.
//
// See codestarconnections package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codestarconnections/
//
// Using the Client
//
// To contact AWS CodeStar connections with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS CodeStar connections client CodeStarConnections for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codestarconnections/#New
package codestarconnections