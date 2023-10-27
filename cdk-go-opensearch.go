package main

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkGoOpensearchStackProps struct {
	cdkProps       awscdk.StackProps
	collectionName string
	principalARN   string
}

func NewCdkGoOpensearchStack(scope constructs.Construct, id string, props *CdkGoOpensearchStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.cdkProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	collectionName := props.collectionName
	collection := awsopensearchserverless.NewCfnCollection(stack, jsii.String("collection"), &awsopensearchserverless.CfnCollectionProps{
		Name: &collectionName,
		Type: jsii.String("SEARCH"),
	})

	encryptionPolicy := awsopensearchserverless.NewCfnSecurityPolicy(stack, jsii.String("encryptionPolicy"), &awsopensearchserverless.CfnSecurityPolicyProps{
		Name: jsii.String("myEncryptionPolicy"),
		Type: jsii.String("encryption"),
		Policy: jsii.String(`{
			"Rules":[
			   {
				  "ResourceType":"collection",
				  "Resource":[
					 "collection/mycollection"
				  ]
			   }
			],
			"AWSOwnedKey":true
		 }`),
	})
	collection.AddDependency(encryptionPolicy)

	networkPolicy := awsopensearchserverless.NewCfnSecurityPolicy(stack, jsii.String("networkPolicy"), &awsopensearchserverless.CfnSecurityPolicyProps{
		Name: jsii.String("myNetworkPolicy"),
		Type: jsii.String("network"),
		Policy: jsii.String(`[
			{
			  "Description": "Dashboards access",
			  "Rules": [
				{
				  "ResourceType": "dashboard",
				  "Resource": [
					"collection/mycollection"
				  ]
				}
			  ],
			  "AllowFromPublic": true
			}
		  ]`),
	})
	collection.AddDependency(networkPolicy)

	dataPolicy := awsopensearchserverless.NewCfnSecurityPolicy(stack, jsii.String("dataPolicy"), &awsopensearchserverless.CfnSecurityPolicyProps{
		Name: jsii.String("myDataPolicy"),
		Type: jsii.String("data"),
		Policy: jsii.String(fmt.Sprintf(`[
			{
			   "Description": "Rule 1",
			   "Rules":[
				  {
					 "ResourceType":"collection",
					 "Resource":[
						"collection/mycollection"
					 ],
					 "Permission":[
						"aoss:CreateCollectionItems",
						"aoss:UpdateCollectionItems",
						"aoss:DescribeCollectionItems"
					 ]
				  },
				  {
					 "ResourceType":"index",
					 "Resource":[
						"index/mycollection/*"
					 ],
					 "Permission":[
						"aoss:*"
					 ]
				  }
			   ],
			   "Principal":[
				  "%s",
			   ]
			}
		 ]`, props.principalARN)),
	})
	collection.AddDependency(dataPolicy)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	accountCtx := app.Node().TryGetContext(jsii.String("account"))
	if accountCtx == nil {
		panic("account is required")
	}
	account := accountCtx.(string)

	regionCtx := app.Node().TryGetContext(jsii.String("region"))
	if regionCtx == nil {
		panic("region is required")
	}
	region := regionCtx.(string)

	collectionNameCtx := app.Node().TryGetContext(jsii.String("collection-name"))
	if collectionNameCtx == nil {
		panic("collection-name is required")
	}
	collectionName := collectionNameCtx.(string)

	principalARNCtx := app.Node().TryGetContext(jsii.String("principal-arn"))
	if principalARNCtx == nil {
		panic("principal-arn is required")
	}
	principalARN := principalARNCtx.(string)

	NewCdkGoOpensearchStack(app, "CdkGoOpensearchStack", &CdkGoOpensearchStackProps{
		cdkProps: awscdk.StackProps{
			Env: env(account, region),
		},
		collectionName: collectionName,
		principalARN:   principalARN,
	})

	app.Synth(nil)
}

func env(account, region string) *awscdk.Environment {
	return &awscdk.Environment{
		Account: &account,
		Region:  &region,
	}
}
