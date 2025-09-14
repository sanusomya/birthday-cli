# birthday CLI

A simple, fast CLI to manage a list of birthdays and send reminders for today’s or upcoming birthdays. Includes subcommands to add, get, edit, delete, remind, and generate shell completions.

### Features

- Add, list, edit, and delete birthday entries with names, dates, and phone numbers.
- Send reminders for today or within a configurable upcoming day range.
- Autocompletion scripts for popular shells.

### Prerequisites

Before running the CLI, make sure the following requirements are set up:

1. **Backend/API Setup**
    - You will need to establish a backend to handle interactions with the CLI.
    - The recommended approach is to reuse the existing backend here:
[birthday-lambda](https://github.com/sanusomya/birthday-lambda).
    - If you have your own backend, make small modifications in the CLI code to integrate with it.
2. **Environment Variables**
The CLI requires the following environment variables to be configured:
    - `backend_url` – The base URL of your backend.
    - `token` – Authorization token (needed if you are using the authorizer lambda).
More details: [birthday-lambda README](https://github.com/sanusomya/birthday-lambda/blob/main/README.md).
    - `mail` – The email address where you want to receive the reminder.
    - `pass` – The credential for the mail service:
        - For Gmail, this must be a Google App Password.
        - For other providers, use the corresponding mail service authentication token.

Here’s the Install section rewritten in Markdown format, ready to paste into your README:

### Install

There are two ways to install the CLI:

1. Recommended: Prebuilt binaries
   - Download the latest release for your OS from the releases page: https://github.com/sanusomya/birthday-cli/releases
   - Unpack the archive and place the binary in a directory on PATH (for example, on Linux/macOS: /usr/local/bin; on Windows: add the folder to PATH)
   - Verify the installation:
     - Linux/macOS: `birthday-cli --help`
     - Windows (PowerShell): `.\birthday-cli.exe --help`

2. Build from source (Go)
   - Prerequisites:
     - Go 1.21+ installed and configured (GOBIN/GOPATH as desired)
     - Git installed
   - Steps:
     - Clone the repository:
       - `git clone https://github.com/sanusomya/birthday-cli.git`
       - `cd birthday-cli`
     - Build:
       - Linux/macOS: `go build -o birthday-cli ./cmd/birthday-cli`
       - Windows (PowerShell): `go build -o birthday-cli.exe .\cmd\birthday-cli`
     - Optionally install to GOBIN:
       - `go install ./cmd/birthday-cli`
     - Verify:
       - Linux/macOS: `./birthday-cli --help` (or `birthday-cli --help` if installed)
       - Windows: `.\birthday-cli.exe --help`

Optional tips
- Make executable (Linux/macOS): `chmod +x birthday-cli`
- If downloading via terminal:
  - Linux/macOS example:
    - `curl -L -o birthday-cli.tar.gz <release-asset-url>`
    - `tar -xzf birthday-cli.tar.gz`
    - `sudo mv birthday-cli /usr/local/bin/`
- On Windows, if SmartScreen warns on a downloaded binary, choose “More info” → “Run anyway” if the binary is trusted and verified.
- To update, repeat the same method: download the latest release or rebuild from source on the new tag/commit.

### Usage

Use the command with flags and args to get started.

```bash
Usage:
  birthday [flags]
  birthday [command]

Aliases:
  birthday, birth, bday

Examples:
  birthday <args> -<flag> <data>

Available Commands:
  add         use this command to add to the list of all birthdays
  completion  Generate the autocompletion script for the specified shell
  delete      use this command to delete the entry from birthdays
  edit        use this command to delete the entry from birthdays
  get         use this command to get the list of all birthdays
  help        Help about any command
  remind      use this command to dend a remainder for all birthdays today

Flags:
  -h, --help      help for birthday
  -v, --version   version for birthday

Use "birthday [command] --help" for more information about a command.
```


### Global flags

- -h, --help: Show contextual help.
- -v, --version: Print CLI version.


### Commands

#### add

Add a birthday entry. Supports name, day, month, and phone; date can also be supplied.

```bash
Usage:
  birthday add [flags]

Examples:
  birthday add -name <data> -day <data> -month <data> -phone <data>

Flags:
  -d, --date string     date
  -h, --help            help for add
  -m, --month string    month
  -n, --name string     name of person
  -p, --phone string    phone number
  -v, --version         version for add
```

Examples:

- birthday add -n "Sophie" -d 14 -m 9 -p +15551234567
- birthday add --name "Ken" --date 1990-04-22 --phone 0712345678


#### get

List birthdays; filter by month or show only today’s.

```bash
Usage:
  birthday get [flags]

Examples:
  birthday get

Flags:
  -h, --help      help for get
  -m, --month     month
  -t, --today     today
  -v, --version   version for get
```

Examples:

- birthday get --today
- birthday get --month


#### delete

Delete a birthday entry by name and optional date/month/phone.

```bash
Usage:
  birthday delete [flags]

Aliases:
  delete, del

Examples:
  birthday delete -name <data> -day <data> -month <data> -mobile <data>

Flags:
  -d, --date string     date
  -h, --help            help for delete
  -m, --month string    month
  -n, --name string     name of person
  -p, --phone string    phone number
  -v, --version         version for delete
```

Examples:

- birthday del -n "Sophie" -m 9 -d 14
- birthday delete --name "Ken" --phone 0712345678


#### edit

Edit an existing entry; supports subcommands to change mobile or name.

```bash
Usage:
  birthday edit [flags]
  birthday edit [command]

Examples:
  birthday edit -name <data> -day <data> -month <data> -phone <data>

Available Commands:
  mobile      use this command to edit name of the entry from birthdays
  name        use this command to edit name of the entry from birthdays

Flags:
  -d, --date string     date
  -h, --help            help for edit
  -m, --month string    month
  -n, --name string     name of person
  -p, --phone string    phone number
  -v, --version         version for edit

Use "birthday edit [command] --help" for more information about a command.
```

Subcommands:

- edit mobile
Change the phone number for a matching entry.

```bash
Usage:
  birthday edit mobile [flags]

Aliases:
  mobile, mob

Examples:
  birthday edit mobile -name <data> -mobile <data> <new number>

Flags:
  -h, --help       help for mobile
  -v, --version    version for mobile

Global Flags:
  -d, --date string     date
  -m, --month string    month
  -n, --name string     name of person
  -p, --phone string    phone number
```

- edit name
Change the name for a matching entry.

```bash
Usage:
  birthday edit name [flags]

Examples:
  birthday edit name -name <data> -mobile <data> <new name>

Flags:
  -h, --help       help for name
  -v, --version    version for name

Global Flags:
  -d, --date string     date
  -m, --month string    month
  -n, --name string     name of person
  -p, --phone string    phone number
```


#### remind

Use this command to send a reminder for all birthdays today, or use a range to include upcoming days.

```bash
Usage:
  birthday remind [flags]

Examples:
  birthday wish

Flags:
  -h, --help        help for remind
      --range int   use this flag to set a range of days to send reminders for,
                    eg --range 6 to send reminders for birthdays within next 7 days,
                    note that remainder will be for each birthday and default range is 0 to 6 default is 0
  -v, --version     version for remind
```

Examples:

- birthday remind               \# today’s birthdays
- birthday remind --range 0     \# explicitly only today
- birthday remind --range 6     \# today + next 6 days (7-day window)

Notes:

- The range value N covers today through N days ahead; each matching entry triggers a reminder.
- If both today and range are supported elsewhere in your workflow, prefer using remind exclusively for notifications.


#### completion

Generate an autocompletion script for a supported shell.

Examples:

- bash: `birthday completion bash > /etc/bash_completion.d/birthday`
- zsh: `birthday completion zsh > "${fpath}/_birthday"` then `exec zsh`[^1]


### Conventions and tips

- Prefer explicit flags over positional args for clarity in scripts.
- Use aliases for speed: `birth`, `bday`, and `del` are available.
- Use `--help` on any subcommand for detailed usage.


### Example workflows

- Add and verify:
    - birthday add -n "Aria" -d 1 -m 10 -p 9990011223
    - birthday get --today
- Correct a typo:
    - birthday edit name -n "Aria" -p 9990011223 "Arya"
    - birthday edit mobile -n "Arya" -p 9990011223 9990011224
- Remind for the week:
    - birthday remind --range 6
- Cleanup:
    - birthday del -n "Arya" -m 10 -d 1


### Contributing

If accepting contributions, add CONTRIBUTING.md, issue templates, and a code of conduct. Keep the README concise and link out for details.

