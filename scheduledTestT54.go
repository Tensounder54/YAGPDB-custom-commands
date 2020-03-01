{{if eq (len .Args) 2}}
    {{$team := (index .CmdArgs 0)}}
    {{$dateTime := currentTime}}
    {{$months := (cslice "January" "Febuaray" "March" "April" "May" "June" "July" "August" "September" "October" "November" "December")}}
    {{$month := 0}}
    {{range $itterOne, $months = (cslice "January" "Febuaray" "March" "April" "May" "June" "July" "August" "September" "October" "November" "December")}} {{if eq $months $dateTime.Month.String}} {{$month = add ($itterOne 1)}} {{end}} {{end}}
    {{if and (or (hasRoleName "Fauna Captain") (hasRoleName "Fauna GM")) (eq $team "Fauna")}}
        {{$channel := 678975645371334656}}
        {{$tuesday := joinStr "" "Tuesday - " "03" "/" (toString $month) "/" (toString $dateTime.Year) " 1900 GMT/2000 CET - Scrim ⚔️"}}
        {{$id := (sendMessageNoEscapeRetID $channel $tuesday)}}
        {{addMessageReactions $channel $id "✅" "❓" "❌"}}
        {{$wednesday := joinStr "" "Wednesday - " "04" "/" (toString $month) "/" (toString $dateTime.Year) " 1900 GMT/2000 CET - VoD Review 📽️"}}
        {{$id := (sendMessageNoEscapeRetID $channel $wednesday)}}
        {{addMessageReactions $channel $id "✅" "❓" "❌"}}
        {{$friday := joinStr "" "Friday - " "06" "/" (toString $month) "/" (toString $dateTime.Year) " 1900 GMT/2000 CET - Practice 🎯/Scrim  ⚔️"}}
        {{$id := (sendMessageNoEscapeRetID $channel $friday)}}
        {{addMessageReactions $channel $id "✅" "❓" "❌"}}
    {{else}}
        {{print $team}}
    {{end}}
{{else}}
    correct syntax is -schedule <team>
{{end}}
