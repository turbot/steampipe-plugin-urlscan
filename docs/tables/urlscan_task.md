---
title: "Steampipe Table: urlscan_task - Query Urlscan Tasks using SQL"
description: "Allows users to query Urlscan Tasks specifically the task details, providing insights into scan tasks and potential security threats."
---

# Table: urlscan_task - Query Urlscan Tasks using SQL

Urlscan Tasks in is a service that allows you to execute and manage scan tasks across your applications and infrastructure. It provides a centralized way to set up and manage tasks for various resources, including websites, web applications, and more. Urlscan Tasks helps you stay informed about the health and security of your resources and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `urlscan_task` table provides insights into tasks within Urlscan. As a security analyst, explore task-specific details through this table, including task status, results, and associated metadata. Utilize it to uncover information about tasks, such as those related to specific URLs, the status of each task, and the verification of scan results.

**Important Notes**
- You must specify the `scan` in the `where` clause to query this table.

## Examples

### Get task information
Explore detailed information about a specific task, such as its status, duration, and associated metadata. This is useful for tracking the progress and outcome of individual tasks, enabling efficient monitoring and management.

```sql
select
  *
from
  urlscan_task
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
```