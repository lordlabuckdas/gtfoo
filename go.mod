module gtfoo

go 1.15

replace (
	github.com/lordlabuckdas/gtfoo/gtfobins => ./gtfobins
	github.com/lordlabuckdas/gtfoo/lolbas => ./lolbas
)

require (
	github.com/lordlabuckdas/gtfoo/gtfobins v0.0.0-00010101000000-000000000000
	github.com/lordlabuckdas/gtfoo/lolbas v0.0.0-00010101000000-000000000000
)
