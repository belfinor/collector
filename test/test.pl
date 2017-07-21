#!/usr/bin/perl


use strict;
use warnings;
use IO::Socket::INET;


srand(time);



my $sock = IO::Socket::INET->new(
      PeerAddr => '127.0.0.1',
      PeerPort => $ARGV[0] || 5001,
      Proto    => 'tcp', 
      Timeout  => 2,
);



while( 1 ) {
  my $str = sprintf( "%d%d", time, int(rand(1000)) );
  my $size = length $str;
  my $data = pack "na*", $size, $str;
  if( $sock->send( $data ) != $size + 2 ) {
    warn 'data not send';
  }
}

