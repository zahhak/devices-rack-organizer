data "aws_region" "eu-west-1" {}

resource "aws_s3_bucket" "ns-devices-rack-organizer" {
  bucket = "ns-devices-rack-organizer"
  acl    = "private"

  tags {
    Name        = "ns-devices-rack-organizer"
    Environment = "Production"
  }
}

resource "aws_dynamodb_table" "ns-devices-rack-organizer" {
  name           = "RackOrganizer"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "DeviceID"

  attribute {
    name = "DeviceID"
    type = "S"
  }

#  attribute {
#    name = "UserID"
#    type = "S"
#  }
#
#  attribute {
#    name = "Date"
#    type = "S"
#  }

  tags {
    Name        = "ns-devices-rack-organizer"
    Environment = "Production"
  }
}
