package constants

var CopySnapshotsModes = []string{
	"copy_unencrypted",
	"copy_unencrypted_encrypt",
	"copy_encrypted",
	"copy_both",
	"copy_both_encrypt",
}

var CopyDbSnapshotsOperators = []string{
	"copy_last",
	"copy_older_than",
}

var CopySnapshotsOperators = []string{
	"copy_last",
	"created_greater_than",
}

var AwsRegions = []string{
	"eu-north-1",
  "ap-south-1",
  "eu-west-3",
  "eu-west-2",
  "eu-west-1",
  "ap-northeast-2",
  "ap-northeast-1",
  "sa-east-1",
  "ca-central-1",
  "ap-southeast-1",
  "ap-southeast-2",
  "eu-central-1",
  "us-east-1",
  "us-east-2",
  "us-west-1",
  "us-west-2",
}
