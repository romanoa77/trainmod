package envdef

import "os"

// filesystem
const envadm = "ADMROOT"
const envadmn = "ADMFNM"

const envdata = "DATAROOT"
const envdtnm = "DATANM"
const envlog = "LOGROOT"

const envlogstrm = "LOGSTREAM"

const envres = "RESROOT"

const envdtrt = "DTVOLROOT"

// network
const envserv = "SERVURL"

//filesystem

var Baseadm = os.Getenv(envadm)
var Baseadmn = os.Getenv(envadmn)

var Basedata = os.Getenv(envdata)
var Datanm = os.Getenv(envdtnm)

var Baselog = os.Getenv(envlog)
var Strmlogn = os.Getenv(envlogstrm)

var Baseres = os.Getenv(envres)

var Basevoldt = os.Getenv(envdtrt)

//network

var Basesrvurl = os.Getenv(envserv)
