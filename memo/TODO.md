
# TODO

- v0.1.0以降
    - [x] version表示オプション追加 (--version, -v)
        - version情報は外部ファイルに`VERSION=xx.yy.zz`の形で記載してgo:embedで読み込む
    - [ ] group名の衝突を回避する手段を実装する
        - taskfile読み込み時にaliasを付けられると言い？
        - そもそもmdtk in mdtkで読み込めば別データ扱いで管理できるのでいらないかもしれない
        - scriptを全部読み込んで展開してバイナリで保持しておく.cache機能
    - [x] command helpとmarkdown helpを外部ファイルに出してgo:embedで読み込むように変更

