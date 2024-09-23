# zconv

### 🤔 这是啥
这是一个 golang 的小工具类。   
没有找到特别适合自己的，就自个儿写了一个。  
目前仅支持字符串转大驼峰、小驼峰和蛇形，有没有别的功能完全看我需不需要（：D）

### 🤔 怎么用

- 转换大驼峰
   ```
  # ToCamel
  dst := zconv.ToCamel(ori)
  # To.Camel
  dst := zconv.ToDotCamel(ori)
  # To Camel
  dst := zconv.ToSpaceCamel(ori)
   ```
- 转换小驼峰
   ```
  # toCamel
  dst := zconv.ToLowerCamel(ori)
  # to.Camel
  dst := zconv.ToLowerDotCamel(ori)
  # to Camel
  dst := zconv.ToLowerSpaceCamel(ori)
   ```
- 转换蛇形
   ```
  # to_camel
  dst := zconv.ToSnack(ori)
   ```