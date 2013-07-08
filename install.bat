rem Please set IUP_HOME first

set CGO_CFLAGS=-I%IUP_HOME%\include
set CGO_LDFLAGS=-L%IUP_HOME%
path=%IUP_HOME%;%path%

go install -x github.com/Archs/go-iup...