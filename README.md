# 新Flutter Analyze

本bitrise stepは`bitrise-steplib/bitrise-step-flutter-analyze` を変更したもので、Flutter projectにおける `flutter analyze` commandで実行される。


## 変更内容

### 1. Bitrise flutter-analyzeの問題点
infoレベル error|warning|info のうちinfoレベルしか存在しない場合でも、出力メッセージに'error' ワードが存在するとerrorありと判定してしまう。

![before log](doc_image/before_log.png)

### 2. Bitrise flutter-analyzeの変更
flutter-analyze結果の標準出力行の先頭infoレベル error|warning|info のみでerror有無を判定する。
この変更の有効/無効は引数で指定可能とする。

![edit log](doc_image/edit_log.png)

## 開発環境

### 1. ベンダリング

カスタムステップでの依存ライブラリをリポジトリ内に含める

Go Module / dep

Go Moduleはそのまま使うとvendorディレクトリには入りませんが、 go mod vendor コマンドを使うとvendorディレクトリを利用してくれます。そのため、こちらを使い、vendorをコミットすることでカスタムステップを実行することができます。

```
go mod init
go mod vendor
```

dep を使った例

公式のカスタムステップではdepを使っているケースが多いため、depを使う方法についても記述しておきます。こちらも使うのは簡単で、以下のコマンドで実行できます。

```
dep init
dep ensure
```

