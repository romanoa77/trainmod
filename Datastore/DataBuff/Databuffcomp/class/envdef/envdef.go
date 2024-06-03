package envdef

import "os"

// filesystem
const envadm = "ADMROOT"
const envadmn = "ADMFNM"

const envdata = "DATAROOT"
const envlog = "LOGROOT"

const envlogstrm = "LOGSTREAM"

// network
const envserv = "SERVURL"

//filesystem

var Baseadm = os.Getenv(envadm)
var Baseadmn = os.Getenv(envadmn)

var Basedata = os.Getenv(envdata)

var Baselog = os.Getenv(envlog)
var Strmlogn = os.Getenv(envlogstrm)

//network

var Basesrvurl = os.Getenv(envserv)
