# Email Checker

This Go program, email-checker, checks the DNS records for a given domain extracted from an email address and displays information about its MX, SPF, and DMARC records.

MX (Mail Exchange) records: MX records specify the mail servers responsible for receiving email for a particular domain.

SPF (Sender Policy Framework): SPF is an email authentication protocol used to prevent email spoofing by specifying which mail servers are authorized to send emails on behalf of a domain.

DMARC (Domain-based Message Authentication, Reporting, and Conformance): DMARC helps prevent email spoofing and phishing attacks by providing a framework for email senders to define how their emails should be handled if they fail authentication checks.

## Usage

1. Clone the repository:

```bash
git clone https://github.com/ashwin-elangovan/email_checker.git
```
2. Navigate to the project directory:
```bash
cd email_checker
```
3. Run the program:
```bash
run main.go
```
4. Enter an email address when prompted.

## Features
1. Extracts the domain from the provided email address.
2. Performs DNS lookups for MX, SPF, and DMARC records of the domain.
3. Displays information about the presence of MX, SPF, and DMARC records along with their respective records if available.

## Examples
<img width="816" alt="Screenshot 2024-02-16 at 4 43 35 PM" src="https://github.com/ashwin-elangovan/email-checker/assets/35291686/4537c5ff-7e91-412a-aca5-a4c7494e2d49">

## Dependencies
```github.com/common-nighthawk/go-figure``` - For ASCII art display.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## References
1. https://www.cloudflare.com/en-gb/learning/email-security/dmarc-dkim-spf/
