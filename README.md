
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
        - 'help'    : Show task help.
        - empty     : If 'default' task is defined, run it. Otherwise, run 'help'.
- args       -> arg_name=arg_value
    - Write after 2 hyphens.

### * options
    --file, -f  [+value]      Specify a task file.
    --nest, -n  [+value]      Set the nest maximum times of embedded comment (embed/task).
                              Default is 20.
    --make-cache, -c          Make taskdata cache.
    --debug                   Show run-script.
    --version, -v             Show version.
    --help, -h                Show command help.
    --manual, -m              Show mdtk manual.
    --task-help-all           Show all tasks that include private groups at task help.



## [mdtk manual]

### * Basics of Task Definition
The definition of a task is written in the following code block.  
Commands in the code block are not independent and run as a single script (run in your SHELL environment).  
In other words, you are merely writing a shell, bash, or other script within the code block.  

~~~markdown
```task:<group>:<task>  <description>

# Write your script...

```
~~~

The characters that can be used in \<group> and \<task> are as follows.
- Lower Alphabets, Upper Alphabets, Numbers, '_', '-' and '.'
- First character is only Lower Alphabets, Upper Alphabets and '_'

\<group> and \<description> can be empty. (ex: ```task::\<task> ~)  
'\_' group is as same as empty.  
Group that first is '_' and the length is over 2 is a private group.   
Private groups cannot run from command directly.


#### example
~~~markdown
```task::hello_world
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


### * Task Cache
You can make a taskdata cache by specifying --make-cache option.  
The cache 'may' speed up task reading.  
If the cache already exists, it will be read automatically with no option.  
However, note that if taskfiles have some updates at this time, the cache will be remake, 
 which is slower than if there was no cache.
Good to use for taskfiles that are being updated less frequently.  
To disable the cache, simply delete the relevant cache.  
