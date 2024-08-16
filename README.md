# Journal 
Journal whatever you like. Keep them easily reachable with search.

## Dependencies
This is a program written in go. As such you need to have Golang installed on
your system to build this.

### Steps to build
1. Clone this repository and `cd` into it
2. Run command `go build`
3. [Optional] Run `sudo install -m 755 ./journal /usr/local/bin`

The last step would install the executable into the directory `usr/local/bin` 
so that you don't need to provide the full path to run the program everytime.

## Init
ENVIRONMENT variable 'JOURNAL_DATA_FILE' needs to be setup and should 
contain the file path where journal would write to.

Here's an example command to setup the env value and persist it for future.

`echo -e '#Journal ENV\nexport JOURNAL_DATA_FILE=~/.local/journal\n' >> ~/.bashrc`

this assumes you have `bash` as your terminal Shell.

## Posting to the journal
Posting to the journal is simple. Just the type in following command.

`journal -p "<YOUR_TEXT_HERE>"`

Journal won't show you any sign of success like a message or such. It
is not made that way. Rest assured though, your post has been saved.

## Listing entries
To see the entry you just made, type in the following command.

`journal`

And you'd see the post you've just made. Journal adds the date and
time to the post. The format is `YYYY-MM-DD HH:MM <YOUR_TEXT_HERE>`.

`journal` command would list you entries you've made today. All the
previous entries are safe but, they are not shown directly.

You need to search for it

## Finding entries
Journal uses extended Regex pattern to search posts. The following
command is the basic structure for the search post feature

`journal <REGEX_PATTERN_HERE> [...<MORE_REGEX_PATTERNS>]`

### Examples

1. list all posts mentioning `bash` or `scripting` type in 
    
    `journal bash scripting`
2. list all posts mentioning exactly `bash scripting` type in 

    `journal 'bash scripting'`
3. list all posts made on 24th of April
    
    `journal '04-24'`

## Moving the datafile
Journal uses plain text file to store journal posts. You can view,
edit, share version control or do whatever you like.

Journal has no restrictions on that. 
