<!doctype html>
<style>
.table {
    border-collapse:collapse;
    border-spacing:0px; 
    border:1px solid #FF0000;
}

.table tr td:nth-child(1) {
    width:300px;
}

.table tr td {
    border:1px solid #000000;
}
</style>
<html>
    <head>
        <title>PRINT PDF TEST</title>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    </head>
    <body>
    <table class="table">
        <tr>
            <td>品名</td>
            <td>口径</td>
        </tr>
        {{range $i, $data := .Datas}}
            <tr>
                <td>{{$data.Name}}</td>
                <td>{{$data.Cal}}</td>
            </tr>
        {{end}}
    </table>
    <br>
    <img src="your favorite img" />
    </body>
</html>
