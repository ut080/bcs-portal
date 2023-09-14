resource "aws_kms_key" "bcs_portal" {
  description = "BCS Portal secrets key"
  policy = <<-EOT
  {
      "Version": "2012-10-17",
      "Id": "bcs-portal-secrets-1",
      "Statement": [
          {
              "Sid": "Enable IAM User Permissions",
              "Effect": "Allow",
              "Principal": {
                  "AWS": "${var.aws_principal_arn}"
              },
              "Action": "kms:*",
              "Resource": "*"
          },
          {
              "Sid": "Allow Terraform Cloud access to this key",
              "Effect": "Allow",
              "Principal": {
                  "AWS": "${var.terraform_cloud_role_arn}"
              },
              "Action": [
                "kms:DescribeKey",
                "kms:Decrypt",
                "kms:CreateDataKey",
                "kms:CreateGrant"
              ],
              "Resource": "*"
          }
      ]
  }
EOT


}

resource "aws_kms_alias" "bcs_portal" {
  name          = "bcs-portal"
  target_key_id = aws_kms_key.bcs_portal.key_id
}

resource "aws_db_instance" "bcs_portal" {
  allocated_storage             = 20
  ca_cert_identifier            = "rds-ca-ecc384-g1"
  db_name                       = "bcs_portal"
  engine                        = "postgres"
  engine_version                = "15.4"
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
