var Groupify = angular.module('Groupify', ['timer']);

Groupify.controller('MainCtrl', function($scope, $http, $timeout) {
  var spotifyApi = new SpotifyWebApi();
  $scope.queue = [];
  $scope.query = "";
  $scope.trackResults = [];
  //$scope.current_track = { artist: "Marisa Monte", name: "Barulhinho Bom", time_remaining: 97 };

  (function tick() {
    $http.get('/api/v1/queue/list')
    .then(function(res){

      // sum time to play up to each track in collection
      var sum = 0;

      $scope.current_track = res.data.now_playing.track;
      if ($scope.current_track && res.data.now_playing.time_remaining) { 
        sum = $scope.current_track.time_remaining = parseInt( res.data.now_playing.time_remaining );
        //console.log( "time remaining: " + $scope.current_track.time_remaining + "s" );
      }

      $scope.queue = res.data.queue;
      for(var i = 0; i < $scope.queue.length; i++) {
        track = $scope.queue[i];
        track.time_to_play = sum;
        sum += parseInt(track.time);
      }

      $timeout(tick, 1000);
    });
  })();

  $scope.search = function(){
    spotifyApi.searchTracks($scope.query, {limit: 10, offset: 0}, function(err, data) {
      $scope.trackResults = data.tracks.items;
    });
  };

  $scope.enqueue = function(track){
    $http.post('/api/v1/queue/add', {
      track_id: track.id
    })
    .then(function(res){
      console.log("Enqueued track " + track.name);
    });
  };

  $scope.dequeue = function(track){
    $http.post('/api/v1/queue/delete', {
      track_id: track.id
    })
    .then(function(res){
      // FIXME: handle errors
      console.log("De-queued track " + track.name);
    });
  };

});

Groupify.filter('secondsToTime', function() {
  // shameless copy/paste
  // http://codeaid.net/javascript/convert-seconds-to-hours-minutes-and-seconds-(javascript)
  return function(secs) {
    var hr  = Math.floor(secs / 3600);
    var min = Math.floor((secs - (hr * 3600))/60);
    var sec = secs - (hr * 3600) - (min * 60);

    if (hr  < 10) { hr  = "0" + hr; }
    if (min < 10) { min = "0" + min;}
    if (sec < 10) { sec = "0" + sec;}
    if (hr)       { hr  = "00"; }
    return hr + ':' + min + ':' + sec;
  };
});

