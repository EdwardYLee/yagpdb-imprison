# yagpdb-imprison

Custom YAGPDB Command to move users to and from dedicated mute channels. This command removes all previous roles and saves them within the yagpdb database.


Usage:
  /imprison @User @MutedRole
  
  /release @User


Current Issues:
  * Assigning random muted roles has yet to be implemented.
  * List of muted roles and muted channels need to be manually added to each script. Will need to create a new command that adds these to the yagpdb database.
