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
    <li class="nav-item"><a class="nav-link active" href="#">Edit</a></li>
    <li class="nav-item"><a class="nav-link" href="../create">Create</a></li>
</ul>
<body>
    <!--曜日のチェックボックスを現在の設定に合わせてチェックするJS-->
    <script type="text/javascript">
        function reflectCheckbox(d, id) {
            var dlist = ['mon', 'tue', 'wed', 'thu', 'fri', 'sat', 'sun'];
            for (var i = 0; i < dlist.length; i++) {
                r = d.indexOf(dlist[i]);
                if (r !== -1) {
                    document.getElementById(dlist[i] + id).checked = true;
                } else {
                    document.getElementById(dlist[i] + id).checked = false;
                }
            }
        }
    </script>
    <h2>Google Homeをしゃべらせるやつを管理するやつ</h2>
    <h3>編集</h3>
    <h4>1個ずつ更新してね<br>（最後の1個はさくじょできません）</h4>

    <table border="1">
        <tr>
            <th>ID</th>
            <th>時</th>
            <th>分</th>
            <th>月</th>
            <th>日</th>
            <th>曜日</th>
            <th>しゃべる内容</th>
            <th>削除</th>
            <th>更新</th>
        </tr>
        {{range .}}
        <form action="/edit" method="POST">
            <tr>
                <td><input type="hidden" name="id" value="{{ .ID }}" />{{ .ID }}</td>
                <td><input type="text" size="5" name="hour" pattern="(0[0-9]|1[0-9]|2[0-3])" value="{{ .Hour }}" required /></td>
                <td><input type="text" size="5" name="min" pattern="([0-5][0-9])" value="{{ .Min }}" required /></td>
                <td><input type="hidden" name="month" value="{{ .Month }}" /> {{ .Month }}</td>
                <td><input type="hidden" name="date" value="{{ .Date }}" />{{ .Date }}</td>
                <td><input type="checkbox" id="mon{{ .ID }}" name="days[]" value="mon" />月
                    <input type="checkbox" id="tue{{ .ID }}" name="days[]" value="tue" />火
                    <input type="checkbox" id="wed{{ .ID }}" name="days[]" value="wed" />水
                    <input type="checkbox" id="thu{{ .ID }}" name="days[]" value="thu" />木
                    <input type="checkbox" id="fri{{ .ID }}" name="days[]" value="fri" />金
                    <input type="checkbox" id="sat{{ .ID }}" name="days[]" value="sat" />土
                    <input type="checkbox" id="sun{{ .ID }}" name="days[]" value="sun" />日</td>
                <td><input type="text" size="120" name="text" value="{{ .Text }}" required /></td>
                <td><input type="checkbox" name="delete" value="true" /></td>
                <td><input type="submit" value="更新" /></td>
            </tr>
            <script type="text/javascript">
                reflectCheckbox({{ .Days }}, {{ .ID }});
            </script>
        </form>
        {{end}}
    </table>
</body>
</html>
