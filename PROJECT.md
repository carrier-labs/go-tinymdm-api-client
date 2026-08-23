---
id: go-tinymdm-api-client
name: TinyMDM client
kind: library
state: active
owner: sam
commercial: none

plan:
  doc: null
  ref: null
  maturity: none
  session: null

# target: the date this milestone should be LIVE by. Set at planning time.
# Seeded null deliberately — an invented date becomes a commitment nobody agreed to.
milestones: []   # no milestones in the plan yet — add them at the next planning session

next: null
blocked_by: null
links:
  trello: null
  live: null
tags: [shared, go]
---

Shared Go client for TinyMDM.

**Load-bearing:** a dependency of `1472_switchkiosks`, which is deployed in stores and earning.
It was originally written for AMP's cost-replacement case; AMP was archived 2026-08-21, but this
library is still in live use and is not orphaned.
