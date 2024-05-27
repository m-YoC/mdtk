
# mdtk: Markdown CodeBlock Task Runner 

mdtk is a markdown task-runner using codeblock.

## [mdtk command help] --help, -h

- command:  
    - mdtk group-task
    - mdtk group-task [--options] -- args...

### * How to write each command
- group-task -> group task, group:task, task
    - Can write without group. In this case, all groups will be searched.
    - Special group-task names are as follows.
        - '_' group : Searchs only empty-name group.
        - 'help'    : Show task help.
        - empty     : If 'default' task is defined, run it. Otherwise, run 'help'.
- args       -> arg_name=arg_value
    - Write after 2 underbars.

### * options

    --file, -f  [+value]      Specify a task file.
    --nest, -n  [+value]      Set the nest maximum times of embedded comment (embed/task).
                              Default is 20.
    --debug                   Show run-script.
    --help, -h                Show command help.
    --md-help                 Show Markdown taskfile help.
    --task-help-all           Show all tasks that include private groups at task help.

## [Markdown taskfile help] --md-help

The definition of a task is written in the following code block.  
Scripts in a code block act as a series of multi-line ShellScript.

~~~markdown
```task:<group>:<task>  <description>

# Write your script...

```
~~~

The letters that can be used in \<group> and \<task> are as follows.
- Lower Alphabets, Upper Alphabets, Numbers, '_', '-' and '.'
- First letter is only Lower Alphabets, Upper Alphabets and '_'

\<group> and \<description> can be empty. (ex: ```task::\<task> ~)  
'\_' group is as same as empty.  
Group that first is '_' and the length is over 2 is a private group.   
Private groups cannot run from command directly.


### * Embedded Comments
Some comments written as '#xxxx> comment' have a special behavior.

~~~
#embed>  <group>:<task>         : The specified task is directly embedded.
#task>   <group>:<task> -- args : The specified task is embedded as a subshell.
#task> @ <group>:<task> -- args : The config once flag is temporarily reset.
                                  The rest is the same as without @.
#config> once                   : When called multiple times, it is called only the first time.
#args>   comments               : Show comments as arguments in task help.
~~~

### * Filename
mdtk uses Taskfile.md in the current directory unless you specify a file path with --file/-f flag.  
Instead of Taskfile.md, you can use *.taskrun.md.  
In this case, however, only one *.taskrun.md file should be placed in the same directory. 

Search Order: --file path -> Taskfile.md -> *.taskrun.md  


### * Sub Taskfile
You can load sub taskfile in the following code block.  
There must be no duplicate group/task combinations throughout all loaded files.

~~~markdown
```taskfile

# Write sub taskfile path...

```
~~~
