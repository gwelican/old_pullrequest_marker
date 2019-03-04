
# Pull request marker

## What's this?

This simple go project will scan your pull requests and mark pull request with 'old' flag, which are more than x days(configureable) open.

The tool will not create the old flag, you have to create it manually.

# Usage

## Parameters:
### Parameters can be provided either as flag or environment variable.
#### Environment variable: https://github.com/namsral/flag#parsing-environment-variables

#### Command line arguments:
* --token this is the github auth token, you can create your own: https://github.com/settings/tokens
* --owner the repo owner, for example your user
* --repo the repository name that you want to check
* --days(default 5) filter the pull request, only checks pullrequest that has x days passed since it was created.
* --serverurl(default: https://api.github.com/) github url, for enterprise, you can use: http://enterprise.github.com/api/v3/
For example:

go run main.go --token=[token]--owner=gwelican --repo=old_pullrequest_marker --days=5

