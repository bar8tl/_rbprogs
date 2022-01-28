cd c:\rbhome\go\src\bar8tl\p\
md bomcomp
cd bomcomp
copy c:\c_portab\01_rb\_rbprogs\go-bomcomp\pkg\bitmlocal.go     .
copy c:\c_portab\01_rb\_rbprogs\go-bomcomp\pkg\bommatch.go      .
copy c:\c_portab\01_rb\_rbprogs\go-bomcomp\pkg\gobals.go        .
copy c:\c_portab\01_rb\_rbprogs\go-bomcomp\pkg\settings.go      .
copy c:\c_portab\01_rb\_rbprogs\go-bomcomp\pkg\sqlstatements.go .
copy c:\c_portab\01_rb\_rbprogs\go-bomcomp\pkg\uom.go           .
go install
pause
