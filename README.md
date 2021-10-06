# githuborg

A tool for managing org wide changes within github

It expects you to have a `GITHUB_TOKEN` set with the correct permissions.
This tool exists because I have automation that opens up pull requests and deals with branches at scale across orgs with large amounts of repos. If I want to approve all of these changes or delete all of these branches as there was a bug doing it by hand is no fun ...

## Delete a Branch everywhere

```bash
githuborg branch delete -n myBranchName -o MyOrgName -t ThisTopicOnAllrepos
```

## Approve Pull Requests everywhere

```bash
githuborg branch approve -n myBranchName -o MyOrgName -t ThisTopicOnAllrepos
```
