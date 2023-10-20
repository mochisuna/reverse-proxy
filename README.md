# reverse-proxy
x-forwarded-forの挙動確認。

ホストマシンを `192.168.1.1` としたときに以下のようなネットワーク設定になっています。

| サーバー名 | 説明 | IP |
| --- | --- | --- |
| Proxy01 | 192.168.1.2 | ホストマシンからのリクエストをProxy02にリダイレクトする |
| Proxy02 | 192.168.1.3 | Proxy01からのリクエストを任意のターゲットに伝搬する |
|   host  | 192.168.1.4 | ターゲット01 |
|   app   | 192.168.1.5 | ターゲット02 |

`./proxy/nginx.conf` の `proxy_pass` を変更して、任意のサーバーを対象にリダイレクトしてください。

`curl localhost:8080` を実行すると以下のようなアクセスログが記録されます（対象が `192.168.1.4` の場合）。

```sh
local_time:25/Aug/2023:04:00:41 +0000	host:192.168.1.3	forwardedfor:192.168.1.1, 192.168.1.2	req:GET / HTTP/1.0	status:200	size:13	referer:-	ua:curl/7.88.1	reqtime:0.000	cache:-	runtime:-	vhost:localhost	method:GET	uri:/
```
