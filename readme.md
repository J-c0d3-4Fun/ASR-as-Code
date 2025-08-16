# ASR-as-Code

**Attack Surface Reduction, codified.**

ASR-as-Code is a Go-powered toolset for detecting and reducing cloud attack surface. Inspired by "infrastructure as code," it treats security guardrails as code — versioned, automated, and reproducible.

## 🔑 Current Focus

### S3 Security Checks
- ✅ Detects if encryption is disabled
- ✅ Detects if public access blocks are missing  
- ✅ Detects weak bucket policies

## 🎯 Vision

- **Codify security** — express attack surface reduction policies as reusable code
- **Automate detection** — continuously scan AWS accounts for misconfigurations
- **Enable remediation** — provide secure-by-default recommendations or code snippets

## 🚀 Why ASR-as-Code?

- **Repeatable:** Same security logic everywhere, no drift
- **Auditable:** Version-controlled checks
- **Scalable:** Start small with S3, expand to IAM, EC2, VPC, and beyond
