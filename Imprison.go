{{/*
	Designed for servers with dedicated mute roles and channels. Will store the original roles
	within a pinned comment on the respecive mute channel and remove all previous roles. 
	Also adds a random unused mute role. If no unused mute role exists, then this would fail out.
	There is an optional command that can be passed to force in a particular role regardless of
	the number of accounts that have that role.
	
	Usage:
		/mute @user. ***NOT YET IMPLEMENTED***
		/mute @user @muteRole
		
	Future Edits:
		Move muted roles to database with special tag so it can be accessed across commands
		
		Add check to see if user is already muted, otherwise the user will have all their old roles
		erase in the db and reassigned the muted role upon being muted
*/}}

{{/* Configurable values */}}
{{ $mutedRoleChannelDict := sdict "882063366259105812" "882068893051543675" "989696938775572521" "989696717874167888" }}
{{ $mutedRoles := cslice "882063366259105812" "989696938775572521" }}
{{ $roleToAssign := "" }}
{{ $continueMuting := 1 }}
{{ $previousRoles := "" }}
{{ $member := "" }}
{{/* End of configurable values */}}

{{if ge (len .CmdArgs) 1}}
	{{if ge (len .CmdArgs) 2}}
		{{ $roleToAssign = (getRole (index .CmdArgs 1)) }}
		
		{{ if not $roleToAssign }}
			**Error:** Invalid Role: ``{{(index .CmdArgs 1) }}``
			{{ $continueMuting = 0 }}
		{{ end }}
	{{ end }}
	
	{{ if $continueMuting }}
		{{ $member = (getMember (index .CmdArgs 0)) }}

		{{ if $member }}
			{{ $previousRoles = $member.Roles }}
			
			{{ range $index, $role := $member.Roles }}
				{{- takeRoleID $member.User.ID $role -}} 
			{{- end }}
			
			{{ dbSet $member.User.ID "previousRoles" $previousRoles }}
			{{ if $roleToAssign }}
				{{ giveRoleID $member.User.ID $roleToAssign.ID }}
			{{ else }}
				**Error:** Assigning random unused role still under construction
			{{ end }}
		{{ else }}
			**Error:** Invalid Member ``{{ (index .CmdArgs 0) }}``
			{{ $continueMuting = 0 }}
		{{ end }}
	{{ end }}
{{ end }}

{{ if $continueMuting }}
	{{ print "<@" $member.User.ID ">" "has been imprisoned in <#" ($mutedRoleChannelDict.Get (print $roleToAssign.ID)) ">" }}
{{ end }}
