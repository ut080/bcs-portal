resource "aws_kms_key" "bcs_portal" {
  description = "BCS Portal secrets key"
}

resource "aws_db_instance" "bcs_portal" {
  allocated_storage             = 20
  db_name                       = "bcs_portal"
  engine                        = "postgres"
  engine_version                = "15.4-R1"
  identifier                    = "bcs-portal-db"
  instance_class                = "db.t3.micro"
  manage_master_user_password   = true
  master_user_secret_kms_key_id = aws_kms_key.bcs_portal.id
  username                      = "postgres"

  tags = {
    Service = "bcs-portal"
  }
}

resource "aws_db_subnet_group" "bcs_portal" {
  name       = "bcs-portal-dbsg"
  subnet_ids = [aws_subnet.bcs_portal_db_aza.id, aws_subnet.bcs_portal_db_azb.id]

  tags = {
    Service = "bcs-portal"
  }
}