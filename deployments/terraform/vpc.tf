resource "aws_vpc" "bcs_portal" {
  cidr_block = "10.0.0.0/24"

  tags = {
    Service = "bcs-portal"
  }
}

resource "aws_subnet" "bcs_portal_db" {
  vpc_id = aws_vpc.bcs_portal.id
  cidr_block = "10.0.0.0/24"

  tags = {
    Service = "bcs-portal"
  }
}
