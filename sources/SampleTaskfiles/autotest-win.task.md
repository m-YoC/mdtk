
# Windows version

~~~taskfile
./autotest.task.md
~~~

# Bash

~~~bash task::test -- [hidden priority:1]
#args> task:(group:task) quiet:(bool)
echo -n '*'
if [ "$quiet" == "true" ]; then
    mdtk.exe -f ./first-sample.taskrun.md $task > /dev/null 2>&1
    if [ $? -ne 0 ]; then echo "failed: $task"; fi
else
    echo " $task"
    mdtk.exe -f ./first-sample.taskrun.md $task
    if [ $? -ne 0 ]; then echo "failed: $task"; fi
fi
~~~

~~~bash task::ntest -- [hidden priority:1]
#args> task:(group:task) quiet:(bool)
echo -n '*'
if [ "$quiet" == "true" ]; then
    mdtk.exe -f ./first-sample.taskrun.md $task > /dev/null 2>&1
    if [ $? -ne 1 ]; then echo "failed: $task"; fi
else
    echo " $task"
    mdtk.exe -f ./first-sample.taskrun.md $task
    if [ $? -ne 1 ]; then echo "failed: $task"; fi
fi
~~~

# PowerShell

~~~powershell task::pstest -- [hidden priority:1]
#args> task:(group:task) quiet:(bool)
Write-Host -NoNewline "*"
if ($quiet -eq 'true') {
    mdtk.exe -f ./pwsh-sample.task.md $task > $null 2>&1
    if (!$?) { echo "failed: $task" }
} else {
    echo " $task"
    mdtk.exe -f ./pwsh-sample.task.md $task
    if (!$?) { echo "failed: $task" }
}
~~~

~~~powershell task::psntest -- [hidden priority:1]
#args> task:(group:task) quiet:(bool)
Write-Host -NoNewline "*"
if ($quiet -eq 'true') {
    mdtk.exe -f ./pwsh-sample.task.md $task > $null 2>&1
    if ($?) { echo "failed: $task" }
} else {
    echo " $task"
    mdtk.exe -f ./pwsh-sample.task.md $task
    if ($?) { echo "failed: $task" }
}
~~~

