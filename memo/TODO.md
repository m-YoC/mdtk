
# TODO

- v0.1.0以降
    - [x] version表示オプション追加 (--version, -v)
        - version情報は外部ファイルに`VERSION=xx.yy.zz`の形で記載してgo:embedで読み込む
    - [ ] group名の衝突を回避する手段を実装する
        - taskfile読み込み時にaliasを付けられると言い？
        - そもそもmdtk in mdtkで読み込めば別データ扱いで管理できるのでいらないかもしれない
    - [x] scriptを全部読み込んでバイナリで保持しておく.cache機能
    - [x] command helpとmarkdown helpを外部ファイルに出してgo:embedで読み込むように変更
    - [x] helpのPAGER表示 (今のところ長くなっているmarkdown helpのみ)
    - [x] PAGER表示の拡張
        - cmd helpとtask helpも
        - 改行数をカウントして、長くなったらPAGER表示に自動的に切り替わるようにする
- v0.2.0以降
    - [x] scriptの書き出し用にdebugとは別に生成されたスクリプトを標準出力に出すだけのオプションを用意する
        - 実行はされない
        - [x] 変数に`{$}`を入れると位置パラメータ(positional parameter) `$1 ~ $9`が順番に入るようにする
    - [ ] pager表示を調整する
        - pager表示しなくてもいいようにする
        - PAGER環境変数がなければ何もしないようにするべき？
        - 起動行数指定もいるため、コンフィグファイルで変えられるようにするべき
    - [x] 指定グループのタスクヘルプだけ見れるようにする
    - [x] shorthandオプションをまとめて書けるように変更
        - `-abc`
    - [x] stdoutに出力しないquietモードを用意する
    