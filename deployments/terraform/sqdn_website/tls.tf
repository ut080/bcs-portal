# Domain Settings: the-hawk.us
## DNS Zone
#resource "aws_route53_zone" "ut080_utwg_cap_gov" {
#  name = "ut080.utwg.cap.gov"
#}

## TLS Certificate
### CloudFront Certificate
resource "aws_acm_certificate" "ut080_utwg_cap_gov" {
  domain_name = "ut080.utwg.cap.gov"
  validation_method = "EMAIL"
}

resource "aws_acm_certificate_validation" "ut080_utwg_cap_gov_cf" {
  certificate_arn         = aws_acm_certificate.ut080_utwg_cap_gov.arn
}
