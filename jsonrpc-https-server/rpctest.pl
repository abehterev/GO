#!/bin/env perl

$ENV{PERL_LWP_SSL_VERIFY_HOSTNAME} = 0;

   use JSON::RPC::Client;
   
   my $client = new JSON::RPC::Client;
   my $uri = 'http://localhost:4443/rpc';
   
   my $callobj = {
      method  => 'HelloService.Sum',
      #params  => [ 17, 25 ],
      params => { a => 20, b => 10 },
      version => '2.0',
      id => 1,
   };
   
   my $res = $client->call($uri, $callobj);
   
   if($res) {
      if ($res->is_error) {
          print "Error : ", $res->error_message, " (", $res->content, ")";
      }
      else {
          print $res->result;
      }
   }
   else {
      print $client->status_line;
   }
