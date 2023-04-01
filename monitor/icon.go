package monitor

import "strings"

const cdn = "https://cdn.megaproaktiv.de/img/icons/"

// Which icon to display
func Icon(t *string) *string {
	var icon string
	switch *t {
	case "AWS::ApiGateway::RestApi":
		icon = "https://icons.terrastruct.com/aws%2FNetworking%20&%20Content%20Delivery%2FAmazon-API-Gateway.svg"
	case "AWS::AppSync::GraphQLApi":
		icon = "https://icons.terrastruct.com/aws%2FApplication%20Integration%2FAmazon-AppSync_light-bg.svg"
	case "AWS::AutoScaling::AutoScalingGroup":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-Auto-Scaling_light-bg.svg"
	case "AWS::CertificateManager::Certificate":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-Certificate-Manager_light-bg.svg"
	case "AWS::CloudFormation::Stack":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-CloudFormation_Stack_light-bg.svg"
	case "AWS::CloudFront::Distribution", "AWS::CloudFront::CloudFrontOriginAccessIdentity":
		icon = "https://icons.terrastruct.com/aws%2FNetworking%20&%20Content%20Delivery%2FAmazon-CloudFront.svg"
	case "AWS::CloudWatch::Alarm", "AWS::CloudWatch::CompositeAlarm":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAmazon-CloudWatch_Alarm_light-bg.svg"
	case "AWS::CloudWatch::InsightRule":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAmazon-CloudWatch_Rule_light-bg.svg"
	case "AWS::CodeBuild::Project":
		icon = "https://icons.terrastruct.com/aws%2FDeveloper%20Tools%2FAWS-CodeBuild.svg"
	case "AWS::CodeCommit::Repository":
		icon = "https://icons.terrastruct.com/aws%2FDeveloper%20Tools%2FAWS-CodeCommit.svg"
	case "AWS::DataSync::Agent":
		icon = "https://icons.terrastruct.com/aws%2FMigration%20&%20Transfer%2FAWS-DataSync_Agent_light-bg.svg"
	case "AWS::DynamoDB::Table":
		icon = "https://icons.terrastruct.com/aws%2FDatabase%2FAmazon-DynamoDB.svg"
	case "AWS::EC2::Instance":
		icon = "https://icons.terrastruct.com/aws%2FCompute%2F_Instance%2FAmazon-EC2_Instance_light-bg.svg"
	case "AWS::EC2::EIP":
		icon = "https://icons.terrastruct.com/aws%2FCompute%2FAmazon-EC2_Elastic-IP-Address_light-bg.svg"
	case "AWS::EC2::FlowLog":
		icon = cdn + "Res_Amazon-VPC_Flow-Logs_48_Light.svg"
	case "AWS::EC2::InternetGateway":
		icon = cdn + "Res_Amazon-VPC_Internet-Gateway_48_Light.svg"
	case "AWS::EC2::VPNGateway":
		icon = "https://icons.terrastruct.com/aws%2FNetworking%20&%20Content%20Delivery%2FAWS-Site-to-Site-VPN.svg"
	case "AWS::EC2::VPC":
		icon = "https://icons.terrastruct.com/aws%2F_Group%20Icons%2FVirtual-private-cloud-VPC_light-bg.svg"
	case "AWS::EC2::Subnet::Public":
		icon = "https://icons.terrastruct.com/aws%2F_Group%20Icons%2FVPC-subnet-public_light-bg.svg"
	case "AWS::EC2::Subnet::Private":
		icon = "https://icons.terrastruct.com/aws%2F_Group%20Icons%2FVPC-subnet-private_light-bg.svg"
	case "AWS::ECR::Repository":
		icon = "https://icons.terrastruct.com/aws%2FCompute%2FAmazon-EC2-Container-Registry.svg"
	case "AWS::ECS::Service":
		icon = "https://icons.terrastruct.com/aws%2FCompute%2FAmazon-Elastic-Container-Service_Service_light-bg.svg"
	case "AWS::ECS::TaskDefinition", "AWS::ECS::TaskSet", "AWS::ECS::PrimaryTaskSet":
		icon = cdn + "Res_Amazon-Elastic-Container-Service_Task_48_Light.svg"
	case "AWS::EFS::FileSystem":
		icon = "https://icons.terrastruct.com/aws%2FStorage%2FAmazon-Elastic-File-System_EFS_File-system_light-bg.svg"
	case "AWS::ElasticLoadBalancingV2::LoadBalancer":
		icon = "https://icons.terrastruct.com/aws%2FNetworking%20&%20Content%20Delivery%2FElastic-Load-Balancing-ELB_Application-load-balancer_light-bg.svg"
	case "AWS::ElasticBeanstalk::Application":
		icon = "https://icons.terrastruct.com/aws%2FCompute%2FAWS-Elastic-Beanstalk_Application_light-bg.svg"
	case "AWS::ElasticLoadBalancing::LoadBalancer":
		icon = "Res_Elastic-Load-Balancing_Classic-Load-Balancer_48_Light.svg"
	case "AWS::EMR::Cluster":
		icon = cdn + "Res_Amazon-EMR_Cluster_48_Light.svg"
	case "AWS::Events::Rule":
		icon = cdn + "Res_Amazon-EventBridge_Rule_48_Light.svg"
	case "AWS::Glue::Crawler":
		icon = cdn + "Res_AWS-Glue_Crawler_48_Light.svg"
	case "AWS::Glue::Database":
		icon = cdn + "Res_AWS-Glue_Data-Catalog_48_Light.svg"
	case "AWS::IAM::User":
		icon = "https://icons.terrastruct.com/aws%2F_General%2FUser_light-bg.svg"
	case "AWS::IAM::Group":
		icon = "https://icons.terrastruct.com/aws%2F_General%2FUsers_light-bg.svg"
	case "AWS::IAM::Policy", "AWS::SNS::TopicPolicy", "AWS::IAM::InstanceProfile":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-Identity-and-Access-Management-IAM_Permissions_light-bg.svg"
	case "AWS::IAM::Role":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-Identity-and-Access-Management-IAM_Role_light-bg.svg"
	case "AWS::AccessAnalyzer::Analyzer":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-Identity-and-Access-Management-IAM_Access-Analyzer_light-bg.svg"
	case "AWS::IoTAnalytics::Dataset":
		icon = "https://icons.terrastruct.com/aws%2FInternet%20of%20Things%2FAWS-IoT-Analytics_Data-Set_light-bg.svg"
	case "AWS::IoTAnalytics::Datastore":
		icon = "https://icons.terrastruct.com/aws%2FInternet%20of%20Things%2FAWS-IoT-Analytics_Data-Store_light-bg.svg"
	case "AWS::IoTSiteWise::Asset":
		icon = cdn + "Res_AWS-IoT-SiteWise_Asset_48_Light.svg"
	case "AWS::IoTSiteWise::AssetModel":
		icon = cdn + "Res_AWS-IoT-SiteWise_Asset-Model_48_Light.svg"
	case "AWS::Kinesis::Stream":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAmazon-Kinesis-Data-Streams.svg"
	case "AWS::KMS::Key":
		icon = cdn + "Res_AWS-Identity-Access-Management_Data-Encryption-Key_48_Light.svg"
		
	case "AWS::Lambda::Function":
		icon = "https://icons.terrastruct.com/aws%2FCompute%2FAWS-Lambda_Lambda-Function_light-bg.svg"
	case "AWS::S3::Bucket":
		icon = "https://icons.terrastruct.com/aws%2FStorage%2FAmazon-Simple-Storage-Service-S3_Bucket-with-Objects_light-bg.svg"
	case "AWS::SNS::Topic":
		icon = "https://icons.terrastruct.com/aws%2FApplication%20Integration%2FAmazon-Simple-Notification-Service-SNS_light-bg.svg"
	case "AWS::SQS::Queue":
		icon = "https://icons.terrastruct.com/aws%2FApplication%20Integration%2FAmazon-Simple-Queue-Service-SQS_Queue_light-bg.svg"
	case "AWS::SSM::Parameter":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-Systems-Manager_Parameter-Store_light-bg.svg"
	case "AWS::WAF::Rule":
		icon ="https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-WAF_Filtering-rule_light-bg.svg"
	case "AWS::StepFunctions::StateMachine":
		icon = "https://icons.terrastruct.com/aws%2FApplication%20Integration%2FAWS-Step-Functions_light-bg.svg"
	// Services
	default:
		icon = *ServiceIcons(t)
	}
	return &icon
}

func ServiceIcons(t *string) *string {
	var icon string
	typePath := strings.Split(*t, "::")
	service := typePath[1]
	switch service {
	case "AppConfig":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-App-Config.svg"
	case "Amplify":
		icon = "https://icons.terrastruct.com/aws%2FMobile%2FAWS-Amplify.svg"
	case "ApiGatewayV2":
		icon = "https://icons.terrastruct.com/aws%2FNetworking%20&%20Content%20Delivery%2FAmazon-API-Gateway_light-bg.svg"
	case "ApiGateway":
		icon = "https://icons.terrastruct.com/aws%2FMobile%2FAmazon-API-Gateway_light-bg.svg"
	case "ApplicationAutoScaling":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-Auto-Scaling_light-bg.svg"
	case "AppMesh":
		icon = "https://icons.terrastruct.com/aws%2FNetworking%20&%20Content%20Delivery%2FAWS-App-Mesh_light-bg.svg"
	case "AppStream":
		icon = "https://icons.terrastruct.com/aws%2FEnd%20User%20Computing%2FAmazon-Appstream-2.0_light-bg.svg"
	case "AppSync":
		icon = "https://icons.terrastruct.com/aws%2FApplication%20Integration%2FAmazon-AppSync_light-bg.svg"
	case "Athena":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAmazon-Athena_light-bg.svg"
	case "AutoScalingPlans":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-Auto-Scaling_light-bg.svg"
	case "Backup":
		icon = "https://icons.terrastruct.com/aws%2FStorage%2FAWS-Backup_light-bg.svg"
	case "Batch":
		icon = "https://icons.terrastruct.com/aws%2FStorage%2FAWS-Backup_light-bg.svg"
	case "Budgets":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-Certificate-Manager_light-bg.svg"
	case "CertificateManager":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-Certificate-Manager_light-bg.svg"
	case "Cloud9":
		icon = "https://icons.terrastruct.com/aws%2FDeveloper%20Tools%2FAWS-Cloud9_light-bg.svg"
	case "CloudFormation":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-CloudFormation.svg"
	case "CloudFront":
		icon = "https://icons.terrastruct.com/aws%2FNetworking%20&%20Content%20Delivery%2FAmazon-CloudFront.svg"
	case "ServiceDiscovery":
		icon = "https://icons.terrastruct.com/aws%2FNetworking%20&%20Content%20Delivery%2FAWS-Cloud-Map_light-bg.svg"
	case "CloudTrail":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-CloudTrail_light-bg.svg"
	case "CloudWatch":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAmazon-CloudWatch.svg"
	case "Logs":
		icon = cdn + "Res_Amazon-CloudWatch_Logs_48_Light.svg"
	case "Synthetics":
		icon = cdn + "Res_Amazon-CloudWatch_Synthetics_48_Light.svg"
	case "CodeArtifact":
		icon = cdn + "Arch_AWS-CodeArtifact_48.svg"
	case "CodeBuild":
		icon = "https://icons.terrastruct.com/aws%2FDeveloper%20Tools%2FAWS-CodeBuild.svg"
	case "CodeDeploy":
		icon = "https://icons.terrastruct.com/aws%2FDeveloper%20Tools%2FAWS-CodeDeploy_light-bg.svg"
	case "CodeGuruProfiler":
		icon = "https://icons.terrastruct.com/aws%2FMachine%20Learning%2FAmazon-CodeGuru_light-bg.svg"
	case "CodePipeline":
		icon = "https://icons.terrastruct.com/aws%2FDeveloper%20Tools%2FAWS-CodePipeline.svg"
	case "CodeStar":
		icon = "https://icons.terrastruct.com/aws%2FDeveloper%20Tools%2FAWS-CodeStar.svg"
	case "Cognito":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAmazon-Cognito.svg"
	case "Config":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-Config.svg"
	case "Connect":
		icon = cdn + "Arch_Amazon-Connect_48.svg"
	case "ControlTower":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-Control-Tower.svg"
	case "CE":
		icon = "https://icons.terrastruct.com/aws%2FAWS%20Cost%20Management%2FAWS-Cost-Explorer.svg"
	case "CUR":
		icon = "https://icons.terrastruct.com/aws%2FAWS%20Cost%20Management%2FAWS-Cost-and-Usage-Report.svg"
	case "DataBrew":
		icon = cdn + "Arch_AWS-Glue-DataBrew_48.svg"
	case "DataPipeline":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAWS-Data-Pipeline.svg"
	case "DataSync":
		icon = "https://icons.terrastruct.com/aws%2FMigration%20&%20Transfer%2FAWS-DataSync.svg"
	case "DAX":
		icon = cdn + "Res_Amazon-DynamoDB_Amazon-DynamoDB-Accelerator_48_Light.svg"
	case "Detective":
		icon = "Arch_Amazon-Detective_48.svg"
	case "DeviceFarm":
		icon = "https://icons.terrastruct.com/aws%2FMobile%2FAWS-Device-Farm.svg"
	case "DevOpsGuru":
		icon = cdn + "Arch_Amazon-DevOps-Guru_48.svg"
	case "DirectoryService":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-Directory-Service.svg"
	case "DMS":
		icon = "https://icons.terrastruct.com/aws%2FMigration%20&%20Transfer%2FAWS-Database-Migration-Service.svg"
	case "DocDB":
		icon = cdn + "Arch_Amazon-DocumentDB_48.svg"
	case "DynamoDB":
		icon = cdn + "Arch_Amazon-DynamoDB_48.svg"
	case "EC2":
		icon = "https://icons.terrastruct.com/aws%2FCompute%2FAmazon-EC2_light-bg.svg"
	case "AutoScaling":
		icon = cdn + "Res_Amazon-EC2_Auto-Scaling_48_Light.svg"
	case "ECR":
		icon = "Arch_Amazon-Elastic-Container-Registry_48.svg"
	case "ECS":
		icon = "Arch_Amazon-Elastic-Container-Service_48.svg"
	case "EFS":
		icon = "https://icons.terrastruct.com/aws%2FStorage%2FAmazon-Elastic-File-System_EFS.svg"
	case "EKS":
		icon = cdn + "Arch_Amazon-EKS-Cloud_48.svg"
	case "ElasticBeanstalk":
		icon = "https://icons.terrastruct.com/aws%2FCompute%2FAWS-Elastic-Beanstalk.svg"
	case "ElasticLoadBalancingV2":
		icon = cdn + "Arch_Elastic-Load-Balancing_48.svg"
	case "EMR":
		icon = cdn + "Arch_Amazon-EMR_48.svg"
	case "ElastiCache":
		icon = cdn + "Arch_Amazon-ElastiCache_48.svg"
	case "Events":
		icon = cdn + "Arch_Amazon-EventBridge_48.svg"
	case "EventSchemas":
		icon = cdn + "Res_Amazon-EventBridge_Schema_48_Dark.svg"
	case "Evidently":
		icon = cdn + "Res_Amazon-CloudWatch_Evidently_48_Dark.svg"
	case "FIS":
		icon = cdn + "Arch_AWS-Fault-Injection-Simulator_48.svg"
	case "FMS":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-Firewall-Manager.svg"
	case "Forecast":
		icon = cdn + "Arch_Amazon-Forecast_48.svg"
	case "FraudDetector":
		icon = cdn + "Arch_Amazon-Fraud-Detector_48.svg"
	case "FSx":
		icon = cdn + "Arch_Amazon-FSx_48.svg"
	case "GlobalAccelerator":
		icon = cdn + "Arch_AWS-Global-Accelerator_48.svg"
	case "Glue":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAWS-Glue.svg"
	case "Grafana":
		icon = cdn + "Arch_Amazon-Managed-Grafana_48.svg"
	case "GuardDuty":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAmazon-GuardDuty.svg"
	case "IAM":
		icon = cdn + "Arch_AWS-IAM-Identity-Center_48.svg"
	case "IdentityStore":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-Single-Sign-On_light-bg.svg"
	case "ImageBuilder":
		icon = cdn + "Arch_Amazon-EC2-Image-Builder_48.svg"
	case "SSMIncidents":
		icon = cdn + "Arch_AWS-Systems-Manager-Incident-Manager_48.svg"
	case "Inspector":
		icon = "Arch_Amazon-Inspector_48.svg"
	case "InspectorV2":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAmazon-Inspector.svg"
	case "InternetMonitor":
		icon = cdn + "Res_Internet-alt2_48_Light.svg"
	case "IoT":
		icon = "https://icons.terrastruct.com/aws%2FInternet%20of%20Things%2FAWS-IoT-Core.svg"
	case "IoT1Click":
		icon = "https://icons.terrastruct.com/aws%2FInternet%20of%20Things%2FAWS-IoT-1-Click.svg"
	case "IoTAnalytics":
		icon = "https://icons.terrastruct.com/aws%2FInternet%20of%20Things%2FAWS-IoT-Analytics.svg"
	case "IoTEvents":
		icon = "https://icons.terrastruct.com/aws%2FInternet%20of%20Things%2FAWS-IoT-Events.svg"
	case "IoTFleetWise":
		icon = cdn + "Arch_AWS-IoT-FleetWise_48.svg"
	case "Greengrass":
		icon = "https://icons.terrastruct.com/aws%2FInternet%20of%20Things%2FAWS-IoT-Greengrass.svg"
	case "GreengrassV2":
		icon = "https://icons.terrastruct.com/aws%2FInternet%20of%20Things%2FAWS-IoT-Greengrass.svg"
	case "IoTSiteWise":
		icon = cdn + "Arch_AWS-IoT-SiteWise_48.svg"
	case "Kendra":
		icon = "https://icons.terrastruct.com/aws%2FMachine%20Learning%2FAmazon-Kendra.svg"
	case "Cassandra":
		icon = "https://icons.terrastruct.com/aws%2FDatabase%2FAmazon-Managed-Apache-Cassandra-Service.svg"
	case "Kinesis":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAmazon-Kinesis.svg"
	case "KinesisAnalytics":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAmazon-Kinesis-Data-Analytics.svg"
	case "KinesisAnalyticsV2":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAmazon-Kinesis-Data-Analytics.svg"
	case "KinesisFirehose":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAmazon-Kinesis-Data-Firehose.svg"
	case "KinesisVideo":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAmazon-Kinesis-Video-Streams.svg"
	case "KMS":
		icon = cdn + "Arch_AWS-Key-Management-Service_48.svg"
	case "LakeFormation":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAWS-Lake-Formation.svg"
	case "Lambda":
		icon = "https://icons.terrastruct.com/aws%2FCompute%2FAWS-Lambda.svg"
	case "Lex":
		icon = "https://icons.terrastruct.com/aws%2FMachine%20Learning%2FAmazon-Lex.svg"
	case "LicenseManager":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-License-Manager.svg"
	case "Macie":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAmazon-Macie.svg"
	case "ManagedBlockchain":
		icon = "https://icons.terrastruct.com/aws%2FBlockchain%2FAmazon-Managed-Blockchain.svg"
	case "MediaConnect":
		icon = "https://icons.terrastruct.com/aws%2FMedia%20Services%2FAWS-Elemental-MediaConnect.svg"
	case "MediaLive":
		icon = "https://icons.terrastruct.com/aws%2FMedia%20Services%2FAWS-Elemental-Live.svg"
	case "MediaPackage":
		icon = "https://icons.terrastruct.com/aws%2FMedia%20Services%2FAWS-Elemental-MediaPackage.svg"
	case "MediaTailor":
		icon = "https://icons.terrastruct.com/aws%2FMedia%20Services%2FAWS-Elemental-MediaTailor.svg"
	case "MediaStore":
		icon = "https://icons.terrastruct.com/aws%2FMedia%20Services%2FAWS-Elemental-MediaStore.svg"
	case "AmazonMQ":
		icon = cdn + "Arch_Amazon-MQ_48.svg"
	case "MemoryDB":
		icon = cdn + "Arch_Amazon-MemoryDB-for-Redis_48.svg"
	case "MSK":
		icon = cdn + "Res_Amazon-MSK_Amazon-MSK-Connect_48_Light.svg"
	case "KafkaConnect":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAmazon-Managed-Streaming-for-Kafka_light-bg.svg"
	case "Neptune":
		icon = "https://icons.terrastruct.com/aws%2FDatabase%2FAmazon-Neptune.svg"
	case "NetworkFirewall":
		icon = cdn+"Arch_AWS-Network-Firewall_48.svg"
	case "Elasticsearch":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAmazon-Elasticsearch-Service.svg"
	case "OpenSearchService":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAmazon-Elasticsearch-Service.svg"
	case "OpsWorks":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-OpsWorks.svg"
	case "Organizations":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-Organizations.svg"
	case "Panorama":
		icon = cdn+"Arch_AWS-Panorama_48.svg"
	case "Personalize":
		icon = "https://icons.terrastruct.com/aws%2FMachine%20Learning%2FAmazon-Personalize.svg"
	case "Pinpoint":
		icon = "https://icons.terrastruct.com/aws%2FCustomer%20Engagement%2FAmazon-Pinpoint.svg"
	case "APS":
		icon = cdn+"Arch_Amazon-Managed-Service-for-Prometheus_48.svg"
	case "QLDB":
		icon = "https://icons.terrastruct.com/aws%2FBlockchain%2FAmazon-Quantum-Ledger-Database-QLDB.svg"
	case "QuickSight":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAmazon-Quicksight.svg"
	case "RAM":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-Resource-Access-Manager.svg"
	case "RDS":
		icon = "https://icons.terrastruct.com/aws%2FDatabase%2FAmazon-RDS.svg"
	case "Redshift":
		icon = "https://icons.terrastruct.com/aws%2FAnalytics%2FAmazon-Redshift.svg"
	case "Rekognition":
		icon = "https://icons.terrastruct.com/aws%2FMachine%20Learning%2FAmazon-Rekognition.svg"
	case "Route53":
		icon = "https://icons.terrastruct.com/aws%2FNetworking%20&%20Content%20Delivery%2FAmazon-Route-53.svg"
	case "S3":
		icon = "https://icons.terrastruct.com/aws%2FStorage%2FAmazon-Simple-Storage-Service-S3.svg"
	case "SecretsManager":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-Secrets-Manager.svg"
	case "ServiceCatalog":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-Service-Catalog.svg"
	case "SecurityHub":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-Security-Hub.svg"
	case "SES":
		icon = "https://icons.terrastruct.com/aws%2FCustomer%20Engagement%2FAmazon-Simple-Email-Service-SES_light-bg.svg"
	case "SNS":
		icon = "https://icons.terrastruct.com/aws%2FApplication%20Integration%2FAmazon-Simple-Notification-Service-SNS.svg"
	case "SQS":
		icon = "https://icons.terrastruct.com/aws%2FApplication%20Integration%2FAmazon-Simple-Queue-Service-SQS.svg"
	case "SSO":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-Single-Sign-On.svg"
	case "StepFunctions":
		icon = "https://icons.terrastruct.com/aws%2FApplication%20Integration%2FAWS-Step-Functions.svg"
	case "SSM":
		icon = "https://icons.terrastruct.com/aws%2FManagement%20&%20Governance%2FAWS-Systems-Manager.svg"
	case "Timestream":
		icon = "https://icons.terrastruct.com/aws%2FDatabase%2FAmazon-Timestream.svg"
	case "Transfer":
		icon = "https://icons.terrastruct.com/aws%2FMigration%20&%20Transfer%2FAWS-Transfer-Family.svg"
	case "WAF":
		icon = "https://icons.terrastruct.com/aws%2FSecurity%2C%20Identity%2C%20&%20Compliance%2FAWS-WAF.svg"
	case "WorkSpaces":
		icon = "https://icons.terrastruct.com/aws%2FEnd%20User%20Computing%2FAmazon-Workspaces.svg"
	case "XRay":
		icon = "https://icons.terrastruct.com/aws%2FDeveloper%20Tools%2FAWS-X-Ray.svg"
	default:
		icon = "https://icons.terrastruct.com/aws%2F_Group%20Icons%2FAWS-Cloud_light-bg.svg"
	}
	return &icon
}
