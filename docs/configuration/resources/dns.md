# DNS(sakuracloud_dns / record)


**全ゾーン共通のグローバルリソースです。**

`sakuracloud_dns`がDNSゾーン設定を、`sakuracloud_dns_record`が対象ゾーン内のレコードを表しています。

### 設定例

```
resource "sakuracloud_dns" "dns" {
    zone = "example.com"
}
resource "sakuracloud_dns_record" "record01" {
    dns_id = "${sakuracloud_dns.dns.id}"
    name = "test1"
    type = "A"
    value = "192.168.0.1"
}
resource "sakuracloud_dns_record" "record02" {
    dns_id = "${sakuracloud_dns.dns.id}"
    name = "test1"
    type = "A"
    value = "192.168.0.2"
}
```

## `sakuracloud_dns`

### パラメーター

|パラメーター         |必須  |名称                |初期値     |設定値                    |補足                                          |
|-------------------|:---:|--------------------|:--------:|------------------------|----------------------------------------------|
| `zone`            | ◯   | 対象DNSゾーン        | -        | 文字列                  | - |
| `description`     | -   | 説明  | - | 文字列 | - |
| `tags`            | -   | タグ | - | リスト(文字列) | - |

### 属性

|属性名          | 名称             | 補足                                        |
|---------------|-----------------|--------------------------------------------|
| `id`          | ID              | -                                          |
| `zone`        | 対象DNSゾーン     | -                                          |
| `description` | 説明             | -                                          |
| `tags`        | タグ             | -                                          |
| `dns_servers` | DNSサーバー       | 対象DNSゾーンの委譲先となるネームサーバーのリスト  |

## `sakuracloud_dns_record`

### パラメーター

|パラメーター  |必須  |名称          |初期値   |設定値                    |補足                                          |
|------------|:---:|--------------|:------:|------------------------|----------------------------------------------|
| `dns_id`   | ◯   | DNSゾーンID   | -      | 文字列                  | 対象DNSゾーンのID |
| `name`     | ◯   | レコード名     | -      | `ホスト名`<br />`@` | 英字（小文字）,数字,一部記号（-.@_*）,1～63文字, @は当該ゾーンを示す。|
| `type`     | ◯   | タイプ        | -      | `A`<br />`AAAA`<br />`NS`<br />`CNAME`<br />`MX`<br />`TXT`<br />`SRV` | - |
| `value`    | ◯   | 値           | -      | 文字列 | タイプ`A`:IPアドレス<br />タイプ`AAAA`:IPv6アドレス<br />`NS`:一部記号（_）, 末尾ピリオド, 1～63文字<br />タイプ`CNAME`:一部記号（_）, 末尾ピリオド, 1～63文字<br />タイプ`MX`:一部記号（_）, 末尾ピリオド, 1～63文字<br />タイプ`TXT`:英字, 数字, 半角スペース, 一部記号, 1～255文字<br />タイプ`SRV`:一部記号（_.-）, 末尾ピリオド, 1～63文字|
| `ttl`      | -   | TTL          | `3600` | 数値 | `10`～`3600000`秒 |
| `priority` | -   | プライオリティ | `10`   | 数値 | タイプが`MX`、`SRV`の場合のみ有効。`1`〜`65535` |
| `weight`   | -   | 重み | -   | 数値 | タイプが`SRV`の場合のみ有効。`0`〜`65535` |
| `port`     | -   | ポート | -   | 数値 | タイプが`SRV`の場合のみ有効。`1`〜`65535` |


### 属性

|属性名       | 名称             | 補足 |
|------------|-----------------|------|
| `id`       | ID              | -  |
| `dns_id`   | DNSゾーンID      | -  |
| `name`     | レコード名        | -  |
| `type`     | タイプ            | - |
| `value`    | 値               | -  |
| `ttl`      | TTL             | -  |
| `priority` | プライオリティ    | -  |
| `weight`   | 重み    | -  |
| `port`     | ポート    | -  |

