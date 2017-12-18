## Regular

>Compile expression<br>

    rgx, _ := regexp.Compile("(br|n|gr)exit")

>Match input string against compiled expression - Choice depends on purpose

    rgx.findAllString("frexit, brexit, grexit, nextit", -1)

_&nbsp;&nbsp;&nbsp;&nbsp;  Array of matches_

    rgx.Match([]byte("brexit"))
    
_&nbsp;&nbsp;&nbsp;&nbsp;  Boolean_

    rgx.MatchString("brexit")

_&nbsp;&nbsp;&nbsp;&nbsp; Boolean_
