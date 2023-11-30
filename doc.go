/*
Package [mkerrors] contains no useable Go source code. Instead, it contains a 
portable script (Perl 5) that is invoked via [go generate] and emits Go source
code containing type definitions, global variable declarations, and default
"Error()" methods for each error parsed from a comment block following the
"go:generate" directive.

Usage

[The script] accepts no arguments. The following rules describe how identifiers
are parsed from Go source code comments:

 - Lines following "//go:generate" are parsed until the first uncommented line.
 - The first word in each commented line is used as the identifier for a
   generated type of error (and corresponding global var of that type).
 - Any words following the first on each commented line are ignored.

Although it may be placed anywhere system-wide, I recommend keeping [the script]
alongside version-controlled source files to share project-specific changes with
your team and to make updates without affecting other projects using it.

Example

Given a Go package "foo" at local directory path "/foo", create a new empty Go
source file named "bar.go" that contains only the following:

		//go:generate perl mkerrors.pl
		// InvalidReceiver
		// InvalidArgument
		// OutOfRange
		// WriteOverflow
		// ReadOverflow

Now, running "go generate" from the command line will overwrite the contents of
"bar.go" with elaborated definitions:

		// Code generated by mkerrors.pl. DO NOT EDIT.
		
		package bar
		
		//go:generate perl mkerrors.pl
		// InvalidReceiver
		// InvalidArgument
		// OutOfRange
		// WriteOverflow
		// ReadOverflow
		
		type (
			InvalidReceiver struct{}
			InvalidArgument struct{}
			OutOfRange      struct{}
			WriteOverflow   struct{}
			ReadOverflow    struct{}
		)
		
		var (
			ErrInvalidReceiver InvalidReceiver
			ErrInvalidArgument InvalidArgument
			ErrOutOfRange      OutOfRange
			ErrWriteOverflow   WriteOverflow
			ErrReadOverflow    ReadOverflow
		)
		
		func (e *InvalidReceiver) Error() string {
			return "invalid receiver"
		}
		
		func (e *InvalidArgument) Error() string {
			return "invalid argument"
		}
		
		func (e *OutOfRange) Error() string {
			return "out of range"
		}
		
		func (e *WriteOverflow) Error() string {
			return "write overflow"
		}
		
		func (e *ReadOverflow) Error() string {
			return "read overflow"
		}

Feel free to modify the generated comment block containing the original error 
identifiers. Running "go generate" a subsequent time will reflect your changes 
accordingly without any unnecessary duplication or outdated content.

[the script]: mkerrors.pl
[go generate]: https://go.dev/blog/generate
*/
package mkerrors
