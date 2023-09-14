resource "aws_route53domains_registered_domain" "bcs_portal" {
  domain_name = var.bcs_portal_domain
}

resource "aws_amplify_app" "bcs_portal" {
  name       = "bcs-portal-web"
  repository = var.amplify_repository

  environment_variables = {
    AMPLIFY_DIFF_DEPLOY       = "false"
    AMPLIFY_MONOREPO_APP_ROOT = "web"
  }

  custom_rule {
    source = "/<*>"
    status = "404-200"
    target = "/index.html"
  }

  custom_rule {
    source = "https://www.${aws_route53domains_registered_domain.bcs_portal.domain_name}"
    status = "302"
    target = "https://${aws_route53domains_registered_domain.bcs_portal.domain_name}"
  }

  build_spec = <<-EOT
    version: 1
    applications:
      - appRoot: web
        frontend:
          phases:
            preBuild:
              commands:
                - yarn install
            build:
              commands:
                - yarn generate
          artifacts:
            baseDirectory: '.output/public'
            files:
              - '**/*'
          cache:
            paths:
              - node_modules/**/*
EOT
}

resource "aws_amplify_branch" "main" {
  app_id      = aws_amplify_app.bcs_portal.id
  branch_name = "main"
}

resource "aws_amplify_domain_association" "bcs_portal" {
  app_id      = aws_amplify_app.bcs_portal.id
  domain_name = aws_route53domains_registered_domain.bcs_portal.domain_name

  sub_domain {
    branch_name = aws_amplify_branch.main.branch_name
    prefix      = ""
  }

  sub_domain {
    branch_name = aws_amplify_branch.main.branch_name
    prefix      = "www"
  }
}