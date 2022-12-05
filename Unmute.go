{{/*
	Designed for servers with dedicated mute roles and channels.
	Based on the mute role setting. Will read the pinned comment and remove the muted role and
	assign back all the previous roles they had. Only works if the user has one of the muted roles.
	
	Usage:
		/unmute @user
	
	Future Edits:
		Move muted roles to database with special tag so it can be accessed across commands
*/}}

{{/* Configurable values */}}
{{ $mutedRoleChannelDict := sdict "1048391685215899768" "1049179984343879741" "1049180079713947720" "1049180201801756692" }}
{{ $mutedRoles := cslice "1048391685215899768" "1049180079713947720" }}
{{ $foundMutedRole := 0 }}
{{ $previousRoles := "" }}
{{ $member := "" }}
{{/* End of configurable values */}}

{{if ge (len .CmdArgs) 1}}
	{{ $member = (getMember (index .CmdArgs 0)) }}
	
	{{ range $mutedRoles }}
		{{- if (targetHasRoleID $member.User.ID .) -}}
			{{- $foundMutedRole = . -}}
			{{- break }}
		{{- end -}}
	{{- end -}}
		
	{{ if $foundMutedRole }}
		{{/* No roles will be saved on unmute*/}}
		{{ takeRoleID $member.User.ID $foundMutedRole }}
		
		{{ $previousRoles =(dbGet $member.User.ID "previousRoles") }}
		
		{{ if $previousRoles }}
			{{ range $previousRoles.Value }}
				{{- giveRoleID $member.User.ID . -}} 
			{{- end }}
		{{ else }}
			**Error:** No previous roles found for user
		{{ end }}
		
	{{ else }}
		**Error:** User not currently muted
	{{ end }}
{{ end }}

{{ if $foundMutedRole }}
	{{ print "<@" $member.User.ID ">" "has been released from <#" ($mutedRoleChannelDict.Get $foundMutedRole) ">" }}
{{ end }}
