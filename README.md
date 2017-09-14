# hermes
## goal
### hermes orchestrates playgrounds with revision control
## motivation
personally as a developer i really want one concrete feature
- on `git checkout -b experiment` I get a sandboxed environment deployed accessible to QA (for example) through a private namespace (e.g. "hermes.www.experiment") and on subsequent changes to that branch the environment behaves as a normal CI/CD pipeline would

## previous implementations
you'll find a host of rather vanilla implementations of ci/cd for the usual suspects like docker and vagrant and special cases like jekyll and launchd for my own personal website in the commit history here

## current design
currently there is a redesign in the works for 0.1.0 that centers everything around fsnotify and the .git directory structure. it's essentially a reverse proxy that watches .git/refs/heads

![demo](demo.png)

## why are you using shell commands for everything
because it's the main api i know how to hook into (i like scripting) and i'm using it as a layer of indirection to the REST of these providers. it shouldn't be permanent, and i plan to refactor them to use their proper compile time api's

## why not use some combination of (tool1,...)
the idealistic answet is that i'd like to have some room for my own free and creative expression, and i don't want to bolt together a non-orthogonal combination of black-boxes together while learning nothing and piling cruft on top of cruft. that answer doesn't really satisfy the incurious. a servile and artistically barren malady grips our industry

the pragmatic answer is that this particular implementation is dependent on a bunch of stuff already, so here are some primitives that i have no plans to build myself
- unix
- git
- go
- fsnotify
- jekyll
- vagrant
- docker
- gce
- aws

the current master only uses the top 5, but i plan to unify all the branches with pluggable interfaces and would like to support the entire board

i am open to the idea that modular parts of this project's goal are best done by some other tools, in particular terraform or docker-machine to abstract over all the infrastructure providers and provision remote resources

i am not open to the idea of "not worrying about machines", i have a personal interest in the entire cloud stack, so most of this work will probably be at the level of "CaaS" and lower

what this aims to be in the short to medium term is a reference implementation of an interface for my own personal website's (again, if you want to call it that) needs and dependencies, then use that as a poc to add support for other development frameworks to make it developer agnostic with as few opinions as possible

## influences
clearly this workflow is inspired by Heroku. they've really done some excellent stuff. true artists

the domain of interest is certainly similar to what Hashicorp products offer, and i plan to use many of them. but i think i see a different gestalt than they, and the focus is on development and not production 

the current design was heavily influenced by adg and bradfitz's hacking session here: https://github.com/golang/tools/blob/master/cmd/tip/tip.go
