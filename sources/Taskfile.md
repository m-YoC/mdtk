# mdtk sample

~~~taskconfig:group-order
embed: 1
~~~

## Base Test

```bash task::hello_world  -- aaa
#desc> hello world
#desc> test sample
THIS=mdtk
echo "Hello $THIS World!"
read num
echo $num
echo ion
#embed> test
```

```bash task::test -- [t] mdtk first test 日本語のテキストサンプルです
#desc> hello mdtk
echo hello mdtk! wd:`pwd`

echo embed
#embed> sub:subtest
echo task
#task> sub:subtest
echo task @
#task> @ sub:subtest

echo func
#func> ttt sub:subtest
ttt

echo mdtk in mdtk
mdtk -f ./SubTaskfiles/Taskfile.md sub subtest
mdtk -F ./SubTaskfiles/Taskfile.md sub subtest
```

```taskfile
./SubTaskfiles/Taskfile.md
```

```md task::test2 -- [t]
#desc> hello
#config> once
echo "hello mdtk! (config once)"
```

## Embedded Comment Test

```task:embed:embed_test   mdtk embed test
echo "- embed test -"
#embed> test
```

```task:embed:subtask_test   mdtk subtask test
echo "- subtask test -"
#task> test
```

```task:embed:configonce_test   mdtk config once test
echo "- config once test -"
#embed>  test2
#embed>  test2
```

```task:embed:configonce_test2   mdtk config once test2 (task)
echo "- config once test -"
#embed>  test2
#task>   test2
#task>   test2
#embed>  test2
```

```task:embed:embed_args_test   embedded coment args is used at help
echo "- embed args test -"
#args> hello args
echo hello mdtk!
```

## Arguments Test

```task:args:arg_test   mdtk arg test (args_ex: -- a1=hello a2=world)
echo "- arg test -"
#args> a1:string a2:string
#desc> test sample
echo a1=$a1 a2=$a2
#embed> embed:embed_args_test
```

```task:args:task_arg_test   mdtk task arg test
echo "- task arg test -"
#task> args:arg_test -- a1=hello a2=world
```

## Recursive mdtk Test (mdtk in mdtk)

```task:rec:rec_test   mdtk recursive test (mdtk in mdtk)
echo "- mdtk recursive test -"
mdtk task_arg_test
```

```task:rec:rec_test2   mdtk recursive test2 (mdtk in mdtk)
echo "- mdtk recursive test -"
nest=10
mdtk rec_test2impl -n $nest -- nest=$nest
```

```task:rec:rec_test2impl   mdtk recursive test2 implement (mdtk in mdtk)
echo nest count $nest
mdtk rec_test2impl -n $((nest - 1)) -- nest=$((nest - 1))
echo h
```

## mdtk Subfile Test

```taskfile
./subfile.task.md
```

```task:subfile:subfiletest mdtk sub taskfile test
#embed> subtest
```


