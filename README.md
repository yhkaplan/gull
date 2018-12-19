## What
Gull is a utility for reviewing your GitHub or GitHub Enterprise activity. You can use it to export activity to markdown or explore interactively via the dashboard.

## Install
```
git clone {this_project}
cd {folder_name}
go install
```

Require Go

## Use

For output, `gull activity`
For dashboard, `gull dashboard`

### Setup
Set an environmental variable with a GitHub token with `Repo` access permissions
```sh
export GITHUB_TOKEN={token}
```
### Subcommands
```sh
activity   Shows GitHub activities
dashboard  Visualizes GitHub activities
help, h    Shows a list of commands or help for one command
```

### Flags
```sh
--from date, -f date          From date (default: "2018-09-21")
--to date, -t date            To date (default: "2018-09-28")
--user username, -u username  Get activities of specified username
--eventType                   Show event type along with output (default: don't show)
--comment -c                  Show comment events as well (default: don't show)
```

### For GitHub Enterprise
Set below to url
```
export GITHUB_API={url}
```

## Contribute
Feel free to submit a pull request! Please use gofmt.

