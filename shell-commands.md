# Shell Commands

This is a quick reference on shell commands that can be used to make requests to an API.

# Bash

## Get Request

```bash
curl <url> -H "Authorization: <token>"
```

## Post Request

```bash
curl <url> -X POST -H "Authorization: <token>" -d '{"json": "data"}'
```

## Put Request

```bash
curl <url> -X PUT -H "Authorization: <token>" -d '{"json": "data"}'
```

## Delete Request

```bash
curl <url> -X DELETE -H "Authorization: <token>"
```


# PowerShell

## Get Request

```powershell
$headers = @{
    "Authorization" = "<token>"
}
Invoke-RestMethod -Uri <url> -Method Get -Headers $headers
```

## Post Request

```powershell
$headers = @{
    "Authorization" = "<token>"
}
$body = @{
    "json" = "data"
} | ConvertTo-Json
Invoke-RestMethod -Uri <url> -Method Post -Headers $headers -Body $body
```

## Put Request

```powershell
$headers = @{
    "Authorization" = "<token>"
}
$body = @{
    "json" = "data"
} | ConvertTo-Json
Invoke-RestMethod -Uri <url> -Method Put -Headers $headers -Body $body
```

## Delete Request

```powershell
$headers = @{
    "Authorization" = "<token>"
}
Invoke-RestMethod -Uri <url> -Method Delete -Headers $headers
```