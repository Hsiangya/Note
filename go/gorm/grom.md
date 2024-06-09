# CURD

## 创建

### 基础创建

```go
// 使用结构体插入数据
user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
result := db.Create(&user) // 通过数据的指针来创建
user.ID             // 返回插入数据的主键
result.Error        // 返回 error
result.RowsAffected // 返回插入记录的条数

// 指定字段创建
db.Select("Name", "Age", "CreatedAt").Create(&user)
// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")

// 忽略字段创建
db.Omit("Name", "Age", "CreatedAt").Create(&user)
// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")

// 批量创建
var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
db.Create(&users)
for _, user := range users {
  user.ID // 1,2,3
}

// 指定批量插入大小
var users = []User{{Name: "jinzhu_1"}, ...., {Name: "jinzhu_10000"}}
db.CreateInBatches(users, 100)	// batch size 100

// CreateBatchSize 选项初始化GORM实例后，此后进行创建& 关联操作时所有的INSERT行为都会遵循初始化时的配置。
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{  CreateBatchSize: 1000})
db := db.Session(&gorm.Session{CreateBatchSize: 1000})
users = [5000]User{{Name: "jinzhu", Pets: []Pet{pet1, pet2, pet3}}...}
db.Create(&users)
// INSERT INTO users xxx (5 batches)
// INSERT INTO pets xxx (15 batches)  一条user包含3条pet


// 根据map创建，当使用map来创建时，钩子方法不会执行，关联不会被保存且不会回写主键。
db.Model(&User{}).Create(map[string]interface{}{"Name": "jinzhu", "Age": 18})
db.Model(&User{}).Create([]map[string]interface{}{
  {"Name": "jinzhu_1", "Age": 18},
  {"Name": "jinzhu_2", "Age": 20},
})	// batch insert from `[]map[string]interface{}{}`
```

### SQL 表达式、Context Valuer 创建

```go
// 使用 SQL 表达式、Context Valuer 创建记录
// Create from map
db.Model(User{}).Create(map[string]interface{}{
  "Name": "jinzhu",
  "Location": clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 100)"}},
})
// INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"));

// Create from customized data type
type Location struct {
    X, Y int
}

// Scan implements the sql.Scanner interface
func (loc *Location) Scan(v interface{}) error {
  // Scan a value into struct from database driver
}

func (loc Location) GormDataType() string {
  return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
  return clause.Expr{
    SQL:  "ST_PointFromText(?)",
    Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
  }
}

type User struct {
  Name     string
  Location Location
}

db.Create(&User{
  Name:     "jinzhu",
  Location: Location{X: 100, Y: 100},
})
// INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"))
```

### 关联创建

创建关联数据时，如果关联值非零，这些关联会被upsert，并且它们的`Hooks`方法也会被调用。

```go
type CreditCard struct {
  gorm.Model
  Number   string
  UserID   uint
}

type User struct {
  gorm.Model
  Name       string
  CreditCard CreditCard
}

db.Create(&User{
  Name: "jinzhu",
  CreditCard: CreditCard{Number: "411111111111"}
})
// INSERT INTO `users` ...
// INSERT INTO `credit_cards` ...


// select，Omit方法跳过关联更新
db.Omit("CreditCard").Create(&user)
db.Omit(clause.Associations).Create(&user) // skip all associations
```

### 默认值

```go

// 传递默认值
type User struct {
  ID   int64
  Name string `gorm:"default:galeone"`
  Age  int64  `gorm:"default:18"`
}
// 当结构体的字段默认值是零值的时候比如 0, '', false，这些字段值将不会被保存到数据库中，使用指针类型或者Scanner/Valuer来避免这种情况。
type User struct {
  gorm.Model
  Name string
  Age  *int           `gorm:"default:18"`
  Active sql.NullBool `gorm:"default:true"`
}
// 若要让字段在数据库中拥有默认值则必须使用defaultTag来为结构体字段设置默认值。如果想要在数据库迁移的时候跳过默认值，可以使用 default:(-)，示例如下：
type User struct {
  ID        string `gorm:"default:uuid_generate_v3()"` // db func
  FirstName string
  LastName  string
  Age       uint8
  FullName  string `gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"`
}
```

### 钩子

- BeforeSave
- BeforeCreate
- AfterSave
- AfterCreate

```go
// 定义钩子
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.UUID = uuid.New()

    if u.Role == "admin" {
        return errors.New("invalid role")
    }
    return
}

// 跳过钩子
DB.Session(&gorm.Session{SkipHooks: true}).Create(&user)
DB.Session(&gorm.Session{SkipHooks: true}).Create(&users)
DB.Session(&gorm.Session{SkipHooks: true}).CreateInBatches(users, 100)
```

## 查询

### 查询单个对象

- 想避免`ErrRecordNotFound`错误，你可以使用`Find`，比如`db.Limit(1).Find(&user)`，`Find`方法可以接受struct和slice的数据。
- 对单个对象使用`Find`而不带limit，`db.Find(&user)`将会查询整个表并且只返回第一个对象，只是性能不高并且不确定的。
- `First` and `Last` 方法会按主键排序找到第一条记录和最后一条记录 (分别)。 只有在目标 struct 是指针或者通过 `db.Model()` 指定 model 时，该方法才有效。 此外，如果相关 model 没有定义主键，那么将按 model 的第一个字段进行排序。 例如：

```go
// 获取第一条记录（主键升序）
db.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1;

// 获取一条记录，没有指定排序字段
db.Take(&user)
// SELECT * FROM users LIMIT 1;

// 获取最后一条记录（主键降序）
db.Last(&user)
// SELECT * FROM users ORDER BY id DESC LIMIT 1;

result := db.First(&user)
result.RowsAffected // 返回找到的记录数
result.Error        // returns error or nil

// 检查 ErrRecordNotFound 错误
errors.Is(result.Error, gorm.ErrRecordNotFound)


var user User
var users []User

// works because destination struct is passed in
db.First(&user)
// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

// works because model is specified using `db.Model()`
result := map[string]interface{}{}
db.Model(&User{}).First(&result)	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1


// doesn't work
result := map[string]interface{}{}
db.Table("users").First(&result)

// works with Take
result := map[string]interface{}{}
db.Table("users").Take(&result)

// no primary key defined, results will be ordered by first field (i.e., `Code`)
type Language struct {
  Code string
  Name string
}
db.First(&Language{})	// SELECT * FROM `languages` ORDER BY `languages`.`code` LIMIT 1

// 根据主键检索
db.First(&user, 10)	// SELECT * FROM users WHERE id = 10;
db.First(&user, "10")	// SELECT * FROM users WHERE id = 10;
db.Find(&users, []int{1,2,3})	// SELECT * FROM users WHERE id IN (1,2,3);

var user = User{ID: 10}
db.First(&user)	// SELECT * FROM users WHERE id = 10;
var result User
db.Model(User{ID: 10}).First(&result)	// SELECT * FROM users WHERE id = 10;
```

