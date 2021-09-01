# Design Draft

## Model & Casbin

* Menu go struct

```go
type Menu struct {
    ID uint64
    ParentID string // 父级
    UUID string // 唯一编号
    Name string // 名字（非前端显示名)
    Description string
    LocaleKey string // i18n键 "menu.users"
    Link string // 链接
    Linkable string // 是否连接
    Position string // scope针对父级
}
```

* Menu casbin model

```ini
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
```

* front end casbin policy

```text
p, uid, menuid, show
```
