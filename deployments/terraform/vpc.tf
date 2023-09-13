/*
resource "aws_vpc" "bcs_portal" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Service = "bcs-portal"
  }
}

resource "aws_subnet" "bcs_portal_db_aza" {
  vpc_id = aws_vpc.bcs_portal.id
  cidr_block = "10.0.1.0/24"
  availability_zone = "${var.region}a"

  tags = {
    Service = "bcs-portal"
  }
}

resource "aws_subnet" "bcs_portal_db_azb" {
  vpc_id = aws_vpc.bcs_portal.id
  cidr_block = "10.0.2.0/24"
  availability_zone = "${var.region}b"

  tags = {
    Service = "bcs-portal"
  }
}
*/