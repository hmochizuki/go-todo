# curl による動作確認

## ユーザ操作

```
# 作成
curl -X POST http://localhost:8080/user \
-H "Content-Type: application/json" \
-d '{"name":"ユーザー名","Email":"test@example.com","password":"password123"}'



```


## TODO操作

```
# 全てのTODO取得
curl -X GET http://localhost:8080/todos

# 作成
curl -X POST http://localhost:8080/todo \
-H "Content-Type: application/json" \
-d '{"name":"Test Task","userId":"user id"}'

# スタータス更新
curl -X POST http://localhost:8080/todo/14/status \
-H "Content-Type: application/json" \
-d '{"status":"doing"}'

# 削除(論理削除)
curl -X DELETE http://localhost:8080/todo/2
```