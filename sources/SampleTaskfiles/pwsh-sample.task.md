
# PowerShell Test

~~~powershell task::ps-hello
echo hello
~~~

~~~powershell task:pseb:t -- [hidden] embedded comment test termination
echo hello
~~~

~~~powershell task:pseb:tco -- [hidden] embedded comment test termination (config once)
#config> once
echo hello
~~~

~~~powershell task:pseb:targ -- [hidden] embedded comment test termination (with args)
#args> a:(string)
try {
    if ($a -eq $null) {
        $a = 'default'
    }
} catch {
    $a = 'default'
}; 
echo "hello / arg: $a"
~~~

## `#embed>`

~~~powershell task:pseb:embed-test
#embed> pseb:t
~~~


## `#func>`

- Embedding as a Function

~~~powershell task:pseb:func-test
#func> tt pseb:t
tt
~~~

- with args

~~~powershell task:pseb:funcarg-test
#func> tt pseb:targ -- a='func-arg'
tt
~~~

- with special parameter (required positional parameter type) 
    - Set `$a = $Arg[x];`

~~~powershell task:pseb:funcarg2-test
#func> tt pseb:targ -- a=<$>
echo '* with positional parameter'
tt func-arg
echo '* no positional parameter'
tt
echo 'end'
~~~

- with special parameter (optional positional parameter type) 
    - Set `$a = try{$Arg[x]}catch{$null};`

~~~powershell task:pseb:funcarg3-test
#func> tt pseb:targ -- a=<?>
echo '* with positional parameter'
tt func-arg
echo '* no positional parameter'
tt
echo 'end'
~~~


## `#config> once`

- The second time is not embedded

~~~powershell task:pseb:once1-test
echo '* first'
#embed> pseb:tco
echo '* second'
#embed> pseb:tco
echo 'end'
~~~

- The second time is also embedded
    - Once flag is temporarily reset in `#func>`

~~~powershell task:pseb:once2-test
echo '* first'
#func> tt pseb:tco
tt
echo '* second'
#func> tt2 pseb:tco
tt2
echo 'end'
~~~

## `#desc>`

- For display in task help

~~~powershell task:pseb:desc-test
#desc> This is a description of '#desc>'.
#desc> 2nd row.
echo hello
~~~

## `#args>`

- For display in task help

~~~powershell task:pseb:args-test
#args> a:(xxx) b:(xxx)
#desc> Text of '#args>' is simply a type of task help description.
echo hello
~~~


