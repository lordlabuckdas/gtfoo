module gtfoo

go 1.15

replace (
	github.com/lordlabuckdas/gtfoo/gtfobins => ./gtfobins
	github.com/lordlabuckdas/gtfoo/utils => ./utils
)

require (
	github.com/lordlabuckdas/gtfoo/gtfobins v0.0.0-00010101000000-000000000000
)
