
~~~task:lib:hello_world hello testlib
#embed> _lib:task1
#embed> _lib:task2

echo $(task1) $(task2) 
~~~

~~~task:_lib:task1
function task1 () {
    echo hello
}
~~~

~~~task:_lib:task2
function task2 () {
    echo world
}
~~~

