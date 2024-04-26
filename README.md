# Journal 
Journal whatever you like. Keep them easily reachable with search.

Journal is a bash script! It uses simple command-line tools to give a bare
bones journaling capabilities.

## Dependencies
This script only requires basic command line tools like `grep` and `echo`
and a shell like `sh`, `bash`, or such. If yours is a Unix-like system, you
are good to go!

On Windows you'd need to install WSL.

## Init
Just download the file `journal` and copy it to your path. This would be 
`/usr/local/bin` in general!

ENVIRONMENT variable 'JOURNAL_DATA_FILE' needs to be setup and should 
contain the file path where journal would write to.

Type in `journal` and if the program exits without an error you're all
set.

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

`journal -f <REGEX_PATTERN_HERE>`

### Examples

1. list all posts mentioning `bash scripting` type in 
    
    `journal -f 'bash scripting'`
2. list all posts made on 24th of April
    
    `journal -f '04-24'`
3. list all posts containing the word 'review' or 'github'
    
    `journal -f 'review|github'`

## Moving the datafile
Journal uses plain text file to store journal posts. You can view,
edit, share version control or do whatever you like.

Journal has no restrictions on that. 
