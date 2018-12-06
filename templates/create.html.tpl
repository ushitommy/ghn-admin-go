<!DOCTYPE html>
<html lang="ja"></html>
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
    <li class="nav-item"><a class="nav-link" href="../">Home</a></li>
    <li class="nav-item"><a class="nav-link" href="../edit">Edit</a></li>
    <li class="nav-item"><a class="nav-link active" href="#">Create</a></li>
    <li>Google Homeをしゃべらせるやつを管理するやつ</li>
</ul>
<body>
    <h4>新規登録</h4>
    ==============<br>
    <b>実行条件：</b>
    <form action="/create" method="POST">
        曜日：
        <input type="checkbox" name="days[]" value="mon" />月
        <input type="checkbox" name="days[]" value="tue" />火
        <input type="checkbox" name="days[]" value="wed" />水
        <input type="checkbox" name="days[]" value="thu" />木
        <input type="checkbox" name="days[]" value="fri" />金
        <input type="checkbox" name="days[]" value="sat" />土
        <input type="checkbox" name="days[]" value="sun" />日<br>
        <input type="hidden" name="date" value="*" />
        <input type="hidden" name="month" value="*" />
        時間：
        <input type="text" size="5" name="hour" pattern="(0[0-9]|1[0-9]|2[0-3])" required />時
        <input type="text" size="5" name="min" pattern="([0-5][0-9])" required />分<br>
        しゃべらせる内容：
        <input type="text" size="120" name="text" required /><br>
        <input type="submit" value="追加" />
    </form>
    ==============<br>
    <br><br>
    <h3>ジョブ一覧</h3>
    <table border="1">
        <tr>
            <th>ID</th>
            <th>分</th>
            <th>時</th>
            <th>日</th>
            <th>月</th>
            <th>曜日</th>
            <th>しゃべる内容</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{ .ID }}</td>
            <td>{{ .Hour }}</td>
            <td>{{ .Min }}</td>
            <td>{{ .Date }}</td>
            <td>{{ .Month }}</td>
            <td>{{ .Days }}</td>
            <td>{{ .Text }}</td>
        </tr>
        {{end}}
    </table>
    <br><br>
    <form>
        <input type="button" onclick="location.href='../'" value="Topにもどる" />
    </form>
</body>
</html>
