# [pflag](https://github.com/spf13/pflag) for [logrus](https://github.com/sirupsen/logrus)

##

```go

logLevel := logrus.InfoLevel

cmd := &cobra.Command{
    Use: "my-command",
    RunE: func(cmd *cobra.Command, args []string) error {
        
        logrus.SetLevel(logLevel)
        
        return cmd.Usage()
    },
}

logruspflag.LevelVar(cmd.Flags(), &logLevel, "log-level", logLevel, "Log level")

```
