# Email Domain Verifier

A lightweight Go command-line tool that checks email domains for essential email security and deliverability configurations including MX, SPF, and DMARC records.

## Features

- **MX Record Verification**: Checks if the domain has mail exchange records configured
- **SPF Record Detection**: Identifies Sender Policy Framework records for email authentication
- **DMARC Record Analysis**: Locates Domain-based Message Authentication, Reporting, and Conformance policies
- **Batch Processing**: Process multiple domains from standard input
- **CSV Output**: Results formatted as comma-separated values for easy analysis

## Demo Video

Watch a screen recording demonstration of the tool: 
> https://github.com/user-attachments/assets/3c57eef5-403e-4735-96b9-c7496914f302


## Installation

### Prerequisites
- Go 1.16 or higher

### Build from Source
```bash
git clone <repository-url>
cd email-verifier
go build -o email-verifier main.go
```

### Or compile directly
```bash
go run main.go
```

## Usage

### Interactive Mode
Run the program and enter domains one by one:
```bash
./email-verifier
google.com
facebook.com
example.com
```

### Batch Processing from File
```bash
cat domains.txt | ./email-verifier
```

### Save Results to File
```bash
echo "google.com" | ./email-verifier > results.csv
```

## Output Format

The tool outputs results in CSV format with the following columns:

```
domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord
```

### Example Output
```csv
domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord
google.com,true,true,v=spf1 include:_spf.google.com ~all,true,v=DMARC1; p=reject; rua=mailto:mailauth-reports@google.com
example.com,true,false,,false,
invalid-domain.xyz,false,false,,false,
```

## Understanding the Results

| Column | Description |
|--------|-------------|
| `domain` | The domain being checked |
| `hasMX` | Boolean indicating if MX records exist |
| `hasSPF` | Boolean indicating if SPF record is configured |
| `spfRecord` | The actual SPF record content (if found) |
| `hasDMARC` | Boolean indicating if DMARC policy exists |
| `dmarcRecord` | The actual DMARC record content (if found) |

## What Each Record Means

### MX Records (Mail Exchange)
- **Purpose**: Specify which mail servers accept email for the domain
- **Impact**: Without MX records, the domain cannot receive email
- **Example**: `10 mail.example.com`

### SPF Records (Sender Policy Framework)
- **Purpose**: Specify which IP addresses/servers are authorized to send email for the domain
- **Impact**: Helps prevent email spoofing and improves deliverability
- **Example**: `v=spf1 include:_spf.google.com ~all`

### DMARC Records (Domain-based Message Authentication)
- **Purpose**: Defines policy for handling emails that fail SPF/DKIM authentication
- **Impact**: Provides additional email security and reporting capabilities
- **Example**: `v=DMARC1; p=reject; rua=mailto:reports@example.com`

## Use Cases

- **Email Security Audits**: Assess domain email security posture
- **Deliverability Analysis**: Check if domains are properly configured for email delivery
- **Bulk Domain Verification**: Process large lists of domains efficiently
- **Compliance Checking**: Verify email authentication standards implementation

## Error Handling

The tool gracefully handles various error conditions:
- Invalid domain names
- DNS resolution failures
- Network connectivity issues
- Malformed DNS records

Errors are logged to stderr while maintaining CSV output format for successful lookups.

## Limitations

- Only checks for basic email DNS records (MX, SPF, DMARC)
- Does not validate DKIM records
- Does not test actual email deliverability
- Requires internet connectivity for DNS lookups

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For issues, questions, or contributions, please open an issue in the repository.

## 📧 Contact

**GitHub** : [abhinav-phi](https://github.com/abhinav-phi)

**Project Link** : [click here](https://github.com/abhinav-phi/email-verifier-tool)

---

⭐ If you found this project helpful, please give it a star!

