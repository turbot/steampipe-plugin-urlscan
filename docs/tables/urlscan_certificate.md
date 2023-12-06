---
title: "Steampipe Table: urlscan_certificate - Query urlscan.io Certificates using SQL"
description: "Allows users to query Certificates from urlscan.io, specifically providing details about the SSL certificates associated with a particular scan, including issuer, subject, validity period, and more."
---

# Table: urlscan_certificate - Query urlscan.io Certificates using SQL

urlscan.io is a service that allows you to scan and analyze web pages to identify security issues and malicious behavior. It provides detailed information about the components of the webpage, including the SSL certificates. These certificates are used to secure communication between the website and the user, and contain details such as the issuer, subject, validity period, and more.

## Table Usage Guide

The `urlscan_certificate` table provides insights into SSL certificates associated with a scan in urlscan.io. As a security analyst, explore certificate-specific details through this table, including the issuer, subject, validity period, and associated metadata. Utilize it to uncover information about certificates, such as their validity, the organizations that issued them, and the organizations they were issued to.

**Important Notes**
- You must specify the `scan` in the `where` clause to query this table.

## Examples

### List certificates found in the scan
Explore the validity of certificates identified in a specific scan. This query helps to track the expiration date of each certificate, allowing you to manage and update them before they expire.

```sql
select
  subject_name,
  issuer,
  valid_to,
  (valid_to::date - current_date) as days_until_expiration
from
  urlscan_certificate
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  valid_to
```