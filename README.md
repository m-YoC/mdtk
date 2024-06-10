
# mdtk: Markdown CodeBlock Task Runner 

mdtk is a task runner using code block of Markdown.

## Install

Download the appropriate binary for your environment and move it to $PATH directory such as `/usr/local/bin/`.




## [mdtk command help]

- command:  
    - mdtk group-task
    - mdtk group-task [--options] -- args...

### * How to write each command
- group-task -> group task, group:task, task
    - Can write without group. In this case, all groups will be searched.
    - Special group-task names are as follows.
        - '_' group : Searchs only empty-name group.
        - 'help' task : Show task help. If write group, show only written group's tasks.
        - empty     : If 'default' task is defined, run it. Otherwise, run 'help'.
- args       -> arg_name=arg_value
    - Write after two hyphens.
    - {$} is a special variable for --script option, replaced by positional parameter $1...$9.

### * options
    --file, -f  [+value]      Select a task file.
    --nest, -n  [+value]      Set the nest maximum depth of embedded comment (embed/task).
                              Default is 20.
    --quiet, -q               Task output is not sent to standard output.
    --all-task, -a            Can select private groups and hidden tasks at the command.
                              Or show all tasks that include private groups and hidden tasks  at task help.
    --script, -s              Display script.
    --make-cache, -c          Make taskdata cache.
    --lib, -l  [+value]       Select a library file.
                              This is a special version of --file option.
                              No need to add an extension '.mdtklib'.
    --make-library  [+value]  Make taskdata library.
                              Value is library name.
    --version, -v             Show version.
    --groups, -g              Show groups.
    --help, -h                Show command help.
    --manual, -m              Show mdtk manual.
    --write-configbase        Write config base file to current directory.



## [mdtk manual]

### * Basics of Task Definition
The definition of a task is written in the following code block.  
Commands in the code block are not independent and run as a single script (run in your SHELL environment).  
In other words, you are merely writing a shell, bash, or other script within the code block.  

~~~markdown
```task:<group>:<task> -- <description>

# Write your script...

```
~~~

It can also be written with language aliases added, as follows.

~~~markdown
```bash  task:<group>:<task> -- <description>

# Write your script...

```
~~~

The characters that can be used in \<group> and \<task> are as follows.
- Lower Alphabets, Upper Alphabets, Numbers, '_', '-' and '.'
- First character is only Lower Alphabets, Upper Alphabets and '_'

\<group> and \<description> can be empty. (ex: ```task::\<task> ~)  
'\_' group is as same as empty.  
Group that first is '_' and the length is over two is a private group.   
Private groups cannot run from command directly.

'--' before \<description> is not necessary, but it is better for visibility.  
Also, the number of '-' need not be two, but any number of one or more.

#### ** Attributes
You can write attributes in the beginning of \<description> using '[...]'.  
  - format: `[attr1 attr2 ...] description...`
  - Set space between each attribute
  - Attributes List
    - `hidden`: The task gets same effects as private group.
    - `t`     : In the task help, the task will be written on the upper side of the group.
    - `b`     : In the task help, the task will be written on the lower side of the group.


#### example
~~~markdown
```task::hello_world -- [t] This is Detail Text.
THIS=mdtk
echo "Hello $THIS World!"
```

# -> Hello mdtk World!
~~~


### * Variables
No special configuration is required to use task variables.  
The variables given in the command are expanded at the beginning of the script as just variable definitions.  
You simply use the variable in your script.  

In the command, as follows, write the variables and its values after '--'.
- mdtk ... -- ARG1=value1 ARG2=value2 ARG3=value3 ...

mdtk does not check for the existence of variables during script generation.  
It leaves this to the SHELL environment at runtime.

{$} is a special variable for --script option, replaced by positional parameter $1...$9.  
It cannot be used in normal task run.


### * Embedded Comments
Some comments written as '#xxxx> comment' have a special behavior.

~~~
#embed>  <group>:<task>                  : The selected task is directly embedded.
#task>   <group>:<task> -- args          : The selected task is embedded as a subshell.
#task> @ <group>:<task> -- args          : The config once flag is temporarily reset.
                                           The rest is the same as without @.
#func> <funcname> <group>:<task> -- args : The selected task is embedded as a subshell type function.
                                           The config once flag is temporarily reset.
                                           If you want the function to have arguments, 
                                           pass positional parameters to the task with <args>.
                                           (ex: #func> hello g:t -- arg1=$1 arg2=$2)
#config> once                            : When called multiple times, it is called only the first time.
#desc>   comments                        : Show comments as additional description in task help.
#args>   comments                        : Show comments as arguments in task help.
~~~


### * Filename
mdtk uses Taskfile.md in the current directory unless you select a file path with --file/-f flag.  
Instead of Taskfile.md, you can use *.taskrun.md.  
In this case, however, only one *.taskrun.md file should be placed in the same directory. 

Search Order: --file path -> Taskfile.md -> *.taskrun.md  


### * Sub Taskfile
You can read sub taskfile in the following code block.  
There must be no duplicate group/task combinations throughout all read files.

~~~markdown
```taskfile

# Write sub taskfile path...

```
~~~


### * Taskfile config
#### ** Group Order
You can set integers that control the order of groups in the following code block.  
This configuration is used in task help.  
In task help, the higher the order value, the higher its group is displayed.  
The order value is zero, if not set.   
However, '_' (nameless) group will be the top of task help unless set number explicitly.  

~~~markdown
```taskconfig:group-order
<groupname>: <integer> 
```
~~~

In this configuration, only what is written to the root file is used.  


### * Task Cache
You can make a taskdata cache by setting --make-cache option.  
The cache 'may' speed up task reading.  
If the cache already exists, it will be read automatically with no option.  
However, note that if taskfiles have some updates at this time, the cache will be remade
 and run-speed is slower than it without cache.  
Good to use for taskfiles that are being updated less frequently.  
To disable the cache, simply delete the relevant cache.  

You can also make a taskdata library. It is similar to cache and use --make-library option.  
Differences of the library and the cache are as follows.  
- The library is not read automatically and not remade automataically.  
- Select library file directly using --file/-f or --lib/-l option.  
- Therefore, can use it alone. Not need to hold '.md' taskfiles at the same place.  
- Embedded comments of all public tasks will be expanded.  
- All private tasks will be removed.  
- All file-path data of tasks will be removed.  

In 'taskfile' code block, you cannot read the cache/library.  
Please use mdtk command in task when you use them at outside.  
