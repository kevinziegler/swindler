# Commands
## Add
Store a new database connection profile
```
sw add <adapter> <name>
   --adapter
   --user
   --host
   --db
   --password
   --wizard
   --adapter-args
   --hook-setup
   --hook-cleanup
```

## Remove
Remove a stored connection
```
sw remove <name>
```

## Rename
Rename a connection profile
```
sw rename <name> <new name>
```

## Upate
Update a setting for a connection profile
```
sw update <name>
```

## Connect
Connect to a database using a given profile
```
sw connect <name or prefix>
```

## List
List connections matching an optional prefix (or show all connections)
```
sw list [ prefix ]
```

## Show
Show a connection profile
```
sw show <name>
```

## Forget host
Destroy all connections associated with a given host
```
sw forget-host <host>
```

## Copy a connection
Copy a configuration profile
```
sw copy <name> <clone-name>
```

## Test a connection
Tests if a given connection profile can connect to a host
```
sw test <name>
```
