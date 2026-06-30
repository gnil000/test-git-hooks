# About

Repository contains examples git hooks for golang projects.

### What doing pre-commit hook?  
1. Checks staged files
2. Start quality checks for `*.go` files

If do not contains `*.go` files, skip quality checks.  
If quality checks failed, prevent commit

### What doing pre-push hook?  
1. Forbids pushing to main
2. Runs Go tests

If tests failed, prevent pushing

## How setup hooks?

```shell
# Unix using bash
bash ./git-hooks/setup-hooks.sh

# Windows using Powershell
PS> .\scripts\win\setup-hooks.ps1
```

