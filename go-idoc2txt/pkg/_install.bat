cd c:\rbhome\go\src\bar8tl\p\
md rbidoc
cd rbidoc
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\config.go   .
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\crtdb.go    .
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\deflts.go   .
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\envmnt.go   .
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\idocdata.go .
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\names.go    .
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\outsqlt.go  .
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\query.go    .
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\settings.go .
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\unfold.go   .
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\uplddefs.go .
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\upldmitm.go .
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\upldsgma.go .
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\upldsgrp.go .
copy c:\c_portab\01_rb\_rbprogs\go-idoc2txt\pkg\upldssgm.go .
go install
pause
