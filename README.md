# go-xdg
This doesn't really have a reason to exist. I wanted to play around and see
how to get go packages to import into another one.

---
Go has a built in function for getting XDG paths ```os.UserConfigDir()``` and
```os.UserCacheDir()``` but it didn't have state or data (I also didn't try
looking for them so maybe they exist) but I also didn't want to use the Darwin
paths. So basically this is just a small package to get specific pseudo-XDG
paths
