# Here
When you are doing some log based debugging there are two things that
you often want to do; log where you are in your code, and log the value
of something. While you can write your own print statements, sometimes
you just want to stick a bunch of useful prints in that you will remove
later. That is where these two functions come in.

## here.Here()
Here() logs the name of the name of the file, line in that file and
the function that it was called from:

```[LOG] 0:00.000 INFO here github.com/juju/juju/worker/cleaner_test.(*CleanerSuite).TestCleaner(0xc20809f110, 0xc2080a83c0)
[LOG] 0:00.000 INFO here 	/home/dooferlad/dev/go/src/github.com/juju/juju/worker/cleaner/cleaner_test.go:58 +0xdd
[LOG] 0:00.000 INFO here github.com/dooferlad/here.Is(0xe02b80, 0xc20809f110)```

## here.Is(v interface{})
Same as Here, but also prints the value of a variable that you give it:

```[LOG] 0:00.000 INFO here 	/home/dooferlad/dev/go/src/github.com/dooferlad/here/here.go:32 +0x1f
[LOG] 0:00.000 INFO here 		&cleaner_test.CleanerSuite{BaseSuite:testing.BaseSuite{CleanupSuite:testing.CleanupSuite{testStack:testing.cleanupStack{(testing.CleanupFunc)(0x4fca60)}, suiteStack:testing.cleanupStack(nil)}, LoggingSuite:testing.LoggingSuite{}, JujuOSEnvSuite:testing.JujuOSEnvSuite{oldJujuHome:"", oldHomeEnv:"/home/dooferlad", oldEnvironment:map[string]string{"JUJU_HOME":"", "JUJU_ENV":"", "JUJU_LOGGING_CONFIG":"", "JUJU_DEV_FEATURE_FLAGS":""}, initialFeatureFlags:""}}, mockState:(*cleaner_test.cleanerMock)(0xc2080b67d0)}```

The other nice thing is that since you are using something in a namespace
that is specific for creating debug logs, you can easily find uses in your
code to delete before checking in.