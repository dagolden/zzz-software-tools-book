#!/usr/bin/env perl
use v5.10;
use strict;
use warnings;
use File::Basename qw/fileparse/;
use File::stat;

unless (@ARGV) {
    die "Usage: $0 <file|directory> [file|directory] ...\n";
}

for my $arg (@ARGV) {
    if ( -f $arg ) {
        _build_go($arg);
    }
    elsif ( -d $arg ) {
        _build_go($_) for <$arg/*.go>;
    }
    else {
        warn "unknown: $arg\n";
    }
}

sub _build_go {
    my $src = shift;
    warn "unknown: $src" unless -f $src;
    my $bin = "bin/" . fileparse( $src, qr/\.go/ );
    if ( stat($bin)->mtime > stat($src)->mtime ) {
        say "$bin up to date";
    }
    else {
        my $cmd = "go build -o $bin $src";
        say($cmd);
        system($cmd);
    }
}

