# go-http-client

#### 
1. Http Client, support following method
Get
Post
Delete
Put
....

2. Easy to add path param, header
`client.Post("/path", "/subPath", "/thirdPath").Header("User-Agent", "My-Client").AddParam("Param", "ParamValue").Execute()`

3. Support Json Entity
`client.Post().JsonEntity(<struct{}>).Execute()`

4. Easy to debug
Request&Response are easy to check
```
request url >>>>> https://www.google.com?key=valu
request header >>>>> map[User-Agent:[my-client]]
response status code <<<<< 200
response header <<<<< map[Cache-Control:[private, max-age=0] Content-Type:[text/html; charset=ISO-8859-1] P3p:[CP="This is not a P3P policy! See https://www.google.com/support/accounts/answer/151657?hl=en for more info."] Server:[gws] X-Frame-Options:[SAMEORIGIN] Set-Cookie:[NID=89=gyjRzw_bfEUEPAgnFKqFsr2oUydsLrJwnhLdIvl-Nxbq4rMa5wIQmWeWQpiR1lgr0tPLQFdD1ZtU8Uj745iBMvc8erFSVH9GFdTKXHeAOWpRRVigQ4MU7HAp5XGyvdIfwQkWl377tRujQA; expires=Fri, 28-Apr-2017 08:34:51 GMT; path=/; domain=.google.com; HttpOnly] Date:[Thu, 27 Oct 2016 08:34:51 GMT] Expires:[-1] X-Xss-Protection:[1; mode=block] Alt-Svc:[quic=":443"; ma=2592000; v="36,35,34"]]
response body <<<<< <!
```
5. Get reponse body
```golang
respone, error := client.GET().Execute()
if error != nil {
	fmt.Println("reponse body", respone.Payload)
}
```

6. Response to Entity
```golang
entity  := <struct{}>
client.Post().JsonEntity(<struct{}>).ExecuteForEntity(&entity)
```

