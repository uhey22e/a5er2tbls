# a5er2tbls

DBドキュメント生成ツールの[tbls](https://github.com/k1LoW/tbls)のために、
[A5:SQL](https://a5m2.mmatsubara.com/)で作成したER図から外部キー制約の情報を補完するためのツールです。

ER図上ではテーブル間に線を引いているが、実際のDBでは外部キー制約を付けていないときに役立つと思います。

## 使い方

`-i` オプションでA5ERファイルを指定します。標準出力に、YAML形式のtbls設定が出力されるので、これをtblsの設定ファイルにコピーします。

```sh
$ a5er2tbls -i erd.a5er
```

`-o` オプションを指定すると、ファイルに出力することも可能です。

```sh
$ a5er2tbls -i erd.a5er -o relations.yml
```
