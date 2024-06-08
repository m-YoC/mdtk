#!/bin/bash
set -euo pipefail

/bin/bash << 'EOS'
echo aaa 
read num
echo $num
EOS
