<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>ghn-admin-go Top Page</title>
    <!-- BootstrapのCSS読み込み -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
    <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.3/umd/popper.min.js" integrity="sha384-ZMP7rVo3mIykV+2+9J3UJ46jBk0WLaUAdn689aCwoqbBJiSnjAK/l8WvCWPIPm49" crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/js/bootstrap.min.js" integrity="sha384-ChfqqxuZUCnJSK3+MXmPNIyE6ZbWh2IMqE241rYiqJxyMiZ6OW/JmZQ5stwEULTy" crossorigin="anonymous"></script>
</head>
<ul class="nav nav-tabs">
    <li class="nav-item"><a class="nav-link active" href="#">Home</a></li>
    <li class="nav-item"><a class="nav-link" href="../edit">Edit</a></li>
    <li class="nav-item"><a class="nav-link" href="../create">Create</a></li>
</ul>
<!--<h4>メニュー</h4>
<form>
    <input type="button" class="btn btn-default" onclick="location.href='../edit'" value="編集" />
    <input type="button" class="btn btn-default" onclick="location.href='../create'" value="新規登録" />
</form>
-->
<br><br>
<h3>ジョブ一覧</h3>
<div class="container">
    <table class="table table-striped table-hover">
        <thead>
            <tr>
                <th>ID</th>
                <th>分</th>
                <th>時</th>
                <th>日</th>
                <th>月</th>
                <th>曜日</th>
                <th>しゃべる内容</th>
            </tr>
        </thead>
        <tbody>
            {{range .}}
            <tr>
                <th scope="row">{{ .ID }}</th>
                <td>{{ .Hour }}</td>
                <td>{{ .Min }}</td>
                <td>{{ .Date }}</td>
                <td>{{ .Month }}</td>
                <td>{{ .Days }}</td>
                <td>{{ .Text }}</td>
            </tr>
            {{end}}
        </tbody>
    </table>
</div>
</body>
</html>
