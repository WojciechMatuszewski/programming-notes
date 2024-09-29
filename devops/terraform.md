# Terraform

## Basics

- Terraform is a language for defining infrastructure.

  - **Given its "providers" models, you can hook up anything that has an API to Terraform** – you "just" need to write a provider for it.

- **Terraform, just like some of AWS-related IaC frameworks, use API calls to deploy the infrastructure**.

  - In the case of AWS, this means that rolling back from a failed deployment might be problematic.

    - Instead, we should consider _rolling forward_ where you fix the mistake that caused failed deployment.

## Testing

- As far as I know, there are two ways to test Terraform "output".

  - Via the [`.tftest.hcl` file](https://developer.hashicorp.com/terraform/language/tests)

  - Inside the provider code (most likely written in Go).

- **You can even [mock the resources](https://developer.hashicorp.com/terraform/tutorials/configuration-language/test#mock-tests)**.

  - I have to say, I'm amazed that they have this in the language.
