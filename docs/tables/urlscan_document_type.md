---
title: "Steampipe Table: urlscan_document_type - Query Urlscan Document Types using SQL"
description: "Allows users to query Document Types in Urlscan, specifically the information about the document type of the web pages scanned by Urlscan."
---

# Table: urlscan_document_type - Query Urlscan Document Types using SQL

Urlscan is a service that scans and analyzes websites to identify malicious content and provide detailed information about the behavior and composition of the scanned web page. It provides insights into the structure, behavior, and content of websites, which can be used to detect and mitigate security threats. Document Types in Urlscan represent the various types of documents that a web page can contain, such as HTML, XML, JSON, etc.

## Table Usage Guide

The `urlscan_document_type` table provides insights into the different types of documents that are part of the web pages scanned by Urlscan. As a cybersecurity analyst, explore document-specific details through this table, including the type, associated web page, and other pertinent information. Utilize it to uncover insights about the various types of documents present on a web page, aiding in the detection and mitigation of potential security threats.

**Important Notes**
- You must specify the `scan` in the `where` clause to query this table.

## Examples

### List document types
Explore which document types are most prevalent in a specific scan, allowing you to identify potential threats or anomalies based on frequency. This can aid in strengthening your security measures by focusing on the most common document types.

```sql+postgres
select
  *
from
  urlscan_document_type
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  count desc;
```

```sql+sqlite
select
  *
from
  urlscan_document_type
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  count desc;
```