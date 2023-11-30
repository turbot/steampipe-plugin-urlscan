---
title: "Steampipe Table: urlscan_global - Query OCI Urlscan Global Data using SQL"
description: "Allows users to query global data from OCI Urlscan, providing insights into the overall scan results and metadata."
---

# Table: urlscan_global - Query OCI Urlscan Global Data using SQL

The OCI Urlscan Global Data is a resource that provides an overview of the scan results from Urlscan. It includes data such as the overall status of the scan, task details, page metadata, and other associated information. This data is valuable for understanding the general outcome and details of the scans performed by Urlscan.

## Table Usage Guide

The `urlscan_global` table provides insights into the global scan data within OCI Urlscan. As a security analyst, explore scan-specific details through this table, including the overall status, task details, and page metadata. Utilize it to uncover information about the scans, such as the overall results, the details of the tasks performed, and the metadata of the scanned pages.

## Examples

### List global variables
Analyze the settings to understand the global variables associated with a particular scan. This can be useful for identifying potential areas of concern or improvement within a specific scan process.

```sql
select
  property,
  type
from
  urlscan_global
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  property
```