# hermes
## goal
### hermes orchestrates playgrounds with revision control
## motivation
personally as a developer i really want one concrete feature
- on `git checkout -b experiment` I get a sandboxed environment deployed accessible to QA (for example) through a private namespace (e.g. "hermes.www.experiment") and on subsequent changes to that branch the environment behaves as a normal CI/CD pipeline would

## current implementation
right now master compiles down to a post-commit git hook which builds a docker container for my personal public website (if you want to call it that) with jekyll and go, then deploys that service locally

![demo](http://atec.pub/demo.png)

## why are you using shell commands for everything
because it's the main api i know how to hook into (i like scripting) and i'm using it as a layer of indirection to the REST of these providers. it shouldn't be permanent, and i plan to refactor them to use their proper compile time api's

## why not use some combination of (tool1,...)
i'd like to have some room for my own free and creative expression, and i don't want to be held hostage by black boxes and bolt together a non-orthogonal combination of them together while learning nothing and piling cruft on top of cruft. that answer doesn't really satisfy the incurious. a servile and artistically barren malady grips our industry

the first answer is that this particular implementation is dependent on a bunch of stuff already, so there are already some primitives that i have no plans to build myself
- git
- go
- docker
- gce
- aws
- launchd

i am open to the idea that modular parts of this project's goal are best done by some other tools, in particular terraform or docker-machine to abstract over all the infrastructure providers and provision remote resources

i am not open to the idea of "not worrying about machines", i have a personal interest in the entire cloud stack, so most of this work will probably be at the level of "CaaS" and lower

what this aims to be in the short to medium term is a reference implementation of an interface for my own personal website's (again, if you want to call it that) needs and dependencies, then use that as a poc to add support for other development frameworks to make it developer agnostic with as few opinions as possible (besides git. use git)


## influences
clearly this workflow is inspired by Heroku. they've really done some excellent stuff. true artists
