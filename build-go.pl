#!/usr/bin/env perl
use v5.10;
use strict;
use warnings;
use File::Basename qw/fileparse/;
use File::Spec;
use File::stat;
use Getopt::Long;

my $verbose;
GetOptions( "verbose|v" => \$verbose ) or die("Error in command line arguments\n");

unless (@ARGV) {
    @ARGV = <./ch*>;
}

for my $arg (@ARGV) {
    if ( -f $arg ) {
        _build_go($arg);
    }
    elsif ( -d $arg ) {
        _build_go($_) for map { File::Spec->canonpath($_) } <$arg/*.go>;
    }
    else {
        warn "unknown: $arg\n";
    }
}

sub _build_go {
    my $src = shift;
    unless ( -f $src ) {
        warn "unknown: $src";
        return;
    }
    unless ( -d 'bin' ) {
        mkdir 'bin';
    }
    my $bin = "bin/" . fileparse( $src, qr/\.go/ );
    if ( -f $bin && stat($bin)->mtime > stat($src)->mtime ) {
        say "$bin up to date" if $verbose;
    }
    else {
        my $cmd = "go build -o $bin $src";
        say($cmd);
        system($cmd);
    }
}

